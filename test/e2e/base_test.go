package e2e_test

import (
	"testing"

	grclient "github.com/lfgrillo83/groselha/client"
	grserver "github.com/lfgrillo83/groselha/server"
)

func TestBase(t *testing.T) {
	tmpPort := availablePort()
	//setup server, client, fileserver
	teardown := simpleSetup(t,
		&grserver.Config{},
		&grclient.Config{
			Remotes: []string{tmpPort + ":$FILEPORT"},
		})
	defer teardown()
	//test remote
	result, err := post("http://localhost:"+tmpPort, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if result != "foo!" {
		t.Fatalf("expected exclamation mark added")
	}
}

func TestReverse(t *testing.T) {
	tmpPort := availablePort()
	//setup server, client, fileserver
	teardown := simpleSetup(t,
		&grserver.Config{
			Reverse: true,
		},
		&grclient.Config{
			Remotes: []string{"R:" + tmpPort + ":$FILEPORT"},
		})
	defer teardown()
	//test remote (this goes through the server and out the client)
	result, err := post("http://localhost:"+tmpPort, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if result != "foo!" {
		t.Fatalf("expected exclamation mark added")
	}
}
