package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdotdev/im.dev/server/internal/story"
	"github.com/imdotdev/im.dev/server/internal/user"
	"github.com/imdotdev/im.dev/server/pkg/common"
	"github.com/imdotdev/im.dev/server/pkg/e"
)

func SubmitPost(c *gin.Context) {
	res, err := story.SubmitPost(c)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(res))
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.RespError(e.ParamInvalid))
		return
	}

	user := user.CurrentUser(c)
	creator, err := story.GetPostCreator(id)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	if user.ID != creator {
		c.JSON(http.StatusForbidden, common.RespError(e.NoPermission))
		return
	}

	err = story.DeletePost(id)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}
func GetPost(c *gin.Context) {
	id := c.Param("id")

	ar, err := story.GetPost(id, "")
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	user := user.CurrentUser(c)
	if user == nil {
		ar.Liked = false
	} else {
		ar.Liked = story.GetLiked(ar.ID, user.ID)

	}

	c.JSON(http.StatusOK, common.RespSuccess(ar))
}

func LikeStory(c *gin.Context) {
	user := user.CurrentUser(c)
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.RespError(e.ParamInvalid))
		return
	}

	err := story.Like(id, user.ID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}
