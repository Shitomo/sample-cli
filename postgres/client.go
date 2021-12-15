package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

type DbConfig struct {
	Host string
	Port int
	UserName string
	Password string
	DbName string
}

func NewDbClient(dbConfig DbConfig) (*bun.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.UserName, dbConfig.Password, dbConfig.DbName)
	engine, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(engine, pgdialect.New())
	return db, err
}

func Close(dbClient *bun.DB) {
	err := dbClient.Close()
	if err != nil {
		fmt.Printf("error while database closing, caused by %s", err)
	}
}

func SelectNameById(db *bun.DB) func(ctx context.Context, id string) ([]string, error) {
	return func(ctx context.Context, id string) ([]string, error) {
		var result []string
		db.NewSelect().
			ColumnExpr("name").
			TableExpr("rooms").
			Scan(ctx, &result)
		return result, nil
	}
}
