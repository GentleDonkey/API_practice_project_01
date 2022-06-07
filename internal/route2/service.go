package route2

import (
	"sort"
	"strings"
)

type Service interface {
	Find(slice []string, val string) bool
	FetchData(t string) *Response
	FilterBy(tag string) *Response
	CombineTags(tags1 *Response, tags2 *Response) *Response
	AddShorterToLonger (shorter *Response, longer *Response) *Response
	SortBy(s string, d string, data *Response) *Response
}

type service struct {
	client Client
}

func NewAPIService(c Client) *service {
	return &service{
		c,
	}
}

func (sr *service) Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (sr *service) FetchData(t string) *Response {
	var result *Response
	// check if there are more than one tag,
	// if so, filter by different tags and combine&de-duplicate
	if strings.Contains(t, ",") {
		array := strings.Split(t, ",")
		for _, s:= range array {
			data := sr.FilterBy(s)
			result = sr.CombineTags(result, data)
		}
	} else {
		result = sr.FilterBy(t)
	}
	return result
}

func (sr *service) FilterBy(tag string) *Response {
	return sr.client.GetPosts(tag)
}

func (sr *service) CombineTags(tags1 *Response, tags2 *Response) *Response {
	// tags1 is empty when run this func first time,
	// if so, return tags2 directly
	if tags1 == nil {
		return tags2
	}
	// otherwise, add the smaller one to the bigger one
	response := Response{}
	if len(tags1.Posts) > len(tags2.Posts) {
		response = *sr.AddShorterToLonger(tags2, tags1)
	} else {
		response = *sr.AddShorterToLonger(tags1, tags2)
	}
	return &response
}

func (sr *service) AddShorterToLonger (shorter *Response, longer *Response) *Response {
	// init a map to record Post.ID of longer as key
	list := make(map[int]Post)
	for _, item := range longer.Posts{
		list[item.Id] = item
	}
	for _, item := range shorter.Posts {
		// check if id of item in shorter is in the map,
		// if not, add the item to longer
		if _, ok := list[item.Id]; !ok {
			longer.Posts = append(longer.Posts, item)
		}
	}
	return longer
}

func (sr *service) SortBy(s string, d string, data *Response) *Response {
	response := Response{}
	list := make(map[int]Post)
	// based on different sort method, add different key in the map
	for _, item := range data.Posts{
		if s == "id" {
			list[item.Id] = item
		} else if s == "reads" {
			list[item.Reads] = item
		} else if s == "likes" {
			list[item.Likes] = item
		} else if s == "popularity" {
			// popularity is float type, here I multiply with 100, and add it as the key in the map
			list[int(item.Popularity*100)] = item
		}
	}
	// add keys in a new map
	keys := make([]int, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	// based on different sort method, sort keys in particular direction
	if d == "asc" {
		sort.Ints(keys)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	}
	// add the item to response
	for _, k := range keys {
		response.Posts = append(response.Posts, list[k])
	}
	return &response
}