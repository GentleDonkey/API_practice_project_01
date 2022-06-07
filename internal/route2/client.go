package route2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Client interface{
	GetPosts(string) *Response
}

type client struct{}

func NewAPIClient() *client {
	return &client{}
}

func (c *client) GetPosts(tag string) *Response {
	url := "https://api.hatchways.io/assessment/blog/posts"+"?tag="+tag
	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	body, err2 := ioutil.ReadAll(data.Body)
	if err2 != nil {
		log.Fatal(err2)
		return nil
	}
	response := Response{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return nil
	}
	return &response
}