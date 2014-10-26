package world

import (
	"github.com/ianremmler/ode"
	"glutil"
	//	"math"
)

type World struct {
	Entities      []Entity
	Actors        []Actor
	gravity       glutil.Point3D
	actor_counter int
	CtGrp         *ode.JointGroup
	Ode_world     *ode.World
	Space         *ode.HashSpace
	simulate      bool
}

func NewWorld() World {
	ode.Init(0, ode.AllAFlag)
	ctGrp := ode.NewJointGroup(100000)
	test_world := ode.NewWorld()
	space := ode.NilSpace().NewHashSpace()

	test_world.SetGravity(ode.V3(0, -9.8, 0))
	space.NewPlane(ode.V4(0, 1, 0, 0))
	//plane.Body().SetLinearDamping(.5)
	//test_world.SetLinearDamping(.25)
	test_world.SetLinearDampingThreshold(5)
	test_world.SetERP(.5)

	new_world := World{nil, nil, glutil.Point3D{0, -.5, 0}, 0,
		&ctGrp, &test_world, &space, false}

	return new_world
}

func (self *World) Tick() {
	self.Space.Collide(0, self.collideCallback)
	self.Ode_world.Step(0.050)
	self.CtGrp.Empty()

	for _, actor := range self.Actors {
		actor.Tick()
	}
}

func (self *World) AddEntity(thing Entity) {
	self.Entities = append(self.Entities, thing)
}

func (self *World) AddActor(thing Actor) {
	thing.SetID(self.actor_counter)
	self.actor_counter++
	self.Actors = append(self.Actors, thing)
}

func (self *World) collideCallback(data interface{}, obj1, obj2 ode.Geom) {
	contact := ode.NewContact()
	body1, body2 := obj1.Body(), obj2.Body()
	if body1 != 0 && body2 != 0 && body1.Connected(body2) {
		return
	}
	//contact.Surface.Mode = 0
	contact.Surface.Mu = 10
	contact.Surface.Mu2 = 0
	contact.Surface.Bounce = 0.0
	contact.Surface.BounceVel = 0.0
	contact.Surface.SoftCfm = 0.01

	cts := obj1.Collide(obj2, 1, 0)
	if len(cts) > 0 {
		contact.Geom = cts[0]
		ct := self.Ode_world.NewContactJoint(*self.CtGrp, contact)
		ct.Attach(body1, body2)
	}
}

func (self *World) Pause() {
	self.simulate = false
}

func (self *World) Start() {
	self.simulate = true
	//	go self.Collider()
}
