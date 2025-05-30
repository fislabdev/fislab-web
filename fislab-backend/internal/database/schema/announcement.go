package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

func (Announcement) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.String("description"),
		field.UUID("author_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

func (Announcement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("announcements").
			Field("author_id").
			Unique().
			Required(),
	}
}
