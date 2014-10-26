package entity

import (
	//	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"glutil"
	"os"
	"world"
)

// Todo: stop hardcoding this
var maxHeight = float64(10)

type XYPlane struct {
	Box      world.BoundingBox
	upLeft   glutil.Point3D
	lowRight glutil.Point3D
	yHeight  float64
	myColor  glutil.Color4D
	offset   float32
	Texture  *glh.Texture
}

func NewXYPlane(upLeft, lowRight glutil.Point3D, yHeight float64) XYPlane {
	box := world.NewBoundingBox(upLeft.Average(lowRight), upLeft.Minus(lowRight))

	dirtTex := glh.NewTexture(512, 512)
	file, _ := os.Open("dirt.png")
	dirtTex.FromPngReader(file, 0)

	return XYPlane{box, upLeft, lowRight, yHeight, glutil.Color4D{0, 0, 0, 0}, 0, dirtTex}
}

func (self *XYPlane) Render() {
	dim := 200
	gl.PushMatrix()
	self.myColor.Set()

	gl.Enable(gl.POLYGON_OFFSET_FILL)

	gl.PolygonOffset(self.offset, self.offset)
	x1, _, z1 := self.upLeft.Unpack()
	x2, _, z2 := self.lowRight.Unpack()

	if self.Texture != nil {
		gl.Color4f(.7, .7, .7, 1)
		self.Texture.Enter()
		gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	}
	for i := 0; i < dim; i++ {
		gl.Begin(gl.QUAD_STRIP)

		x_offset := float64(i) / float64(dim)
		x_offset = x1 + (x2-x1)*x_offset

		for j := 0; j < dim; j++ {
			gl.Normal3f(0.0, 1.0, 0.0)
			z_offset := float64(j) / float64(dim)
			z_offset = z1 + (z2-z1)*z_offset
			x2_offset := x_offset + (1/float64(dim))*(x2-x1)

			gl.TexCoord2d(x_offset/10, z_offset/10)
			gl.Vertex3d(x_offset, self.yHeight, z_offset)
			gl.TexCoord2d(x2_offset/10, z_offset/10)
			gl.Vertex3d(x2_offset, self.yHeight, z_offset)
		}
		gl.End()
	}
	if self.Texture != nil {
		self.Texture.Exit()
		gl.Disable(gl.TEXTURE_2D)
	}
	//	gl.Begin(gl.QUADS)
	//
	//	gl.Normal3d(0.0, 1.0, 0.0)
	//	gl.Vertex3d(x1, self.yHeight, z1)
	//	gl.Vertex3d(x1, self.yHeight, z2)
	//	gl.Vertex3d(x2, self.yHeight, z2)
	//	gl.Vertex3d(x2, self.yHeight, z1)
	//	gl.End()
	gl.PopMatrix()

	gl.Disable(gl.POLYGON_OFFSET_FILL)

}

func (self *XYPlane) SetPolygonOffset(offset float32) {
	self.offset = offset
}

func (self *XYPlane) SetColor(color glutil.Color4D) {
	self.myColor = color
}

// If we assume a specific rotation, we only require two opposing
// corners to actually define a 3D prism.
func SimplePrism(corner1, corner2 glutil.Point3D, col_offset float64) {
	x1, y1, z1 := corner1.Unpack()
	x2, y2, z2 := corner2.Unpack()
	dx, dy, dz := (x2 - x1), (y2 - y1), (z2 - z1)
	top_color, bot_color := float64(y1)/maxHeight, float64(y2)/maxHeight
	if bot_color > top_color {
		bot_color, top_color = top_color, bot_color
	}
	bot_color += col_offset
	top_color += col_offset

	gl.PushMatrix()

	// 2. Put the prism in the right place
	gl.Translated(x1, y1, z1)
	// 1. Scale our cube to look like the prism we desire
	gl.Scaled(dx, dy, dz)

	// Code below is based on the simpleCube code in ex9.c
	gl.Begin(gl.QUADS)

	//  Front
	gl.Color3d(bot_color, bot_color, bot_color)
	gl.Normal3f(0, 0, 1)
	gl.Vertex3f(0, 0, 1)
	gl.Vertex3f(+1, 0, 1)
	gl.Color3d(top_color, top_color, top_color)
	gl.Vertex3f(+1, +1, 1)
	gl.Vertex3f(0, +1, 1)

	//  Back
	gl.Color3d(bot_color, bot_color, bot_color)
	gl.Normal3f(0, 0, -1)
	gl.Vertex3f(+1, 0, 0)
	gl.Vertex3f(0, 0, 0)
	gl.Color3d(top_color, top_color, top_color)
	gl.Vertex3f(0, +1, 0)
	gl.Vertex3f(+1, +1, 0)

	//  Right
	gl.Color3d(bot_color, bot_color, bot_color)
	gl.Normal3f(+1, 0, 0)
	gl.Vertex3f(+1, 0, +1)
	gl.Vertex3f(+1, 0, 0)
	gl.Color3d(top_color, top_color, top_color)
	gl.Vertex3f(+1, +1, 0)
	gl.Vertex3f(+1, +1, +1)

	//  Left
	gl.Color3d(bot_color, bot_color, bot_color)
	gl.Normal3f(-1, 0, 0)
	gl.Vertex3f(0, 0, 0)
	gl.Vertex3f(0, 0, +1)
	gl.Color3d(top_color, top_color, top_color)
	gl.Vertex3f(0, +1, +1)
	gl.Vertex3f(0, +1, 0)

	//  Top
	gl.Color3d(top_color, top_color, top_color)
	gl.Normal3f(0, +1, 0)
	gl.Vertex3f(0, +1, +1)
	gl.Vertex3f(+1, +1, +1)
	gl.Vertex3f(+1, +1, 0)
	gl.Vertex3f(0, +1, 0)

	//  Bottom
	gl.Color3d(bot_color, bot_color, bot_color)
	gl.Normal3f(0, -1, 0)
	gl.Vertex3f(0, 0, 0)
	gl.Vertex3f(+1, 0, 0)
	gl.Vertex3f(+1, 0, +1)
	gl.Vertex3f(0, 0, +1)

	gl.End()
	gl.PopMatrix()
}

func SimpleCylinder(bot_center glutil.Point3D, radius, height float64, bot_col, top_col glutil.Color4D) {
	x, y, z := bot_center.Unpack()
	gl.PushMatrix()

	gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.Translated(x, y, z)
	gl.Scaled(radius, height, radius)

	//Cylinder body
	gl.Begin(gl.QUAD_STRIP)
	for theta := int32(0); theta <= 360; theta += 5 {
		pnt := glutil.CreatePointFromPolar(theta, 0)

		gl.Normal3d(pnt.X, 0, pnt.Z)
		bot_col.Set()
		gl.TexCoord2d(float64(theta)/180, 0)
		gl.Vertex3d(pnt.X, 0, pnt.Z)
		top_col.Set()
		gl.TexCoord2d(float64(theta)/180, height/8)
		gl.Vertex3d(pnt.X, 1, pnt.Z)
	}
	gl.End()

	//Cylinder top
	gl.Begin(gl.TRIANGLE_FAN)
	top_col.Set()
	gl.Vertex3d(0, 1, 0)
	for theta := int32(0); theta <= 360; theta += 5 {
		pnt := glutil.CreatePointFromPolar(theta, 0)
		//gl.Normal3d(0, 1, 0)
		gl.Vertex3d(pnt.X, 1, pnt.Z)
	}
	gl.End()

	//Cylinder bottom
	gl.Begin(gl.TRIANGLE_FAN)
	bot_col.Set()
	gl.Vertex3d(0, 0, 0)
	for theta := int32(0); theta <= 360; theta += 5 {
		pnt := glutil.CreatePointFromPolar(theta, 0)
		//gl.Normal3d(0, -1, 0)
		gl.Vertex3d(pnt.X, 0, pnt.Z)
	}
	gl.End()

	gl.PopMatrix()
}

func SimpleColumn(bot_center glutil.Point3D, radius, height, col_offset float64) {
	x, y, z := bot_center.Unpack()

	bot_color, top_color := float64(y)/maxHeight, (float64(y)+height)/maxHeight

	bot_color += col_offset
	top_color += col_offset

	gl.PushMatrix()

	gl.Translated(x, y, z)
	gl.Scaled(radius, height, radius)

	SimpleCylinder(glutil.Point3D{0, 0, 0}, 1, 1,
		glutil.Color4D{bot_color, bot_color, bot_color, 0},
		glutil.Color4D{top_color, top_color, top_color, 0})

	skip := false

	//Top Decorations
	for theta := int32(0); theta <= 360; theta += 36 {
		if !skip {
			gl.Begin(gl.QUAD_STRIP)
			gl.Color3d(top_color-.05, top_color-.05, top_color-.05)
			for piece := int32(0); piece < 36; piece += 6 {
				pnt := glutil.CreatePointFromPolar(theta+piece, 0)
				gl.Vertex3d(pnt.X, 1, pnt.Z)
				gl.Vertex3d(pnt.X, 1.15, pnt.Z)
			}
			gl.End()
		}
		skip = !skip

	}

	gl.PopMatrix()
}

func SimpleCone(bot_center glutil.Point3D, radius, height float64, bot_col, top_col glutil.Color4D) {
	x, y, z := bot_center.Unpack()
	gl.PushMatrix()

	gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.Translated(x, y, z)
	gl.Scaled(radius, height, radius)

	//Cylinder body
	gl.Begin(gl.QUAD_STRIP)
	for theta := int32(0); theta <= 360; theta += 5 {
		pnt := glutil.CreatePointFromPolar(theta, 0)

		gl.Normal3d(pnt.X, .5, pnt.Z)
		bot_col.Set()
		gl.TexCoord2d(float64(theta)/90, height/4)
		gl.Vertex3d(pnt.X, 0, pnt.Z)
		top_col.Set()
		gl.TexCoord2d(float64(theta)/90, 0)
		gl.Vertex3d(0, 1, 0)
	}
	gl.End()

	//Cylinder bottom
	gl.Begin(gl.TRIANGLE_FAN)
	bot_col.Set()
	gl.Vertex3d(0, 0, 0)
	for theta := int32(0); theta <= 360; theta += 5 {
		gl.Normal3d(0, -1, 0)
		pnt := glutil.CreatePointFromPolar(theta, 0)
		gl.Vertex3d(pnt.X, 0, pnt.Z)
	}
	gl.End()

	gl.PopMatrix()
}
