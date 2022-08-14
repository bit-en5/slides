package main

import (
	"fmt"
	"time"
)

//START1 OMIT

type Scheduler struct{ m mover }

type mover interface {
	TimeToGo(from, to string) time.Duration
}

//END1 OMIT

type Taxi struct{}

func (t Taxi) TimeToGo(from, to string) time.Duration {
	return 90 * time.Minute
}

func main() {
	t := Taxi{}
	s := Scheduler{m: t}
	fmt.Println(s.m.TimeToGo("Bern", "Basel"))
}
