package gttag

import (
	"encoding/json"
	"fmt"
	"github.com/tumb1er/gtoggl-api/gthttp"
)

// Toggl Tag Definition
type Tag struct {
	Id   uint64 `json:"id"`
	WId  uint64 `json:"wid"`
	Name string `json:"name"`
}

type Tags []Tag

const Endpoint = "/tags"

//Return a TagClient. An error is also returned when some configuration option is invalid
//    thc,err := gtoggl.NewClient("token")
//    tc,err := gtag.NewClient(thc)
func NewClient(thc *gthttp.TogglHttpClient) *TagClient {
	tc := &TagClient{
		thc: thc,
	}
	tc.endpoint = thc.Url + Endpoint
	return tc
}

type TagClient struct {
	thc      *gthttp.TogglHttpClient
	endpoint string
}

func (tc *TagClient) List() (Tags, error) {
	body, err := tc.thc.GetRequest(tc.endpoint)
	var tags []Tag
	if err != nil {
		return nil, err
	}
	if body == nil {
		return tags, nil
	}
	err = json.Unmarshal(*body, &tags)
	return tags, err
}

func (tc *TagClient) Get(id uint64) (*Tag, error) {
	return tagResponse(tc.thc.GetRequest(fmt.Sprintf("%s/%d", tc.endpoint, id)))
}

func (tc *TagClient) Create(t *Tag) (*Tag, error) {
	put := map[string]interface{}{"tag": t}
	return tagResponse(tc.thc.PostRequest(tc.endpoint, put))
}

func (tc *TagClient) Update(t *Tag) (*Tag, error) {
	put := map[string]interface{}{"tag": t}
	return tagResponse(tc.thc.PutRequest(fmt.Sprintf("%s/%d", tc.endpoint, t.Id), put))
}

func (tc *TagClient) Delete(id uint64) error {
	_, err := tc.thc.DeleteRequest(fmt.Sprintf("%s/%d", tc.endpoint, id), nil)
	return err
}

func tagResponse(response *json.RawMessage, error error) (*Tag, error) {
	if error != nil {
		return nil, error
	}
	if response == nil {
		return nil, nil
	}
	var tResp gthttp.TogglResponse
	err := json.Unmarshal(*response, &tResp)
	if err != nil {
		return nil, err
	}
	var t Tag
	err = json.Unmarshal(*tResp.Data, &t)
	if err != nil {
		return nil, err
	}
	return &t, err
}
