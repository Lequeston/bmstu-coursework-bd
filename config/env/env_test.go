package env

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnv(t *testing.T) {
	ConfigInit()

	Convey("MODE_ENV should be test", t, func() {
		So(getEnv("MODE_ENV", ""), ShouldEqual, TEST_MODE)
	})
}
