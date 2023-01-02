package expenses

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
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
