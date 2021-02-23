package posts

import (
	"database/sql"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/imdotdev/im.dev/server/pkg/db"
	"github.com/imdotdev/im.dev/server/pkg/e"
	"github.com/imdotdev/im.dev/server/pkg/models"
	"github.com/imdotdev/im.dev/server/pkg/utils"
)

func SubmitTag(tag *models.Tag) *e.Error {
	if strings.TrimSpace(tag.Title) == "" {
		return e.New(http.StatusBadRequest, "title格式不合法")
	}

	if strings.TrimSpace(tag.Name) == "" {
		return e.New(http.StatusBadRequest, "name格式不合法")
	}

	if strings.TrimSpace(tag.Name) == "new" {
		return e.New(http.StatusBadRequest, "name不能为new")
	}

	if strings.TrimSpace(tag.Cover) != "" && !govalidator.IsURL(tag.Cover) {
		return e.New(http.StatusBadRequest, "图片链接格式不正确")
	}

	if strings.TrimSpace(tag.Icon) != "" && !govalidator.IsURL(tag.Icon) {
		return e.New(http.StatusBadRequest, "图片链接格式不正确")
	}

	now := time.Now()

	md := utils.Compress(tag.Md)

	if tag.ID == 0 {
		//create
		_, err := db.Conn.Exec("INSERT INTO tags (creator,name, title, md, icon, cover, created, updated) VALUES(?,?,?,?,?,?,?,?)",
			tag.Creator, tag.Name, tag.Title, md, tag.Icon, tag.Cover, now, now)
		if err != nil {
			if e.IsErrUniqueConstraint(err) {
				return e.New(http.StatusConflict, "同样的Tag name已存在")
			}

			logger.Warn("submit post error", "error", err)
			return e.New(http.StatusInternalServerError, e.Internal)
		}
	} else {
		_, err := db.Conn.Exec("UPDATE tags SET name=?, title=?, md=?, icon=?, cover=?, updated=? WHERE id=?",
			tag.Name, tag.Title, md, tag.Icon, tag.Cover, now, tag.ID)
		if err != nil {
			logger.Warn("upate post error", "error", err)
			return e.New(http.StatusInternalServerError, e.Internal)
		}
	}

	return nil
}

func GetTags() (models.Tags, *e.Error) {
	tags := make(models.Tags, 0)

	rows, err := db.Conn.Query("SELECT id,creator,title,name,icon,cover,created,updated from tags")
	if err != nil {
		if err == sql.ErrNoRows {
			return tags, nil
		}
		logger.Warn("get tags error", "error", err)
		return tags, e.New(http.StatusInternalServerError, e.Internal)
	}

	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.ID, &tag.Creator, &tag.Title, &tag.Name, &tag.Icon, &tag.Cover, &tag.Created, &tag.Updated)
		if err != nil {
			logger.Warn("scan tags error", "error", err)
			continue
		}

		tags = append(tags, tag)
	}

	sort.Sort(tags)

	return tags, nil
}

func DeleteTag(id int64) *e.Error {
	_, err := db.Conn.Exec("DELETE FROM tags WHERE id=?", id)
	if err != nil {
		logger.Warn("delete post error", "error", err)
		return e.New(http.StatusInternalServerError, e.Internal)
	}

	return nil
}

func GetTag(name string) (*models.Tag, *e.Error) {
	tag := &models.Tag{}
	var rawmd []byte
	err := db.Conn.QueryRow("SELECT id,creator,title,name,icon,cover,created,updated,md from tags where name=?", name).Scan(
		&tag.ID, &tag.Creator, &tag.Title, &tag.Name, &tag.Icon, &tag.Cover, &tag.Created, &tag.Updated, &rawmd,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.New(http.StatusNotFound, e.NotFound)
		}
		logger.Warn("get post error", "error", err)
		return nil, e.New(http.StatusInternalServerError, e.Internal)
	}

	md, _ := utils.Uncompress(rawmd)
	tag.Md = string(md)

	return tag, nil
}
