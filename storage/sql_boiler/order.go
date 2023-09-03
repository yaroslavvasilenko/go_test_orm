package sql_boiler_storage

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yaroslavvasilenko/go_test_orm/orm/sql_boiler/dbmodels"
)

func (b *ClientSQLBoiler) CreateOrderForUser(order *dbmodels.OrdersBoler) error {
	err := order.InsertG(context.Background(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
