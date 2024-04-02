package main

import (
	"fmt"
	"sync"
	"time"
)

type lift struct {
	currentFloor int
	direction    int // 1 - up, 0 - down, -1 - idle
	liftFloorSet *OrderedSet
	lock         sync.Mutex
}

func Lift() *lift {
	return &lift{
		currentFloor: 0,
		direction:    -1,
		liftFloorSet: NewOrderedSet(),
	}
}
func (l *lift) addFloor(sourceFloor, destinationFloor int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.liftFloorSet.Add(FloorRequest{Source: sourceFloor, Destination: destinationFloor})
}

func (l *lift) getNextRequest() (FloorRequest, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	items := l.liftFloorSet.Items()

	// If lift is idle or there are no requests, return false
	if l.direction == -1 || len(items) == 0 {
		return FloorRequest{}, false
	}

	//If going upwards
	if l.direction == 1 {
		for _, req := range items {
			if req.Source >= l.currentFloor {
				return req, true
			}
		}
	}

	//If going downwards
	if l.direction == 0 {
		for i := len(items) - 1; i >= 0; i-- {
			if items[i].Source <= l.currentFloor {
				return items[i], true
			}
		}
	}

	return FloorRequest{}, false

}

func (l *lift) StartLift() {
	go func() {
		for {
			req, exists := l.getNextRequest()

			if !exists {
				l.direction = -1
				time.Sleep(1 * time.Second)
				continue
			}

			//Update direction based on destination
			if req.Destination > l.currentFloor {
				l.direction = 1
			} else {
				l.direction = 0
			}

			//Simulate lift movement
			if req.Source != l.currentFloor {
				time.Sleep(2 * time.Second)
				l.currentFloor = req.Source
			}

			if req.Destination != l.currentFloor {
				time.Sleep(2 * time.Second)
				l.currentFloor = req.Destination
			}
			l.liftFloorSet.Remove(req)
		}
	}()
}

func (l *lift) DisplayCurrentFloor() {
	go func() {
		for {
			fmt.Printf("\n-------------------------\nCurrent floor: %d\nCurrent Direction: %d\nCurrent Queue: %v\n-------------------------\n", l.GetCurrentFloor(), l.GetCurrentDirection(), l.GetCurrentQueue())
			time.Sleep(1 * time.Second)
		}
	}()
}

func (l *lift) GetCurrentFloor() int {
	return l.currentFloor
}

func (l *lift) GetCurrentQueue() *OrderedSet {
	return l.liftFloorSet
}

func (l *lift) GetCurrentDirection() int {
	return l.direction
}
