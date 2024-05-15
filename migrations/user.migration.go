package main

import (
	"github.com/aliftech/go-bookingapp/models"
	"github.com/aliftech/go-bookingapp/utils"
)

func init() {
	utils.Setup()
	utils.ConnectDB()
}

func main() {
	utils.DB.AutoMigrate(&models.User{})
}
