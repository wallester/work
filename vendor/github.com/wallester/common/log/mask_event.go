package log

import (
	"encoding/json"
	"net/url"

	"github.com/juju/errors"
	"github.com/wallester/common/strings"
)

func maskEvent(event map[string]interface{}) (map[string]interface{}, error) {
	event, _ = maskMap(event)
	for key, value := range event {
		masked, err := maskValue(value)
		if err != nil {
			return nil, errors.Annotate(err, "masking value failed")
		}

		event[key] = masked
	}

	return event, nil
}

func maskValue(value interface{}) (interface{}, error) {
	s, ok := value.(string)
	if !ok {
		return value, nil
	}

	masked, err := maskString(s)
	if err != nil {
		return nil, errors.Annotate(err, "masking string failed")
	}

	return masked, nil
}

func maskString(value string) (interface{}, error) {
	// Attempt to mask field value as JSON
	masked, err := maskJSON(value)
	if err != nil {
		return nil, errors.Annotate(err, "masking JSON failed")
	}

	if masked != nil {
		return masked, nil
	}

	// Attempt to mask field value as query
	masked, err = maskQuery(value)
	if err != nil {
		return nil, errors.Annotate(err, "masking query failed")
	}

	if masked != nil {
		return masked, nil
	}

	// If both failed, return the original unmasked value
	return value, nil
}

func maskMap(m map[string]interface{}) (map[string]interface{}, bool) {
	changed := false
	for _, keyword := range sensitiveKeywords {
		if m[keyword] != nil {
			m[keyword] = strings.Mask
			changed = true
		}
	}

	return m, changed
}

func maskJSON(data string) (interface{}, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		// It's not a map, we don't know how to mask it.
		return nil, nil
	}

	masked, changed := maskMap(m)
	if !changed {
		return data, nil
	}

	b, err := json.Marshal(masked)
	if err != nil {
		return nil, errors.Annotate(err, "marshalling to JSON failed")
	}

	return string(b), nil
}

func maskQuery(data string) (interface{}, error) {
	query, err := url.ParseQuery(data)
	if err != nil {
		return nil, errors.Annotate(err, "parsing query failed")
	}

	var changed bool
	for _, keyword := range sensitiveKeywords {
		if len(query[keyword]) != 0 {
			query.Set(keyword, strings.Mask)
			changed = true
		}
	}

	if !changed {
		return data, nil
	}

	return query.Encode(), nil
}
