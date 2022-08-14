package main

type Instance struct{}

//START OMIT

var instance *Instance

func New() *Instance {
	if instance == nil {
		instance = new(Instance)
	}

	return instance
}

//END OMIT
