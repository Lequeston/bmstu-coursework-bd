package main

import (
	"github.com/Lequeston/bmstu-coursework-bd/app"
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
)

func init() {
	env.ConfigInit()
}
func main() {
	config := env.New()
	app.App(config)
}
