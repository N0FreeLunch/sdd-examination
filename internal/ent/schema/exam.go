package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exam holds the schema definition for the Exam entity.
type Exam struct {
	ent.Schema
}

// Fields of the Exam.
func (Exam) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Text("description").Optional(),
		field.Int("time_limit").Comment("Time limit in minutes"),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the Exam.
func (Exam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sections", Section.Type),
		edge.To("topics", Topic.Type),
		edge.To("units", Unit.Type),
		edge.To("version_rules", VersionRule.Type),
	}
}
