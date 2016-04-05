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

func isSorted(l []int) bool {
	if len(l) == 0 {
		return true
	}
	previous := l[0]
	for i := 0; i < len(l); i++ {
		if previous > l[i] {
			return false
		}
	}
	return true
}

var sortTestCases = [][]int{
	[]int{},
	[]int{1},
	[]int{1, 2, 3, 4, 5},
	[]int{10, 3, 6, 100, -1, -10, 0, 5},
}

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	for _, tc := range sortTestCases {
		InsertionSort(&tc)
		if !isSorted(tc) {
			t.Errorf("list %#v is not sorted", tc)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	for _, tc := range sortTestCases {
		SelectionSort(&tc)
		if !isSorted(tc) {
			t.Errorf("list %#v is not sorted", tc)
		}
	}
}
