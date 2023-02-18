package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the Project schema.
		BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("givenname").Default("unknown"),
		field.String("email").Default("unknown"),
		field.Bool("owned").Default(false),
		field.String("age").Default("unknown").Optional(),
		field.String("homedir").Default("unknown"),
		field.Bool("enabled").Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("devices", Device.Type).Ref("users"),
		edge.To("rodents", Rodent.Type),
		edge.From("groups", Group.Type).Ref("users"),
		edge.From("domain", Domain.Type).Ref("users").Unique(),
	}
}
