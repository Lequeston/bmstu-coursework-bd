package dbService

import (
	"fmt"
	"testing"

	"github.com/Lequeston/bmstu-coursework-bd/config/db/postgres"
	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	"github.com/Lequeston/bmstu-coursework-bd/config/logger"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateReadOnlyRole(t *testing.T) {
	env.ConfigInit()
	conf := env.New()
	logger.LoggerInit(conf.Application)

	conn, err := postgres.InitPostgres(conf.Database)

	conn.Exec(fmt.Sprintf("DROP ROLE %s;", readonlyRoleName))
	Convey("Connected to DB ", t, func() {
		Convey("should be success", func() {
			So(err, ShouldBeNil)
		})
	})

	err = RecreateSchema(conn)

	Convey("RecreateSchema ", t, func() {
		Convey("should be success", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("CreateReadOnlyRole when empty array tables", t, func() {
		err = CreateReadOnlyRole(conn, conf.Database, []string{})
		Convey("should be success", func() {
			So(err, ShouldBeNil)
		})
	})
}
