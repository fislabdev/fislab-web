package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("schedule_id"),
		field.UUID("user_id", uuid.UUID{}),
		field.Enum("status").
			Values(
				AttendanceStatusHadir,
				AttendanceStatusTidakHadir,
				AttendanceStatusSakit,
				AttendanceStatusIzin).
			Default(AttendanceStatusTidakHadir),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

func (Attendance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("attendance").
			Unique().
			Required().
			Field("user_id"),
	}
}
