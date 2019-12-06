package main

import (
	"fmt"
	"github.com/akado2009/amazing-stl/mail"
	"github.com/akado2009/amazing-stl/models"
	"log"
)

func main() {

	if err := models.LoadConfig(&models.AppConfig); err != nil {
		log.Fatal("err loading config: ", err)
	}

	fmt.Printf("%+v\n", models.AppConfig)

	if err := mail.SendEmail("prosvirov.k@gmail.com", "go.mod"); err != nil {
		log.Fatal(err)
	}
}