package myplayground

import (
	"testing"
	"time"
)

var mockDate = time.Date(2016, 3, 23, 0, 0, 0, 0, time.UTC)

func TestMock1(t *testing.T) {
	i := iNeedMocking1{
		now: func() time.Time { return mockDate },
	}
	got := i.somethingWithNow()
	if got != mockDate {
		t.Errorf("wrong now. exp: %#v, got: %#v", got, mockDate)
	}
}

type mockNow struct{}

func (m mockNow) Now() time.Time {
	return mockDate
}

func TestMock2(t *testing.T) {
	i := iNeedMocking2{nower: mockNow{}}
	got := i.somethingWithNow()
	if got != mockDate {
		t.Errorf("wrong now. exp: %#v, got: %#v", got, mockDate)
	}
}
