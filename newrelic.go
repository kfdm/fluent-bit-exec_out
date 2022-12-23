package main

import (
	"C"
)

// Taken from https://github.com/newrelic/newrelic-fluent-bit-output/blob/304303f8912a3d6680e497a3a0a68006d62ba0fc/record/record.go
// parseRecord transforms a log record emitted by FluentBit into a LogRecord
// domain type: a map of string keys and arbitrary (int, string, etc.) values.
// No value modification is performed by this method (except casting).
func ParseRecord(inputRecord map[interface{}]interface{}) map[string]interface{} {
	return parseValue(inputRecord).(map[string]interface{})
}

func parseValue(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		return string(value)
	case map[interface{}]interface{}:
		remapped := make(map[string]interface{})
		for k, v := range value {
			remapped[k.(string)] = parseValue(v)
		}
		return remapped
	case []interface{}:
		remapped := make([]interface{}, len(value))
		for i, v := range value {
			remapped[i] = parseValue(v)
		}
		return remapped
	default:
		return value
	}
}
