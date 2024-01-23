package booru

import (
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/db/sqlite3"
	"applemango/boorutan/backend/utils/http"
	"encoding/json"
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

type DanbooruTag struct {
	UpdatedAt string `json:"updated_at"`
	IsLocked  bool   `json:"is_locked"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	PostCount string `json:"post_count"`
	ID        string `json:"id"`
}

func SearchTags(name string) []DanbooruTag {
	var tags []DanbooruTag
	_, values, _ := redis.SearchKV(fmt.Sprintf("cache:tag:%v*", name), 10)
	for _, v := range values {
		var tag *DanbooruTag
		_ = json.Unmarshal([]byte(v), &tag)
		tags = append(tags, *tag)
	}
	return tags
}

func SearchTagsFast(name string) []DanbooruTag {
	var tags []DanbooruTag
	rows, err := sqlite3.TagDB.Query("SELECT id, name, category, post_count FROM tag WHERE name LIKE ? OR translated_name LIKE ? ORDER BY post_count DESC LIMIT 30", "%"+name+"%", "%"+name+"%")
	for rows.Next() {
		tag := DanbooruTag{}
		if err = rows.Scan(&tag.ID, &tag.Name, &tag.Category, &tag.PostCount); err != nil {
			break
		}
		tags = append(tags, tag)
	}
	return tags
}

type TranslatedTag struct {
	id             int
	name           string
	translated     bool
	translatedName string
}

func TranslateTags(tags []string) []string {
	var translated []string
	for _, v := range tags {
		row := sqlite3.TagDB.QueryRow("SELECT id, name, translated, translated_name FROM tag WHERE translated = true AND name = ? OR alias LIKE '%' || ? || '%'", v, v)
		var tag TranslatedTag
		err := row.Scan(&tag.id, &tag.name, &tag.translated, &tag.translatedName)
		if err != nil || !tag.translated {
			translated = append(translated, v)
			continue
		}
		translated = append(translated, tag.translatedName)
	}
	return translated
}
