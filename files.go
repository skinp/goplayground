package myplayground

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// ScannerReadFileLines read a file using bufio.Scanner. It returns a slice of
// byte slice each containing 1 line.
func ScannerReadFileLines(filename string) ([][]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// this is what we return
	var bytes [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bytes = append(bytes, scanner.Bytes())
	}
	if err = scanner.Err(); err != nil {
		return bytes, err
	}

	return bytes, nil
}

// IoutilReadFile simply wraps ioutil.ReadFile.
func IoutilReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// FileReadFile reads a file using os.File.Read. It's pretty dumb and there are
// much better ways to do that in Go like using bytes.Buffer.ReadFrom or the
// above...
func FileReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// this is what we return
	var bytes []byte

	// temp buffer for each read
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return bytes, err
		}

		// EOF
		if n == 0 && err == io.EOF {
			break
		}
		bytes = append(bytes, buf[:n]...)
	}

	return bytes, nil
}
