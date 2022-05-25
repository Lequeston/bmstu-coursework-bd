package dbRunnerService

import (
	"github.com/Lequeston/bmstu-coursework-bd/config/db/postgres"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

func RunSQL(conn postgres.DatabaseConnect, sql string) ([]pgx.FieldDescription, [][]interface{}, error) {
	table := make([][]interface{}, 0)

	logConfig := log.Fields{
		"file":     "dbRunnerService",
		"function": "RunSQL",
	}

	rows, err := conn.Query(sql)
	if err != nil {
		log.WithFields(logConfig).Errorf("Failed to run SQL query: %s", err)
		return nil, table, err
	}
	defer rows.Close()

	fieldDescription := rows.FieldDescriptions()

	for rows.Next() {
		row, err := rows.Values()
		if err != nil {
			log.WithFields(logConfig).Errorf("Failed to read row from SQL query: %s", err)
			return fieldDescription, table, err
		}
		table = append(table, row)
	}

	err = rows.Err()
	if err != nil {
		log.WithFields(logConfig).Errorf("Error occured while reading rows from SQL query: %s", err)
		return fieldDescription, table, err
	}

	return fieldDescription, table, nil
}
