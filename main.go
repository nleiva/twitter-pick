package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

var (
	token  = getenv("TWITTER_TOKEN")
	userID = getenv("TWITTER_USER")
	convID = getenv("TWITTER_CID")
	max    = 30
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Panicf("%s environment variable not set.", name)
	}
	return v
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func main() {
	client := &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.UserMentionTimelineOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAuthorID, twitter.TweetFieldConversationID, twitter.TweetFieldPublicMetrics, twitter.TweetFieldContextAnnotations},
		UserFields:  []twitter.UserField{twitter.UserFieldUserName},
		Expansions:  []twitter.Expansion{twitter.ExpansionAuthorID},
		MaxResults:  max,
	}

	timeline, err := client.UserMentionTimeline(context.Background(), userID, opts)
	if err != nil {
		log.Panicf("user mention timeline error: %v", err)
	}

	tweets := timeline.Raw.TweetDictionaries()
	id := 0

	for _, t := range tweets {
		if t.Tweet.ConversationID == convID && t.Tweet.PublicMetrics.Likes >= 1 {
			id++
			fmt.Printf("Raffle ID: %v\n", id)
			fmt.Printf("Text: %v\n", t.Tweet.Text)
			fmt.Printf("Author ID: %v\n", t.Tweet.AuthorID)
			fmt.Println("-------------------------")
		} 
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("\n\n -->  Winner: %v  !!!  <--- \n\n", random.Intn(id)+1)
}
