package postgres

import (
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

const fileName = "postgres.go"

type DatabaseConnect = *pgx.Conn

func InitPostgres(envConfig env.DatabaseConfig) (*pgx.Conn, error) {
	config := &pgx.ConnConfig{
		Host:     envConfig.Host,
		Port:     uint16(envConfig.Port),
		User:     envConfig.User,
		Password: envConfig.Password,
		Database: envConfig.DatabaseName,
	}

	logConfig := log.Fields{
		"host":     config.Host,
		"port":     config.Port,
		"user":     envConfig.User,
		"database": envConfig.DatabaseName,
		"file":     fileName,
		"function": "InitPostgres",
	}

	conn, err := pgx.Connect(*config)
	if err != nil {
		log.WithFields(logConfig).Errorf("Failed to connect to the database %s", err)
	}

	log.WithFields(logConfig).Info("Connection to the database is successful")
	return conn, err
}

func CheckConnection(envConfig env.DatabaseConfig) {
	_, err := InitPostgres(envConfig)
	logConfig := log.Fields{
		"file":     fileName,
		"function": "CheckConnection",
	}
	if err != nil {
		log.WithFields(logConfig).Fatalf("Check connection to postgres failed %s", err)
	}
}
