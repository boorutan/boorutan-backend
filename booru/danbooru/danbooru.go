package danbooru

import "applemango/boorutan/backend/booru"

func CreateDanBooru(base string) booru.Booru {
	return booru.Booru{
		Base: base,
		Url: booru.BooruUrl{
			Post: "/posts.json",
			Tag:  "/tags.json?search[order]=count",
		},
		BooruType: booru.DanBooru,
	}
}
