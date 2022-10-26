package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadConfig(t *testing.T) {
	Convey("load config", t, func() {
		config := &Config{}
		config.Load()
		t.Log("config: ", config)
	})
}
