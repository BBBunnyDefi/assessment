package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BBBunnyDefi/assessment/rest/expenses"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
	fmt.Println("docker database url:", os.Getenv("DATABASE_URL"))

	h := expenses.NewApp(expenses.InitDB(os.Getenv("DATABASE_URL")))

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", h.HealthHandler)
	// EXP01: POST /expenses - with json body
	e.POST("/expenses", h.CreateExpensesHandler)
	// EXP02: GET /expenses/:id
	e.GET("/expenses/:id", h.GetExpensesHandler)
	// EXP03: PUT /expenses/:id - with json body
	e.PUT("/expenses/:id", h.UpdateExpensesHandler)
	// EXP04: GET /expenses
	e.GET("/expenses", h.GetAllExpensesHandler)

	// Bonus middleware check Autorization
	// EXP04: GET /expenses
	// http://localhost:2565/expenses
	// - Autorization: November 10, 2009wrong_token ?
	// Note: other story
	// - Autorization: November 10, 2009

	// implement graceful shutdown
	// https://echo.labstack.com/cookbook/graceful-shutdown/
	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("graceful shutting down...แล้วมั้งนะ")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
