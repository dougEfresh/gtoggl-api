package gtwebsocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tumb1er/gtoggl-api/gttimentry"
	"log"
	"net/http"
)

const (
	DefaultUrl   = "wss://stream.toggl.com/ws'"
	ActionInsert = "INSERT"
	ActionUpdate = "UPDATE"
	ActionDelete = "DELETE"
)

func (mp MsgPing) OK() bool {
	return mp.Type == "ping"
}

func defaultOnPing(msg MsgPing) error {
	if !msg.OK() {
		return fmt.Errorf("ping: something wrong: %s", msg.Type)
	}
	log.Println("ping")
	return nil
}

func defaultOnTimeEntryAction(action string, entry gttimeentry.TimeEntry) error {
	log.Printf("TimeEntry action: %s %+v", action, entry)
	return nil
}

type TogglWebsocketClient struct {
	url               string
	token             string
	ws                *websocket.Conn
	onPing            OnPingCallback
	onTimeEntryAction OnTimeEntryActionCallback
}

func NewClient(token string) *TogglWebsocketClient {
	return &TogglWebsocketClient{
		url:               DefaultUrl,
		token:             token,
		onPing:            defaultOnPing,
		onTimeEntryAction: defaultOnTimeEntryAction,
	}
}

func (c *TogglWebsocketClient) OnPing(callback OnPingCallback) {
	c.onPing = callback
}

func (c *TogglWebsocketClient) OnTimeEntryAction(callback OnTimeEntryActionCallback) {
	c.onTimeEntryAction = callback
}

func (c *TogglWebsocketClient) Dial() error {
	log.Println("Dialing WS...")
	ws, resp, err := websocket.DefaultDialer.Dial(c.url, http.Header{
		"Origin": {"https://jw-toggl.com/app"},
	})
	if err != nil {
		log.Printf("Failed WS connect:%+v %+v", err, resp)
		return err
	}
	c.ws = ws
	log.Println("WS connected")
	return nil
}

func (c *TogglWebsocketClient) Listen(ctx context.Context) error {
	for {
		if err := c.Dial(); err != nil {
			return err
		}
		if err := c.authenticate(); err != nil {
			return err
		}

		disconnected := make(chan error)
		go func() {
			disconnected <- c.listen(ctx)
		}()

		select {
		case <-disconnected:
			continue
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *TogglWebsocketClient) listen(ctx context.Context) error {
	webSocketClosed := make(chan error)
	messages := make(chan []byte)
	go func() {
		defer close(webSocketClosed)
		for {
			if mt, data, err := c.ws.ReadMessage(); err != nil {
				log.Println("Read Websocket error:", err)
				webSocketClosed <- err
				if err = c.ws.Close(); err != nil {
					log.Println("Close Websocket error:", err)

				}
				return
			} else {
				if mt != websocket.TextMessage {
					log.Println("Unexpected message type:", mt)
				}
				messages <- data
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			if err := c.ws.Close(); err != nil {
				log.Println("Websocket close error:", err)
			}
			return nil
		case err := <-webSocketClosed:
			log.Println("Websocket closed")
			return err
		case msg := <-messages:
			if err := c.handleMessage(msg); err != nil {
				log.Println("Message handle error:", err)
				if err := c.ws.Close(); err != nil {
					log.Println("Websocket close error:", err)
				}
				return err
			}
		}
	}
}

func (c *TogglWebsocketClient) handleMessage(msg []byte) error {
	var guess map[string]interface{}
	if err := json.Unmarshal(msg, &guess); err != nil {
		return err
	}
	if _, ok := guess["type"]; ok {
		return c.handlePingMessage(msg)
	}
	if _, ok := guess["action"]; ok {
		return c.handleActionMessage(msg)
	}
	return errors.New("unknown message")
}

func (c *TogglWebsocketClient) handlePingMessage(msg []byte) error {
	var pingMsg MsgPing
	if err := json.Unmarshal(msg, &pingMsg); err != nil {
		return err
	}
	if err := c.onPing(pingMsg); err != nil {
		return err
	} else {
		pingMsg.Type = "pong"
		return c.ws.WriteJSON(pingMsg)
	}
}

func (c *TogglWebsocketClient) authenticate() error {
	type auth struct {
		Type     string `json:"type"`
		ApiToken string `json:"api_token"`
	}

	data := auth{
		Type:     "authenticate",
		ApiToken: c.token,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = c.ws.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		return err
	}

	var msg struct {
		Session string `json:"session_id"`
	}

	mt, bytes, err := c.ws.ReadMessage()
	if err != nil {
		return err
	}
	if mt != websocket.TextMessage {
		return fmt.Errorf("unexpected auth message type: %d", mt)
	}
	err = json.Unmarshal(bytes, &msg)
	if err != nil {
		return err
	}
	log.Printf("Auth: %+v\n", msg)
	if msg.Session == "" {
		return errors.New("no session after auth WS message")
	}
	return err
}

func (c *TogglWebsocketClient) handleActionMessage(msg []byte) error {
	var actionMsg MsgAction
	if err := json.Unmarshal(msg, &actionMsg); err != nil {
		return err
	}
	if actionMsg.Model == "time_entry" {
		var timeEntryActionMsg MsgTimeEntryAction
		if err := json.Unmarshal(msg, &timeEntryActionMsg); err != nil {
			return err
		}
		return c.onTimeEntryAction(actionMsg.Action, timeEntryActionMsg.Data)
	}
	return fmt.Errorf("unknown model action: %s", actionMsg.Model)
}
