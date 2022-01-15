package main

import (
	"document-server/configs"
	"document-server/web/route"
)

func init() {
	configs.Initialize()
}

func main() {
	port := configs.Conf.Server.Port
	if len(port) == 0 {
		port = ":80"
	} else {
		port = ":" + port
	}
	r := route.RegisterRouter()
	r.Run(port)

}
