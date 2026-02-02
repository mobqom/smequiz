package main

import (
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/pkg/server"
)

const (
	WSPort = ":8080"
)

func main() {
	s := server.NewWSServer()
	s.Run(WSPort)
	log.Fatal(http.ListenAndServe(WSPort, nil))

}
