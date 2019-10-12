package app

import (
	"github.com/rodsher/selectel/pkg/router"
	"github.com/rodsher/selectel/pkg/config"
	"github.com/rodsher/selectel/pkg/database"
)

func Run() {
	config.Init()
	database.Init()
	r := router.Init()

	r.Run()
}
