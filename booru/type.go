package booru

import "applemango/boorutan/backend/utils/image"

type Type int

const (
	DanBooru Type = iota
	MoeBooru
)

type Url struct {
	Post       string
	Tag        string
	TagSummary any
}

type Booru struct {
	Base      string
	Url       Url
	BooruType Type
}

type Post struct {
	ID                  int           `json:"id"`
	Tags                string        `json:"tags"`
	CreatedAt           int           `json:"created_at"`
	CreatorID           int           `json:"creator_id"`
	Author              string        `json:"author"`
	Change              int           `json:"change"`
	Source              string        `json:"source"`
	Score               int           `json:"score"`
	Md5                 string        `json:"md5"`
	FileSize            int           `json:"file_size"`
	FileURL             string        `json:"file_url"`
	IsShownInIndex      bool          `json:"is_shown_in_index"`
	PreviewURL          string        `json:"preview_url"`
	PreviewWidth        int           `json:"preview_width"`
	PreviewHeight       int           `json:"preview_height"`
	ActualPreviewWidth  int           `json:"actual_preview_width"`
	ActualPreviewHeight int           `json:"actual_preview_height"`
	SampleURL           string        `json:"sample_url"`
	SampleWidth         int           `json:"sample_width"`
	SampleHeight        int           `json:"sample_height"`
	SampleFileSize      int           `json:"sample_file_size"`
	JpegURL             string        `json:"jpeg_url"`
	JpegWidth           int           `json:"jpeg_width"`
	JpegHeight          int           `json:"jpeg_height"`
	JpegFileSize        int           `json:"jpeg_file_size"`
	Rating              string        `json:"rating"`
	HasChildren         bool          `json:"has_children"`
	ParentID            interface{}   `json:"parent_id"`
	Status              string        `json:"status"`
	Width               int           `json:"width"`
	Height              int           `json:"height"`
	IsHeld              bool          `json:"is_held"`
	FramesPendingString string        `json:"frames_pending_string"`
	FramesPending       []interface{} `json:"frames_pending"`
	FramesString        string        `json:"frames_string"`
	Frames              []interface{} `json:"frames"`

	UploaderID          int         `json:"uploader_id"`
	LastCommentBumpedAt interface{} `json:"last_comment_bumped_at"`
	ImageWidth          int         `json:"image_width"`
	ImageHeight         int         `json:"image_height"`
	TagString           string      `json:"tag_string"`
	FavCount            int         `json:"fav_count"`
	FileExt             string      `json:"file_ext"`
	LastNotedAt         interface{} `json:"last_noted_at"`
	ApproverID          interface{} `json:"approver_id"`
	TagCountGeneral     int         `json:"tag_count_general"`
	TagCountArtist      int         `json:"tag_count_artist"`
	TagCountCharacter   int         `json:"tag_count_character"`
	TagCountCopyright   int         `json:"tag_count_copyright"`
	UpScore             int         `json:"up_score"`
	DownScore           int         `json:"down_score"`
	IsPending           bool        `json:"is_pending"`
	IsFlagged           bool        `json:"is_flagged"`
	IsDeleted           bool        `json:"is_deleted"`
	TagCount            int         `json:"tag_count"`
	UpdatedAt           string      `json:"updated_at"`
	IsBanned            bool        `json:"is_banned"`
	PixivID             int         `json:"pixiv_id"`
	LastCommentedAt     interface{} `json:"last_commented_at"`
	HasActiveChildren   bool        `json:"has_active_children"`
	BitFlags            int         `json:"bit_flags"`
	TagCountMeta        int         `json:"tag_count_meta"`
	HasLarge            bool        `json:"has_large"`
	HasVisibleChildren  bool        `json:"has_visible_children"`
	MediaAsset          struct {
		ID          int         `json:"id"`
		CreatedAt   string      `json:"created_at"`
		UpdatedAt   string      `json:"updated_at"`
		Md5         string      `json:"md5"`
		FileExt     string      `json:"file_ext"`
		FileSize    int         `json:"file_size"`
		ImageWidth  int         `json:"image_width"`
		ImageHeight int         `json:"image_height"`
		Duration    interface{} `json:"duration"`
		Status      string      `json:"status"`
		FileKey     string      `json:"file_key"`
		IsPublic    bool        `json:"is_public"`
		PixelHash   string      `json:"pixel_hash"`
		Variants    []struct {
			Type    string `json:"type"`
			URL     string `json:"url"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			FileExt string `json:"file_ext"`
		} `json:"variants"`
	} `json:"media_asset"`
	TagStringGeneral   string `json:"tag_string_general"`
	TagStringCharacter string `json:"tag_string_character"`
	TagStringCopyright string `json:"tag_string_copyright"`
	TagStringArtist    string `json:"tag_string_artist"`
	TagStringMeta      string `json:"tag_string_meta"`
	LargeFileURL       string `json:"large_file_url"`
	PreviewFileURL     string `json:"preview_file_url"`
	BooruType          string `json:"booru_type"`

	Summary []image.Color `json:"summary"`
}

type Tag struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	Type      int    `json:"type"`
	Ambiguous bool   `json:"ambiguous"`

	PostCount    int      `json:"post_count"`
	Category     int      `json:"category"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	IsDeprecated bool     `json:"is_deprecated"`
	Words        []string `json:"words"`
}
