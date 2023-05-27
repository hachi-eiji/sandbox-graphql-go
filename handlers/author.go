package handlers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"http-go-sandbox/sqls"
	"http-go-sandbox/types"
	"net/http"
	"strconv"
)

func ListAuthors(c echo.Context) error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dataSourceName())

	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	q := sqls.New(conn)
	list, err := q.ListAuthors(ctx)

	if err != nil {
		return err
	}

	var users []types.User

	for _, author := range list {
		user := types.User{
			Id:   author.ID,
			Name: author.Name,
		}
		users = append(users, user)
	}
	err = c.JSON(http.StatusOK, users)

	if err != nil {
		return err
	}
	return nil
}

func FindAuthor(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return err
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dataSourceName())

	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	queries := sqls.New(conn)
	author, err := queries.GetAuthor(ctx, id)

	if err != nil {
		return err
	}

	user := types.User{
		Id:   author.ID,
		Name: author.Name,
	}

	err = c.JSON(http.StatusOK, user)

	if err != nil {
		return err
	}
	return nil
}

type DataSourceName struct {
	Host     string
	Port     int32
	User     string
	Password string
	Dbname   string
}

func dataSourceName() string {
	config := DataSourceName{
		User:     "postgres",
		Host:     "127.0.0.1",
		Password: "example",
		Dbname:   "postgres",
		Port:     5431,
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
}
