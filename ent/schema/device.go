package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Mixin of the Device.
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname"),
		field.String("os").Default("unknown"),
		field.String("arch").Default("unknown"),
		field.String("version").Default("unknown"),
		field.Strings("net_interfaces").Optional(),
		field.String("machinepass").Optional(),
		field.Strings("certificates").Optional(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("rodents", Rodent.Type),
		edge.From("groups", Group.Type).Ref("devices"),
		edge.From("domain", Domain.Type).Ref("devices").Unique(),
		edge.From("subnets", Subnet.Type).Ref("hosts"),
		edge.From("services", Services.Type).Ref("devices"),
	}
}
