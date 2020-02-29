package gtworkspace

import (
	"encoding/json"
	"fmt"
	"github.com/tumb1er/gtoggl-api/gthttp"
)

type Workspace struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Premium bool   `json:"premium"`
}

type Workspaces []Workspace

const Endpoint = "/workspaces"

//Return a Workspace Client. An error is also returned when some configuration option is invalid
//    tc,err := gtoggl.NewClient("token")
//    wsc,err := gtoggl.NewWorkspaceClient(tc)
func NewClient(tc *gthttp.TogglHttpClient) *WorkspaceClient {
	ws := &WorkspaceClient{
		tc: tc,
	}
	ws.endpoint = tc.Url + Endpoint
	return ws
}

type WorkspaceClient struct {
	tc       *gthttp.TogglHttpClient
	endpoint string
}

//GET https://www.toggl.com/api/v8/workspaces/123213
func (wc *WorkspaceClient) Get(id uint64) (*Workspace, error) {
	return workspaceResponse(wc.tc.GetRequest(fmt.Sprintf("%s/%d", wc.endpoint, id)))
}

//PUT https://www.toggl.com/api/v8/workspaces
func (wc *WorkspaceClient) Update(ws *Workspace) (*Workspace, error) {
	put := map[string]interface{}{"workspace": ws}
	body, err := json.Marshal(put)
	if err != nil {
		return nil, err
	}
	return workspaceResponse(wc.tc.PutRequest(fmt.Sprintf("%s/%d", wc.endpoint, ws.Id), body))
}

//GET https://www.toggl.com/api/v8/workspaces
func (wc *WorkspaceClient) List() (Workspaces, error) {
	body, err := wc.tc.GetRequest(wc.endpoint)
	var workspaces Workspaces
	if err != nil {
		return workspaces, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &workspaces)
	return workspaces, err
}

func workspaceResponse(response *json.RawMessage, error error) (*Workspace, error) {
	if error != nil {
		return nil, error
	}
	if response == nil {
		return nil, nil
	}
	var tResp gthttp.TogglResponse
	var ws Workspace
	err := json.Unmarshal(*response, &tResp)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*tResp.Data, &ws)
	if err != nil {
		return nil, err
	}
	return &ws, err
}
