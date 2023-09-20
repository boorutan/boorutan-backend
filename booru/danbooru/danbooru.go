package danbooru

import "applemango/boorutan/backend/booru"

func CreateDanBooru(base string) booru.Booru {
	return booru.Booru{
		Base: base,
		Url: booru.BooruUrl{
			Post: "/posts.json",
		},
		BooruType: booru.DanBooru,
	}
}
