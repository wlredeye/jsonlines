package jsonlines

import (
	"testing"
	"strings"
	"bytes"
)

const data = `{"Name":"Paul","Age":20}
{"Name":"John","Age":30}`

type Person struct {
	Name string
	Age int64
}

func TestDecode(t *testing.T) {
	people := []Person{}
	err := Decode(strings.NewReader(data), &people)
	if err != nil {
		t.Fatalf("Decode returns error: ", err)
	}
	if len(people) != 2 {
		t.Fatalf("Expected 2 objects in slice, got %s", len(people))
	}
	if people[0].Name != "Paul" {
		t.Fatalf("Unexpected value in first object Name field: ", people[0].Name)
	}
}

func TestDecodeWrongTypes(t *testing.T) {
	people := map[Person]int{}
	err := Decode(strings.NewReader(data), &people)
	if err == nil {
		t.Fatal("Decode doesnt returns error")
	}
	if err.Error() != "expected pointer to slice, got pointer to map" {
		t.Fatalf("Decode return wrong error: ", err)
	}
}

func TestEncode(t *testing.T) {
	people := []*Person{
		&Person{Name: "Paul", Age: 20},
		&Person{Name: "John", Age: 30},
	}

	var buf bytes.Buffer
	err := Encode(&buf, &people)
	if err != nil {
		t.Fatal("Encode returns error: ", err)
	}
	if !strings.Contains(buf.String(), data) {
        	t.Fatalf("Encode return wrong data, expected: %s, got: %s",
			data, buf.String())
	}
}