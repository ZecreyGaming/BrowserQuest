package db

import (
	"github.com/ZecreyGaming/zecrey_treasure_hunt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
	Game   game
	Player player
}

type db struct {
	db *gorm.DB
}

func NewDb(dsn string) *Db {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	err = gdb.AutoMigrate(&model.Game{}, &model.Player{})
	if err != nil {
		panic(err)
	}

	return &Db{DB: gdb, Game: game{db: gdb}, Player: player{db: gdb}}
}
