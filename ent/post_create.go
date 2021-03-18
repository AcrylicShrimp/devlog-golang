// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/category"
	"devlog/ent/post"
	"devlog/ent/postimage"
	"devlog/ent/postthumbnail"
	"devlog/ent/postvideo"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (pc *PostCreate) SetUUID(s string) *PostCreate {
	pc.mutation.SetUUID(s)
	return pc
}

// SetSlug sets the "slug" field.
func (pc *PostCreate) SetSlug(s string) *PostCreate {
	pc.mutation.SetSlug(s)
	return pc
}

// SetAccessLevel sets the "access_level" field.
func (pc *PostCreate) SetAccessLevel(pl post.AccessLevel) *PostCreate {
	pc.mutation.SetAccessLevel(pl)
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pc *PostCreate) SetNillableContent(s *string) *PostCreate {
	if s != nil {
		pc.SetContent(*s)
	}
	return pc
}

// SetHTMLContent sets the "html_content" field.
func (pc *PostCreate) SetHTMLContent(s string) *PostCreate {
	pc.mutation.SetHTMLContent(s)
	return pc
}

// SetNillableHTMLContent sets the "html_content" field if the given value is not nil.
func (pc *PostCreate) SetNillableHTMLContent(s *string) *PostCreate {
	if s != nil {
		pc.SetHTMLContent(*s)
	}
	return pc
}

// SetPreviewContent sets the "preview_content" field.
func (pc *PostCreate) SetPreviewContent(s string) *PostCreate {
	pc.mutation.SetPreviewContent(s)
	return pc
}

// SetNillablePreviewContent sets the "preview_content" field if the given value is not nil.
func (pc *PostCreate) SetNillablePreviewContent(s *string) *PostCreate {
	if s != nil {
		pc.SetPreviewContent(*s)
	}
	return pc
}

// SetAccumulatedImageIndex sets the "accumulated_image_index" field.
func (pc *PostCreate) SetAccumulatedImageIndex(u uint64) *PostCreate {
	pc.mutation.SetAccumulatedImageIndex(u)
	return pc
}

// SetNillableAccumulatedImageIndex sets the "accumulated_image_index" field if the given value is not nil.
func (pc *PostCreate) SetNillableAccumulatedImageIndex(u *uint64) *PostCreate {
	if u != nil {
		pc.SetAccumulatedImageIndex(*u)
	}
	return pc
}

// SetAccumulatedVideoIndex sets the "accumulated_video_index" field.
func (pc *PostCreate) SetAccumulatedVideoIndex(u uint64) *PostCreate {
	pc.mutation.SetAccumulatedVideoIndex(u)
	return pc
}

// SetNillableAccumulatedVideoIndex sets the "accumulated_video_index" field if the given value is not nil.
func (pc *PostCreate) SetNillableAccumulatedVideoIndex(u *uint64) *PostCreate {
	if u != nil {
		pc.SetAccumulatedVideoIndex(*u)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetModifiedAt sets the "modified_at" field.
func (pc *PostCreate) SetModifiedAt(t time.Time) *PostCreate {
	pc.mutation.SetModifiedAt(t)
	return pc
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (pc *PostCreate) SetCategoryID(id int) *PostCreate {
	pc.mutation.SetCategoryID(id)
	return pc
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableCategoryID(id *int) *PostCreate {
	if id != nil {
		pc = pc.SetCategoryID(*id)
	}
	return pc
}

// SetCategory sets the "category" edge to the Category entity.
func (pc *PostCreate) SetCategory(c *Category) *PostCreate {
	return pc.SetCategoryID(c.ID)
}

// SetThumbnailID sets the "thumbnail" edge to the PostThumbnail entity by ID.
func (pc *PostCreate) SetThumbnailID(id int) *PostCreate {
	pc.mutation.SetThumbnailID(id)
	return pc
}

// SetNillableThumbnailID sets the "thumbnail" edge to the PostThumbnail entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableThumbnailID(id *int) *PostCreate {
	if id != nil {
		pc = pc.SetThumbnailID(*id)
	}
	return pc
}

// SetThumbnail sets the "thumbnail" edge to the PostThumbnail entity.
func (pc *PostCreate) SetThumbnail(p *PostThumbnail) *PostCreate {
	return pc.SetThumbnailID(p.ID)
}

// AddImageIDs adds the "images" edge to the PostImage entity by IDs.
func (pc *PostCreate) AddImageIDs(ids ...int) *PostCreate {
	pc.mutation.AddImageIDs(ids...)
	return pc
}

// AddImages adds the "images" edges to the PostImage entity.
func (pc *PostCreate) AddImages(p ...*PostImage) *PostCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddImageIDs(ids...)
}

// AddVideoIDs adds the "videos" edge to the PostVideo entity by IDs.
func (pc *PostCreate) AddVideoIDs(ids ...int) *PostCreate {
	pc.mutation.AddVideoIDs(ids...)
	return pc
}

// AddVideos adds the "videos" edges to the PostVideo entity.
func (pc *PostCreate) AddVideos(p ...*PostVideo) *PostCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddVideoIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.AccumulatedImageIndex(); !ok {
		v := post.DefaultAccumulatedImageIndex
		pc.mutation.SetAccumulatedImageIndex(v)
	}
	if _, ok := pc.mutation.AccumulatedVideoIndex(); !ok {
		v := post.DefaultAccumulatedVideoIndex
		pc.mutation.SetAccumulatedVideoIndex(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if v, ok := pc.mutation.UUID(); ok {
		if err := post.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New("ent: missing required field \"slug\"")}
	}
	if v, ok := pc.mutation.Slug(); ok {
		if err := post.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf("ent: validator failed for field \"slug\": %w", err)}
		}
	}
	if _, ok := pc.mutation.AccessLevel(); !ok {
		return &ValidationError{Name: "access_level", err: errors.New("ent: missing required field \"access_level\"")}
	}
	if v, ok := pc.mutation.AccessLevel(); ok {
		if err := post.AccessLevelValidator(v); err != nil {
			return &ValidationError{Name: "access_level", err: fmt.Errorf("ent: validator failed for field \"access_level\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := pc.mutation.PreviewContent(); ok {
		if err := post.PreviewContentValidator(v); err != nil {
			return &ValidationError{Name: "preview_content", err: fmt.Errorf("ent: validator failed for field \"preview_content\": %w", err)}
		}
	}
	if _, ok := pc.mutation.AccumulatedImageIndex(); !ok {
		return &ValidationError{Name: "accumulated_image_index", err: errors.New("ent: missing required field \"accumulated_image_index\"")}
	}
	if _, ok := pc.mutation.AccumulatedVideoIndex(); !ok {
		return &ValidationError{Name: "accumulated_video_index", err: errors.New("ent: missing required field \"accumulated_video_index\"")}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pc.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New("ent: missing required field \"modified_at\"")}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: post.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := pc.mutation.Slug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldSlug,
		})
		_node.Slug = value
	}
	if value, ok := pc.mutation.AccessLevel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: post.FieldAccessLevel,
		})
		_node.AccessLevel = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := pc.mutation.HTMLContent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldHTMLContent,
		})
		_node.HTMLContent = value
	}
	if value, ok := pc.mutation.PreviewContent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldPreviewContent,
		})
		_node.PreviewContent = value
	}
	if value, ok := pc.mutation.AccumulatedImageIndex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: post.FieldAccumulatedImageIndex,
		})
		_node.AccumulatedImageIndex = value
	}
	if value, ok := pc.mutation.AccumulatedVideoIndex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: post.FieldAccumulatedVideoIndex,
		})
		_node.AccumulatedVideoIndex = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.ModifiedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldModifiedAt,
		})
		_node.ModifiedAt = value
	}
	if nodes := pc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.CategoryTable,
			Columns: []string{post.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ThumbnailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   post.ThumbnailTable,
			Columns: []string{post.ThumbnailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: postthumbnail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ImagesTable,
			Columns: []string{post.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: postimage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.VideosTable,
			Columns: []string{post.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: postvideo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
