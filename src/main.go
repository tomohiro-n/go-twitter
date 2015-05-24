package main
import (
	"./server"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

var QUERY_PARAMETERS url.Values = url.Values{"count": {"100"}}

var api *anaconda.TwitterApi

func main() {
	server.LoadConfig()
	server.ConnectDatabase()
	defer server.DatabaseSession.Close()
	c := server.Database.C(server.TWEETS_COLLECTION)

	anaconda.SetConsumerKey(server.ServerConfig.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(server.ServerConfig.Twitter.ConsumerSecret)
	api := anaconda.NewTwitterApi(server.ServerConfig.Twitter.AccessToken, server.ServerConfig.Twitter.AccessSecret)

	for _, term := range server.ServerConfig.SearchTerms {
		fmt.Println(term)
		searchResult, _ := api.GetSearch(term, QUERY_PARAMETERS)
		for _, tweet := range searchResult.Statuses {
			fmt.Println(tweet)
			c.Insert(tweet)
		}
	}
}