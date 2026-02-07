package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Unit holds the schema definition for the Unit entity.
type Unit struct {
	ent.Schema
}

// Fields of the Unit.
func (Unit) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Int("seq").Comment("Sequence order"),
		field.Int("exam_id"),
		field.Int("section_id").Optional().Nillable(),
		field.Int("topic_id").Optional().Nillable(),
	}
}

// Edges of the Unit.
func (Unit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam", Exam.Type).
			Ref("units").
			Field("exam_id").
			Unique().
			Required(),
		edge.From("section", Section.Type).
			Ref("units").
			Field("section_id").
			Unique(),
		edge.From("topic", Topic.Type).
			Ref("units").
			Field("topic_id").
			Unique(),
		edge.To("problems", Problem.Type),
	}
}
