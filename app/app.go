package app

import (
	"fmt"

	"github.com/Lequeston/bmstu-coursework-bd/config/env"
)

func App(config *env.Config) {
	fmt.Println(config.Database.Login)
	fmt.Println(config.Database.Password)
	fmt.Println(config.Database.DatabaseName)
}
