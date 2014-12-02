package actor

import (
	"entity"
	//	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/ianremmler/ode"
	"glutil"
	"math"
	"world"
)

type Tree struct {
	*BasicActor
	pos      glutil.Point3D
	x_scale  float64
	y_scale  float64
	z_scale  float64
	TrunkTex *glh.Texture
	TreeTex  *glh.Texture
}

func NewTree(myworld *world.World, pos glutil.Point3D, x_scale, y_scale, z_scale float64) Tree {
	radius := 3 * math.Max(x_scale, z_scale)
	//cylinder := myworld.Space.NewCylinder(radius, 100*y_scale)
	cylinder := myworld.Space.NewBox(ode.V3(3*x_scale, 10*y_scale, 3*z_scale))

	pos = pos.Minus(glutil.Point3D{radius / 2, 0, radius / 2})
	cylinder.SetPosition(pos.ToODE())

	actor := NewBasicActor(myworld, &cylinder)
	cylinder.SetData(&actor)

	return Tree{&actor, pos, x_scale, y_scale, z_scale, nil, nil}
}

func (self *Tree) Render() {
	cyl := self.Model.(*ode.Box)
	//x, y, z := self.pos.Unpack()
	pnt := glutil.NewODEPoint3D(cyl.Position())
	x, y, z := pnt.Unpack()
	gl.PushMatrix()

	gl.Translated(x, y, z)

	gl.Scaled(self.x_scale, self.y_scale, self.z_scale)

	mat := self.Model.Rotation()
	rot := [16]float64{
		mat[0][0], mat[1][0], mat[2][0], 0,
		mat[0][1], mat[1][1], mat[2][1], 0,
		mat[0][2], mat[1][2], mat[2][2], 0,
		0, 0, 0, 1,
	}
	gl.MultMatrixd(&rot)

	white := glutil.Color4D{1, 1, 1, 1}

	if self.TrunkTex != nil {
		gl.Color4f(1, 1, 1, 1)
		self.TrunkTex.Enter()
	}

	entity.SimpleCylinder(glutil.Point3D{0, 0, 0}, 1, 3, white, white)
	if self.TrunkTex != nil {
		self.TrunkTex.Exit()
	}
	if self.TreeTex != nil {
		gl.Color4f(1, 1, 1, 1)
		self.TreeTex.Enter()
	}
	entity.SimpleCone(glutil.Point3D{0, 3, 0}, 3, 5, white, white)
	if self.TreeTex != nil {
		self.TreeTex.Exit()
	}

	gl.PopMatrix()
}

func (self *Tree) CanCache() bool {
	return true
}
