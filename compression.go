package store

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type Compress interface {
	Compress([]byte) ([]byte, error)
}

type Decompress interface {
	Decompress([]byte) ([]byte, error)
}

// Primary interface for all kind of compression.
type Compression interface {
	Compress
	Decompress
}

type genericCompression struct {
	cf func([]byte) ([]byte, error)
	df func([]byte) ([]byte, error)
}

func (g *genericCompression) Compress(dst []byte) ([]byte, error) {
	return g.cf(dst)
}

func (g *genericCompression) Decompress(src []byte) ([]byte, error) {
	return g.df(src)
}

// NewGzipCompression returns a Gzip-based Compression.
// Used at store package only for transparent compression.
func NewGzipCompression() Compression {
	return &genericCompression{
		cf: CompressGzip,
		df: DecompressGzip,
	}
}

func CompressGzip(dst []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(dst); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecompressGzip(src []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	if err := zr.Close(); err != nil {
		return nil, err
	}
	return b, nil
}
