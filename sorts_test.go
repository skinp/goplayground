package myplayground

import (
	"reflect"
	"sort"
	"testing"
)

func TestPersonSorter(t *testing.T) {
	persons := []Person{
		{Name: "leonardo", Age: 30},
		{Name: "donatello", Age: 10},
		{Name: "michelangelo", Age: 100},
		{Name: "raphael", Age: 50},
	}
	expected := []Person{
		{Name: "donatello", Age: 10},
		{Name: "leonardo", Age: 30},
		{Name: "raphael", Age: 50},
		{Name: "michelangelo", Age: 100},
	}
	if reflect.DeepEqual(persons, expected) {
		t.Errorf("people should not be in their right order")
	}
	sort.Sort(PersonSorter(persons))
	if !reflect.DeepEqual(persons, expected) {
		t.Errorf("people are not in their right order")
	}
}
