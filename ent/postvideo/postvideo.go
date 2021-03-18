// Code generated by entc, DO NOT EDIT.

package postvideo

import (
	"time"
)

const (
	// Label holds the string label denoting the postvideo type in the database.
	Label = "post_video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"

	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"

	// Table holds the table name of the postvideo in the database.
	Table = "post_videos"
	// PostTable is the table the holds the post relation/edge.
	PostTable = "post_videos"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_videos"
)

// Columns holds all SQL columns for postvideo fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldIndex,
	FieldURL,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the PostVideo type.
var ForeignKeys = []string{
	"post_videos",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
