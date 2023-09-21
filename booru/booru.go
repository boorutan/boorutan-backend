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

type GetPostOption struct {
	Page  int
	Cache bool
}

func (b *Booru) GetPost(option GetPostOption) (*[]BooruPost, error) {
	var post *[]BooruPost
	var url string
	url = fmt.Sprintf("%v%v?page=%v", b.Base, b.Url.Post, option.Page)
	if option.Page == 1 {
		url = fmt.Sprintf("%v%v", b.Base, b.Url.Post)
	}
	err := http.RequestJSON(http.RequestOption{
		Data:   &post,
		Url:    url,
		Method: "GET",
		Body:   nil,
		Cache:  option.Cache,
	})
	return post, err
}

type GetTagOption struct {
	Cache bool
}

func (b *Booru) GetTag(option GetTagOption) (*[]BooruTag, error) {
	var tag *[]BooruTag
	var url = fmt.Sprintf("%v%v", b.Base, b.Url.Tag)
	err := http.RequestJSON(http.RequestOption{
		Data:   &tag,
		Url:    url,
		Method: "GET",
		Body:   nil,
		Cache:  option.Cache,
	})
	return tag, err
}
