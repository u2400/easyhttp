package io_util

import (
	"io"
	"io/ioutil"
)

func F_io_reader_to_string(reader *io.ReadCloser) (string, error) {
	s, err := ioutil.ReadAll(*reader)

	if err == nil {
		return string(s), nil
	}
	return "", err
}
