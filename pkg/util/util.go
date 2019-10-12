package util

import (
	"regexp"
	"reflect"
	"strings"
)

func StructToMap(s interface{}) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	pattern := `[[:upper:]]*[[:lower:]|[[:digit:]]*`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, nil
	}

	e := reflect.ValueOf(s)

	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i).Name
		words := re.FindAllString(field, -1)
		words = wordsToLower(words)
		snakeCaseName := toSnakeCase(words)
		m[snakeCaseName] = e.Field(i).Interface()
	}

	return m, nil
}

func wordsToLower(words []string) []string {
	var lws []string
	for _, w := range words {
		lws = append(lws, strings.ToLower(w))
	}
	return lws
}

func toSnakeCase(words []string) string {
	var word string

	if len(words) == 1 {
		word = words[0]
	} else {
		for i := 0; i < len(words) - 1; i++ {
			word += words[i] + "_"
		}
		word += words[len(words)-1]
	}

	return word
}
