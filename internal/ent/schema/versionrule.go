package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VersionRule holds the schema definition for the VersionRule entity.
type VersionRule struct {
	ent.Schema
}

// Fields of the VersionRule.
func (VersionRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("year").Optional().Nillable(),
		field.Int("round").Optional().Nillable(),
		field.String("category").Optional().Nillable(),
		field.Enum("operator").Values("Greater", "GreaterEqual", "Less", "LessEqual", "Equal", "NotEqual"),
		field.Enum("status").Values("ACTIVE", "DEPRECATED").Default("ACTIVE"),
		field.Int("exam_id"),
		field.Int("problem_id").Optional().Nillable(),
	}
}

// Edges of the VersionRule.
func (VersionRule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam", Exam.Type).
			Ref("version_rules").
			Field("exam_id").
			Unique().
			Required(),
		edge.From("problem", Problem.Type).
			Ref("versions").
			Field("problem_id").
			Unique(),
	}
}
