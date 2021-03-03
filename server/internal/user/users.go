package user

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/imdotdev/im.dev/server/internal/cache"
	"github.com/imdotdev/im.dev/server/internal/tags"
	"github.com/imdotdev/im.dev/server/pkg/db"
	"github.com/imdotdev/im.dev/server/pkg/e"
	"github.com/imdotdev/im.dev/server/pkg/models"
)

func GetUsers(q string) ([]*models.User, *e.Error) {
	allUsers := cache.Users

	users := make([]*models.User, 0)
	for _, u := range allUsers {
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

func GetUserDetail(id int64, username string) (*models.User, *e.Error) {
	user := &models.User{}
	err := user.Query(id, username, "")
	if err != nil {
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
	user.Skills = make([]int64, 0)
	user.RawSkills = make([]*models.Tag, 0)
	rows, err := db.Conn.Query("SELECT skill_id from user_skills WHERE user_id=?", user.ID)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("query user skills error", "error", err)
	}
	for rows.Next() {
		var skill int64
		rows.Scan(&skill)
		user.Skills = append(user.Skills, skill)

		rawTag, err := tags.GetTag(skill, "")
		if err != nil {
			logger.Warn("get tag error", "error", err)
			continue
		}
		user.RawSkills = append(user.RawSkills, rawTag)
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

	var nid int64
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
	_, err = db.Conn.Exec("DELETE FROM user_skills WHERE user_id=?", u.ID)
	if err != nil {
		logger.Warn("delete user skills error", "error", err)
	}

	for _, skill := range u.Skills {
		_, err = db.Conn.Exec("INSERT INTO user_skills (user_id,skill_id) VALUES (?,?)", u.ID, skill)
		if err != nil {
			logger.Warn("add user skill error", "error", err)
		}
	}

	return nil
}
