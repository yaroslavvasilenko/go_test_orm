package my_ent

import (
	"context"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent"
)

func (e *ClientEnt) Create(user *ent.User) (*ent.User, error) {
	newUser, err := e.Ent.User.Create().SetName(user.Name).
		SetAge(user.Age).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
func (e *ClientEnt) Update(user *ent.User) (*ent.User, error) {
	newUser, err := e.Ent.User.UpdateOneID(user.ID).
		SetName(user.Name).
		SetAge(user.Age).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return newUser, nil

}
func (e *ClientEnt) Read(user *ent.User) (*ent.User, error) {
	userNew, err := e.Ent.User.Get(context.Background(), user.ID)
	//e.Ent.User.Query().Where(func(selector *sql.Selector) {
	//	selector.Where(sql.Like(user.Name, "%Niko%"))
	//
	//})
	if err != nil {
		return nil, err
	}

	return userNew, nil
}

func (e *ClientEnt) ReadAll() ([]*ent.User, error) {

	users, err := e.Ent.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (e *ClientEnt) Delete(userID int) error {
	err := e.Ent.User.DeleteOneID(userID).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
