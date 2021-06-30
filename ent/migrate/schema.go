// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "password", Type: field.TypeString, Size: 60},
		{Name: "joined_at", Type: field.TypeTime},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:        "admins",
		Columns:     AdminsColumns,
		PrimaryKey:  []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// AdminSessionsColumns holds the columns for the "admin_sessions" table.
	AdminSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true, Size: 256},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "admin_sessions", Type: field.TypeInt, Nullable: true},
	}
	// AdminSessionsTable holds the schema information for the "admin_sessions" table.
	AdminSessionsTable = &schema.Table{
		Name:       "admin_sessions",
		Columns:    AdminSessionsColumns,
		PrimaryKey: []*schema.Column{AdminSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "admin_sessions_admins_sessions",
				Columns:    []*schema.Column{AdminSessionsColumns[4]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:        "categories",
		Columns:     CategoriesColumns,
		PrimaryKey:  []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "access_level", Type: field.TypeEnum, Enums: []string{"public", "unlisted", "private"}},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "html_content", Type: field.TypeString, Size: 2147483647},
		{Name: "preview_content", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "admin_posts", Type: field.TypeInt, Nullable: true},
		{Name: "category_posts", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_admins_posts",
				Columns:    []*schema.Column{PostsColumns[10]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_categories_posts",
				Columns:    []*schema.Column{PostsColumns[11]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostAttachmentsColumns holds the columns for the "post_attachments" table.
	PostAttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Size: 64},
		{Name: "size", Type: field.TypeUint64},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "mime", Type: field.TypeString, Size: 64},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_attachments", Type: field.TypeInt, Nullable: true},
	}
	// PostAttachmentsTable holds the schema information for the "post_attachments" table.
	PostAttachmentsTable = &schema.Table{
		Name:       "post_attachments",
		Columns:    PostAttachmentsColumns,
		PrimaryKey: []*schema.Column{PostAttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_attachments_posts_attachments",
				Columns:    []*schema.Column{PostAttachmentsColumns[7]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostImagesColumns holds the columns for the "post_images" table.
	PostImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Size: 64},
		{Name: "width", Type: field.TypeUint32},
		{Name: "height", Type: field.TypeUint32},
		{Name: "hash", Type: field.TypeString, Size: 64},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_images", Type: field.TypeInt, Nullable: true},
	}
	// PostImagesTable holds the schema information for the "post_images" table.
	PostImagesTable = &schema.Table{
		Name:       "post_images",
		Columns:    PostImagesColumns,
		PrimaryKey: []*schema.Column{PostImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_images_posts_images",
				Columns:    []*schema.Column{PostImagesColumns[8]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostThumbnailsColumns holds the columns for the "post_thumbnails" table.
	PostThumbnailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "width", Type: field.TypeUint32},
		{Name: "height", Type: field.TypeUint32},
		{Name: "hash", Type: field.TypeString, Size: 64},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_thumbnail", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PostThumbnailsTable holds the schema information for the "post_thumbnails" table.
	PostThumbnailsTable = &schema.Table{
		Name:       "post_thumbnails",
		Columns:    PostThumbnailsColumns,
		PrimaryKey: []*schema.Column{PostThumbnailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_thumbnails_posts_thumbnail",
				Columns:    []*schema.Column{PostThumbnailsColumns[6]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostVideosColumns holds the columns for the "post_videos" table.
	PostVideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Size: 64},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_videos", Type: field.TypeInt, Nullable: true},
	}
	// PostVideosTable holds the schema information for the "post_videos" table.
	PostVideosTable = &schema.Table{
		Name:       "post_videos",
		Columns:    PostVideosColumns,
		PrimaryKey: []*schema.Column{PostVideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_videos_posts_videos",
				Columns:    []*schema.Column{PostVideosColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UnsavedPostsColumns holds the columns for the "unsaved_posts" table.
	UnsavedPostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "slug", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "access_level", Type: field.TypeEnum, Nullable: true, Enums: []string{"public", "unlisted", "private"}},
		{Name: "title", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "admin_unsaved_posts", Type: field.TypeInt, Nullable: true},
		{Name: "category_unsaved_posts", Type: field.TypeInt, Nullable: true},
	}
	// UnsavedPostsTable holds the schema information for the "unsaved_posts" table.
	UnsavedPostsTable = &schema.Table{
		Name:       "unsaved_posts",
		Columns:    UnsavedPostsColumns,
		PrimaryKey: []*schema.Column{UnsavedPostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "unsaved_posts_admins_unsaved_posts",
				Columns:    []*schema.Column{UnsavedPostsColumns[8]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "unsaved_posts_categories_unsaved_posts",
				Columns:    []*schema.Column{UnsavedPostsColumns[9]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UnsavedPostAttachmentsColumns holds the columns for the "unsaved_post_attachments" table.
	UnsavedPostAttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Size: 64},
		{Name: "validity", Type: field.TypeEnum, Enums: []string{"pending", "valid", "invalid"}, Default: "pending"},
		{Name: "size", Type: field.TypeUint64, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "mime", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "unsaved_post_attachments", Type: field.TypeInt, Nullable: true},
	}
	// UnsavedPostAttachmentsTable holds the schema information for the "unsaved_post_attachments" table.
	UnsavedPostAttachmentsTable = &schema.Table{
		Name:       "unsaved_post_attachments",
		Columns:    UnsavedPostAttachmentsColumns,
		PrimaryKey: []*schema.Column{UnsavedPostAttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "unsaved_post_attachments_unsaved_posts_attachments",
				Columns:    []*schema.Column{UnsavedPostAttachmentsColumns[8]},
				RefColumns: []*schema.Column{UnsavedPostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UnsavedPostImagesColumns holds the columns for the "unsaved_post_images" table.
	UnsavedPostImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Size: 64},
		{Name: "validity", Type: field.TypeEnum, Enums: []string{"pending", "valid", "invalid"}, Default: "pending"},
		{Name: "width", Type: field.TypeUint32, Nullable: true},
		{Name: "height", Type: field.TypeUint32, Nullable: true},
		{Name: "hash", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "title", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "unsaved_post_images", Type: field.TypeInt, Nullable: true},
	}
	// UnsavedPostImagesTable holds the schema information for the "unsaved_post_images" table.
	UnsavedPostImagesTable = &schema.Table{
		Name:       "unsaved_post_images",
		Columns:    UnsavedPostImagesColumns,
		PrimaryKey: []*schema.Column{UnsavedPostImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "unsaved_post_images_unsaved_posts_images",
				Columns:    []*schema.Column{UnsavedPostImagesColumns[9]},
				RefColumns: []*schema.Column{UnsavedPostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UnsavedPostThumbnailsColumns holds the columns for the "unsaved_post_thumbnails" table.
	UnsavedPostThumbnailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "validity", Type: field.TypeEnum, Enums: []string{"pending", "valid", "invalid"}, Default: "pending"},
		{Name: "width", Type: field.TypeUint32, Nullable: true},
		{Name: "height", Type: field.TypeUint32, Nullable: true},
		{Name: "hash", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "unsaved_post_thumbnail", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UnsavedPostThumbnailsTable holds the schema information for the "unsaved_post_thumbnails" table.
	UnsavedPostThumbnailsTable = &schema.Table{
		Name:       "unsaved_post_thumbnails",
		Columns:    UnsavedPostThumbnailsColumns,
		PrimaryKey: []*schema.Column{UnsavedPostThumbnailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "unsaved_post_thumbnails_unsaved_posts_thumbnail",
				Columns:    []*schema.Column{UnsavedPostThumbnailsColumns[7]},
				RefColumns: []*schema.Column{UnsavedPostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UnsavedPostVideosColumns holds the columns for the "unsaved_post_videos" table.
	UnsavedPostVideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Size: 64},
		{Name: "validity", Type: field.TypeEnum, Enums: []string{"pending", "valid", "invalid"}, Default: "pending"},
		{Name: "title", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "unsaved_post_videos", Type: field.TypeInt, Nullable: true},
	}
	// UnsavedPostVideosTable holds the schema information for the "unsaved_post_videos" table.
	UnsavedPostVideosTable = &schema.Table{
		Name:       "unsaved_post_videos",
		Columns:    UnsavedPostVideosColumns,
		PrimaryKey: []*schema.Column{UnsavedPostVideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "unsaved_post_videos_unsaved_posts_videos",
				Columns:    []*schema.Column{UnsavedPostVideosColumns[6]},
				RefColumns: []*schema.Column{UnsavedPostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		AdminSessionsTable,
		CategoriesTable,
		PostsTable,
		PostAttachmentsTable,
		PostImagesTable,
		PostThumbnailsTable,
		PostVideosTable,
		UnsavedPostsTable,
		UnsavedPostAttachmentsTable,
		UnsavedPostImagesTable,
		UnsavedPostThumbnailsTable,
		UnsavedPostVideosTable,
	}
)

func init() {
	AdminSessionsTable.ForeignKeys[0].RefTable = AdminsTable
	PostsTable.ForeignKeys[0].RefTable = AdminsTable
	PostsTable.ForeignKeys[1].RefTable = CategoriesTable
	PostAttachmentsTable.ForeignKeys[0].RefTable = PostsTable
	PostImagesTable.ForeignKeys[0].RefTable = PostsTable
	PostThumbnailsTable.ForeignKeys[0].RefTable = PostsTable
	PostVideosTable.ForeignKeys[0].RefTable = PostsTable
	UnsavedPostsTable.ForeignKeys[0].RefTable = AdminsTable
	UnsavedPostsTable.ForeignKeys[1].RefTable = CategoriesTable
	UnsavedPostAttachmentsTable.ForeignKeys[0].RefTable = UnsavedPostsTable
	UnsavedPostImagesTable.ForeignKeys[0].RefTable = UnsavedPostsTable
	UnsavedPostThumbnailsTable.ForeignKeys[0].RefTable = UnsavedPostsTable
	UnsavedPostVideosTable.ForeignKeys[0].RefTable = UnsavedPostsTable
}
