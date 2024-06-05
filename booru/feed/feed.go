package feed

import (
	"applemango/boorutan/backend/booru"
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/handler"
	"applemango/boorutan/backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func generateFeedCore(b booru.Booru, page int, f func(p booru.Post)) error {
	posts, err := b.GetPosts(booru.GetPostsOption{
		Page:  page,
		Tags:  nil,
		Cache: false,
	})
	if err != nil {
		return err
	}
	now := time.Now()
	lastUpdate := time.Now()
	for _, p := range *posts {
		if &p == nil {
			return errors.New("")
		}
		updateAt, err := utils.ParseDanbooruTime(p.UpdatedAt)
		if err != nil {
			return err
		}
		diff := updateAt.Sub(now)
		if diff.Minutes() > 2 {
			return nil
		}
		lastUpdate = updateAt
		cacheKey := fmt.Sprintf("feed_cache:post:%v:%v", b.Base, p.ID)
		cache, err := redis.Get(cacheKey)
		if err == nil || cache != "" {
			return nil
		}
		v, err := p.ToString()
		if err != nil {
			return err
		}
		err = redis.Push(cacheKey, v)
		if err != nil {
			return err
		}
		f(p)
	}
	if lastUpdate.Sub(now).Minutes() > 2 && page < 10 {
		return generateFeedCore(b, page+1, f)
	}
	return nil
}

type DiscordWebhookEmbedsFooter struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}
type DiscordWebhookEmbedsImage struct {
	Url string `json:"url,omitempty"`
}
type DiscordWebhookEmbeds struct {
	Color       int                        `json:"color,omitempty"`
	Title       string                     `json:"title,omitempty"`
	Description string                     `json:"description,omitempty"`
	Image       DiscordWebhookEmbedsImage  `json:"image,omitempty"`
	Timestamp   int                        `json:"timestamp"`
	Footer      DiscordWebhookEmbedsFooter `json:"footer"`
}
type DiscordWebhookEmbedsMessage struct {
	Embeds []DiscordWebhookEmbeds `json:"embeds"`
}

func SendWebhook(post booru.Post) error {
	webhook := "https://discord.com/api/webhooks/1237821446097735710/YGTizIU1mht3xjonySY_clOrIwu2jxF_Wnoonsvc5TPZ7WYc6XPyV84RaK9lkzw29Sju"
	embeds := DiscordWebhookEmbedsMessage{
		Embeds: []DiscordWebhookEmbeds{
			{
				Color:       45300,
				Title:       fmt.Sprintf("%v ( %v )", post.TagStringCharacter, post.TagStringArtist),
				Description: fmt.Sprintf("https://booru.i32.jp/danbooru/%v", post.ID),
				Image: DiscordWebhookEmbedsImage{
					Url: fmt.Sprintf("https://api-booru.i32.jp/image?url=%v", post.FileURL),
				},
				//Timestamp: int(time.Now().UnixMilli()),
				Footer: DiscordWebhookEmbedsFooter{
					Text:    "Danbooru",
					IconUrl: "https://api-booru.i32.jp/image?url=https://cdn.discordapp.com/emojis/1209826715631624243.webp?size=96&quality=lossless",
				},
			},
		},
	}
	str, err := json.Marshal(embeds)
	if err != nil {
		panic(err.Error())
	}
	reader := strings.NewReader(string(str))
	fmt.Println(string(str))
	res, err := http.Post(webhook, "application/json", reader)
	bodyBytes, err := io.ReadAll(res.Body)
	fmt.Println(string(bodyBytes), res.StatusCode)
	return err
}

func GenerateFeed() {
	//_ = SendWebhook("https://cdn.donmai.us/original/2c/8c/2c8cbf3acb2c9f03321c93b848ee8969.jpg")
	b := handler.GetBooruFromString("danbooru")
	err := generateFeedCore(b, 1, func(p booru.Post) {
		err := SendWebhook(p)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("push")
	})
	if err != nil {
		panic(err.Error())
	}
}
