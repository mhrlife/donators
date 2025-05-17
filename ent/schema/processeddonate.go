package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProcessedDonate holds the schema definition for the ProcessedDonate entity.
type ProcessedDonate struct {
	ent.Schema
}

// Fields of the ProcessedDonate.
func (ProcessedDonate) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("display_name"),
		field.Int64("created_at"),
		field.Int64("amount"),
		field.String("currency"),
	}
}

// Edges of the ProcessedDonate.
func (ProcessedDonate) Edges() []ent.Edge {
	return nil
}
