package gtproject

import (
	"encoding/json"
	"fmt"
	"github.com/dougEfresh/gtoggl-api/gthttp"
)

// Toggl Project Definition
type Project struct {
	Id        uint64 `json:"id"`
	WId       uint64 `json:"wid"`
	CId       uint64 `json:"cid"`
	Name      string `json:"name"`
	IsPrivate *bool  `json:"is_private,omitempty"`
}

type Projects []Project

const Endpoint = "/projects"

//Return a ProjectClient. An error is also returned when some configuration option is invalid
//    thc,err := gtoggl.NewClient("token")
//    pc,err := gproject.NewClient(tc)
func NewClient(thc *gthttp.TogglHttpClient) *ProjectClient {
	tc := &ProjectClient{
		thc: thc,
	}
	tc.endpoint = thc.Url + Endpoint
	return tc
}

type ProjectClient struct {
	thc      *gthttp.TogglHttpClient
	endpoint string
}

func (tc *ProjectClient) List() (Projects, error) {
	body, err := tc.thc.GetRequest(tc.endpoint)
	var projects []Project
	if err != nil {
		return nil, err
	}
	if body == nil {
		return projects, nil
	}
	err = json.Unmarshal(*body, &projects)
	return projects, err
}

func (tc *ProjectClient) Get(id uint64) (*Project, error) {
	return projectResponse(tc.thc.GetRequest(fmt.Sprintf("%s/%d", tc.endpoint, id)))
}

func (tc *ProjectClient) Create(p *Project) (*Project, error) {
	put := map[string]interface{}{"project": p}
	return projectResponse(tc.thc.PostRequest(tc.endpoint, put))
}

func (tc *ProjectClient) Update(p *Project) (*Project, error) {
	put := map[string]interface{}{"project": p}
	return projectResponse(tc.thc.PutRequest(fmt.Sprintf("%s/%d", tc.endpoint, p.Id), put))
}

func (tc *ProjectClient) Delete(id uint64) error {
	_, err := tc.thc.DeleteRequest(fmt.Sprintf("%s/%d", tc.endpoint, id), nil)
	return err
}

func projectResponse(response *json.RawMessage, error error) (*Project, error) {
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
	var p Project
	err = json.Unmarshal(*tResp.Data, &p)
	if err != nil {
		return nil, err
	}
	return &p, err
}
