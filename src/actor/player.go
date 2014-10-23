package actor

import (
	"glutil"
	"world"
)

type Player struct {
	BasicActor
	camera  *glutil.DisplayViewport
	height  float64
	gravity glutil.Point3D
}

func NewPlayer(myworld *world.World, position glutil.Point3D, height float64) Player {
	size := glutil.Point3D{height, height, height}
	box := world.NewBoundingBox(position, size)
	camera := glutil.CreateDefaultViewport()
	gravity := glutil.Point3D{0, -1, 0}
	return Player{NewBasicActor(myworld, box), &camera, height, gravity}
}

func (self *Player) Translate(dx float64, dy float64, dz float64) {
	rel_dx := (-glutil.Sin(self.camera.Theta) * dx) +
		(glutil.Cos(self.camera.Theta) * dz)
	rel_dz := (glutil.Cos(self.camera.Theta) * dx) +
		(glutil.Sin(self.camera.Theta) * dz)

	self.World.TranslateActor(self, rel_dx, dy, rel_dz)
	self.updateCameraFromPos()
}

func (self *Player) Rotate(dtheta, dphi int32) {
	self.camera.Rotate(dtheta, dphi)
}
func (self *Player) GetViewAngle() (int32, int32) {
	return self.camera.GetViewAngle()
}
func (self *Player) ImmediateLook() {
	self.updateCameraFromPos()
}
func (self *Player) ImmediateJump() {
	self.Position.Center.X = self.camera.Position.X
	self.Position.Center.Y = self.camera.Position.Y
	self.Position.Center.Z = self.camera.Position.Z
}
func (self *Player) PositionSelf() {
	self.camera.PositionSelf()
}
func (self *Player) ToString() string {
	return self.camera.ToString()
}

func (self *Player) updateCameraFromPos() {
	self.camera.Position = *self.GetCenter()
}

func (self *Player) GetProjection() *glutil.DisplayProjection {
	return &self.camera.Projection
}

func (self *Player) CanObstruct() bool {
	return true
}
