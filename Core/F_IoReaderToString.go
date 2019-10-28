package Core

import (
	"io"
	"io/ioutil"
)

func ioReaderToString(reader *io.ReadCloser) (string, error) {
	s, err := ioutil.ReadAll(*reader)

	if err == nil {
		return string(s), nil
	}
	return "", err
}
