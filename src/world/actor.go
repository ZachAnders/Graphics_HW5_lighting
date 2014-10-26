package world

import (
	"glutil"
)

type Actor interface {
	GetPosition() glutil.Point3D
	SetPosition(glutil.Point3D)
	Translate(dx, dy, dz float64)
	CanObstruct() bool
	CanFall() bool
	CanClip() bool
	SetID(id int)
	GetID() int
	Tick()
}
