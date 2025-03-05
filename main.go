package main

import (
	"flag"

	"proto.proxy/httpserver"
)

func main() {
	host := flag.String("host", "127.0.0.1", "ip addres/host to run on")
	port := flag.String("port", "3000", "port")
	flag.Parse()

	httpserver.StartServer(host, port)
}
