package storage

import "dirScanner/internal/server/models"

type Storage struct {
	path  string
	files []models.File
}

var customMap = make(map[string][]models.File)

func New(path string) Storage {
	return Storage{
		path: path,
	}
}
