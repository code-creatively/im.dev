package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imdotdev/im.dev/server/internal/tags"
	"github.com/imdotdev/im.dev/server/internal/user"
	"github.com/imdotdev/im.dev/server/pkg/common"
	"github.com/imdotdev/im.dev/server/pkg/e"
	"github.com/imdotdev/im.dev/server/pkg/models"
)

func GetTag(c *gin.Context) {
	name := c.Param("name")
	res, err := tags.GetTag("", name)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(res))
}

func GetTags(c *gin.Context) {
	res, err := tags.GetTags()
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(res))
}

func GetTagsByIDs(c *gin.Context) {
	ids := make([]string, 0)
	err := c.Bind(&ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.RespError(e.ParamInvalid))
		return
	}

	ts, err0 := tags.GetTagsByIDs(ids)
	if err != nil {
		c.JSON(err0.Status, common.RespError(err0.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(ts))
}

func SubmitTag(c *gin.Context) {
	user := user.CurrentUser(c)
	if !user.Role.IsAdmin() {
		c.JSON(http.StatusForbidden, common.RespError(e.NoEditorPermission))
		return
	}

	tag := &models.Tag{}
	err := c.Bind(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.RespError(e.ParamInvalid))
		return
	}

	tag.Creator = user.ID
	err1 := tags.SubmitTag(tag)
	if err1 != nil {
		c.JSON(err1.Status, common.RespError(err1.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id == 0 {
		c.JSON(http.StatusBadRequest, common.RespError(e.ParamInvalid))
		return
	}

	user := user.CurrentUser(c)
	if !user.Role.IsAdmin() {
		c.JSON(http.StatusForbidden, common.RespError(e.NoPermission))
	}

	err := tags.DeleteTag(id)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}

func GetUserTags(c *gin.Context) {
	userID := c.Param("userID")
	res, err := tags.GetUserTags(userID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(res))
}
