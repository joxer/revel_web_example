package controllers

import (
	"example_app/models"
	"fmt"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(models.InitDB)

	fmt.Println("HERE INIT DB")
}
