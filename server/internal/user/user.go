package user

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/imdotdev/im.dev/server/internal/interaction"
	"github.com/imdotdev/im.dev/server/internal/org"
	"github.com/imdotdev/im.dev/server/internal/tags"
	"github.com/imdotdev/im.dev/server/pkg/db"
	"github.com/imdotdev/im.dev/server/pkg/e"
	"github.com/imdotdev/im.dev/server/pkg/models"
)

func GetUsers(q string) ([]*models.User, *e.Error) {
	users := make([]*models.User, 0)
	for _, u := range models.UsersCache {
		if strings.HasPrefix(strings.ToLower(u.Nickname), strings.ToLower(q)) {
			users = append(users, u)
			continue
		}

		if strings.HasPrefix(strings.ToLower(u.Username), strings.ToLower(q)) {
			users = append(users, u)
			continue
		}
	}

	return users, nil
}

func GetUsersByIDs(ids []string) []*models.User {
	users := make([]*models.User, 0)
	for _, id := range ids {
		u, ok := models.UsersMapCache[id]
		if ok {
			users = append(users, u)
		}
		u.Followed = true
	}

	return users
}

func GetUserDetail(id string, username string) (*models.User, *e.Error) {
	user := &models.User{}
	err := user.Query(id, username, "")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.New(http.StatusNotFound, e.NotFound)
		}
		logger.Warn("query user error", "error", err)
		return nil, e.New(http.StatusInternalServerError, e.Internal)
	}

	err = db.Conn.QueryRow("SELECT tagline,cover,location,avail_for,about,website,twitter,github,zhihu,weibo,facebook,stackoverflow from user_profile WHERE id=?", user.ID).Scan(
		&user.Tagline, &user.Cover, &user.Location, &user.AvailFor, &user.About, &user.Website, &user.Twitter,
		&user.Github, &user.Zhihu, &user.Weibo, &user.Facebook, &user.Stackoverflow,
	)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("query user profile error", "error", err)
		return nil, e.New(http.StatusInternalServerError, e.Internal)
	}

	if user.Cover == "" {
		user.Cover = models.DefaultCover
	}

	// get user skills
	skills, rawSkills, err := tags.GetTargetTags(user.ID)
	if err != nil {
		logger.Warn("get user skills error", "error", err)
		return nil, e.New(http.StatusInternalServerError, e.Internal)
	}
	user.RawSkills = rawSkills
	user.Skills = skills

	user.Follows = interaction.GetFollows(user.ID)
	if user.Type == models.IDTypeUser {
		user.Followings = interaction.GetFollowings(user.ID, models.IDTypeUser)
	} else {
		user.Followings = org.GetMemberCount(user.ID)
	}

	return user, nil
}

func UpdateUser(u *models.User) *e.Error {
	_, err := db.Conn.Exec("UPDATE user SET nickname=?,avatar=?,email=?,updated=? WHERE id=?", u.Nickname, u.Avatar, u.Email, time.Now(), u.ID)
	if err != nil {
		if e.IsErrUniqueConstraint(err) {
			return e.New(http.StatusConflict, "email已经存在")
		}
		logger.Warn("update user error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	var nid string
	err = db.Conn.QueryRow("SELECT id FROM user_profile WHERE id=?", u.ID).Scan(&nid)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("update user profile error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	now := time.Now()
	if err == sql.ErrNoRows {
		_, err = db.Conn.Exec("INSERT INTO user_profile (id,tagline,cover,location,avail_for,about,website,twitter,github,zhihu,weibo,facebook,stackoverflow,updated) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			u.ID, u.Tagline, u.Cover, u.Location, u.AvailFor, u.About, u.Website, u.Twitter, u.Github, u.Zhihu, u.Weibo, u.Facebook, u.Stackoverflow, now)
	} else {
		_, err = db.Conn.Exec("UPDATE user_profile SET tagline=?,cover=?,location=?,avail_for=?,about=?,website=?,twitter=?,github=?,zhihu=?,weibo=?,facebook=?,stackoverflow=?,updated=? WHERE id=?",
			u.Tagline, u.Cover, u.Location, u.AvailFor, u.About, u.Website, u.Twitter, u.Github, u.Zhihu, u.Weibo, u.Facebook, u.Stackoverflow, now, u.ID)
	}

	if err != nil {
		logger.Warn("update user profile error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	//update user skills
	err = tags.UpdateTargetTags("", u.ID, u.Skills)
	if err != nil {
		logger.Warn("upate tags error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	return nil
}

func NameExist(name string) (bool, *e.Error) {
	var username string
	err := db.Conn.QueryRow("SELECT username FROM user  WHERE username=?", name).Scan(&username)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("check name exist  error", "error", err)
		return false, e.New(http.StatusInternalServerError, e.Internal)
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
