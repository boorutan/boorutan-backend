package moebooru

import "applemango/boorutan/backend/booru"

func CreateMoeBooru(base string) booru.Booru {
	return booru.Booru{
		Base: base,
		Url: booru.BooruUrl{
			Post: "/post.json",
			Tag:  "/tag.json?order=count",
		},
		BooruType: booru.MoeBooru,
	}
}
