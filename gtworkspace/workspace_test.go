package gtworkspace

import (
	"github.com/dougEfresh/gtoggl-api/test"
	"testing"
)

func workspaceClient(t *testing.T) *WorkspaceClient {
	tu := &gttest.TestUtil{}
	client := tu.MockClient(t)
	return NewClient(client)
}

func TestWorkspaceList(t *testing.T) {
	workspaceClient := workspaceClient(t)
	workspaces, err := workspaceClient.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(workspaces) != 2 {
		t.Fatal("Workspace is not 2")
	}
	if workspaces[0].Id != 1 {
		t.Error("Workspace Id is not 1")
	}
	if workspaces[0].Name != "Id 1" {
		t.Error("Workspace name not Id ")
	}
	if !workspaces[0].Premium {
		t.Error("Workspace is not premium ")
	}

	if workspaces[1].Id != 2 {
		t.Error("Workspace Id is not 2")
	}
	if workspaces[1].Name != "Id 2" {
		t.Error("Workspace name")
	}
	if workspaces[1].Premium {
		t.Error("Workspace is not premium ")
	}

}

func TestWorkspaceGet(t *testing.T) {
	workspaceClient := workspaceClient(t)

	workspace, err := workspaceClient.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if workspace.Id != 1 {
		t.Error("Workspace id != 1")
	}
}

func TestWorkspaceUpdate(t *testing.T) {
	workspaceClient := workspaceClient(t)
	workspace, err := workspaceClient.Get(1)

	if err != nil {
		t.Error(err)
	}
	workspace.Name = "new name"
	workspace, err = workspaceClient.Update(workspace)

	if err != nil {
		t.Error(err)
	}

	if workspace.Name != "new name" {
		t.Error("Wrong name")
	}

}
