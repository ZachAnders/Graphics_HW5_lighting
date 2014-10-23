package world

import (
	"glutil"
)

type Entity interface {
	HeightAt(position glutil.Point3D) (bool, float64)
	Collides(other Entity) bool
	GetCenter() *glutil.Point3D
	GetSize() *glutil.Point3D
}
