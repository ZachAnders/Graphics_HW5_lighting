package actor

import (
	//	"github.com/go-gl/gl"
	"github.com/ianremmler/ode"
	"glutil"
	"world"
)

type OrbitActor struct {
	BasicActor
	child  world.Actor
	center glutil.Point3D
	radius float64
	angle  int32
	speed  int32
	moving bool
	axis   glutil.Point3D
}

func NewOrbitActor(myworld *world.World, child world.Actor, center glutil.Point3D, radius float64) OrbitActor {
	box := myworld.Space.NewBox(ode.V3(0, 0, 0))
	basic := NewBasicActor(myworld, &box)
	return OrbitActor{basic, child, center, radius, 0, 1, true, glutil.Point3D{0, 1, 0}}
}

func (self *OrbitActor) Tick() {
	var new_point glutil.Point3D
	if self.axis.Y == 1 {
		new_point = glutil.Point3D{self.radius * glutil.Cos(self.angle),
			self.center.Y,
			self.radius * glutil.Sin(self.angle)}
	} else if self.axis.Z == 1 {
		new_point = glutil.Point3D{self.radius * glutil.Cos(self.angle),
			self.radius * glutil.Sin(self.angle),
			self.center.Z}
	} else {
		new_point = glutil.Point3D{self.center.X,
			self.radius * glutil.Cos(self.angle),
			self.radius * glutil.Sin(self.angle)}
	}

	self.child.SetPosition(new_point)
	//*self.child.GetCenter() = new_point
	self.center = new_point
	if self.moving {
		self.angle += self.speed
	}
}

func (self *OrbitActor) SetSpeed(speed int32) {
	self.speed = speed
}

func (self *OrbitActor) Toggle() {
	self.moving = !self.moving
}

func (self *OrbitActor) SetAngle(angle int32) {
	self.angle = angle
}

func (self *OrbitActor) AdjustAngle(dtheta int32) {
	self.angle += dtheta
}

func (self *OrbitActor) SetAxis(axis glutil.Point3D) {
	self.axis = axis
}

func (self *OrbitActor) CanObstruct() bool {
	return false
}
func (self *OrbitActor) CanFall() bool {
	return false
}
func (self *OrbitActor) CanClip() bool {
	return false
}
