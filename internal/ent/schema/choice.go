package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Choice holds the schema definition for the Choice entity.
type Choice struct {
	ent.Schema
}

// Fields of the Choice.
func (Choice) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").NotEmpty(),
		field.Bool("is_correct").Default(false),
		field.Text("explanation").Optional(),
		field.Int("seq").Comment("Sequence order"),
		field.Int("problem_translation_id"),
	}
}

// Edges of the Choice.
func (Choice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("problem_translation", ProblemTranslation.Type).
			Ref("choices").
			Field("problem_translation_id").
			Unique().
			Required(),
	}
}
