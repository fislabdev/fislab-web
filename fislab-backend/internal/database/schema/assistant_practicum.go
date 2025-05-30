package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type AssistantPracticum struct {
	ent.Schema
}

func (AssistantPracticum) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("assistant_id", uuid.UUID{}),
		field.String("practicum_code"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (AssistantPracticum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assistant", User.Type).
			Field("assistant_id").
			Unique().
			Required(),
		edge.To("practicum", Practicum.Type).
			Field("practicum_code").
			Unique().
			Required(),
	}
}
