package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/syumai/til/sqlc/sqlc-tutorial/model"
)

type server struct {
	queries *model.Queries
}

func (s *server) ListAuthors(c echo.Context) error {
	authors, err := s.queries.ListAuthors(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authors)
}

func (s *server) GetAuthor(c echo.Context) error {
	authorID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	author, err := s.queries.GetAuthor(c.Request().Context(), authorID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, author)
}

func (s *server) CreateAuthor(c echo.Context) error {
	params := new(model.CreateAuthorParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	author, err := s.queries.CreateAuthor(c.Request().Context(), *params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, author)
}

func newHandler(s *server) http.Handler {
	e := echo.New()
	e.GET("/authors", s.ListAuthors)
	e.GET("/authors/:id", s.GetAuthor)
	e.POST("/authors", s.CreateAuthor)
	return e
}

func run() error {
	db, err := sql.Open("sqlite3", "file:db.sqlite")
	if err != nil {
		return err
	}

	queries := model.New(db)
	s := &server{queries: queries}
	handler := newHandler(s)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9876"
	}
	fmt.Printf("Listening on port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
}
