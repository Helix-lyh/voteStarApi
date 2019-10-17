package main

import (
	_ "voteStarApi/boot"
	_ "voteStarApi/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
