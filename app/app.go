package app

import (
	"github.com/Lequeston/bmstu-coursework-bd/config/db/postgres"
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	log "github.com/sirupsen/logrus"
)

func App(config *env.Config) {
	log.Info("Application is run!")
	postgres.CheckConnection(config.Database)
}
