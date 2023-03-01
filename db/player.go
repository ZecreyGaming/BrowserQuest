package db

import "github.com/ZecreyGaming/zecrey_treasure_hunt/model"

type player db

func (p *player) Create(game *model.Player) error {
	return p.db.Create(game).Error
}

func (p *player) Update(game *model.Player) error {
	return p.db.Updates(game).Error
}
func (p *player) Get(name string) (*model.Player, error) {
	var player *model.Player
	err := p.db.First(&player, "account_name = ?", name).Error
	return player, err
}
