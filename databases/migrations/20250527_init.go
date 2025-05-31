package migrations

import (
	"context"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		gamesCollection := core.NewBaseCollection("games")
		gamesCollection.Fields.Add(&core.TextField{
			Name:   "name",
			System: true,
		}, &core.TextField{
			Name:   "description",
			System: true,
		}, &core.URLField{
			Name:   "image",
			System: true,
		}, &core.TextField{
			Name:   "slug",
			System: true,
		}, &core.SelectField{
			Name:   "status",
			System: true,
			Values: []string{"active", "inactive"},
		})

		ctx := context.Background()
		if err := app.SaveWithContext(ctx, gamesCollection); err != nil {
			return err
		}

		// game items
		gameItemsCollection := core.NewBaseCollection("game_items")

		gameItemsCollection.Fields.Add(&core.TextField{
			Name:   "name",
			System: true,
		}, &core.NumberField{
			Name:   "price",
			System: true,
		}, &core.SelectField{
			Name:   "currency",
			System: true,
			Values: []string{"IDR", "USD"},
		}, &core.SelectField{
			Name:   "status",
			System: true,
			Values: []string{"active", "inactive"},
		})

		if err := app.SaveWithContext(ctx, gameItemsCollection); err != nil {
			return err
		}

		// transactions
		transactionsCollection := core.NewBaseCollection("transactions")

		transactionsCollection.Fields.Add(&core.SelectField{
			Name:   "type",
			System: true,
			Values: []string{"topup", "withdraw", "transfer"},
		}, &core.NumberField{
			Name:   "amount",
			System: true,
		}, &core.JSONField{
			Name:   "game_metadata",
			System: true,
		}, &core.JSONField{
			Name:   "payment_metadata",
			System: true,
		}, &core.SelectField{
			Name:   "payment_status",
			System: true,
			Values: []string{"success", "failed", "unknown"},
		}, &core.SelectField{
			Name:   "status",
			System: true,
			Values: []string{"pending", "completed", "failed"},
		})

		if err := app.SaveWithContext(ctx, transactionsCollection); err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		return nil
	})
}
