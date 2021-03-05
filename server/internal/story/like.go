package story

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/imdotdev/im.dev/server/pkg/db"
	"github.com/imdotdev/im.dev/server/pkg/e"
)

func Like(storyID string, userId string) *e.Error {
	exist := Exist(storyID)
	if !exist {
		return e.New(http.StatusNotFound, e.NotFound)
	}

	// 查询当前like状态
	liked := GetLiked(storyID, userId)

	var count int
	err := db.Conn.QueryRow("SELECT count FROM likes_count WHERE story_id=?", storyID).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("query likes count error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}
	exist = !(err == sql.ErrNoRows)

	tx, err := db.Conn.Begin()
	if err != nil {
		logger.Warn("start like transaction error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	if liked {
		// 已经喜欢过该篇文章，更改为不喜欢
		_, err := tx.Exec("DELETE FROM likes WHERE story_id=? and user_id=?", storyID, userId)
		if err != nil {
			return e.New(http.StatusInternalServerError, e.Internal)
		}
		count = count - 1
	} else {
		_, err := tx.Exec("INSERT INTO likes (story_id,user_id,created) VALUES (?,?,?)", storyID, userId, time.Now())
		if err != nil {
			logger.Warn("add like error", "error", err)
			return e.New(http.StatusInternalServerError, e.Internal)
		}
		count = count + 1
	}

	var err0 error
	if !exist {
		_, err0 = tx.Exec("INSERT INTO likes_count (story_id,count) VALUES (?,?)", storyID, count)
	} else {
		_, err0 = tx.Exec("UPDATE likes_count SET count=? WHERE story_id=?", count, storyID)
	}

	if err0 != nil {
		logger.Warn("add like error", "error", err0)
		tx.Rollback()
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	tx.Commit()
	return nil
}

func GetLiked(storyID string, userID string) bool {
	liked := false
	var nid string
	err := db.Conn.QueryRow("SELECT story_id FROM likes WHERE story_id=? and user_id=?", storyID, userID).Scan(&nid)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("query story like error", "error", err)
		return false
	}

	if nid == storyID {
		liked = true
	}

	return liked
}

func GetLikes(storyID string) int {
	var likes int
	err := db.Conn.QueryRow("SELECT count FROM likes_count WHERE story_id=?", storyID).Scan(&likes)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn("get like count error", "error", err)
	}

	return likes
}
