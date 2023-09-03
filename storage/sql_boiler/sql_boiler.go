package sql_boiler_storage

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yaroslavvasilenko/go_test_orm/orm/sql_boiler"
)

type ClientSQLBoiler struct {
	db *sql.DB
}

func NewSQLBoiler() *ClientSQLBoiler {
	conn := sql_boiler.SQLConnect()
	err := sql_boiler.CreateTableForSQLBoiler(conn)
	if err != nil {
		panic(err)
	}

	boil.SetDB(conn)

	return &ClientSQLBoiler{db: conn}
}
