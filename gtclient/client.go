package gtclient

import (
	"encoding/json"
	"fmt"
	"github.com/dougEfresh/gtoggl-api/gthttp"
)

// Toggl Client Definition
type Client struct {
	Id       uint64 `json:"id"`
	WId      uint64 `json:"wid"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

type Clients []Client

const Endpoint = "/clients"

//Return a Toggl Client. An error is also returned when some configuration option is invalid
//    thc,err := gtoggl.NewClient("token")
//    tc,err := gclient.NewClient(tc)
func NewClient(thc *gthttp.TogglHttpClient) *TClient {
	tc := &TClient{
		thc: thc,
	}
	tc.endpoint = thc.Url + Endpoint
	return tc
}

type TClient struct {
	thc      *gthttp.TogglHttpClient
	endpoint string
}

func (tc *TClient) List() (Clients, error) {
	body, err := tc.thc.GetRequest(tc.endpoint)
	if err != nil {
		return nil, err
	}
	var clients Clients
	err = json.Unmarshal(*body, &clients)
	return clients, err
}

func (tc *TClient) Get(id uint64) (*Client, error) {
	return clientResponse(tc.thc.GetRequest(fmt.Sprintf("%s/%d", tc.endpoint, id)))
}

func (tc *TClient) Create(c *Client) (*Client, error) {
	body := map[string]interface{}{"client": c}
	return clientResponse(tc.thc.PostRequest(tc.endpoint, body))
}

func (tc *TClient) Update(c *Client) (*Client, error) {
	body := map[string]interface{}{"client": c}
	return clientResponse(tc.thc.PutRequest(fmt.Sprintf("%s/%d", tc.endpoint, c.Id), body))
}

func (tc *TClient) Delete(id uint64) error {
	_, err := tc.thc.DeleteRequest(fmt.Sprintf("%s/%d", tc.endpoint, id), nil)
	return err
}

func clientResponse(response *json.RawMessage, error error) (*Client, error) {
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
	var cl Client
	err = json.Unmarshal(*tResp.Data, &cl)
	if err != nil {
		return nil, err
	}
	return &cl, err
}
