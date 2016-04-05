package myplayground

import "sort"

// Person is used to show how we can sort collections of complex objects. We
// need to implement the 3 methods Len, Swap and Less on a collection type.
// Whenever we have an array/slice []Person, we can sort it as follow:
// sort.Sort(PersonSorter([]Person{...})
type Person struct {
	Name string
	Age  int
}

// Less returns whether the Person object is smaller than the other based on
// their age.
func (p Person) Less(another Person) bool {
	return p.Age < another.Age
}

// PersonSorter sorts Persons by age.
type PersonSorter []Person

func (s PersonSorter) Len() int           { return len(s) }
func (s PersonSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSorter) Less(i, j int) bool { return s[i].Less(s[j]) }

// PersonPair is an intermediate data structure that helps us sort a map of
// Persons.
type PersonPair struct {
	Name   string
	Person Person
}

// PersonPairSorter sorts a map of Persons by age.
type PersonPairSorter []PersonPair

func (s PersonPairSorter) Len() int           { return len(s) }
func (s PersonPairSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonPairSorter) Less(i, j int) bool { return s[i].Person.Less(s[j].Person) }

// SortPersonMapByKey sorts a map of name->Person by key. We just return the
// list of keys because we assume the caller can just iterate over the keys
// and use map index to do whatever it wants to do with the order of keys.
func SortPersonMapByKey(pm map[string]Person) []string {
	keys := []string{}
	for k, _ := range pm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// SortPersonMapByValue sorts a map of name->Person by value. Because go doesn't
// allow for sort by value built in like Python, we need to use an intermediate
// list of pair data structure and sort the list. This list is what we return
// because map a never sorted in Go.
func SortPersonMapByValue(pm map[string]Person) []PersonPair {
	pp := []PersonPair{}
	for k, v := range pm {
		pp = append(pp, PersonPair{k, v})
	}
	sort.Sort(PersonPairSorter(pp))
	return pp
}

// InsertionSort sorts a list of integers using the insertion sort algorithm.
func InsertionSort(toSort *[]int) {
	for i := 1; i < len(*toSort); i++ {
		pos := i
		current := (*toSort)[pos]

		for pos > 0 && (*toSort)[pos-1] > current {
			(*toSort)[pos] = (*toSort)[pos-1]
			pos--
		}
		(*toSort)[pos] = current
	}
}

// SelectionSort sorts a list of integers using the selection sort algorithm.
func SelectionSort(toSort *[]int) {
	for i := 0; i < len(*toSort); i++ {
		min := i
		for j := i + 1; j < len(*toSort); j++ {
			if (*toSort)[j] < (*toSort)[min] {
				min = j
			}
		}
		if min != i {
			i, min = min, i
		}
	}
}
