package main

import (
	"fmt"
	"log"
	"strings"

	YtChat "github.com/ultraderek/youtube-live-chat-downloader"
)

func main() {
	/*customCookies := []*http.Cookie{
		{Name: "PREF",
			Value:  "tz=Europe.Rome",
			MaxAge: 300},
		{Name: "CONSENT",
			Value:  fmt.Sprintf("YES+yt.432048971.it+FX+%d", 100+rand.Intn(999-100+1)),
			MaxAge: 300},
	}
	// Google would sometimes ask you to solve a CAPTCHA before accessing it's websites.
	// or ask for your CONSENT if you are an EU user
	// You can add those cookies here.
	// Adding cookies is OPTIONAL
	YtChat.AddCookies(customCookies)
	*/
	continuation, cfg, err := YtChat.ParseInitialData("https://www.youtube.com/watch?v=bOPCzAvBkQg")
	if err != nil {
		log.Fatal(err)
	}
	for {
		chat, newContinuation, err := YtChat.FetchContinuationChat(continuation, cfg)
		if err == YtChat.ErrLiveStreamOver {
			log.Fatal("Live stream over")
		}
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "400") {
				return
			}
			continue
		}
		// set the newly received continuation
		continuation = newContinuation

		for _, msg := range chat {
			fmt.Print(msg.Timestamp, " | ")
			fmt.Println(msg.AuthorName, ": ", msg.Message)
		}
	}
}
