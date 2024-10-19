package Work

import (
	"github.com/ThePiachu/Go/Log"
	"golang.org/x/net/context"
	gobcy "github.com/ThePiachu/gobcy/v2"
)

var BCY gobcy.API

func Work(c context.Context) {
	Log.Infof(c, "Work!")
}