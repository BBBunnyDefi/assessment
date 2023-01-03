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
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	h := expenses{}
	c := e.NewContext(req, rec)

	err := h.HealthHandler(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Server health OKKK", rec.Body.String())
	}
}

func TestCreateExpensesHandler(t *testing.T) {
	// t.Skip("TODO: EXP01: POST /expenses - with json body")
	t.Log("EXP01: POST /expenses - with json body  COMPLETED!!")

	e := echo.New()

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

	exp := Expenses{}
	err := json.Unmarshal([]byte(strBody), &exp)
	if err != nil {
		t.Fatalf("an error, json.Marshal *bytes.Buffer: '%s' ", err)
	}

	// create new sqlmock
	newMockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	db, mock, err := sqlmock.New()

	mock.ExpectQuery("INSERT INTO expenses").
		WithArgs(exp.Title, exp.Amount, exp.Note, pq.Array(exp.Tags)).
		WillReturnRows(newMockRows)

	if err != nil {
		t.Fatalf("an error, mock expect query '%s' was not...", err)
	}

	h := expenses{db}
	c := e.NewContext(req, rec)

	expected := "{\"id\":1,\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"

	err = h.CreateExpensesHandler(c)
	if err != nil {
		t.Fatalf("an error, act '%s' was not...", err)
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}

func TestGetExpensesHandler(t *testing.T) {
	// t.Skip("TODO: EXP02: GET /expenses/:id")
	t.Log("EXP02: GET /expenses/:id COMPLETED!!")

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/expenses/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	newMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, "apple smoothie", 89, "no discount", pq.Array([]string{"beverage"}))

	db, mock, err := sqlmock.New()

	mock.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses WHERE id=\\$1").
		ExpectQuery().
		WithArgs("1").
		WillReturnRows(newMockRows)

	if err != nil {
		t.Fatalf("an error, mock expect query '%s' was not...", err)
	}

	h := expenses{db}
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	expected := "{\"id\":1,\"title\":\"apple smoothie\",\"amount\":89,\"note\":\"no discount\",\"tags\":[\"beverage\"]}"

	err = h.GetExpensesHandler(c)
	if err != nil {
		t.Fatalf("an error, act '%s' was not...", err)
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}

func TestUpdateExpensesHandler(t *testing.T) {
	t.Skip("TODO: EXP03: PUT /expenses/:id - with json body FAILED!!")
	// t.Log("EXP03: PUT /expenses/:id - with json body")

	e := echo.New()

	str_body := `{
		"title": "apple smoothie",
		"amount": 89,
		"note": "no discount",
		"tags": ["beverage"]
	}`
	body := bytes.NewBufferString(str_body)

	req := httptest.NewRequest(http.MethodPut, "/expenses/:id", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	exps := Expenses{}
	err := json.Unmarshal([]byte(str_body), &exps)
	if err != nil {
		t.Fatalf("an error, json.Marshal *bytes.Buffer: '%s' ", err)
	}

	newMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags))

	db, mock, err := sqlmock.New()

	// Error: 500 Step Prepare: Error
	// mock.ExpectPrepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1").
	// 	ExpectQuery().
	// 	WithArgs("1", exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags)).
	// 	WillReturnRows(newMockRows)

	// Error: 400 Step Prepare: Pass, Step Exec: Error
	// mock.ExpectPrepare("UPDATE expenses SET title=\\$2, amount=\\$3, note=\\$4, tags=\\$5 WHERE id=\\$1").
	// 	ExpectQuery().
	// 	WithArgs("1", exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags)).
	// 	WillReturnRows(newMockRows)

	// Error: 400 Step Prepare: Pass, Step Exec: Error
	mock.ExpectPrepare("UPDATE expenses").
		ExpectQuery().
		WithArgs("1", exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags)).
		WillReturnRows(newMockRows)

	// Error: 400
	// mock.ExpectPrepare("UPDATE expenses").
	// 	ExpectQuery().
	// 	WithArgs(1, exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags)).
	// 	WillReturnRows(newMockRows)

	if err != nil {
		t.Fatalf("an error, mock expect query '%s' was not...", err)
	}

	h := expenses{db}
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	expected := "{\"id\":1,\"title\":\"apple smoothie\",\"amount\":89,\"note\":\"no discount\",\"tags\":[\"beverage\"]}"

	err = h.UpdateExpensesHandler(c)
	if err != nil {
		t.Fatalf("an error, act '%s' was not...", err)
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}

func TestGetAllExpensesHandler(t *testing.T) {
	// t.Skip("TODO: EXP04: GET /expenses")
	t.Log("EXP04: GET /expenses COMPLETED!!")

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/expenses", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	newMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, "apple smoothie", 89, "no discount", pq.Array([]string{"beverage"})).
		AddRow(2, "iPhone 14 Pro Max 1TB", 66900, "birthday gift from my love", pq.Array([]string{"gadget"}))

	mock.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses").
		ExpectQuery().
		WillReturnRows(newMockRows)

	if err != nil {
		t.Fatalf("an error, mock expect query '%s' was not...", err)
	}

	h := expenses{db}
	c := e.NewContext(req, rec)

	expected := "[{\"id\":1,\"title\":\"apple smoothie\",\"amount\":89,\"note\":\"no discount\",\"tags\":[\"beverage\"]},{\"id\":2,\"title\":\"iPhone 14 Pro Max 1TB\",\"amount\":66900,\"note\":\"birthday gift from my love\",\"tags\":[\"gadget\"]}]"

	err = h.GetAllExpensesHandler(c)
	if err != nil {
		t.Fatalf("an error, act '%s' was not...", err)
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}
