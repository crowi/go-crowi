package crowi

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

const (
	baseURL = "http://localhost:3000"
	token   = "abcdefghijklmnopqrstuvwxyz0123456789="
)

func TestNewClient_url(t *testing.T) {
	client, err := NewClient(baseURL, token)
	if err != nil {
		t.Fatal(err)
	}

	if client.URL.String() != baseURL {
		t.Fatalf("expected %q to be %q", client.URL.String(), baseURL)
	}
}

func TestNewClient_badURL(t *testing.T) {
	_, err := NewClient("", token)
	if err == nil {
		t.Fatal("expected error, but nothing was returned")
	}

	expected := "missing api url"
	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected %q to contain %q", err.Error(), expected)
	}
}

func TestNewClient_badToken(t *testing.T) {
	_, err := NewClient(baseURL, "")
	if err == nil {
		t.Fatal("expected error, but nothing was returned")
	}

	expected := "missing token"
	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected %q to contain %q", err.Error(), expected)
	}
}

func TestNewClient_parsesURL(t *testing.T) {
	client, err := NewClient(baseURL, token)
	if err != nil {
		t.Fatal(err)
	}

	expected := &url.URL{
		Scheme: "http",
		Host:   "localhost:3000",
		Path:   "",
	}
	if !reflect.DeepEqual(client.URL, expected) {
		t.Fatalf("expected %q to equal %q", client.URL, expected)
	}
}
