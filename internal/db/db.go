package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/bismastr/anti-judol-regex/internal/config"
	_ "github.com/microsoft/go-mssqldb"
)

var server = "vedex-sql-server.database.windows.net"
var port = 1433
var user = "vedex"
var password = config.Envs.SQLPassword
var database = "anti-judol"

type Db struct {
	Conn *sql.DB
}

func NewDb() (*Db, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("DB Connected!")

	return &Db{
		Conn: db,
	}, nil
}
