package main

import (
	"github.com/moreira0102/zootoma/internal/server/server"
)

const (
	HOST      = "localhost"
	PORT      = "9000"
	CONN_TYPE = "tcp"
)

func main() {
	server.StartListen(HOST, PORT, CONN_TYPE)
}
