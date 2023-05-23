package templates

import (
	"encoding/json"
	"io/ioutil"
)

// TemplateExtras defines extra annotations
// that can be used in the templates
type TemplateExtras struct {
	Extras map[string]interface{} `json:"extras"`
}

func BindFromFile(file string) (map[string]interface{}, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	extras := map[string]interface{}{}
	if err := json.Unmarshal(f, &extras); err != nil {
		return nil, err
	}

	return extras, nil
}
