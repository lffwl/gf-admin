package main

import (
	_ "bieshu-oa/boot"
	_ "bieshu-oa/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
