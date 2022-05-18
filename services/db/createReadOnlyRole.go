package dbService

import (
	"fmt"

	"github.com/Lequeston/bmstu-coursework-bd/config/db/postgres"
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	log "github.com/sirupsen/logrus"
)

const fileName = "createReadOnlyRole"

func createReadOnlyRoleSQL(config env.DatabaseConfig) []string {
	const roleName = "bmstu_readonly"
	return []string{
		fmt.Sprintf("CREATE ROLE %s;", roleName),
		fmt.Sprintf("GRANT CONNECT ON DATABASE %s TO %s;", config.DatabaseName, roleName),
		fmt.Sprintf("GRANT USAGE ON SCHEMA public TO %s;", roleName),
	}
}

/**
* Функция которая создает в бд роль только для чтения
 */
func CreateReadOnlyRole(conn postgres.DatabaseConnect, config env.DatabaseConfig) error {
	commandCreateRole := createReadOnlyRoleSQL(config)

	logConfig := log.Fields{
		"file":     fileName,
		"function": "CheckConnection",
	}

	for _, val := range commandCreateRole {
		_, err := conn.Exec(val)

		if err != nil {
			log.WithFields(logConfig).WithField("querySQL", val).Fatalf("Command return error: %s", err)
			return err
		}
	}

	return nil
}
