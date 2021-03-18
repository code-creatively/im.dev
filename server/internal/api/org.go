package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdotdev/im.dev/server/internal/org"
	"github.com/imdotdev/im.dev/server/internal/user"
	"github.com/imdotdev/im.dev/server/pkg/common"
	"github.com/imdotdev/im.dev/server/pkg/db"
	"github.com/imdotdev/im.dev/server/pkg/e"
	"github.com/imdotdev/im.dev/server/pkg/models"
)

func CreateOrg(c *gin.Context) {
	o := &models.User{}
	c.Bind(&o)

	exist, err := user.NameExist(o.Username)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	if exist {
		c.JSON(http.StatusConflict, common.RespError(e.AlreadyExist))
		return
	}

	u := user.CurrentUser(c)
	err = org.Create(o, u.ID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}

func UpdateOrg(c *gin.Context) {
	u := &models.User{}
	c.Bind(&u)

	currentUser := user.CurrentUser(c)
	if !org.IsOrgAdmin(currentUser.ID, u.ID) {
		c.JSON(http.StatusForbidden, common.RespError(e.NoPermission))
		return
	}

	err := user.UpdateUser(u)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}

func GetOrgByUserID(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "0" {
		u := user.CurrentUser(c)
		if u == nil {
			c.JSON(http.StatusBadRequest, common.RespError(e.BadRequest))
			return
		}

		userID = u.ID
	}

	orgs, err := org.GetOrgByUserID(userID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(orgs))
}

func GetOrgMembers(c *gin.Context) {
	orgID := c.Param("id")

	u := user.CurrentUser(c)
	users, err := org.GetMembers(u, orgID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(users))
}

func GenOrgSecret(c *gin.Context) {
	orgID := c.Param("id")
	currentUser := user.CurrentUser(c)
	if !org.IsOrgAdmin(currentUser.ID, orgID) {
		c.JSON(http.StatusForbidden, common.RespError(e.NoPermission))
		return
	}

	secret, err := user.GenSecret(orgID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(secret))
}

func GetOrgSecret(c *gin.Context) {
	orgID := c.Param("id")
	currentUser := user.CurrentUser(c)
	if !org.IsOrgAdmin(currentUser.ID, orgID) {
		c.JSON(http.StatusForbidden, common.RespError(e.NoPermission))
		return
	}

	secret, err := user.GetSecret(orgID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(secret))
}

type JoinOrgReq struct {
	Secret string `json:"secret"`
}

func JoinOrg(c *gin.Context) {
	req := &JoinOrgReq{}
	c.Bind(&req)

	u := user.CurrentUser(c)
	err := org.Join(req.Secret, u.ID)
	if err != nil {
		c.JSON(err.Status, common.RespError(err.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}

type UpdateMemberReq struct {
	OrgID    string          `json:"orgID"`
	MemberID string          `json:"memberID"`
	Role     models.RoleType `json:"role"`
}

func UpdateOrgMember(c *gin.Context) {
	req := &UpdateMemberReq{}
	c.Bind(&req)

	if !req.Role.IsValid() {
		c.JSON(http.StatusBadRequest, common.RespError(e.BadRequest))
		return
	}

	if req.Role == models.ROLE_SUPER_ADMIN || req.Role == models.ROLE_EDITOR {
		c.JSON(http.StatusBadRequest, common.RespError(e.BadRequest))
		return
	}
	u := user.CurrentUser(c)
	var role models.RoleType
	err := db.Conn.QueryRow("SELECT role FROM org_member WHERE org_id=? and user_id=?", req.OrgID, u.ID).Scan(&role)
	if err != nil {
		logger.Warn("select role error", "error", err)
		c.JSON(http.StatusBadRequest, common.RespError(e.BadRequest))
		return
	}

	// 修改角色要求至少是管理员
	if !role.IsAdmin() {
		c.JSON(http.StatusForbidden, common.RespError("你需要成为组织的管理员"))
		return
	}

	// 目标角色修改为管理员，要求当前操作用户必须是超级管理员
	if req.Role == models.ROLE_ADMIN {
		if role != models.ROLE_SUPER_ADMIN {
			c.JSON(http.StatusForbidden, common.RespError("你需要成为组织的超级管理员"))
			return
		}
	}

	// 若目标角色之前是管理员，那么必须是超级管理员才能对其修改
	targetRole, err := org.GetMemberRole(req.OrgID, req.MemberID)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.RespError(e.BadRequest))
		return
	}
	if targetRole == models.ROLE_SUPER_ADMIN {
		c.JSON(http.StatusForbidden, common.RespError("超级管理员不能被修改"))
		return
	}
	if targetRole == models.ROLE_ADMIN {
		if role != models.ROLE_SUPER_ADMIN {
			c.JSON(http.StatusForbidden, common.RespError("你需要成为组织的超级管理员"))
			return
		}
	}

	err0 := org.UpdateMember(req.OrgID, req.MemberID, req.Role)
	if err0 != nil {
		c.JSON(err0.Status, common.RespError(err0.Message))
		return
	}

	c.JSON(http.StatusOK, common.RespSuccess(nil))
}
