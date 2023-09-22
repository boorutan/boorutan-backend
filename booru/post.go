package booru

import (
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/utils/http"
	"encoding/json"
	"fmt"
)

type GetPostsOption struct {
	Page  int
	Tags  any
	Cache bool
}

func (b *Booru) GetPosts(option GetPostsOption) (*[]Post, error) {
	var post *[]Post
	var url string
	url = fmt.Sprintf("%v%v?page=%v", b.Base, b.Url.Post, option.Page)
	if option.Page == 1 {
		url = fmt.Sprintf("%v%v", b.Base, b.Url.Post)
	}
	if option.Tags != nil {
		a := "&"
		if option.Page == 1 {
			a = "?"
		}
		url = fmt.Sprintf("%v%vtags=%v", url, a, option.Tags)
	}
	err := http.RequestJSON(http.RequestOption{
		Data:   &post,
		Url:    url,
		Method: "GET",
		Body:   nil,
		Cache:  option.Cache,
	})
	for _, p := range *post {
		postJson, err := json.Marshal(p)
		if err != nil {
			println(err.Error())
			continue
		}
		v := string(postJson)
		_ = redis.Push(fmt.Sprintf("cache:post:%v:%v", b.Base, p.ID), v)
	}
	return post, err
}

// GetPostOption 仕組み上Cacheは強制的にTrue
type GetPostOption struct {
	ID int
}

func (b *Booru) GetPost(option GetPostOption) (*Post, error) {
	var post *Post
	cache, err := redis.Get(fmt.Sprintf("cache:post:%v:%v", b.Base, option.ID))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(cache), &post); err != nil {
		return nil, err
	}
	return post, nil
}
