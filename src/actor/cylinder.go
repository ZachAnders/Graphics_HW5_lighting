package actor

import (
	"entity"
	"glutil"
	"world"
)

type Cylinder struct {
	BasicActor
	radius float64
	height float64
}

func NewCylinder(my_world *world.World, position glutil.Point3D, radius, height float64) Cylinder {
	center := glutil.Point3D{position.X, position.Y + (height / 2), position.Z}
	size := glutil.Point3D{radius, height / 2, radius}
	box := world.NewBoundingBox(center, size)
	return Cylinder{NewBasicActor(my_world, box), radius, height}
}

func (self *Cylinder) Render() {
	bot := self.Position.Center
	bot = bot.Minus(glutil.Point3D{0, self.height, 0})
	color := glutil.Color4D{.25, .25, .25, 1}
	entity.SimpleCylinder(bot, self.radius, self.height, color, color)
}

func (self *Cylinder) CanFall() bool {
	return false
}

func (self *Cylinder) CanObstruct() bool {
	return true
}
