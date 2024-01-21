package handler

import (
	"applemango/boorutan/backend/booru"
	"applemango/boorutan/backend/db/sqlite3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LikePost(c *gin.Context) {
	type Body struct {
		Like bool `json:"like"`
	}
	var b Body
	err := c.Bind(&b)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	booruname := c.Param("booru")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	_, _ = sqlite3.DB.Exec("DELETE FROM like WHERE booru = ? AND post_id = ? AND user_id = ?", booruname, id, 1)
	if b.Like {
		_, _ = sqlite3.DB.Exec("INSERT INTO like (booru, post_id, user_id) VALUES ( ?, ?, ? )", booruname, id, 1)
	}
	type msg struct {
		Msg string `json:"msg"`
	}
	c.JSON(http.StatusOK, msg{Msg: "success"})
	return
}

func GetLikedPost(c *gin.Context) {
	rows, err := sqlite3.DB.Query("SELECT id, booru, post_id, user_id FROM like WHERE user_id = ? ORDER BY id DESC", 1)
	type like struct {
		ID     int64  `json:"id"`
		Booru  string `json:"booru"`
		PostId int64  `json:"post_id"`
		UserId int64  `json:"user_id"`
	}
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var posts []*booru.Post
	for rows.Next() {
		l := like{}
		if err = rows.Scan(&l.ID, &l.Booru, &l.PostId, &l.UserId); err != nil {
			println(err.Error())
			break
		}
		b := GetBooruFromString(l.Booru)
		post, err := b.GetPost(booru.GetPostOption{
			ID: int(l.PostId),
		})
		if err != nil {
			println(err.Error())
			break
		}
		post.BooruType = l.Booru
		posts = append(posts, post)
	}
	c.JSON(http.StatusOK, posts)
	return
}
