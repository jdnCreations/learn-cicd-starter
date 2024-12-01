package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	req, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal("new req no worky")
	}
	req.Header.Set("Authorization", "ApiKey fdsa432lkjfds")
	want := "fdsa432lkjfds"

	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("error moment %v", err)
	}

	if want != apiKey {
		t.Errorf("got API key %q, want %q", apiKey, want)
	}
}

func TestInvalidAPIKey(t *testing.T) {
	req, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal("new req no worky")
	}
	req.Header.Set("Authorization", "ApiKey")
	want := ErrMalformedHeader

	_, err = GetAPIKey(req.Header)

	if err.Error() != want.Error() {
		t.Errorf("got error %q, want %q", err.Error(), want.Error())
	}
}

func TestNoAuthHeader(t *testing.T) {
	req, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal("new req no worky")
	}
	want := ErrNoAuthHeaderIncluded

	_, err = GetAPIKey(req.Header)
	if err.Error() != want.Error() {
		t.Fatalf("got error %v, expected %v", err.Error(), want.Error())
	}
}

func TestInvalidAuthPrefix(t *testing.T) {
	req, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal("new req no worky")
	}
	req.Header.Set("Authorization", "Bearer 412321as")

	want := ErrMalformedHeader

	_, err = GetAPIKey(req.Header)
	if err.Error() != "bob" {
		t.Fatalf("expected error %v, got %v", want, err)
	}
}
