package client

import (
	"net/url"
	"testing"
)

var TestBaseUrls = [...]string{
	"https://apic.host.cisco",
	"https://apic.host.cisco/",
	"https://apic.host.cisco//",
	"https://apic.host.cisco///",
	"https://apic.host.cisco/test",
	"https://apic.host.cisco/test/",
	"https://apic.host.cisco/test//",
}

func AssertFullUrl(t *testing.T, baseUrl string, path string, expected string, preserveBaseUrl bool) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	aciClient := &Client{
		BaseURL:  url,
		preserveBaseUrlRef: preserveBaseUrl,
	}

	actual, _, err := aciClient.MakeFullUrl(path)
	if actual != expected || err != nil {
		t.Errorf(`MakeFullUrl("%s") = %q, %v, expected %#q`, path, actual, err, expected)
	}
}

func TestMakeFullUrl(t *testing.T) {
	expected := "https://apic.host.cisco/api/mo/uni/tn-test_tenant"
	paths := [...]string {
		"api/mo/uni/tn-test_tenant",
		"/api/mo/uni/tn-test_tenant",
		"///api/mo/uni/tn-test_tenant",
	}
	for _, baseUrl := range TestBaseUrls {
		for _, path := range paths {
			AssertFullUrl(t, baseUrl, path, expected, false)
		}
	}
}

func TestMakeFullUrl_PreserveBaseUrl(t *testing.T) {
	expected_path := "/api/mo/uni/tn-test_tenant"
	paths := [...]string {
		"api/mo/uni/tn-test_tenant",
		"/api/mo/uni/tn-test_tenant",
		"///api/mo/uni/tn-test_tenant",
	}
	for _, baseUrl := range TestBaseUrls {
		for _, path := range paths {
			expected := baseUrl + expected_path
			AssertFullUrl(t, baseUrl, path, expected, true)
		}
	}
}

func TestMakeFullUrl_QueryParams(t *testing.T) {
	expected := "https://apic.host.cisco/api/mo/uni/tn-test_tenant?rsp-subtree=full&rsp-prop-include=config-only"
	path := "api/mo/uni/tn-test_tenant?rsp-subtree=full&rsp-prop-include=config-only"
	for _, baseUrl := range TestBaseUrls {
		AssertFullUrl(t, baseUrl, path, expected, false)
	}
}