package libhoney

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMarkerClient_SendMarker(t *testing.T) {
	type test struct{
		name string
		message string
		markerType string
		startTime time.Time
		expectErr bool
		id string
		handler func(http.ResponseWriter, *http.Request)
	}

	tests := []test{
		{
			name: "happy path",
			message: "foo",
			markerType: "test",
			id: "mymarker",
			handler: func(rw http.ResponseWriter, r *http.Request) {
				_, _ = rw.Write([]byte(`{"id":"mymarker"}`))
			},
		},
		{
			name: "sad path: we hang up",
			message: "foo",
			markerType: "test",
			expectErr: true,
			handler: func(rw http.ResponseWriter, r *http.Request) {
				<-time.After(markerAPITimeout + time.Millisecond)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// DO NOT RUN PARALLEL

			server := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer server.Close()
			mc = NewMarkerClient(MarkerClientConfig{
				APIKey: "boop",
				APIHost: server.URL + "/",
				Dataset: "test",
			})

			marker := &Marker{
				Type:      tt.markerType,
				Message:   tt.message,
				StartTime: tt.startTime,
			}
			id, err := marker.Send()
			if err != nil && !tt.expectErr {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectErr {
				t.Fatal("expected error, got none")
			}
			if id != tt.id {
				t.Fatalf("expected id '%s', got '%s", tt.id, id)
			}
		})
	}
}

func TestMarkerClient_DeleteMarker(t *testing.T) {
	// out of time
}
