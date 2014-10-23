package world

import ()

type Actor interface {
	Entity
	Translate(dx, dy, dz float64)
	CanObstruct() bool
	CanFall() bool
	CanClip() bool
	SetID(id int)
	GetID() int
	Tick()
}
