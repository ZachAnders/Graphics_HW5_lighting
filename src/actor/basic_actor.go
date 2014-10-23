package actor

import (
	"world"
)

type BasicActor struct {
	world.Entity
	World    *world.World
	Position world.BoundingBox
	id       int
}

func NewBasicActor(world *world.World, pos world.BoundingBox) BasicActor {
	newPos := pos
	return BasicActor{&newPos, world, pos, 0}
}

func (self *BasicActor) Translate(dx, dy, dz float64) {
	self.World.TranslateActor(self, dx, dy, dz)
}

func (self *BasicActor) CanObstruct() bool {
	// By default, this actor can obstruct other objects
	return true
}

func (self *BasicActor) CanFall() bool {
	// By default, this actor is under the effect of gravity
	return true
}

func (self *BasicActor) CanClip() bool {
	// By default, this actor cannot pass through obstructions
	return false
}

func (self *BasicActor) GetID() int {
	return self.id
}
func (self *BasicActor) SetID(id int) {
	self.id = id
}

func (self *BasicActor) Tick() {

}
