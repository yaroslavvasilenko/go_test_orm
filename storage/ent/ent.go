package my_ent

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent"
)

type ClientEnt struct {
	Ent *ent.Client
}

func NewEnt() (*ClientEnt, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=root sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		return nil, err
	}

	return &ClientEnt{Ent: client}, err
}
