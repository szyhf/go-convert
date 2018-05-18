package convert

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)
import "fmt"

func TestMustInt(t *testing.T) {
	fmt.Print("a ...interface{}")
	Convey("测试字符串“-1.2”", t, func() {
		So(MustInt("-1.2"), ShouldEqual, -1)
	})
}
