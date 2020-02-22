package gtclient

import (
	"github.com/tumb1er/gtoggl-api/test"
	"testing"
)

func togglClient(t *testing.T) *TClient {
	tu := &gttest.TestUtil{}
	return NewClient(tu.MockClient(t))
}

func TestClientCreate(t *testing.T) {
	tClient := togglClient(t)
	c := &Client{Name: "Very Big Company", WId: 777}
	nc, err := tClient.Create(c)
	if err != nil {
		t.Fatal(err)
	}

	if nc.Name != "Very Big Company" {
		t.Fatal("!= Very Big Company")
	}

	if nc.Id != 1239455 {
		t.Fatal("!= 1239455")
	}

	if nc.WId != 777 {
		t.Fatal("!= 777")
	}
}

func TestClientUpdate(t *testing.T) {
	tClient := togglClient(t)
	c := &Client{Id: 1, Name: "new name", WId: 777}
	nc, err := tClient.Update(c)
	if err != nil {
		t.Fatal(err)
	}

	if nc.Name != "new name" {
		t.Fatal("!= new name")
	}
}

func TestClientDelete(t *testing.T) {
	tClient := togglClient(t)
	c := &Client{Id: 1, Name: "new name", WId: 777}
	err := tClient.Delete(c.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientList(t *testing.T) {
	tClient := togglClient(t)
	clients, err := tClient.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(clients) != 2 {
		t.Fatal("Workspace is not 2")
	}
	if clients[0].Id != 1 {
		t.Error("Workspace Id is not 1")
	}
	if clients[0].Name != "Id 1" {
		t.Error("Workspace name not Id ")
	}

	if clients[1].Id != 2 {
		t.Error("Workspace Id is not 2")
	}
	if clients[1].Name != "Id 2" {
		t.Error("Workspace name")
	}
}

func TestClientGet(t *testing.T) {
	tClient := togglClient(t)

	client, err := tClient.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if client.Id != 1 {
		t.Error("!= 1")
	}

	if client.Name != "Id 1" {
		t.Error("!= Id 1:  " + client.Name)
	}
}
