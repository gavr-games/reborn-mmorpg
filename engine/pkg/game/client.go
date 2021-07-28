package game

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"

	"github.com/gorilla/websocket"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
    origin := r.Header.Get("Origin")
    return origin == os.Getenv("SERVER_ORIGIN")
	},
}

// Client is a middleman between the websocket connection and the engine.
type Client struct {
	engine *Engine

	// The websocket connection.
	conn *websocket.Conn

	// Character info
	character *utils.Character

	// Buffered channel of outbound messages.
	send chan []byte
}

// ClientCommand is a message sent to engine from clients
type ClientCommand struct {
	// Character Id to identify from which player the command is coming from
	characterId int

	// player command in json
	command map[string]interface{}
}

// readPump pumps messages from the websocket connection to the engine.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.engine.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var command map[string]interface{}
		json.Unmarshal(message, &command)
		cc := &ClientCommand{characterId: c.character.Id, command: command}
		c.engine.commands <- cc
	}
}

// writePump pumps messages from the engine to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The engine closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(engine *Engine, w http.ResponseWriter, r *http.Request) {
	character, ok := utils.GetCharacter(r)
	if !ok {
		log.Println("Unable to get character for chat ws")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{engine: engine, conn: conn, character: character, send: make(chan []byte, 256)}
	client.engine.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
