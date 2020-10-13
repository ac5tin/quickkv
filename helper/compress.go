package helper

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/ulikunitz/xz"
)

// Gzip - gzip compress binary
func Gzip(data *[]byte) ([]byte, error) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	if _, err := gz.Write(*data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// UnGzip - gzip decompress
func UnGzip(input *[]byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewBuffer(*input))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Xz - xz compress binary
func Xz(data *[]byte) ([]byte, error) {
	var b bytes.Buffer

	xw, err := xz.NewWriter(&b)
	if err != nil {
		return nil, err
	}
	if _, err := xw.Write(*data); err != nil {
		return nil, err
	}
	if err := xw.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// UnXz - xz decompress
func UnXz(input *[]byte) ([]byte, error) {
	xr, err := xz.NewReader(bytes.NewBuffer(*input))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(xr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Compress - compress data
func Compress(data *[]byte) ([]byte, error) {
	switch os.Getenv("COMPRESSION") {
	case "xz":
		return Xz(data)
	default:
		return Gzip(data)
	}
}

// Decompress - decompress data
func Decompress(input *[]byte) ([]byte, error) {
	switch os.Getenv("COMPRESSION") {
	case "xz":
		return UnXz(input)
	default:
		return UnGzip(input)
	}
}
