package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
		"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func main() {
    app := &cli.App{
        Name:  "engine",
        Usage: "Different engine commands",
				Commands: []*cli.Command{
					{
						Name:    "gm:set",
						Usage:   "Set game object as Game Master (GM) by id",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "game_object_id", Aliases: []string{"goid"}},
						},
						Action: func(cCtx *cli.Context) error {
							id := cCtx.String("game_object_id")
							if id != "" {
								gameObj := storage.GetClient().GetGameObject(id)
								gameObj.SetProperty("game_master", true)
								storage.GetClient().SaveGameObject(gameObj)
								fmt.Println("Player with id=", id, "is now a Game Master")
								return nil
							} else {
								return cli.Exit("Please provide game_object_id", 2)
							}
						},
					},
				},
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}