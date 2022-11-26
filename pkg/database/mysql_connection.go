package database

import (
	"github.com/BounkBU/kurester/config"
	"github.com/BounkBU/kurester/pkg/util"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func NewMySQLDatabaseConnection(config *config.Config) (*sqlx.DB, error) {
	var stuff string
	if config.App.Env == "production" {
		stuff = "dns"
	} else {
		stuff = "mysql"
	}
	mysqlUrl := util.NewConnectionUrlBuilder(stuff, config.Database)
	db, err := sqlx.Connect("mysql", mysqlUrl)
	if err != nil {
		log.Errorf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("error, can't ping the database, %s", err.Error())
		return nil, err
	}

	log.Info("Connected to mysql database successfully")
	return db, nil
}
