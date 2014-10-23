package entity

import (
	"github.com/go-gl/gl"
	"glutil"
)

type Tree struct {
	pos     glutil.Point3D
	x_scale float64
	y_scale float64
	z_scale float64
}

func NewTree(pos glutil.Point3D, x_scale, y_scale, z_scale float64) Tree {
	return Tree{pos, x_scale, y_scale, z_scale}
}

func (self *Tree) Render() {
	x, y, z := self.pos.Unpack()
	gl.PushMatrix()

	gl.Translated(x, y, z)
	gl.Scaled(self.x_scale, self.y_scale, self.z_scale)
	l_brown := glutil.Color4D{0.4313725490 - .1, 0.24705882 - .1, 0.098039215 - .1, 1}
	d_brown := glutil.Color4D{0.4313725490 - .15, 0.24705882 - .15, 0.098039215 - .15, 1}
	l_green := glutil.Color4D{.15, .55, .15, 1}
	d_green := glutil.Color4D{.15 - .15, .55 - .15, .15 - .15, 1}
	SimpleCylinder(glutil.Point3D{0, 0, 0}, 1, 3, d_brown, l_brown)
	SimpleCone(glutil.Point3D{0, 3, 0}, 3, 5, d_green, l_green)

	gl.PopMatrix()
}
