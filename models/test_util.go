package models

import (
	"github.com/shiyicode/gin-framework/common"
	"sync"
)

var once sync.Once

func InitAllInTest() {
	once.Do(func() {
		common.Init("../cfg/cfg.toml.debug")
	})
}
