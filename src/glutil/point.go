package glutil

import (
	"fmt"
	"github.com/go-gl/gl"
)

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func (pnt *Point3D) Unpack() (float64, float64, float64) {
	return pnt.X, pnt.Y, pnt.Z
}

func (self *Point3D) ToString() string {
	return fmt.Sprintf("(%f, %f, %f)", self.X, self.Y, self.Z)
}

func (self *Point3D) Average(other Point3D) Point3D {
	return Point3D{
		(self.X + other.X) / 2,
		(self.Y + other.Y) / 2,
		(self.Z + other.Z) / 2}
}

func (self *Point3D) Minus(other Point3D) Point3D {
	return Point3D{
		(self.X - other.X),
		(self.Y - other.Y),
		(self.Z - other.Z)}
}

func (self *Point3D) Add(other Point3D) Point3D {
	return Point3D{
		(self.X + other.X),
		(self.Y + other.Y),
		(self.Z + other.Z)}
}

func CreatePointFromPolar(theta, phi int32) Point3D {
	return Point3D{Sin(theta) * Cos(phi), Sin(phi), Cos(theta) * Cos(phi)}
}

type Color4D struct {
	ColR float64
	ColG float64
	ColB float64
	ColA float64
}

func (self *Color4D) Set() {
	gl.Color4d(self.ColR, self.ColG, self.ColB, self.ColA)
}
