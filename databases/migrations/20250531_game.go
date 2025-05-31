package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {

		gamesCollection, err := app.FindCollectionByNameOrId("games")
		if err != nil {
			return err
		}

		gameItemsCollection, err := app.FindCollectionByNameOrId("game_items")
		if err != nil {
			return err
		}

		gameItemsCollection.Fields.Add(&core.RelationField{
			Name:         "game",
			CollectionId: gamesCollection.Id,
			Required:     true,
		})

		if err := app.Save(gameItemsCollection); err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		return nil
	})
}
