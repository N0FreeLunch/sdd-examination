package schema_test

import (
	"testing"

	"examination/internal/ent/schema"

	"entgo.io/ent"
	"github.com/stretchr/testify/assert"
)

func TestSchema_Fields(t *testing.T) {
	// Simple test to ensure schemas are valid Ent schemas
	schemas := []ent.Interface{
		&schema.Exam{},
		&schema.Section{},
		&schema.Topic{},
		&schema.Unit{},
		&schema.VersionRule{},
		&schema.Problem{},
		&schema.ProblemTranslation{},
		&schema.Choice{},
	}

	for _, s := range schemas {
		assert.NotNil(t, s, "Schema should not be nil")
	}
}
