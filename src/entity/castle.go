package entity

import (
	"github.com/go-gl/gl"
	"glutil"
	"util"
)

type Castle struct {
	drawbridge_angle   int32
	drawbridge_closing bool
	x_pos              float64
	y_pos              float64
	z_pos              float64
	children           util.RenderQueue
}

func NewCastle() Castle {
	moat := NewXYPlane(glutil.Point3D{-3.75, 0, -3.75}, glutil.Point3D{18.75, 0, 18.75}, 0)
	moat.SetColor(glutil.Color4D{0, 0, .5, 0})

	return Castle{90, false, 0, 0, 0, util.NewRenderQueue(&moat)}
}

func (self *Castle) Render() {
	height := float64(5)
	gl.PushMatrix()

	gl.Translated(self.x_pos, self.y_pos, self.z_pos)

	// Bigger castle!
	gl.Scaled(2, 2, 1.65)

	self.children.RenderAll()

	//Floor
	SimplePrism(glutil.Point3D{0 - .1, 0 - .1, 0 - .1}, glutil.Point3D{12 + .1, .25 + .1, 15 + .1}, .15)

	// Ceiling
	SimplePrism(glutil.Point3D{0 - .1, height - .1, 0 - .1}, glutil.Point3D{12 + .1, height + .1, 15 + .1}, .15)

	//Left wall
	SimplePrism(glutil.Point3D{0, 0, 0}, glutil.Point3D{12, height, 0 + 1}, .35)

	//Right wall
	SimplePrism(glutil.Point3D{0, 0, 15}, glutil.Point3D{12, height, 15 - 1}, .35)

	//gl.Translated(-6.0, 0, 7.5)
	//Back wall
	SimplePrism(glutil.Point3D{12, 0, 0}, glutil.Point3D{12 - 1, height, 15}, .35)

	//Front Wall
	//SimplePrism(glutil.Point3D{0, 0, 0}, glutil.Point3D{0 + 1, height, 15}, .35)
	SimplePrism(glutil.Point3D{0, 0, 0}, glutil.Point3D{0 + 1, height, 4}, .35)
	SimplePrism(glutil.Point3D{0, 0, 11}, glutil.Point3D{0 + 1, height, 15}, .35)
	SimplePrism(glutil.Point3D{0, 0, 5}, glutil.Point3D{0 + 1, height, 6}, .35)
	SimplePrism(glutil.Point3D{0, 0, 9}, glutil.Point3D{0 + 1, height, 10}, .35)

	//	Right Window
	SimplePrism(glutil.Point3D{0, 0, 4}, glutil.Point3D{0 + 1, 2, 5}, .35)

	//	Left Window
	SimplePrism(glutil.Point3D{0, 0, 10}, glutil.Point3D{0 + 1, 2, 11}, .35)

	//	Ceiling beam
	SimplePrism(glutil.Point3D{0, 4, 4}, glutil.Point3D{0 + 1, 5, 11}, .35)

	//Front left Cylinder
	SimpleColumn(glutil.Point3D{0, -.25, 0}, 3, height+2, .15)

	//Front right Cylinder
	SimpleColumn(glutil.Point3D{0, -.25, 15}, 3, height+2, .15)

	self.DrawDrawbridge(self.drawbridge_angle)

	gl.PopMatrix()
}

func (self *Castle) DrawDrawbridge(angle int32) {
	if self.drawbridge_closing && self.drawbridge_angle > 0 {
		self.drawbridge_angle -= 5
	} else if !self.drawbridge_closing && self.drawbridge_angle < 90 {
		self.drawbridge_angle += 5
	}

	gl.PushMatrix()

	gl.Translated(0, .25, 6)            // Put in place
	gl.Translated(.125, 0, 1.5)         // Undo rotate translate
	gl.Rotated(float64(angle), 0, 0, 1) // Rotate around hinge
	gl.Translated(-.125, 0, -1.5)       // Translate bridge 'hinge' to origin

	SimplePrism(glutil.Point3D{-.25, 0, 0}, glutil.Point3D{0, 4, 3}, 0)

	gl.PopMatrix()
}

func (self *Castle) OpenCloseDoor() {
	self.drawbridge_closing = !self.drawbridge_closing
}
