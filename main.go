package main

import (
	"flag"
	"github.com/shiyicode/gin-framework/common"
	"github.com/shiyicode/gin-framework/router"
	"github.com/TV4/graceful"
	"net/http"
	"github.com/shiyicode/gin-framework/common/g"
	"fmt"
)

func main() {
	cfgFile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	common.Init(*cfgFile)
	defer common.Close()

	router := router.GetRouter()

	graceful.LogListenAndServe(&http.Server{
		Addr:    fmt.Sprintf(":%d", g.Conf().Run.HTTPPort),
		Handler: router,
	})
}
