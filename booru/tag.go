package booru

import (
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/utils/http"
	"fmt"
	"strings"
)

type Tags struct {
	Version int    `json:"version"`
	Data    string `json:"data"`
}

func (b *Booru) GetTags() error {
	var tags *Tags
	var url string
	if b.Url.TagSummary != nil {
		url = fmt.Sprintf("%v%v", b.Base, b.Url.TagSummary)
	} else {
		url = "https://konachan.com/tag/summary.json"
	}
	err := http.RequestJSON(http.RequestOption{
		Data:   &tags,
		Url:    url,
		Method: "POST",
		Body:   nil,
		Cache:  true,
	})
	if err != nil {
		return err
	}
	tagStr := strings.Split(tags.Data, " ")
	for _, t := range tagStr {
		if len(t) <= 3 {
			continue
		}
		category := string(t[0])
		name := t[2 : len(t)-1]
		_ = redis.Push(fmt.Sprintf("cache:tag:%v:%v", url, name), category)
		//fmt.Printf("%v, %v\n", category, name)
	}
	fmt.Printf("All tags saved, length: %v\n", len(tagStr))
	return nil
}

func (b *Booru) GetTagCategory(name string) (string, error) {
	var url string
	if b.Url.TagSummary != nil {
		url = fmt.Sprintf("%v%v", b.Base, b.Url.TagSummary)
	} else {
		url = "https://konachan.com/tag/summary.json"
	}
	fmt.Printf("cache:tag:%v:%v\n", url, name)
	cache, err := redis.Get(fmt.Sprintf("cache:tag:%v:%v", url, name))
	if err != nil {
		return "", err
	}
	return cache, nil
}

type TagCategory struct {
	Name     string
	Category string
}

func (b *Booru) GetTagsCategory(names []string) ([]TagCategory, error) {
	var tags []TagCategory
	for _, name := range names {
		category, err := b.GetTagCategory(name)
		if err != nil {
			category = "0"
		}
		tags = append(tags, TagCategory{
			Name:     name,
			Category: category,
		})
	}
	return tags, nil
}

func (b *Booru) GetTagsCategoryFromString(names string) ([]TagCategory, error) {
	tags := strings.Split(names, " ")
	return b.GetTagsCategory(tags)
}
