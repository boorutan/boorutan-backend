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

type TagCategory struct {
	Category string
	Name     string
}

func (b *Booru) GetTagCategory() TagCategory {

}
