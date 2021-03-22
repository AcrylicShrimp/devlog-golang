package admin

import (
	"devlog/common"
	"devlog/ent"
	"devlog/ent/category"
	post2 "devlog/ent/post"
	"github.com/labstack/echo"
	"net/http"
)

func AttachPost(group *echo.Group) {

}

func NewPostHandler(c echo.Context) error {
	type PostInfo struct {
		Slug        string `json:"slug" form:"slug" query:"slug" validate:"required,max=255"`
		AccessLevel string `json:"access-level" form:"access-level" query:"access-level" validate:"required,oneof=public unlisted private"`
		Category    string `json:"category" form:"category" query:"category"`
		Title       string `json:"title" form:"title" query:"title" validate:"required,max=255"`
		Content     string `json:"content" form:"content" query:"content" validate:"required"`
	}

	postInfo := new(PostInfo)
	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	var cat *ent.Category = nil
	if postInfo.Category != "" {
		category, err := client.Category.Query().Where(category.NameEQ(postInfo.Category)).Select(category.FieldID).First()
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		cat = category
	}

	post, err := client.Post.Create().SetSlug(postInfo.Slug).SetAccessLevel(post2.AccessLevel()).Save()
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return echo.NewHTTPError(http.StatusConflict)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
