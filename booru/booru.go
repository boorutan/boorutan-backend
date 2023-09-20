package booru

import (
	"applemango/boorutan/backend/utils/http"
	"fmt"
)

func CreateBooru(t BooruType, u string) Booru {
	if t == MoeBooru {
		return Booru{
			BooruType: MoeBooru,
		}
	}
	return Booru{
		BooruType: DanBooru,
	}
}

func (b *Booru) GetPost(cache bool) (*[]BooruPost, error) {
	var post *[]BooruPost
	err := http.RequestJSON(http.RequestOption{
		Data:   &post,
		Url:    fmt.Sprintf("%v%v", b.Base, b.Url.Post),
		Method: "GET",
		Body:   nil,
		Cache:  cache,
	})
	return post, err
}
