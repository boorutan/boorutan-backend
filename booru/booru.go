package booru

import (
	"applemango/boorutan/backend/utils/http"
	"fmt"
)

func CreateBooru(t Type, u string) Booru {
	if t == MoeBooru {
		return Booru{
			BooruType: MoeBooru,
		}
	}
	return Booru{
		BooruType: DanBooru,
	}
}

type GetTagOption struct {
	Cache bool
}

func (b *Booru) GetTag(option GetTagOption) (*[]Tag, error) {
	var tag *[]Tag
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
