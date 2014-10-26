package entity

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"glutil"
)

type Tree struct {
	pos      glutil.Point3D
	x_scale  float64
	y_scale  float64
	z_scale  float64
	TrunkTex *glh.Texture
	TreeTex  *glh.Texture
}

func NewTree(pos glutil.Point3D, x_scale, y_scale, z_scale float64) Tree {
	return Tree{pos, x_scale, y_scale, z_scale, nil, nil}
}

func (self *Tree) Render() {
	x, y, z := self.pos.Unpack()
	gl.PushMatrix()

	gl.Translated(x, y, z)
	gl.Scaled(self.x_scale, self.y_scale, self.z_scale)
	//l_brown := glutil.Color4D{0.4313725490 - .1, 0.24705882 - .1, 0.098039215 - .1, 1}
	//d_brown := glutil.Color4D{0.4313725490 - .15, 0.24705882 - .15, 0.098039215 - .15, 1}
	//l_green := glutil.Color4D{.15, .55, .15, 1}
	//d_green := glutil.Color4D{.15 - .15, .55 - .15, .15 - .15, 1}
	white := glutil.Color4D{1, 1, 1, 1}

	if self.TrunkTex != nil {
		gl.Color4f(1, 1, 1, 1)
		self.TrunkTex.Enter()
	}

	SimpleCylinder(glutil.Point3D{0, 0, 0}, 1, 3, white, white)
	if self.TrunkTex != nil {
		self.TrunkTex.Exit()
	}
	if self.TreeTex != nil {
		gl.Color4f(1, 1, 1, 1)
		self.TreeTex.Enter()
	}
	SimpleCone(glutil.Point3D{0, 3, 0}, 3, 5, white, white)
	if self.TreeTex != nil {
		self.TreeTex.Exit()
	}

	gl.PopMatrix()
}
