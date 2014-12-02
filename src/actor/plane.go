package actor

import (
	"entity"
	"fmt"
	//	"github.com/go-gl/gl"
	//	"github.com/go-gl/glh"
	"github.com/ianremmler/ode"
	"glutil"
	"world"
)

type Plane struct {
	*BasicActor
	ground *entity.XYPlane
}

func NewPlane(myworld *world.World, pos glutil.Point3D) Plane {
	plane := myworld.Space.NewPlane(ode.V4(0, 1, 0, 0))
	// Create ground
	c1 := pos.Add(glutil.Point3D{500, 0, 500})
	c2 := pos.Add(glutil.Point3D{-500, 0, -500})
	ground := entity.NewXYPlane(c1, c2, 0)
	ground.SetColor(glutil.Color4D{0.4313725490 * .7, 0.24705882 * .7, 0.098039215 * .7, 1})
	ground.SetPolygonOffset(1.0)

	//renderQueue.Add(&ground)
	//my_world.AddEntity(&ground.Box
	actor := NewBasicActor(myworld, &plane)
	new_plane := Plane{&actor, &ground}
	myworld.AddActor(&new_plane)
	return new_plane
}

func (self *Plane) Render() {
	self.ground.Render()
}

func (self *Plane) ToString() string {
	return fmt.Sprintf("Ground-plane Actor %d", self.GetID())
}

func (self *Plane) CanCache() bool {
	return true
}
