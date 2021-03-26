package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64).Unique(),
		field.String("slug").MaxLen(255).Unique(),
		field.Enum("access_level").Values("public", "unlisted", "private"),
		field.String("title").MaxLen(255),
		field.Text("content"),
		field.Text("html_content"),
		field.String("preview_content").MaxLen(255),
		field.Time("created_at").Default(time.Now),
		field.Time("modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Admin.Type).Ref("posts").Unique().Required(),
		edge.From("category", Category.Type).Ref("posts").Unique(),
		edge.To("thumbnail", PostThumbnail.Type).Unique(),
		edge.To("images", PostImage.Type),
		edge.To("videos", PostVideo.Type),
		edge.To("attachments", PostAttachment.Type),
	}
}
