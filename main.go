package main

import (
	"github.com/yaroslavvasilenko/go_test_orm/poligon/ent"
	sql_boiler_polygon "github.com/yaroslavvasilenko/go_test_orm/poligon/sql_boiler"
	sql_boiler_storage "github.com/yaroslavvasilenko/go_test_orm/storage/sql_boiler"
)

func main() {
	// консольные команды необходимые для генерации кода
	//   for ent go run -mod=mod entgo.io/ent/cmd/ent new User
	//	go gerate ./orm/ent
	ent.TestEnt()

	//  sqlboiler psql
	sqlBoilerTest()

	//	go run github.com/stephenafamo/bob/gen/bobgen-psql
}

func sqlBoilerTest() {
	db := sql_boiler_storage.NewSQLBoiler()
	sql_boiler_polygon.TestSQLBoiler(db)
}
