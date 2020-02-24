package gtwebsocket

import "github.com/tumb1er/gtoggl-api/gttimentry"

type MsgPing struct {
	Type string `json:"type"`
}

type OnPingCallback func(ping MsgPing) error

type MsgAction struct {
	Action string `json:"action"`
	Model  string `json:"model"`
}

type MsgTimeEntryAction struct {
	MsgAction
	Data gttimeentry.TimeEntry `json:"data"`
}

type OnTimeEntryActionCallback func(action string, entry gttimeentry.TimeEntry) error
