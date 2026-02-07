package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Int("seq").Comment("Sequence order"),
		field.Int("exam_id"),
		field.Int("section_id").Optional().Nillable(),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam", Exam.Type).
			Ref("topics").
			Field("exam_id").
			Unique().
			Required(),
		edge.From("section", Section.Type).
			Ref("topics").
			Field("section_id").
			Unique(),
		edge.To("units", Unit.Type),
	}
}
