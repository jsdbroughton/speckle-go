package serialization

import (
	"encoding/json"
	"fmt"

	"github.com/specklesystems/speckle-go/internal/objects"
)

// ToJSON converts a Speckle object to its JSON representation
func ToJSON(obj *objects.Base) ([]byte, error) {
	return json.Marshal(obj)
}

// FromJSON converts a JSON representation to a Speckle object
func FromJSON(data []byte) (*objects.Base, error) {
	var obj objects.Base
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return &obj, nil
}

// TODO: Implement more complex serialization methods, e.g.:
// func SerializeDetached(obj *objects.Base) (string, []objects.Base, error) {}
// func DeserializeDetached(data string, objects map[string]*objects.Base) (*objects.Base, error) {}

