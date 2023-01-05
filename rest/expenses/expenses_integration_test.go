//go:build integration
// +build integration

package expenses

import (
	"context"
	"database/sql"
	"encoding/json"
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

// use for run docker compose
// const serverPort = 80

// use for run in terminal connect with database
// run database & server for testing outside sandbox
const serverPort = 2565

const databaseURL = "postgres://root:root@db/assessment?sslmode=disable"

func TestITHealthHandler(t *testing.T) {
	// t.Skip("TODO: implement integration Health GET /")
	// t.Log("TODO: implement integration Health GET /")

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

	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "Server health OKKK", string(byteBody))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestITCreateExpensesHandler(t *testing.T) {
	// t.Log("TODO: implement integration EXP01: POST /expenses - with json body")
	t.Skip("TODO: implement integration EXP01: POST /expenses - with json body")
}

func TestITGetExpensesHandler(t *testing.T) {
	// t.Log("TODO: implement integration EXP02 GET /expenses/:id")
	// t.Skip("TODO: implement integration EXP02 GET /expenses/:id")

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

	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/1", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}

	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

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

func TestITUpdateExpensesHandler(t *testing.T) {
	// t.Log("TODO: implement integration EXP03: PUT /expenses/:id - with json body")
	t.Skip("TODO: implement integration EXP03: PUT /expenses/:id - with json body")
}

func TestITGetAllExpensesHandler(t *testing.T) {
	// t.Log("TODO: implement integration EXP04 GET /expenses")
	// t.Skip("TODO: implement integration EXP04 GET /expenses")
	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			log.Fatal(err)
		}

		h := NewApp(db)

		e.GET("/expenses", h.GetAllExpensesHandler)
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

	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	exps := []Expenses{}
	err = json.Unmarshal(byteBody, &exps)
	if err != nil {
		log.Println(err)
	}

	expected := 1

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		// assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
		assert.GreaterOrEqual(t, len(exps), expected)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
