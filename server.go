package main

import (
	"github.com/aliftech/go-bookingapp/routes"
	"github.com/aliftech/go-bookingapp/utils"
)

func init() {
	utils.Setup()
	utils.ConnectDB()
}

func main() {
	routes.AuthRouter()
}
