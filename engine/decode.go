package engine

import (
	"encoding/json"
	"io"
)

// Decoder defines the signature for response decoder functions
type Decoder func(io.Reader, *ResponseContext) error

// JSONDecoder decodes the reader content and puts it into the Data property of the
// injected ResponseContext
func JSONDecoder(r io.Reader, c *ResponseContext) error {
	var target map[string]interface{}

	if err := json.NewDecoder(r).Decode(&target); err != nil {
		return err
	}
	c.Data = target
	return nil
}

// JSONDecoder decodes the reader content and puts it into the Array property of the
// injected ResponseContext
func JSONArrayDecoder(r io.Reader, c *ResponseContext) error {
	var target []map[string]interface{}
	if err := json.NewDecoder(r).Decode(&target); err != nil {
		return err
	}
	c.Array = target
	return nil
}
