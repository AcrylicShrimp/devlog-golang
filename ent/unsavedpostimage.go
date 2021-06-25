// Code generated by entc, DO NOT EDIT.

package ent

import (
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostimage"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UnsavedPostImage is the model entity for the UnsavedPostImage schema.
type UnsavedPostImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Validity holds the value of the "validity" field.
	Validity unsavedpostimage.Validity `json:"validity,omitempty"`
	// Width holds the value of the "width" field.
	Width *uint32 `json:"width,omitempty"`
	// Height holds the value of the "height" field.
	Height *uint32 `json:"height,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash *string `json:"hash,omitempty"`
	// Title holds the value of the "title" field.
	Title *string `json:"title,omitempty"`
	// URL holds the value of the "url" field.
	URL *string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UnsavedPostImageQuery when eager-loading is set.
	Edges               UnsavedPostImageEdges `json:"edges"`
	unsaved_post_images *int
}

// UnsavedPostImageEdges holds the relations/edges for other nodes in the graph.
type UnsavedPostImageEdges struct {
	// UnsavedPost holds the value of the unsaved_post edge.
	UnsavedPost *UnsavedPost `json:"unsaved_post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UnsavedPostOrErr returns the UnsavedPost value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UnsavedPostImageEdges) UnsavedPostOrErr() (*UnsavedPost, error) {
	if e.loadedTypes[0] {
		if e.UnsavedPost == nil {
			// The edge unsaved_post was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: unsavedpost.Label}
		}
		return e.UnsavedPost, nil
	}
	return nil, &NotLoadedError{edge: "unsaved_post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UnsavedPostImage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case unsavedpostimage.FieldID, unsavedpostimage.FieldWidth, unsavedpostimage.FieldHeight:
			values[i] = new(sql.NullInt64)
		case unsavedpostimage.FieldUUID, unsavedpostimage.FieldValidity, unsavedpostimage.FieldHash, unsavedpostimage.FieldTitle, unsavedpostimage.FieldURL:
			values[i] = new(sql.NullString)
		case unsavedpostimage.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case unsavedpostimage.ForeignKeys[0]: // unsaved_post_images
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UnsavedPostImage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UnsavedPostImage fields.
func (upi *UnsavedPostImage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case unsavedpostimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			upi.ID = int(value.Int64)
		case unsavedpostimage.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				upi.UUID = value.String
			}
		case unsavedpostimage.FieldValidity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field validity", values[i])
			} else if value.Valid {
				upi.Validity = unsavedpostimage.Validity(value.String)
			}
		case unsavedpostimage.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				upi.Width = new(uint32)
				*upi.Width = uint32(value.Int64)
			}
		case unsavedpostimage.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				upi.Height = new(uint32)
				*upi.Height = uint32(value.Int64)
			}
		case unsavedpostimage.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				upi.Hash = new(string)
				*upi.Hash = value.String
			}
		case unsavedpostimage.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				upi.Title = new(string)
				*upi.Title = value.String
			}
		case unsavedpostimage.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				upi.URL = new(string)
				*upi.URL = value.String
			}
		case unsavedpostimage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				upi.CreatedAt = value.Time
			}
		case unsavedpostimage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field unsaved_post_images", value)
			} else if value.Valid {
				upi.unsaved_post_images = new(int)
				*upi.unsaved_post_images = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUnsavedPost queries the "unsaved_post" edge of the UnsavedPostImage entity.
func (upi *UnsavedPostImage) QueryUnsavedPost() *UnsavedPostQuery {
	return (&UnsavedPostImageClient{config: upi.config}).QueryUnsavedPost(upi)
}

// Update returns a builder for updating this UnsavedPostImage.
// Note that you need to call UnsavedPostImage.Unwrap() before calling this method if this UnsavedPostImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (upi *UnsavedPostImage) Update() *UnsavedPostImageUpdateOne {
	return (&UnsavedPostImageClient{config: upi.config}).UpdateOne(upi)
}

// Unwrap unwraps the UnsavedPostImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (upi *UnsavedPostImage) Unwrap() *UnsavedPostImage {
	tx, ok := upi.config.driver.(*txDriver)
	if !ok {
		panic("ent: UnsavedPostImage is not a transactional entity")
	}
	upi.config.driver = tx.drv
	return upi
}

// String implements the fmt.Stringer.
func (upi *UnsavedPostImage) String() string {
	var builder strings.Builder
	builder.WriteString("UnsavedPostImage(")
	builder.WriteString(fmt.Sprintf("id=%v", upi.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(upi.UUID)
	builder.WriteString(", validity=")
	builder.WriteString(fmt.Sprintf("%v", upi.Validity))
	if v := upi.Width; v != nil {
		builder.WriteString(", width=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := upi.Height; v != nil {
		builder.WriteString(", height=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := upi.Hash; v != nil {
		builder.WriteString(", hash=")
		builder.WriteString(*v)
	}
	if v := upi.Title; v != nil {
		builder.WriteString(", title=")
		builder.WriteString(*v)
	}
	if v := upi.URL; v != nil {
		builder.WriteString(", url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", created_at=")
	builder.WriteString(upi.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UnsavedPostImages is a parsable slice of UnsavedPostImage.
type UnsavedPostImages []*UnsavedPostImage

func (upi UnsavedPostImages) config(cfg config) {
	for _i := range upi {
		upi[_i].config = cfg
	}
}
