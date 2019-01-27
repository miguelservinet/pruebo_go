package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type grpcLogger struct {
	*log.Logger
}

func main() {
	fmt.Print("Hola Mundo")
}
