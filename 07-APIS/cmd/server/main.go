package main

import (
	"fmt"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/configs"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.DBDriver)
}
