package main

import "os"

type tape struct {
	file *os.File //os.File has a truncate function that will let us effectively empty the file
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0) // empty the file
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
