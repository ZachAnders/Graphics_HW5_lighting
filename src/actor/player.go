package actor

import (
	"github.com/ianremmler/ode"
	"glutil"
	"world"
)

type Player struct {
	BasicActor
	camera  *glutil.DisplayViewport
	height  float64
	gravity glutil.Point3D
}

func (self *Player) Tick() {
	self.updateCameraFromPos()
}

func NewPlayer(myworld *world.World, position glutil.Point3D, height float64) Player {
	box := myworld.Space.NewBox(ode.V3(height, height, height))

	mass := ode.NewMass()
	mass.SetBox(1, ode.V3(height, height, height))
	mass.Adjust(1)
	body := myworld.Ode_world.NewBody()
	//body.SetLinearDamping(.45)
	body.SetMass(mass)
	body.SetPosition(position.ToODE())
	body.SetMass(mass)
	box.SetBody(body)

	box.SetPosition(position.ToODE())
	camera := glutil.CreateDefaultViewport()
	gravity := glutil.Point3D{0, -1, 0}
	return Player{NewBasicActor(myworld, &box), &camera, height, gravity}
}

func (self *Player) Translate(dx float64, dy float64, dz float64) {
	rel_dx := (-glutil.Sin(self.camera.Theta) * dx) +
		(glutil.Cos(self.camera.Theta) * dz)
	rel_dz := (glutil.Cos(self.camera.Theta) * dx) +
		(glutil.Sin(self.camera.Theta) * dz)

		//	pos := self.Model.Position()
	//self.Model.SetPosition(ode.V3(rel_dx+pos[0], pos[1], rel_dz+pos[2]))
	self.Model.Body().SetForce(ode.V3(20*rel_dx, 20*dy, 20*rel_dz))
	//self.Model.Body().SetForce(ode.V3(20*rel0*rel_dx, 20*dy, 20*rel_dz))
	//self.Model.Body().SetForce(ode.V3(0, 20*dy, 0))
	//self.World.TranslateActor(self, rel_dx, dy, rel_dz)
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
	//	self.Position.Center.X = self.camera.Position.X
	//	self.Position.Center.Y = self.camera.Position.Y
	//	self.Position.Center.Z = self.camera.Position.Z
}
func (self *Player) PositionSelf() {
	self.camera.PositionSelf()
}
func (self *Player) ToString() string {
	return self.camera.ToString()
}

func (self *Player) updateCameraFromPos() {
	self.camera.Position = glutil.NewODEPoint3D(self.Model.Position())
}

func (self *Player) GetProjection() *glutil.DisplayProjection {
	return &self.camera.Projection
}

func (self *Player) CanObstruct() bool {
	return true
}
