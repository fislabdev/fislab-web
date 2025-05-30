package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Grade struct {
	ent.Schema
}

func (Grade) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.Int("schedule_id").Immutable().Positive(),
		field.UUID("user_id", uuid.UUID{}).Immutable(),
		field.UUID("graded_by", uuid.UUID{}).Immutable(),
		field.Float32("punctuality").
			Optional(),
		field.Float32("pre_exam").
			Optional(),
		field.Float32("oral_test").
			Optional(),
		field.Float32("skills_and_attitude").
			Optional(),
		field.Float32("abstract").
			Optional(),
		field.Float32("introduction").
			Optional(),
		field.Float32("methodology").
			Optional(),
		field.Float32("discussion").
			Optional(),
		field.Float32("data_processing").
			Optional(),
		field.Float32("conclusion").
			Optional(),
		field.Float32("formatting").
			Optional(),
		field.String("feedback").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Grade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", Schedule.Type).Ref("grades"),
		edge.From("grader", User.Type).Ref("graded_grades"),
		edge.From("graded", User.Type).Ref("user_grades"),
	}
}
