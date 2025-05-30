package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Practicum holds the schema definition for the Practicum entity.
type Practicum struct {
	ent.Schema
}

func (Practicum) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			StorageKey("code"),
		field.String("title"),
		field.String("description").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Practicum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assistants", AssistantPracticum.Type).
			Ref("practicum"),
		edge.To("schedules", Schedule.Type),
	}
}
