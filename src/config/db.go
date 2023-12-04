package config

import (
	"blog-api/entities"
	"sync"
)

var (
	Id     int = 1
	IdLock sync.Mutex
)

var (
	Db     map[int]entities.BlogEntity = map[int]entities.BlogEntity{}
	DbLock sync.Mutex
)
