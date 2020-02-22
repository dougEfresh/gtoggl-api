package gtproject

import (
	"github.com/tumb1er/gtoggl-api/test"
	"testing"
)

func togglClient(t *testing.T) *ProjectClient {
	tu := &gttest.TestUtil{}
	client := tu.MockClient(t)
	return NewClient(client)
}

func TestProjectCreate(t *testing.T) {
	tClient := togglClient(t)
	c := &Project{Name: "Very Big Company", WId: 777}
	nc, err := tClient.Create(c)
	if err != nil {
		t.Fatal(err)
	}

	if nc.Name != "An awesome project" {
		t.Fatal("!= An awesome project")
	}

	if nc.Id != 3 {
		t.Fatal("!= 3")
	}

	if nc.WId != 777 {
		t.Fatal("!= 777")
	}
}

func TestProjectUpdate(t *testing.T) {
	tClient := togglClient(t)
	c := &Project{Id: 1, Name: "new name", WId: 777}
	nc, err := tClient.Update(c)
	if err != nil {
		t.Fatal(err)
	}

	if nc.Name != "new name" {
		t.Fatal("!= new name")
	}
}

func TestProjectDelete(t *testing.T) {
	tClient := togglClient(t)
	c := &Project{Id: 1, Name: "new name", WId: 777}
	err := tClient.Delete(c.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProjectList(t *testing.T) {
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

func TestProjectGet(t *testing.T) {
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
