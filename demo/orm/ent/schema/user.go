package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Default(0),
		field.Text("name").Default("unknown").Comment("用户名"),
		field.Ints("tree").Default([]int{}).Comment("树"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
