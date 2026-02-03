package main

import (
	"github.com/ibezgin/mobqom-smequiz/config"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

func main() {
	cfg := config.InitConfig()
	server.Run(cfg)
}
