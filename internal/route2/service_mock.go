package route2

import (
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m MockService) Find(slice []string, val string) bool {
	args := m.Called(slice, val)
	res, _ := args.Get(0).(bool)
	return res
}

func (m MockService) FetchData(t string) *Response {
	args := m.Called(t)
	res, _ := args.Get(0).(*Response)
	return res
}

func (m MockService) FilterBy(tag string) *Response {
	args := m.Called(tag)
	res, _ := args.Get(0).(*Response)
	return res
}

func (m MockService) CombineTags(tags1 *Response, tags2 *Response) *Response {
	args := m.Called(tags1, tags2)
	res, _ := args.Get(0).(*Response)
	return res
}

func (m MockService) AddShorterToLonger(shorter *Response, longer *Response) *Response {
	args := m.Called(shorter, longer)
	res, _ := args.Get(0).(*Response)
	return res
}

func (m MockService) SortBy(s string, d string, data *Response) *Response {
	args := m.Called(s, d, data)
	res, _ := args.Get(0).(*Response)
	return res
}
