package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Operator holds the schema definition for the Operator entity.
type Operator struct {
	ent.Schema
}

// Fields of the Operator.
func (Operator) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.Bytes("privkey").Optional(),
		field.Bytes("cert").Optional(),
	}
}

func (Operator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}

// Edges of the Operator.
func (Operator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("projects", Project.Type).
			Ref("operators"),
		edge.To("tasks", Task.Type),
	}
}

/*
// Policy defines the privacy policy of the User.
func (Operator) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Deny if not set otherwise.
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			// Allow any viewer to read anything.
			privacy.AlwaysAllowRule(),
		},
	}
}
*/
