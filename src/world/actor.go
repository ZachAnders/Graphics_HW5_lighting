package world

import (
	"github.com/ianremmler/ode"
	"glutil"
)

type Actor interface {
	GetPosition() glutil.Point3D
	SetPosition(glutil.Point3D)
	Translate(dx, dy, dz float64)
	Render()
	CanObstruct() bool
	CanFall() bool
	CanClip() bool
	CanCache() bool
	SetID(id int)
	GetID() int
	Tick()
	ToString() string
	Interact(ode.Geom, ode.Geom) bool
}
