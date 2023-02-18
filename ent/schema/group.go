package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").Default("unknown"),
		field.String("permissions").Default("unknown"),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("users", User.Type),
		edge.From("domain", Domain.Type).Ref("groups").Unique(),
	}
}
