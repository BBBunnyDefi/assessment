//go:build integration
// +build integration

package expenses

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// Docker-Compose: run with integration tests
// PORT & DATABASE_URL
const serverPort = 80

const databaseURL = "postgres://root:root@db/assessment?sslmode=disable"

func TestITHealthHandler(t *testing.T) {
	// t.Skip("TODO: implement integration Health GET /")
	t.Log("TODO: implement integration Health GET /")

	// Setup server
	eh := echo.New()
	go func(e *echo.Echo) {
		h := NewApp(nil)

		e.GET("/", h.HealthHandler)
		e.Start(fmt.Sprintf(":%d", serverPort))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "Server health OKKK", string(byteBody))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestITGetExpensesHandler(t *testing.T) {
	t.Log("TODO: implement integration EXP02 GET /expenses/:id")
	// t.Skip("TODO: implement integration EXP02 GET /expenses/:id")
	// Setup server
	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			log.Fatal(err)
		}

		h := NewApp(db)

		e.GET("/expenses/:id", h.GetExpensesHandler)
		e.Start(fmt.Sprintf(":%d", serverPort))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/1", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// q := req.URL.Query()
	// q.Add("id", "1")

	// fmt.Println("##### dump q #####")
	// fmt.Printf("%T\n", q)
	// fmt.Println(q)
	// fmt.Println("##### dump q #####")

	client := http.Client{}

	// fmt.Println("##### dump client #####")
	// fmt.Printf("%T\n", client)
	// fmt.Println(client)
	// fmt.Println("##### dump client #####")

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// fmt.Println("##### dump byteBody #####")
	// fmt.Printf("%T\n", strings.TrimSpace(string(byteBody)))
	// fmt.Println(strings.TrimSpace(string(byteBody)))
	// fmt.Println("##### dump byteBody #####")

	// Assertions
	expected := "{\"id\":1,\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
