package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("amount"),
		field.Int("price"),
		field.Time("created_at").Default(time.Now),
		field.Int("user_orders_id").Unique(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
