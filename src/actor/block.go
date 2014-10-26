package actor

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/ianremmler/ode"
	"glutil"
	"world"
)

type Block struct {
	BasicActor
	color    glutil.Color4D
	Offset   int
	Texture  *glh.Texture
	TexScale float64
}

func Tick() {
}

func NewBlock(myworld *world.World, position, size glutil.Point3D, color glutil.Color4D, hasMass bool) Block {
	box := myworld.Space.NewBox(ode.V3(size.Unpack()))
	if hasMass {
		mass := ode.NewMass()
		mass.SetBox(1, ode.V3(size.X, size.Y, size.Z))
		mass.Adjust(1)
		body := myworld.Ode_world.NewBody()
		body.SetMass(mass)
		body.SetPosition(position.ToODE())
		body.SetMass(mass)
		box.SetBody(body)

	}
	box.SetPosition(position.ToODE())

	actor := NewBasicActor(myworld, &box)
	return Block{actor, color, 0, nil, 1}
}

func (self *Block) Render() {
	center_pnt, size_pnt := self.Model.Position(), self.Model.Lengths()
	//	bot := center_pnt.Minus(*size_pnt)
	//	size := size_pnt.Add(*size_pnt)

	x1, y1, z1 := center_pnt[0], center_pnt[1], center_pnt[2]
	dx, dy, dz := size_pnt[0], size_pnt[1], size_pnt[2]

	gl.PushMatrix()

	gl.Enable(gl.POLYGON_OFFSET_FILL)

	gl.PolygonOffset(float32(self.Offset/2), float32(self.Offset/2))
	//	fmt.Println(self.Offset)

	// 2. Put the prism in the right place
	gl.Translated(x1, y1, z1)
	// 1. Scale our cube to look like the prism we desire
	gl.Scaled(dx/2, dy/2, dz/2)

	mat := self.Model.Rotation()
	rot := [16]float64{
		mat[0][0], mat[1][0], mat[2][0], 0,
		mat[0][1], mat[1][1], mat[2][1], 0,
		mat[0][2], mat[1][2], mat[2][2], 0,
		0, 0, 0, 1,
	}
	gl.MultMatrixd(&rot)
	// Code below is based on the simpleCube code in ex9.c

	if self.Texture != nil {
		gl.Color4f(1, 1, 1, 1)
		self.Texture.Enter()
		gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	} else {
		self.color.Set()
	}
	gl.Begin(gl.QUADS)

	tscale := self.TexScale

	//  Front
	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2d(dx/tscale, 0)
	gl.Vertex3f(+1, -1, 1)
	gl.TexCoord2d(dx/tscale, dy/tscale)
	gl.Vertex3f(+1, +1, 1)
	gl.TexCoord2d(0, dy/tscale)
	gl.Vertex3f(-1, +1, 1)

	//  Back
	gl.Normal3f(0, 0, -1)
	gl.TexCoord2d(0, 0)
	gl.Vertex3f(+1, -1, -1)
	gl.TexCoord2d(dx/tscale, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2d(dx/tscale, dy/tscale)
	gl.Vertex3f(-1, +1, -1)
	gl.TexCoord2d(0, dy/tscale)
	gl.Vertex3f(+1, +1, -1)

	//  Right
	gl.Normal3f(+1, 0, 0)
	gl.TexCoord2d(0, 0)
	gl.Vertex3f(+1, -1, +1)
	gl.TexCoord2d(dy/tscale, 0)
	gl.Vertex3f(+1, -1, -1)
	gl.TexCoord2d(dy/tscale, dz/tscale)
	gl.Vertex3f(+1, +1, -1)
	gl.TexCoord2d(0, dz/tscale)
	gl.Vertex3f(+1, +1, +1)

	//  Left
	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2d(0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2d(dz/tscale, 0)
	gl.Vertex3f(-1, -1, +1)
	gl.TexCoord2d(dz/tscale, dy/tscale)
	gl.Vertex3f(-1, +1, +1)
	gl.TexCoord2d(0, dy/tscale)
	gl.Vertex3f(-1, +1, -1)

	//  Top
	gl.Normal3f(0, +1, 0)
	gl.TexCoord2d(0, 0)
	gl.Vertex3f(-1, +1, +1)
	gl.TexCoord2d(dx/tscale, 0)
	gl.Vertex3f(+1, +1, +1)
	gl.TexCoord2d(dx/tscale, dz/tscale)
	gl.Vertex3f(+1, +1, -1)
	gl.TexCoord2d(0, dz/tscale)
	gl.Vertex3f(-1, +1, -1)

	//  Bottom
	gl.Normal3f(0, -1, 0)
	gl.TexCoord2d(0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2d(dx/tscale, 0)
	gl.Vertex3f(+1, -1, -1)
	gl.TexCoord2d(dx/tscale, dz/tscale)
	gl.Vertex3f(+1, -1, +1)
	gl.TexCoord2d(0, dz/tscale)
	gl.Vertex3f(-1, -1, +1)

	gl.End()

	if self.Texture != nil {
		self.Texture.Exit()
		gl.Disable(gl.TEXTURE_2D)
	}
	gl.PopMatrix()
}
