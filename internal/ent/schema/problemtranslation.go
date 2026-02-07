package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProblemTranslation holds the schema definition for the ProblemTranslation entity.
type ProblemTranslation struct {
	ent.Schema
}

// Fields of the ProblemTranslation.
func (ProblemTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.String("locale").NotEmpty(),
		field.String("title").NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Text("explanation").Optional(),
		field.Int("problem_id"),
	}
}

// Edges of the ProblemTranslation.
func (ProblemTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("problem", Problem.Type).
			Ref("translations").
			Field("problem_id").
			Unique().
			Required(),
		edge.To("choices", Choice.Type),
	}
}

// Indexes of the ProblemTranslation.
func (ProblemTranslation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("problem_id", "locale").Unique(),
	}
}
