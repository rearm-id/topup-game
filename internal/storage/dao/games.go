package dao

import (
	"github.com/pocketbase/pocketbase/core"
)

// ensures that the Game struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Game)(nil)

func NewGame(record *core.Record) *Game {
	return &Game{
		Record: record,
	}
}

// FindGameById finds a game by its ID
func FindGameById(app core.App, id string) (*Game, error) {
	record, err := app.FindRecordById("games", id)
	if err != nil {
		return nil, err
	}

	return NewGame(record), nil
}

// FindGameBySlug finds a game by its slug
func FindGameBySlug(app core.App, slug string) (*Game, error) {
	record, err := app.FindFirstRecordByData("games", "slug", slug)
	if err != nil {
		return nil, err
	}

	return NewGame(record), nil
}

// FindActiveGames finds all active games
func FindActiveGames(app core.App) ([]*Game, error) {
	records, err := app.FindRecordsByFilter(
		"games",
		"status = 'active'",
		"name",
		100,
		0,
		nil,
	)
	if err != nil {
		return nil, err
	}

	games := make([]*Game, len(records))
	for i, record := range records {
		games[i] = NewGame(record)
	}

	return games, nil
}

type Game struct {
	*core.Record
}

// ProxyRecord implements the core.RecordProxy interface
func (g *Game) ProxyRecord() *core.Record {
	return g.Record
}

// SetProxyRecord implements the core.RecordProxy interface
func (g *Game) SetProxyRecord(r *core.Record) {
	g.Record = r
}

// GetName returns the game name
func (g *Game) GetName() string {
	return g.GetString("name")
}

// SetName sets the game name
func (g *Game) SetName(name string) {
	g.Set("name", name)
}

// GetDescription returns the game description
func (g *Game) GetDescription() string {
	return g.GetString("description")
}

// SetDescription sets the game description
func (g *Game) SetDescription(description string) {
	g.Set("description", description)
}

// GetImage returns the game image URL
func (g *Game) GetImage() string {
	return g.GetString("image")
}

// SetImage sets the game image URL
func (g *Game) SetImage(imageURL string) {
	g.Set("image", imageURL)
}

// GetSlug returns the game slug
func (g *Game) GetSlug() string {
	return g.GetString("slug")
}

// SetSlug sets the game slug
func (g *Game) SetSlug(slug string) {
	g.Set("slug", slug)
}

// GetStatus returns the game status
func (g *Game) GetStatus() string {
	return g.GetString("status")
}

// SetStatus sets the game status
func (g *Game) SetStatus(status string) {
	g.Set("status", status)
}

// IsActive returns true if the game status is "active"
func (g *Game) IsActive() bool {
	return g.GetStatus() == "active"
}

// IsInactive returns true if the game status is "inactive"
func (g *Game) IsInactive() bool {
	return g.GetStatus() == "inactive"
}
