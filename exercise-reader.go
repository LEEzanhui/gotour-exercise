package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func (r MyReader) Read(l []byte) (int, error) {
	for i := 0; i < len(l); i++ {
		l[i] = 'A'
	}
	return len(l), nil
}

func main() {
	reader.Validate(MyReader{})
}
