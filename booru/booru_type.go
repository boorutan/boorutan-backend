package booru

type BooruType int

const (
	DanBooru BooruType = iota
	MoeBooru
)

type BooruUrl struct {
	base string
	post string
}

type Booru struct {
	url       BooruUrl
	booruType BooruType
}
