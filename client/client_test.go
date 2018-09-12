package client

import (
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
)

func TestClientAuthenticate(t *testing.T) {

	client := GetTestClient()
	err := client.Authenticate()
	if err != nil {
		t.Error(err)
	}

	if client.AuthToken.Token == "" {
		t.Error("Token is empty")
	}
}

func GetTestClient() *client.Client {
	return client.GetClient("https://192.168.10.102", "admin", client.Password("cisco123"), client.Insecure(true))

}
