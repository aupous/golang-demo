package db

import (
	"awesomeProject/internal/configs"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type DB struct {
	DB *pg.DB
}

func (db *DB) Connect(cf *configs.PostgresConfig) {
	db.DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf(":%v", cf.Port),
		User:     cf.User,
		Password: cf.Password,
		Database: cf.Database,
	})
}