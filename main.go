package main

import (
	"fmt"
	"suceanne/config"
	"suceanne/vaccapi"
)

func main() {
	config.InitializeConfig()
	fmt.Println(vaccapi.Encrypt("TEST"))
}