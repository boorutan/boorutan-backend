package image

import (
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/utils/http"
	"fmt"

	"github.com/google/uuid"
	r "github.com/redis/go-redis/v9"
)

func DownloadImage(url string) (string, error) {
	uuid := uuid.NewString()
	path := fmt.Sprintf("./static/images/%s", uuid)
	err := http.RequestDownloadImage(url, path)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func GetImage(url string) (string, error) {
	uuid, err := redis.Get(fmt.Sprintf("image:%s", url))
	if err != nil && err != r.Nil {
		return "", err
	}
	if err == r.Nil {
		uuid, err := DownloadImage(url)
		redis.Push(fmt.Sprintf("image:%s", url), uuid)
		return uuid, err
	}
	return uuid, err
}
