package main

import (
	"beiyi/conf"
	"beiyi/server"
)

func main() {
	conf.Init()

	r := server.NewRouter()
	r.Run(":3000")
}
