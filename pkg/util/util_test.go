package util

import "testing"

func TestStructureToMap(t *testing.T) {
	type User struct {
		ID int
		Email string
		FirstName string
	}

	u := User{1, "example@gmail.com", "John"}

	result, err := StructToMap(u)
	if err != nil {
		t.Error("not expect error")
	}

	m := map[string]interface{}{
		"id": 1,
		"email": "example@gmail.com",
		"first_name": "John",
	}


	if result["id"] != m["id"] {
		t.Errorf("expected: %d, get: %d", m["id"], result["id"])
	}

	if result["email"] != m["email"] {
		t.Errorf("expected: %s, get: %s", m["email"], result["email"])
	}

	if result["first_name"] != m["first_name"] {
		t.Errorf("expected: %s, get: %s", m["first_name"], result["first_name"])
	}
}

func TestWordsToLower(t *testing.T) {
	result := wordsToLower([]string{"First", "Name"})
	words := []string{"first", "name"}

	if result[0] != words[0] {
		t.Errorf("expected: %s, get: %s", words[0], result[0])
	}

	if result[1] != words[1] {
		t.Errorf("expected: %s, get: %s", words[1], result[1])
	}
}

func TestToSnakeCase_SingleWord(t *testing.T) {
	result := toSnakeCase([]string{"id"})
	word := "id"

	if result != word {
		t.Errorf("expected: %s, get: %s", word, result)
	}
}

func TestToSnakeCase_DoubleWord(t *testing.T) {
	result := toSnakeCase([]string{"first", "name"})
	word := "first_name"

	if result != word {
		t.Errorf("expected: %s, get: %s", word, result)
	}
}
