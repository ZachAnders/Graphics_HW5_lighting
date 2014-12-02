package actor

import (
	"fmt"
	"github.com/ianremmler/ode"
	"glutil"
	"world"
)

type BasicActor struct {
	Model ode.Geom
	World *world.World
	id    int
}

func NewBasicActor(world *world.World, model ode.Geom) BasicActor {
	return BasicActor{model, world, 0}
}

func (self *BasicActor) Translate(dx, dy, dz float64) {
	pos := self.Model.Position()
	self.Model.SetPosition(ode.V3(dx+pos[0], dy+pos[1], dz+pos[2]))
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

func (self *BasicActor) CanCache() bool {
	// By default, you cannot cache actor renderings
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

func (self *BasicActor) Render() {
}

func (self *BasicActor) GetPosition() glutil.Point3D {
	return glutil.NewODEPoint3D(self.Model.Position())
}

func (self *BasicActor) SetPosition(newVal glutil.Point3D) {
	self.Model.SetPosition(newVal.ToODE())
}

func (self *BasicActor) ToString() string {
	return fmt.Sprintf("Basic Actor %d", self.id)
}

func (self *BasicActor) Interact(my_geom, other_geom ode.Geom) bool {
	return true
}
