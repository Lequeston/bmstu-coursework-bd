package postgres

import (
	"testing"

	"github.com/Lequeston/bmstu-coursework-bd/config/env"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPostgres(t *testing.T) {
	env.ConfigInit()
	conf := env.New()

	Convey("Connected to DB ", t, func() {
		conn, err := InitPostgres(conf.Database)

		Convey("should be success", func() {
			So(err, ShouldBeNil)
		})

		Convey("connection object shouldn't be nil", func() {
			So(conn, ShouldNotBeNil)
		})
	})
}
