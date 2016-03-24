package myplayground

// Person is used to show how we can sort collections of complex objects. We
// need to implement the 3 methods Len, Swap and Less on a collection type.
// Whenever we have an array/slice []Person, we can sort it as follow:
// sort.Sort(PersonSorter([]Person{...})
type Person struct {
	Name string
	Age  int
}

//PersonSorter sorts Persons by age
type PersonSorter []Person

func (s PersonSorter) Len() int           { return len(s) }
func (s PersonSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSorter) Less(i, j int) bool { return s[i].Age < s[j].Age }
