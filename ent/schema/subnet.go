package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subnet holds the schema definition for the Subnet entity.
type Subnet struct {
	ent.Schema
}

func (Subnet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the Project schema.
		BaseMixin{},
	}
}

// Fields of the Subnet.
func (Subnet) Fields() []ent.Field {
	return []ent.Field{
		field.String("cidr"),
		field.Bytes("mask").Optional(),
		field.Strings("outbound_tcpports").Optional(),
		field.Strings("outbound_udpports").Optional(),
		field.Bool("proxy").Optional(),
	}
}

// Edges of the Subnet.
func (Subnet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hosts", Device.Type),
	}
}
