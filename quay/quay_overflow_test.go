package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"k8s.io/client-go/rest"
)

type MockHTTPClient struct {
	wantErr  bool
	mutateFn func(res *http.Response)
}

func (m MockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if m.wantErr {
		return nil, errors.New("oops")
	}

	resp := &http.Response{}
	if m.mutateFn != nil {
		m.mutateFn(resp)
	}

	return resp, nil
}

var _ rest.HTTPClient = &MockHTTPClient{}

func Test_fetchTags(t *testing.T) {
	t.Run("test error making request", func(t *testing.T) {
		tags, err := fetchTags(&MockHTTPClient{wantErr: true})

		if err == nil {
			t.Error("error expected")
		}

		if err.Error() != "error making request: oops" {
			t.Errorf("error expected, got %s", err.Error())
		}

		if tags != nil {
			t.Error("expected nil tags")
		}
	})

	t.Run("test error for non-200 status codes", func(t *testing.T) {
		tags, err := fetchTags(&MockHTTPClient{mutateFn: func(res *http.Response) {
			res.Status = string(rune(400))
			res.Body = io.NopCloser(bytes.NewReader(nil))
		}})

		if err == nil {
			t.Error("error expected")
		}

		if strings.Contains(err.Error(), "tags, error: received status code 400") {
			t.Errorf("error expected, got %s", err.Error())
		}

		if tags != nil {
			t.Error("expected nil tags")
		}
	})

	t.Run("test error parsing json", func(t *testing.T) {
		tags, err := fetchTags(&MockHTTPClient{mutateFn: func(res *http.Response) {
			res.Status = string(rune(200))
			res.Body = io.NopCloser(bytes.NewReader([]byte("{notTags: error}")))
		}})

		if err == nil {
			t.Error("error expected")
		}

		if strings.Contains(err.Error(), "error unmarshalling response:") {
			t.Errorf("error expected, got %s", err.Error())
		}

		if tags != nil {
			t.Error("expected nil tags")
		}
	})

	// TODO - test successful response containing the right json with tags
}

// TODO - do the same kind of testing for deleteTag() like the above

// TODO - unit test for filterTags - no need to use the mock
