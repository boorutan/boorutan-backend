package booru

import (
	http "applemango/boorutan/backend/utils"
	"fmt"
)

func (b *Booru) GetPost() (*[]BooruPost, error) {
	var post *[]BooruPost
	err := http.RequestJSON(
		&post,
		fmt.Sprintf("%v%v", b.Url.Base, b.Url.Post),
		"GET",
		nil,
	)
	return post, err
}
