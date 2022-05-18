package dbService

import (
	"fmt"

	"github.com/Lequeston/bmstu-coursework-bd/config/db/postgres"
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	log "github.com/sirupsen/logrus"
)

func createReadOnlyRoleSQL(config env.DatabaseConfig) []string {
	return []string{
		fmt.Sprintf("CREATE ROLE %s;", readonlyRoleName),
		fmt.Sprintf("GRANT CONNECT ON DATABASE %s TO %s;", config.DatabaseName, readonlyRoleName),
		fmt.Sprintf("GRANT USAGE ON SCHEMA %s TO %s;", schema, readonlyRoleName),
	}
}

func createGrantPermissionsNewRole(tables []string) []string {
	sqlQueryArray := make([]string, 0, len(tables))

	for _, val := range tables {
		query := fmt.Sprintf("GRANT SELECT ON TABLE \"%s\" TO %s;", val, readonlyRoleName)
		sqlQueryArray = append(sqlQueryArray, query)
	}

	return sqlQueryArray
}

func checkingReadOnlyRoleExist(conn postgres.DatabaseConnect) (bool, error) {
	logConfig := log.Fields{
		"package":  packageName,
		"function": "CheckingReadOnlyRoleExist",
	}

	query := fmt.Sprintf("SELECT * FROM pg_roles WHERE rolname = '%s';", readonlyRoleName)
	commandTag, err := conn.Exec(query)
	if err != nil {
		log.WithFields(logConfig).WithField("querySQL", query).Fatalf("Command return error: %s", err)
		return false, err
	}

	return commandTag.RowsAffected() == 1, nil
}

/**
* Функция которая создает в бд роль только для чтения
 */
func CreateReadOnlyRole(conn postgres.DatabaseConnect, config env.DatabaseConfig, tables []string) error {
	commandCreateRole := createReadOnlyRoleSQL(config)

	logConfig := log.Fields{
		"package":  packageName,
		"function": "CreateReadOnlyRole",
	}

	isExist, err := checkingReadOnlyRoleExist(conn)
	if err != nil {
		return err
	}
	if isExist {
		return nil
	}

	for _, val := range commandCreateRole {
		_, err := conn.Exec(val)

		if err != nil {
			log.WithFields(logConfig).WithField("querySQL", val).Fatalf("Command return error: %s", err)
			return err
		}
	}

	commandsGrantPermissionsNewRole := createGrantPermissionsNewRole(tables)

	for _, val := range commandsGrantPermissionsNewRole {
		_, err := conn.Exec(val)

		if err != nil {
			log.WithFields(logConfig).WithField("querySQL", val).Fatalf("Command return error: %s", err)
			return err
		}
	}

	return nil
}
