package events

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// Marshaler provides a default JSON marshaler.
var Marshaler = &protojson.MarshalOptions{
	UseEnumNumbers:    false,
	EmitDefaultValues: false,
	Indent:            "\t",
}

// Marshal marshals the event to JSON.
func Marshal(ev *Event) (string, error) {
	if b, err := Marshaler.Marshal(ev); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

// Unmarshal unmarshals the event from JSON.
func Unmarshal(b []byte, ev *Event) error {
	return protojson.Unmarshal(b, ev)
}
