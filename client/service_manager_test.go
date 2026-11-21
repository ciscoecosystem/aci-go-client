package client

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	emptyListBody = `{"totalCount":"0","imdata":[]}`
	oneSubnetBody = `{"totalCount":"1","imdata":[{"l3extSubnet":{"attributes":{"ip":"10.0.0.1/32","dn":"uni/tn-t/out-o/instP-p/extsubnet-[10.0.0.1/32]"}}}]}`
	forbiddenBody = `{"totalCount":"1","imdata":[{"error":{"attributes":{"code":"403","text":"Unauthorized request"}}}]}`
)

func testClient(t *testing.T, body string) *Client {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	baseURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	aciClient := &Client{
		BaseURL:          baseURL,
		username:         "test",
		password:         "test",
		MOURL:            DefaultMOURL,
		httpClient:       server.Client(),
		AuthToken:        &Auth{Token: "test-token", Expiry: time.Now().Add(time.Hour)},
		maxRetries:       0,
		maxReAuthRetries: 0,
	}
	aciClient.ServiceManager = NewServiceManager(aciClient.MOURL, aciClient)

	return aciClient
}

// An empty class query is a valid answer: list helpers must not fail on it.
func TestListAllowEmpty(t *testing.T) {
	aciClient := testClient(t, emptyListBody)

	subnets, err := aciClient.ListL3L3ExtSubnet("epg", "l3out", "tenant")
	if err != nil {
		t.Fatalf("expected no error for an empty list, got: %v", err)
	}
	if len(subnets) != 0 {
		t.Fatalf("expected an empty list, got %d entries", len(subnets))
	}
}

// A point read keeps its current behaviour: an empty response means the object
// does not exist.
func TestGetViaURLStillErrorsOnEmpty(t *testing.T) {
	aciClient := testClient(t, emptyListBody)

	_, err := aciClient.GetViaURL("/api/mo/uni/tn-tenant/out-l3out/instP-epg.json")
	if err == nil {
		t.Fatal("expected an error for an empty point read, got nil")
	}
	if !strings.Contains(err.Error(), "Object may not exist") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A permission failure must still be reported, not swallowed as an empty list.
func TestListStillErrorsOnPermissionFailure(t *testing.T) {
	aciClient := testClient(t, forbiddenBody)

	_, err := aciClient.ListL3L3ExtSubnet("epg", "l3out", "tenant")
	if err == nil {
		t.Fatal("expected an error for a 403 response, got nil")
	}
	if !strings.Contains(err.Error(), "Unauthorized request") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestListParsesObjects(t *testing.T) {
	aciClient := testClient(t, oneSubnetBody)

	subnets, err := aciClient.ListL3L3ExtSubnet("epg", "l3out", "tenant")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(subnets) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(subnets))
	}
	if subnets[0].Ip != "10.0.0.1/32" {
		t.Fatalf("unexpected ip: %q", subnets[0].Ip)
	}
}

func TestCheckForErrorsAllowEmpty(t *testing.T) {
	empty, err := container.ParseJSON([]byte(emptyListBody))
	if err != nil {
		t.Fatal(err)
	}

	if err := CheckForErrors(empty, "GET", true); err == nil {
		t.Fatal("CheckForErrors must keep failing on an empty response")
	}
	if err := checkForErrors(empty, "GET", true, true); err != nil {
		t.Fatalf("checkForErrors must accept an empty response when allowEmpty is set, got: %v", err)
	}

	forbidden, err := container.ParseJSON([]byte(forbiddenBody))
	if err != nil {
		t.Fatal(err)
	}
	if err := checkForErrors(forbidden, "GET", true, true); err == nil {
		t.Fatal("checkForErrors must still report an APIC error when allowEmpty is set")
	}
}
