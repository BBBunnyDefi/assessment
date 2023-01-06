package expenses

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type expenses struct {
	DB *sql.DB
}

func NewApp(db *sql.DB) *expenses {
	return &expenses{db}
}

type Expenses struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

func (h *expenses) HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Server health OKKK")
}

// EXP01: POST /expenses - with json body
func (h *expenses) CreateExpensesHandler(c echo.Context) error {
	e := Expenses{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind: " + err.Error()})
	}

	row := h.DB.QueryRow("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id", e.Title, e.Amount, e.Note, pq.Array(e.Tags))
	err = row.Scan(&e.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}

// EXP02: GET /expenses/:id
func (h *expenses) GetExpensesHandler(c echo.Context) error {
	id := c.Param("id")

	stmt, err := h.DB.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query statment with id: " + err.Error()})
	}

	row := stmt.QueryRow(id)
	e := Expenses{}
	err = row.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expenses not found: " + err.Error()})
	case nil:
		return c.JSON(http.StatusOK, e)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan: " + err.Error()})
	}
}

// EXP03: PUT /expenses/:id - with json body
func (h *expenses) UpdateExpensesHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	e := Expenses{
		ID: idInt,
	}

	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't bind param expenses: " + err.Error()})
	}

	statement := `UPDATE expenses SET 
		title=$2, amount=$3, note=$4, tags=$5 
		WHERE id=$1
		RETURNING id, title, amount, note, tags;
	`
	err = h.DB.QueryRow(statement, id, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags)).
		Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't query row returning expenses: " + err.Error()})
	}

	return c.JSON(http.StatusOK, e)
}

// EXP04: GET /expenses
func (h *expenses) GetAllExpensesHandler(c echo.Context) error {
	stmt, err := h.DB.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all expenses statment: " + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses: " + err.Error()})
	}

	exps := []Expenses{}
	for rows.Next() {
		e := Expenses{}
		err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan: " + err.Error()})
		}

		exps = append(exps, e)
	}

	return c.JSON(http.StatusOK, exps)
}
