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

func (e *expenses) HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Server health OKKK")
}

// EXP01: POST /expenses - with json body
func (e *expenses) CreateExpensesHandler(c echo.Context) error {
	// Query all
	exp := Expenses{}
	err := c.Bind(&exp)
	if err != nil {
		// log.Fatal("can't insert data", err)
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind : " + err.Error()})
	}

	// fmt.Println("##### dump pq.Array(exp.Tags) #####")
	// fmt.Printf("%T\n", pq.Array(exp.Tags))
	// fmt.Println(pq.Array(exp.Tags))
	// fmt.Println("##### dump pq.Array(exp.Tags) #####")

	row := e.DB.QueryRow("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id", exp.Title, exp.Amount, exp.Note, pq.Array(exp.Tags))
	err = row.Scan(&exp.ID)

	// fmt.Println(exps)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan : " + err.Error()})
	}

	return c.JSON(http.StatusCreated, exp)
}

// EXP02: GET /expenses/:id
func (e *expenses) GetExpensesHandler(c echo.Context) error {
	// Query one
	stmt, err := e.DB.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1")
	// fmt.Println(c)
	rowId := c.Param("id")

	if err != nil {
		// log.Fatal("can't prepare query statement:", err)
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query statment with id: " + err.Error()})
	}

	row := stmt.QueryRow(rowId)
	// var u User
	exp := Expenses{}
	// pq.StringArray
	// pq.Array(exp.Tags)
	err = row.Scan(&exp.ID, &exp.Title, &exp.Amount, &exp.Note, pq.Array(&exp.Tags))

	// fmt.Println(exp)

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expenses not found" + err.Error()})
	case nil:
		return c.JSON(http.StatusOK, exp)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan : " + err.Error()})
	}
}

// EXP03: PUT /expenses/:id - with json body
func (e *expenses) UpdateExpensesHandler(c echo.Context) error {
	stmt, err := e.DB.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare update statment: " + err.Error()})
	}

	rowId := c.Param("id")
	rowIdInt, _ := strconv.Atoi(rowId)

	exp := Expenses{
		ID: rowIdInt,
	}
	err = c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't bind param expenses: " + err.Error()})
	}

	// fmt.Println(pq.Array(&exp.Tags))
	// fmt.Println(&exp)

	_, err = stmt.Exec(rowId, &exp.Title, &exp.Amount, &exp.Note, pq.Array(&exp.Tags))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't execute expenses: " + err.Error()})
	}

	return c.JSON(http.StatusOK, exp)
}

// EXP04: GET /expenses
func (e *expenses) GetAllExpensesHandler(c echo.Context) error {
	// Query all
	stmt, err := e.DB.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		// log.Fatal("can't prepare query all users statement:", err)
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all expenses statment: " + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		// log.Fatal("can't query all ", err)
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses: " + err.Error()})
	}

	// loop query all
	expenses := []Expenses{}
	for rows.Next() {
		exps := Expenses{}
		err := rows.Scan(&exps.ID, &exps.Title, &exps.Amount, &exps.Note, pq.Array(&exps.Tags))
		if err != nil {
			// log.Fatal("can't Scan row into variable", err)
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan : " + err.Error()})
		}

		expenses = append(expenses, exps)
	}

	return c.JSON(http.StatusOK, expenses)
}
