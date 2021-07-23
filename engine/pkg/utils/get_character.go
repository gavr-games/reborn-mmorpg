package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Character struct {
	Id int
	Name string
	Gender string
}

func GetCharacter(r *http.Request) (char *Character, ok bool) {
	keys, ok := r.URL.Query()["token"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Chat Url Param 'token' is missing")
		return nil, false
	}
	token := keys[0]
	keys2, ok2 := r.URL.Query()["character_id"]
	if !ok2 || len(keys2[0]) < 1 {
		log.Println("Chat Url Param 'character_id' is missing")
		return nil, false
	}
	character_id := keys2[0]

	url := "http://api:4567/api/v1/characters/" + character_id
	apiClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	req.Header.Set("Authorization", token)

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Println(err)
		return nil, false
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(err)
		return nil, false
	}

	character := Character{}
	jsonErr := json.Unmarshal(body, &character)
	if jsonErr != nil {
		log.Println(err)
		return nil, false
	}

	return &character, true
}