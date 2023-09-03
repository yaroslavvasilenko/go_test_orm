package sql_boiler_polygon

import (
	"fmt"
	"github.com/volatiletech/null/v8"
	"github.com/yaroslavvasilenko/go_test_orm/orm/sql_boiler/dbmodels"
	sql_boiler_storage "github.com/yaroslavvasilenko/go_test_orm/storage/sql_boiler"
)

func TestSQLBoiler(clientSQLBoiler *sql_boiler_storage.ClientSQLBoiler) {
	user := &dbmodels.UsersBoiler{
		Name: "test",
		Age:  10,
	}

	ord := &dbmodels.OrdersBoler{
		ID:           1,
		Name:         null.StringFrom("qweqe"),
		Amount:       null.IntFrom(23),
		Price:        null.IntFrom(55),
		UserOrdersID: null.Int64From(1),
	}

	err := clientSQLBoiler.CreateUser(user)
	if err != nil {
		panic(err)
	}

	err = clientSQLBoiler.CreateOrderForUser(ord)
	if err != nil {
		panic(err)
	}

	user.Age = 20

	err = clientSQLBoiler.UpdateUser(user)
	if err != nil {
		panic(err)
	}

	userBD, err := clientSQLBoiler.ReadUser(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userBD)

	err = clientSQLBoiler.CreateUser(&dbmodels.UsersBoiler{
		Name: "test2",
		Age:  242,
	})
	if err != nil {
		panic(err)
	}

	usersBD, err := clientSQLBoiler.ReadAllUser()
	if err != nil {
		panic(err)
	}
	fmt.Println(usersBD)
}
