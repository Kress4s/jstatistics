package minio_sdk

import "io"

type buffer struct {
	content []byte
	offset  int
}

func (f *buffer) Size() int64 { return int64(len(f.content)) }

func (f *buffer) Read(p []byte) (int, error) {
	if f.offset >= len(f.content) {
		return 0, io.EOF
	}
	n := copy(p, f.content[f.offset:])
	f.offset += n
	return n, nil
}

func (f *buffer) Close() error {
	f.offset = 0
	return nil
}
