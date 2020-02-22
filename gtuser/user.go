package gtuser

import (
	"encoding/json"
	"github.com/tumb1er/gtoggl-api/gtclient"
	"github.com/tumb1er/gtoggl-api/gthttp"
	"github.com/tumb1er/gtoggl-api/gtproject"
	"github.com/tumb1er/gtoggl-api/gtworkspace"
)

// Toggl User Definition
type User struct {
	Id         uint64                 `json:"id"`
	ApiToken   string                 `json:"api_token"`
	Email      string                 `json:"email"`
	FullName   string                 `json:"fullname"`
	Timezone   string                 `json:"timezone"`
	Clients    gtclient.Clients       `json:"clients"`
	Projects   gtproject.Projects     `json:"projects"`
	Workspaces gtworkspace.Workspaces `json:"workspaces"`
}
type UserUpdate struct {
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}
type UserCreate struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Timezone    string `json:"timezone"`
	CreatedWith string `json:"created_with"`
}

type Users []User

const Endpoint = "/me"
const SignupEndpoint = "/signups"
const ResetEndpoint = "/reset_token"
const MeWithRelatedData = "/me?with_related_data=true"

//Return a UserClient. An error is also returned when some configuration option is invalid
//    thc,err := gtoggl.NewClient("token")
//    uc,err := guser.NewClient(thc)
func NewClient(thc *gthttp.TogglHttpClient) *UserClient {
	tc := &UserClient{
		thc: thc,
	}
	tc.endpoint = thc.Url + Endpoint
	tc.signupEndpoint = thc.Url + SignupEndpoint
	tc.resetEndpoint = thc.Url + ResetEndpoint
	tc.relatedEndpoint = thc.Url + MeWithRelatedData
	return tc
}

type UserClient struct {
	thc             *gthttp.TogglHttpClient
	endpoint        string
	resetEndpoint   string
	signupEndpoint  string
	relatedEndpoint string
}

func (c *UserClient) Get(realatedData bool) (*User, error) {
	if realatedData {
		return userResponse(c.thc.GetRequest(c.relatedEndpoint))
	}
	return userResponse(c.thc.GetRequest(c.endpoint))
}

func (c *UserClient) Create(email, password, timezone string) (*User, error) {
	up := &UserCreate{Password: password, Email: email, Timezone: timezone, CreatedWith: "gtoggl"}
	put := map[string]interface{}{"user": up}
	return userResponse(c.thc.PostRequest(c.signupEndpoint, put))
}

func (c *UserClient) Update(u *User) (*User, error) {
	up := &UserUpdate{FullName: u.FullName, Email: u.Email}
	put := map[string]interface{}{"user": up}
	return userResponse(c.thc.PutRequest(c.endpoint, put))
}

func (c *UserClient) ResetToken() (string, error) {
	response, err := c.thc.PostRequest(c.resetEndpoint, nil)
	if err != nil {
		return "", err
	}
	var aux string
	err = json.Unmarshal(*response, &aux)
	return aux, nil

}

func userResponse(response *json.RawMessage, error error) (*User, error) {
	if error != nil {
		return nil, error
	}
	var tResp gthttp.TogglResponse
	err := json.Unmarshal(*response, &tResp)
	if err != nil {
		return nil, err
	}
	var u User
	err = json.Unmarshal(*tResp.Data, &u)
	if err != nil {
		return nil, err
	}
	return &u, err
}
