package Work

import (
	"github.com/ThePiachu/Go/Log"
	"golang.org/x/net/context"
)

func Work(c context.Context) {
	Log.Infof(c, "Work!")
}