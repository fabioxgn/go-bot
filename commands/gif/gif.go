package gif

import (
	"fmt"
	"github.com/fabioxgn/go-bot"
	"github.com/fabioxgn/go-bot/web"
	"math/rand"
	"net/url"
	"time"
)

const (
	URL = "http://api.giphy.com/v1/gifs/search?q=%s&api_key=dc6zaTOxFJmzC&limit=50"
)

type Giphy struct {
	Data []struct {
		BitlyUrl string `json:"bitly_url"`
		Images   struct {
			FixedHeight struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height"`
			FixedHeightDownsampled struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_downsampled"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			FixedWidth struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width"`
			FixedWidthDownsampled struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			Original struct {
				Frames string `json:"frames"`
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original"`
		} `json:"images"`
		Type        string `json:"type"`
		Username    string `json:"username"`
		BitlyGifUrl string `json:"bitly_gif_url"`
		EmbedUrl    string `json:"embed_url"`
		Id          string `json:"id"`
		Rating      string `json:"rating"`
		Source      string `json:"source"`
		Url         string `json:"url"`
	} `json:"data"`
	Meta struct {
		Msg    string `json:"msg"`
		Status int64  `json:"status"`
	} `json:"meta"`
	Pagination struct {
		Count      int64 `json:"count"`
		Offset     int64 `json:"offset"`
		TotalCount int64 `json:"total_count"`
	} `json:"pagination"`
}

func gif(command *bot.Cmd) (msg string, err error) {
	data := &Giphy{}
	err = web.GetJSON(fmt.Sprintf(URL, url.QueryEscape(command.FullArg)), data)
	if err != nil {
		return "", err
	}

	if len(data.Data) == 0 {
		return ":cry:", nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Len(data.Data):", len(data.Data))
	index := r.Intn(len(data.Data))

	return fmt.Sprintf(data.Data[index].Images.FixedHeight.Url), nil
}

func init() {
	bot.RegisterCommand(&bot.CustomCommand{
		Cmd:         "gif",
		CmdFunc:     gif,
		Description: "Searchs and posts a random gif url from Giphy.",
		Usage:       "hell yes",
	})
}