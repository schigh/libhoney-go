package libhoney

import (
	"encoding/json"
	"time"
)

var (
	mc *markerClient
)

type Marker struct {
	Type string
	Message string
	StartTime time.Time
}

func (m *Marker) MarshalJSON() ([]byte, error) {
	type alias struct {
		Type string `json:"type"`
		Message string `json:"message"`
		StartTime int64 `json:"start_time"`
	}
	return json.Marshal(&alias{
		Type: m.Type,
		Message: m.Message,
		StartTime: m.StartTime.Unix(),
	})
}

func (m *Marker) Send() (string, error) {
	if m.StartTime.IsZero() {
		m.StartTime = time.Now()
	}

	return mc.SendMarker(m)
}

func DeleteMarker(id string) error {
	return mc.DeleteMarker(id)
}
