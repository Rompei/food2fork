package food2fork

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BaseURL is url of the API.
const BaseURL = "http://food2fork.com/api"

// Food2Fork is api structure of Food2fork.
type Food2Fork struct {
	key        string
	queryQueue chan query
	baseURL    string
}

type query struct {
	url        string
	params     url.Values
	data       interface{}
	responseCh chan response
}

type response struct {
	data interface{}
	err  error
}

// New is initializer.
func New(apiKey string) *Food2Fork {
	queue := make(chan query)
	api := &Food2Fork{
		key:        apiKey,
		queryQueue: queue,
		baseURL:    BaseURL,
	}

	go api.requestRoop()
	return api
}

func (f *Food2Fork) sendGetRequest(url string, params url.Values, data interface{}) error {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode == 204 {
		fmt.Println("err204")
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("err")
		// TODO
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(data)
}

func (f *Food2Fork) cleanValue(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	v.Add("key", f.key)
	return v
}

func (f *Food2Fork) requestRoop() {
	for q := range f.queryQueue {
		url := q.url
		params := q.params
		data := q.data
		responseCh := q.responseCh
		err := f.sendGetRequest(url, params, data)

		if err != nil {
			fmt.Println(err)
			continue
		}

		responseCh <- response{data, err}
	}
}

// Close is close queue
func (f *Food2Fork) Close() {
	close(f.queryQueue)
}
