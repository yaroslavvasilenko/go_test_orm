package sql_boiler_storage

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yaroslavvasilenko/go_test_orm/orm/sql_boiler/dbmodels"
)

func (b *ClientSQLBoiler) CreateUser(user *dbmodels.UsersBoiler) error {
	err := user.InsertG(context.Background(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (b *ClientSQLBoiler) ReadAllUser() (dbmodels.UsersBoilerSlice, error) {
	users, err := dbmodels.UsersBoilers().AllG(context.Background())
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (b *ClientSQLBoiler) ReadUser(userID int64) (*dbmodels.UsersBoiler, error) {
	//	user, err := dbmodels.UsersBoilers(dbmodels.UsersBoilerWhere.ID.EQ(int64(userID))).OneG(context.Background())
	user, err := dbmodels.FindUsersBoilerG(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *ClientSQLBoiler) UpdateUser(user *dbmodels.UsersBoiler) error {
	_, err := user.UpdateG(context.Background(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (b *ClientSQLBoiler) DeleteUser(user *dbmodels.UsersBoiler) error {
	_, err := user.DeleteG(context.Background())
	if err != nil {
		return err
	}
	return nil
}
