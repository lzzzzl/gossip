package neo4j

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorld(t *testing.T) {
	Convey("hello world", t, func() {
		msg, err := HelloWorld("neo4j://10.254.61.243:7687", "neo4j", "starlink2025")
		So(err, ShouldBeNil)
		t.Log(msg)
	})
}
