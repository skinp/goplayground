package myplayground

import "time"

// iNeedMocking1 shows a first way to mock in Go. The object using the mock
// has an optional function it can take when we instantiate, otherwise we call
// a default function that's always part of the using method.
// Pros:
//	- super simple to mock and no boilerplate when using without mocking
// Cons:
//	- gets messy if we have a lot of functions to mock
type iNeedMocking1 struct {
	now func() time.Time
}

func (i iNeedMocking1) somethingWithNow() time.Time {
	now := time.Now
	if i.now != nil {
		now = i.now
	}
	return now()
}

type nower interface {
	Now() time.Time
}

type timeNow struct{}

func (n timeNow) Now() time.Time {
	return time.Now()
}

// iNeedMocking1 shows a second way to mock in Go. It uses an interface type
// nower that requires 1 function Now that has the same signature as time.Now().
// Pros:
//	- we can mock multiple function calls with 1 interface
//	- simple to call the function itself, only type.interface.Call()
// Cons:
//	- more boilerplate, we need to declare an interface and an extra struct
//	- need to instantiate the implementing struct when creting the object that
//	  uses the mocked function
type iNeedMocking2 struct {
	nower nower
}

func (i iNeedMocking2) somethingWithNow() time.Time {
	return i.nower.Now()
}
