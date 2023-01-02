//go:build unit
// +build unit

package expenses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	// Arrange
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	h := expenses{}
	c := e.NewContext(req, rec)

	// Act
	err := h.HealthHandler(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Server health OKKK", rec.Body.String())
	}
}

func TestCreateExpensesHandler(t *testing.T) {
	// t.Skip("TODO: EXP01: POST /expenses - with json body")
	t.Log("EXP01: POST /expenses - with json body  COMPLETED!!")
	// Arrange
	e := echo.New()
	// ref: echo api sql seedUser
	strBody := `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`
	body := bytes.NewBufferString(strBody)

	req := httptest.NewRequest(http.MethodPost, "/expenses", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// fmt.Println("##### dump req #####")
	// fmt.Printf("%T\n", req)
	// fmt.Println(req)
	// fmt.Println("##### dump req #####")

	// fmt.Println("##### dump body #####")
	// fmt.Printf("%T\n", body)
	// fmt.Println(body)
	// fmt.Println("##### dump body #####")

	exps := Expenses{}
	err := json.Unmarshal([]byte(strBody), &exps)
	if err != nil {
		t.Fatalf("an error, json.Marshal *bytes.Buffer: '%s' ", err)
	}

	// exps.ID = 1

	// fmt.Println("##### dump exps #####")
	// fmt.Printf("%T\n", exps)
	// fmt.Println(exps)
	// fmt.Println("##### dump exps #####")

	// create new sqlmock
	newMockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	db, mock, err := sqlmock.New()

	// mock.ExpectQuery("INSERT INTO expenses").
	// 	WithArgs("strawberry smoothie", float64(79), "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"})).
	// 	WillReturnRows(newMockRows)

	mock.ExpectQuery("INSERT INTO expenses").
		WithArgs(exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags)).
		WillReturnRows(newMockRows)

	if err != nil {
		t.Fatalf("an error, mock expect query '%s' was not...", err)
	}

	h := expenses{db}
	c := e.NewContext(req, rec)
	// Epected
	expected := "{\"id\":1,\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"

	// Act
	err = h.CreateExpensesHandler(c)
	if err != nil {
		t.Fatalf("an error, act '%s' was not...", err)
	}

	// fmt.Println("##### dump req #####")
	// fmt.Printf("%T\n", req)
	// fmt.Println(req)
	// fmt.Println("##### dump req #####")

	// fmt.Println("##### dump rec #####")
	// fmt.Printf("%T\n", rec)
	// fmt.Println(rec)
	// fmt.Println("##### dump rec #####")

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}
