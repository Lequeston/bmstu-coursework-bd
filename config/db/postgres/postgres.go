package postgres

import (
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

func InitPostgres(envConfig env.DatabaseConfig) (*pgx.Conn, error) {
	var config = &pgx.ConnConfig{
		Host:     envConfig.User,
		Port:     uint16(envConfig.Port),
		User:     envConfig.User,
		Password: envConfig.Password,
		Database: envConfig.DatabaseName,
	}
	conn, err := pgx.Connect(*config)
	if err != nil {
		log.Error("Failed to connect to the database", err)
	}

	return conn, err
}
