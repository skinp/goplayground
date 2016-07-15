package myplayground

import (
	"reflect"
	"testing"
)

const singleLineFile = "test_data/single_line.txt"
const multiLinesFile = "test_data/multi_lines.txt"

func TestScannerReadFileLines(t *testing.T) {
	t.Parallel()

	bytes, err := ScannerReadFileLines(singleLineFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := [][]byte{[]byte("test")}
	if !reflect.DeepEqual(bytes, expected) {
		t.Fatal(bytes, expected)
	}

	bytes, err = ScannerReadFileLines(multiLinesFile)
	if err != nil {
		t.Fatal(err)
	}

	expected = [][]byte{[]byte("test"), []byte("test")}
	if !reflect.DeepEqual(bytes, expected) {
		t.Fatal(bytes, expected)
	}
}

func testReadFileBytes(t *testing.T, readF func(string) ([]byte, error)) {
	bytes, err := readF(singleLineFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte("test")
	if !reflect.DeepEqual(bytes, expected) {
		t.Fatal(bytes, expected)
	}

	bytes, err = readF(multiLinesFile)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte("test\ntest\n")
	if !reflect.DeepEqual(bytes, expected) {
		t.Fatal(bytes, expected)
	}
}

func TestIoutilReadFile(t *testing.T) {
	t.Parallel()

	testReadFileBytes(t, IoutilReadFile)
}

func TestFileReadFile(t *testing.T) {
	t.Parallel()

	testReadFileBytes(t, FileReadFile)
}

func TestFileErrors(t *testing.T) {

}
