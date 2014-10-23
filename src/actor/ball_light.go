package actor

import (
	"entity"
	"github.com/go-gl/gl"
	"glutil"
	"world"
)

type BallLight struct {
	BasicActor
}

func NewBallLight(myWorld *world.World, position glutil.Point3D) BallLight {
	box := world.NewBoundingBox(position, glutil.Point3D{0, 0, 0})
	actor := NewBasicActor(myWorld, box)
	return BallLight{actor}
}

func (self *BallLight) Render() {
	x, y, z := self.Position.Center.Unpack()
	ambient := []float32{0.25, 0.25, 0.25, 1.0}
	diffuse := []float32{.55, .55, .55, 1.0}
	specular := []float32{0.00, 0.00, 0.00, 1.0}
	position := []float32{float32(x), float32(y), float32(z), 1}

	gl.ShadeModel(gl.SMOOTH)
	gl.PushMatrix()

	gl.Enable(gl.NORMALIZE)
	gl.Enable(gl.LIGHTING)

	gl.LightModeli(gl.LIGHT_MODEL_LOCAL_VIEWER, 1)

	gl.ColorMaterial(gl.FRONT_AND_BACK, gl.AMBIENT_AND_DIFFUSE)
	gl.Enable(gl.COLOR_MATERIAL)
	//.  Enable light 0
	gl.Enable(gl.LIGHT0)
	//.  Set ambient, diffuse, specular components and position of light 0
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, ambient)
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, diffuse)
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, specular)
	gl.Lightfv(gl.LIGHT0, gl.POSITION, position)

	entity.SimplePrism(self.Position.Center, self.Position.Center.Add(glutil.Point3D{2, 2, 2}), 0)

	gl.PopMatrix()
}

func (self *BallLight) CanObstruct() bool {
	return false
}

func (self *BallLight) CanFall() bool {
	return false
}

func (self *BallLight) CanClip() bool {
	return true
}
