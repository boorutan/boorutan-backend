package image

import (
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/utils/http"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	r "github.com/redis/go-redis/v9"
)

func DownloadImage(url string) (string, error) {
	id := uuid.NewString()
	path := fmt.Sprintf("./static/images/%s", id)
	err := http.RequestDownloadImage(url, path)
	if err != nil {
		return "", err
	}
	_, _ = GetImageMock(id)
	return id, nil
}

func GetImage(url string) (string, error) {
	id, err := redis.Get(fmt.Sprintf("image:%s", url))
	if err != nil && err != r.Nil {
		return "", err
	}
	if err == r.Nil {
		id, err := DownloadImage(url)
		redis.Push(fmt.Sprintf("image:%s", url), id)
		return id, err
	}
	return id, err
}

func GetImageUuid(url string) (string, error) {
	id, err := redis.Get(fmt.Sprintf("image:%s", url))
	if err != nil && err != r.Nil {
		return "", err
	}
	if err == r.Nil {
		return "", err
	}
	return id, nil
}

func GetImageMock(uuid string) ([]Color, error) {
	cache, err := redis.Get(fmt.Sprintf("image:color:%s", uuid))
	if err != nil && err != r.Nil {
		return []Color{}, err
	}
	if err == r.Nil {
		summary, err := GetImageSummary(uuid, 5, 5)
		if err != nil {
			return []Color{}, err
		}
		summaryJson, err := json.Marshal(summary)
		if err != nil {
			return []Color{}, err
		}
		_ = redis.Push(fmt.Sprintf("image:color:%s", uuid), string(summaryJson))
		return summary, nil
	}
	var colors *[]Color
	if err := json.Unmarshal([]byte(cache), &colors); err != nil {
		return []Color{}, err
	}
	return *colors, nil
}
