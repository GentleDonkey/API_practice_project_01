package route2

type Query struct {
	Tags string `json:"tags1"`
	SortBy string `json:"sortBy"`
	Direction string `json:"direction"`
}

type errResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Id         int      `json:"id"`
	Author     string   `json:"author"`
	AuthorId   int      `json:"authorId"`
	Likes      int      `json:"likes"`
	Popularity float64  `json:"popularity"`
	Reads      int      `json:"reads"`
	Tags       []string `json:"tags"`
}



