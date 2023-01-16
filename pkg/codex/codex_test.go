package codex

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCompletion(t *testing.T) {
	// Create a test server that will return a response.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request parameters
		if r.URL.String() != "/completions" {
			t.Errorf("expected request to /completions, got %s", r.URL.String())
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected request method to be POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected request Content-Type to be application/json, got %s", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Errorf("expected request Authorization to be Bearer token, got %s", r.Header.Get("Authorization"))
		}

		// Send response to be tested
		w.Write([]byte(`{"id": "id", "status": "status", "result": "result"}`))
	}))
	defer ts.Close()

	// Create a new client and inject the test server URL
	c := New("token", ts.URL)

	// Create a request to pass to the function you want to test
	args := &CompletionArgs{
		Model:       "model",
		Prompt:      "prompt",
		MaxTokens:   10,
		Temperature: 0.5,
		TopP:        0.9,
		N:           5,
		Stream:      false,
		Stop:        "",
	}

	// Perform the request
	res, err := c.CreateCompletion(args)

	// Check the response
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	if res.Object != "text_completion" {
		t.Errorf("expected object to be text_completion, got %s", res.Object)
	}
	if res.Model != c.Model {
		t.Errorf("expected model to be %s, got %s", c.Model, res.Object)
	}
}
