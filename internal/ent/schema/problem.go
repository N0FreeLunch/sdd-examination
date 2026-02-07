package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("SOURCE", "VARIANT").Default("SOURCE"),
		field.Int("difficulty").Default(1),
		field.Time("created_at"),
		field.Int("unit_id"),
		field.Int("parent_id").Optional().Nillable(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unit", Unit.Type).
			Ref("problems").
			Field("unit_id").
			Unique().
			Required(),
		edge.To("versions", VersionRule.Type),
		edge.To("translations", ProblemTranslation.Type),
		edge.To("children", Problem.Type).
			From("parent").
			Field("parent_id").
			Unique(),
	}
}
