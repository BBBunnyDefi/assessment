package main

import (
	"fmt"
	"os"

	"github.com/BBBunnyDefi/assessment/rest/expenses"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	h := expenses.NewApp(nil)

	// use echo lib for server
	e := echo.New()

	e.GET("/", h.HealthHandler)

	serverPort := ":2565"
	e.Logger.Fatal(e.Start(serverPort))
}
