package dbService

import (
	"fmt"

	"github.com/Lequeston/bmstu-coursework-bd/config/db/postgres"
	log "github.com/sirupsen/logrus"
)

func CreateSchema(conn postgres.DatabaseConnect) error {
	logConfig := log.Fields{
		"package":  packageName,
		"function": "CreateSchema",
	}

	query := fmt.Sprintf("CREATE SCHEMA %s;", schema)
	_, err := conn.Exec(query)

	if err != nil {
		log.WithFields(logConfig).WithField("querySQL", query).Fatalf("Command return error: %s", err)
		return err
	}

	return nil
}

func RemoveSchema(conn postgres.DatabaseConnect) error {
	logConfig := log.Fields{
		"package":  packageName,
		"function": "RemoveSchema",
	}

	query := fmt.Sprintf("DROP SCHEMA %s CASCADE;", schema)
	_, err := conn.Exec(query)

	if err != nil {
		log.WithFields(logConfig).WithField("querySQL", query).Fatalf("Command return error: %s", err)
		return err
	}

	return nil
}

func CheckingSchemaExist(conn postgres.DatabaseConnect) (bool, error) {
	logConfig := log.Fields{
		"package":  packageName,
		"function": "CheckingSchemaExist",
	}

	query := fmt.Sprintf("SELECT schema_name FROM information_schema.schemata WHERE schema_name = '%s';", schema)
	commandTag, err := conn.Exec(query)

	if err != nil {
		log.WithFields(logConfig).WithField("querySQL", query).Fatalf("Command return error: %s", err)
		return false, err
	}

	return commandTag.RowsAffected() == 1, nil
}

func RecreateSchema(conn postgres.DatabaseConnect) error {
	isExist, err := CheckingSchemaExist(conn)
	if err != nil {
		return err
	}

	if isExist {
		err = RemoveSchema(conn)
		if err != nil {
			return err
		}
	}

	err = CreateSchema(conn)
	if err != nil {
		return err
	}

	return nil
}
