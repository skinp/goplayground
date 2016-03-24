package myplayground

import (
	"reflect"
	"sort"
	"testing"
)

func TestPersonSorter(t *testing.T) {
	t.Parallel()
	persons := []Person{
		{"leonardo", 30},
		{"donatello", 10},
		{"michelangelo", 100},
		{"raphael", 50},
	}
	expected := []Person{
		{"donatello", 10},
		{"leonardo", 30},
		{"raphael", 50},
		{"michelangelo", 100},
	}
	if reflect.DeepEqual(persons, expected) {
		t.Errorf("people should not be in their right order")
	}
	sort.Sort(PersonSorter(persons))
	if !reflect.DeepEqual(persons, expected) {
		t.Errorf("people are not in their right order")
	}
}

func TestPersonPairSorter(t *testing.T) {
	t.Parallel()
	persons := []PersonPair{
		{"l", Person{"leonardo", 30}},
		{"d", Person{"donatello", 10}},
		{"m", Person{"michelangelo", 100}},
		{"r", Person{"raphael", 50}},
	}
	expected := []PersonPair{
		{"d", Person{"donatello", 10}},
		{"l", Person{"leonardo", 30}},
		{"r", Person{"raphael", 50}},
		{"m", Person{"michelangelo", 100}},
	}
	if reflect.DeepEqual(persons, expected) {
		t.Errorf("people should not be in their right order")
	}
	sort.Sort(PersonPairSorter(persons))
	if !reflect.DeepEqual(persons, expected) {
		t.Errorf("people are not in their right order")
	}
}

func TestSortMapByKey(t *testing.T) {
	t.Parallel()
	persons := map[string]Person{
		"l": Person{"leonardo", 30},
		"d": Person{"donatello", 10},
		"m": Person{"michelangelo", 100},
		"r": Person{"raphael", 50},
	}
	expected := []string{"d", "l", "m", "r"}
	got := SortPersonMapByKey(persons)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("keys are not in their right order")
	}
}

func TestSortMapByValue(t *testing.T) {
	t.Parallel()
	persons := map[string]Person{
		"l": Person{"leonardo", 30},
		"d": Person{"donatello", 10},
		"m": Person{"michelangelo", 100},
		"r": Person{"raphael", 50},
	}
	expected := []PersonPair{
		{"d", Person{"donatello", 10}},
		{"l", Person{"leonardo", 30}},
		{"r", Person{"raphael", 50}},
		{"m", Person{"michelangelo", 100}},
	}
	got := SortPersonMapByValue(persons)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("person pairs are not in their right order")
	}
}
