package main

import "sync"

type Instance struct{}

//START OMIT

var (
	instance Instance
	once     sync.Once
)

func New() Instance {
	once.Do(func() {
		instance = Instance{}
	})

	return instance
}

//END OMIT
