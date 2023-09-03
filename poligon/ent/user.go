package ent

import (
	"fmt"
	gen_ent "github.com/yaroslavvasilenko/go_test_orm/orm/ent"
	my_ent "github.com/yaroslavvasilenko/go_test_orm/storage/ent"
)

func TestEnt() {
	clientEnt, err := my_ent.NewEnt()
	if err != nil {
		panic(err)
	}

	testUser(clientEnt)

}

func testUser(clientEnt *my_ent.ClientEnt) {
	us := gen_ent.User{
		Name: "test",
		Age:  10,
	}

	ord := &gen_ent.Order{
		ID:     1,
		Name:   "qweqe",
		Amount: 23,
		Price:  55,
	}

	userBD, err := clientEnt.Create(&us)
	if err != nil {
		panic(err)
	}
	fmt.Println(userBD)

	newOrd, err := clientEnt.CreateOrderForUser(ord, userBD.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(newOrd)

	userBD.Age = 20

	userBD, err = clientEnt.Update(userBD)
	if err != nil {
		panic(err)
	}
	fmt.Println(userBD)

	userBD, err = clientEnt.Read(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(userBD)

	_, err = clientEnt.Create(&gen_ent.User{
		Name: "test2",
		Age:  242,
	})
	if err != nil {
		panic(err)
	}

	usersBD, err := clientEnt.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(usersBD)
}
