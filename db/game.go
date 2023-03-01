package db

import "github.com/ZecreyGaming/zecrey_treasure_hunt/model"

type game db

func (g *game) Create(game *model.Game) error {
	return g.db.Create(game).Error
}

func (g *game) Update(game *model.Game) error {
	return g.db.Updates(game).Error
}
