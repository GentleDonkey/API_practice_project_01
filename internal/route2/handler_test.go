package route2

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIHandler_posts(t *testing.T) {
	type fields struct {
		service *service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests = []struct {
		name                 string
		url				     string
		service              func() Service
		expectedResponseCode int
		expectedBody         *Response
	}{
		{
			name: "200-1",
			url: "http://localhost:8000/api/posts?tags=health,tech",
			service: func() Service {
				healthResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
					},
				}
				techResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
					},
				}
				mockClient := MockClient{}
				mockClient.On("GetPosts", "health").Return(healthResponse).Once()
				mockClient.On("GetPosts", "tech").Return(techResponse).Once()
				service := NewAPIService(mockClient)
				return service
			},
			expectedResponseCode: 200,
			expectedBody: &Response{
				Posts: []Post{
					{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
					{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
					{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
				},
			},
		},
		{
			name: "200-2",
			url: "http://localhost:8000/api/posts?tags=health,tech&sortBy=popularity&direction=desc",
			service: func() Service {
				healthResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
					},
				}
				techResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
					},
				}
				mockClient := MockClient{}
				mockClient.On("GetPosts", "health").Return(healthResponse).Once()
				mockClient.On("GetPosts", "tech").Return(techResponse).Once()
				service := NewAPIService(mockClient)
				return service
			},
			expectedResponseCode: 200,
			expectedBody: &Response{
				Posts: []Post{
					{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
					{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
					{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := tt.service()
			api := NewAPIHandler(service)
			r, _ := http.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()
			api.Posts(w, r)
			require.Equal(t, tt.expectedResponseCode, w.Code)
			if tt.expectedBody != nil {
				expectedJSON, err := json.Marshal(tt.expectedBody)
				require.NoError(t, err)
				require.Equal(t, string(expectedJSON), strings.Trim(string(w.Body.Bytes()), "\n"))
			}
		})
	}
	var errs = []struct {
		name                 string
		url				     string
		service              func() Service
		expectedResponseCode int
		expectedBody         *errResponse
	}{
		{
			name: "400-1",
			url:  "http://localhost:8000/api/posts?sortBy=id&direction=desc",
			service: func() Service {
				//response := &Response{
				//	Posts: []Post{
				//		{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
				//		{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
				//		{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
				//	},
				//}
				mockClient := MockClient{}
				mockClient.On("GetPosts", "").Return(nil).Once()
				service := NewAPIService(mockClient)
				return service
			},
			expectedResponseCode: 400,
			expectedBody: &errResponse{
				Error: "Tags parameter is required",
			},
		},
		{
			name: "400-2",
			url:  "http://localhost:8000/api/posts?tags=health,tech&sortBy=popularityyyyyyyyyyyy&direction=desc",
			service: func() Service {
				healthResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
					},
				}
				techResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
					},
				}
				mockClient := MockClient{}
				mockClient.On("GetPosts", "health").Return(healthResponse).Once()
				mockClient.On("GetPosts", "tech").Return(techResponse).Once()
				service := NewAPIService(mockClient)
				return service
			},
			expectedResponseCode: 400,
			expectedBody: &errResponse{
				Error: "SortBy parameter is invalid",
			},
		},
		{
			name: "400-3",
			url:  "http://localhost:8000/api/posts?tags=health,tech&sortBy=popularity&direction=desccccccccccccc",
			service: func() Service {
				healthResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{95, "Jon Abbott", 4, 985, 0.42, 55875, []string{"politics", "tech", "health", "history"}},
					},
				}
				techResponse := &Response{
					Posts:[]Post{
						{1, "Rylee Paul", 9, 960, 0.13, 50361, []string{"tech", "health"}},
						{18, "Jaden Bryant", 3, 983, 0.09, 33952, []string{"tech", "history"}},
					},
				}
				mockClient := MockClient{}
				mockClient.On("GetPosts", "health").Return(healthResponse).Once()
				mockClient.On("GetPosts", "tech").Return(techResponse).Once()
				service := NewAPIService(mockClient)
				return service
			},
			expectedResponseCode: 400,
			expectedBody: &errResponse{
				Error: "Direction parameter is invalid",
			},
		},
	}
	for _, ee := range errs {
		t.Run(ee.name, func(t *testing.T) {
			service := ee.service()
			api := NewAPIHandler(service)
			r, _ := http.NewRequest("GET", ee.url, nil)
			w := httptest.NewRecorder()
			api.Posts(w, r)
			require.Equal(t, ee.expectedResponseCode, w.Code)
			if ee.expectedBody != nil {
				expectedJSON, err := json.Marshal(ee.expectedBody)
				require.NoError(t, err)
				require.Equal(t, string(expectedJSON), strings.Trim(string(w.Body.Bytes()), "\n"))
			}
		})
	}
}

