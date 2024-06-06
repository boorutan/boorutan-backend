package booru

import (
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/utils/http"
	"applemango/boorutan/backend/utils/image"
	"encoding/json"
	"fmt"
)

type GetPostsOption struct {
	Page  int
	Tags  any
	Cache bool
}

func (p *Post) ToString() (string, error) {
	postJson, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	v := string(postJson)
	return v, nil
}

func StringToPost(sp string) (*Post, error) {
	var post *Post
	if err := json.Unmarshal([]byte(sp), &post); err != nil {
		return nil, err
	}
	summary, err := post.GetImageMock()
	if err == nil {
		(*post).Summary = summary
	}
	return post, nil
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
	fmt.Println(url)
	err := http.RequestJSON(http.RequestOption{
		Data:   &post,
		Url:    url,
		Method: "GET",
		Body:   nil,
		Cache:  option.Cache,
	})
	if err != nil || post == nil {
		return nil, err
	}
	for _, p := range *post {
		v, err := p.ToString()
		if err != nil {
			println(err.Error())
			continue
		}
		_ = redis.Push(fmt.Sprintf("cache:post:%v:%v", b.Base, p.ID), v)
	}
	for i, p := range *post {
		summary, err := p.GetImageMock()
		if err != nil {
			continue
		}
		(*post)[i].Summary = summary
	}
	return post, err
}

func (p *Post) GetImageMock() ([]image.Color, error) {
	url, err := p.GetPostSampleImage()
	if err != nil {
		return []image.Color{}, nil
	}
	println(url)
	uuid, err := image.GetImageUuid(url)
	if err != nil {
		return []image.Color{}, nil
	}
	summary, err := image.GetImageMock(uuid)
	if err != nil {
		return []image.Color{}, nil
	}
	return summary, nil
}

func (p *Post) GetPostSampleImage() (string, error) {
	if p.PreviewURL != "" {
		return p.PreviewURL, nil
	}
	println(p.PreviewFileURL)
	if p.PreviewFileURL != "" {
		return p.PreviewFileURL, nil
	}
	for _, media := range p.MediaAsset.Variants {
		if media.Type == "sample" {
			return media.URL, nil
		}
	}
	return p.FileURL, nil
}

// GetPostOption 仕組み上Cacheは強制的にTrue
type GetPostOption struct {
	ID int
}

func (b *Booru) GetPost(option GetPostOption) (*Post, error) {
	cache, err := redis.Get(fmt.Sprintf("cache:post:%v:%v", b.Base, option.ID))
	if err != nil {
		return nil, err
	}
	return StringToPost(cache)
}
