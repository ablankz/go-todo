package main

import (
	"fmt"
	"go-todoo/app/controllers"
	"go-todoo/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
