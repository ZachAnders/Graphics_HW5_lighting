package glutil

import (
	//	"math"
	"fmt"
	//	"github.com/go-gl/gl"
	"github.com/go-gl/glu"
	//	"time"
)

// Represents the current viewport, contains two (theta, phi) pairs
// One is the current value, the other is the value to approach
type DisplayViewport struct {
	Theta      int32
	Phi        int32
	Position   Point3D
	Projection DisplayProjection
}

func CreateDefaultViewport() DisplayViewport {
	origin := Point3D{0, 0, 0}
	projection := NewPerspectiveProjection()
	return DisplayViewport{0, 0, origin, projection}
}

func (self *DisplayViewport) PositionSelf() {
	self.Projection.updateProjection()

	c_theta, c_phi := self.GetViewAngle()
	dim := self.Projection.Dim

	//Perspective calculations, based on ex9.c
	var Ex float64 = -2 * dim * Sin(c_theta) * Cos(c_phi)
	var Ey float64 = 2 * dim * Sin(c_phi)
	var Ez float64 = 2 * dim * Cos(c_theta) * Cos(c_phi)

	self.Projection.SetProjectionMatrix()

	tx, ty, tz := self.Position.X, self.Position.Y, self.Position.Z
	glu.LookAt(tx, ty, tz, tx+Ex, ty+Ey, tz+Ez, 0, Cos(c_phi), 0)
	//glu.LookAt(0, 0, 0, Ex, Ey, Ez, 0, Cos(c_phi), 0)
	//gl.Translated(-self.Position.X, -self.Position.Y, -self.Position.Z)

}

func (self *DisplayViewport) GetViewAngle() (int32, int32) {
	//  Keep angles to +/-360 degrees
	return self.Theta % 360, self.Phi % 360
}

func (self *DisplayViewport) Translate(dx float64, dy float64, dz float64) {
	// TODO: Translate y?
	self.Position.X += (-Sin(self.Theta) * dx) +
		(Cos(self.Theta) * dz)
	self.Position.Y += dy
	self.Position.Z += (Cos(self.Theta) * dx) +
		(Sin(self.Theta) * dz)
}

func (self *DisplayViewport) Rotate(dtheta, dphi int32) {
	self.Theta += dtheta
	self.Phi += dphi
}

func (self *DisplayViewport) ToString() string {
	proj := "Perspective"
	if !self.Projection.perspective_increasing {
		proj = "Orthogonal"
	}
	return fmt.Sprintf("View Angle=%d,%d X=%f, Y=%f, Z=%f Proj=%s",
		self.Theta, self.Phi,
		self.Position.X, self.Position.Y, self.Position.Z, proj)
}
