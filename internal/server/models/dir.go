package models

type Dir struct {
	Path  string
	Files []File
}

type File struct {
	FileName string
	Size     int64
}
