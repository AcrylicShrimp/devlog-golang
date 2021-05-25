package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// UnsavedPost holds the schema definition for the UnsavedPost entity.
type UnsavedPost struct {
	ent.Schema
}

// Fields of the UnsavedPost.
func (UnsavedPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64).Unique(),
		field.String("slug").MaxLen(255).Optional(),
		field.Enum("access_level").Values("public", "unlisted", "private").Optional(),
		field.String("title").MaxLen(255).Optional(),
		field.Text("content").Optional(),
		field.Text("html_content").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UnsavedPost.
func (UnsavedPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Admin.Type).Ref("unsaved_posts").Unique().Required(),
		edge.To("thumbnail", UnsavedPostThumbnail.Type).Unique(),
		edge.To("images", UnsavedPostImage.Type),
		edge.To("videos", UnsavedPostVideo.Type),
		edge.To("attachments", UnsavedPostAttachment.Type),
	}
}
