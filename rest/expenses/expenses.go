package expenses

import (
	"database/sql"
	"net/http"

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
	exps := Expenses{}
	err := c.Bind(&exps)
	if err != nil {
		// log.Fatal("can't insert data", err)
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind : " + err.Error()})
	}

	// fmt.Println("##### dump pq.Array(exps.Tags) #####")
	// fmt.Printf("%T\n", pq.Array(exps.Tags))
	// fmt.Println(pq.Array(exps.Tags))
	// fmt.Println("##### dump pq.Array(exps.Tags) #####")

	row := e.DB.QueryRow("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id", exps.Title, exps.Amount, exps.Note, pq.Array(exps.Tags))
	err = row.Scan(&exps.ID)

	// fmt.Println(exps)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan : " + err.Error()})
	}

	return c.JSON(http.StatusCreated, exps)
}
