package moebooru

import "applemango/boorutan/backend/booru"

func CreateMoeBooru(base string) booru.Booru {
	return booru.Booru{
		Base: base,
		Url: booru.Url{
			Post:       "/post.json",
			Tag:        "/tag.json?order=count",
			TagSummary: "/tag/summary.json",
		},
		BooruType: booru.MoeBooru,
	}
}
