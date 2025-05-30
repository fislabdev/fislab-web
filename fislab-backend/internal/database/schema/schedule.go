package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.UUID("assistant_id", uuid.UUID{}), // Foreign key ke User
		field.String("practicum_code"),          // Foreign key ke Practicum
		field.UUID("group_id", uuid.UUID{}),     // Foreign key ke Group
		field.Uint8("week").
			Optional(),
		field.Time("start_time").
			Optional(),
		field.Time("end_time").
			Optional(),
		field.Enum("status").
			Values(
				StatusUnscheduled,
				StatusScheduled,
				StatusFinished,
				StatusCompleted,
				StatusCancelled,
			).
			Default(StatusUnscheduled),
	}
}

func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("practicum", Practicum.Type).
			Ref("schedules").
			Field("practicum_code"). // Field foreign key
			Unique().
			Required(),

		edge.From("group", Group.Type).
			Ref("schedules").
			Field("group_id"). // Field foreign key
			Unique().
			Required(),

		edge.From("assistant", User.Type).
			Ref("assistant_schedules").
			Field("assistant_id"). // Field foreign key
			Unique().
			Required(),

		edge.To("grades", Grade.Type), // Relasi ke Grade
	}
}
