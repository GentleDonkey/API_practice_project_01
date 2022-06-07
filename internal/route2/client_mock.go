package route2

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (m MockClient) GetPosts(tag string) *Response {
	args := m.Called(tag)
	res, _ := args.Get(0).(*Response)
	return res
}