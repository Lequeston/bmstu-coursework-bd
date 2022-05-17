package main

import (
	"github.com/Lequeston/bmstu-coursework-bd/app"
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	"github.com/Lequeston/bmstu-coursework-bd/config/logger"
)

func init() {
	env.ConfigInit()
	logger.LoggerInit()
}
func main() {
	config := env.New()
	app.App(config)
}
