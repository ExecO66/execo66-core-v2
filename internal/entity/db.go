package entity

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DbClient *DbClientStruct
var once sync.Once

func NewDbClient() *DbClientStruct {
	once.Do(func() {
		DbClient = &DbClientStruct{}
	})
	return DbClient
}

type DbClientStruct struct {
	Db *sql.DB
}

func (ctx *DbClientStruct) Connect(connString string) {
	db, err := sql.Open("pgx", connString)

	if err != nil {
		log.Fatal("unable to connect to database: ", connString)
	}
	ctx.Db = db
}
