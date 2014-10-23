package actor

import (
	"github.com/go-gl/gl"
	"glutil"
	"world"
)

type Block struct {
	BasicActor
	color  glutil.Color4D
	Offset int
}

func NewBlock(myworld *world.World, position, size glutil.Point3D, color glutil.Color4D) Block {
	box := world.NewBoundingBox(position, size)
	actor := NewBasicActor(myworld, box)
	return Block{actor, color, 0}
}

func (self *Block) Render() {
	center_pnt, size_pnt := self.GetCenter(), self.GetSize()
	bot := center_pnt.Minus(*size_pnt)
	size := size_pnt.Add(*size_pnt)

	x1, y1, z1 := (&bot).Unpack()
	dx, dy, dz := size.Unpack()

	gl.PushMatrix()

	gl.Enable(gl.POLYGON_OFFSET_FILL)

	gl.PolygonOffset(float32(self.Offset/2), float32(self.Offset/2))
	//	fmt.Println(self.Offset)

	// 2. Put the prism in the right place
	gl.Translated(x1, y1, z1)
	// 1. Scale our cube to look like the prism we desire
	gl.Scaled(dx, dy, dz)

	// Code below is based on the simpleCube code in ex9.c
	gl.Begin(gl.QUADS)

	//  Front
	self.color.Set()
	gl.Normal3f(0, 0, 1)
	gl.Vertex3f(0, 0, 1)
	gl.Vertex3f(+1, 0, 1)
	self.color.Set()
	gl.Vertex3f(+1, +1, 1)
	gl.Vertex3f(0, +1, 1)

	//  Back
	self.color.Set()
	gl.Normal3f(0, 0, -1)
	gl.Vertex3f(+1, 0, 0)
	gl.Vertex3f(0, 0, 0)
	self.color.Set()
	gl.Vertex3f(0, +1, 0)
	gl.Vertex3f(+1, +1, 0)

	//  Right
	self.color.Set()
	gl.Normal3f(+1, 0, 0)
	gl.Vertex3f(+1, 0, +1)
	gl.Vertex3f(+1, 0, 0)
	self.color.Set()
	gl.Vertex3f(+1, +1, 0)
	gl.Vertex3f(+1, +1, +1)

	//  Left
	self.color.Set()
	gl.Normal3f(-1, 0, 0)
	gl.Vertex3f(0, 0, 0)
	gl.Vertex3f(0, 0, +1)
	self.color.Set()
	gl.Vertex3f(0, +1, +1)
	gl.Vertex3f(0, +1, 0)

	//  Top
	self.color.Set()
	gl.Normal3f(0, +1, 0)
	gl.Vertex3f(0, +1, +1)
	gl.Vertex3f(+1, +1, +1)
	gl.Vertex3f(+1, +1, 0)
	gl.Vertex3f(0, +1, 0)

	//  Bottom
	self.color.Set()
	gl.Normal3f(0, -1, 0)
	gl.Vertex3f(0, 0, 0)
	gl.Vertex3f(+1, 0, 0)
	gl.Vertex3f(+1, 0, +1)
	gl.Vertex3f(0, 0, +1)

	gl.End()
	gl.PopMatrix()
}
