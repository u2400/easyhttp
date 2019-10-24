package io_util

import (
	"io"
	"io/ioutil"
)

func IoReaderToString(reader *io.ReadCloser) (string, error) {
	s, err := ioutil.ReadAll(*reader)

	if err == nil {
		return string(s), nil
	}
	return "", err
}
