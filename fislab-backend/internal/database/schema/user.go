package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id",
			uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("nrp").
			Unique(),
		field.String("name").
			Optional(),
		field.Enum("role").
			Values(
				RoleSuperAdmin,
				RoleAdmin,
				RoleAsisten,
				RolePraktikan,
			),

		field.String("email").
			Unique().
			Optional(),
		field.String("phone").
			Unique().
			Optional(),
		field.String("password").
			Sensitive().
			NotEmpty(),
		field.String("about").
			Optional(),
		field.String("profile_picture").
			Optional(),
		field.Bool("email_verified").
			Default(false),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assistant_practicums", AssistantPracticum.Type),
		edge.To("attendance", Attendance.Type),
		edge.To("assistant_schedules", Schedule.Type),
		edge.To("graded_grades", Grade.Type),
		edge.To("user_grades", Grade.Type),
		edge.To("announcements", Announcement.Type),

		edge.From("member_group", Group.Type).Ref("members"),
	}
}
