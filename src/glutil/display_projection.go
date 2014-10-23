package glutil

import (
	//	"math"
	"github.com/go-gl/gl"
	"github.com/go-gl/glu"
	//	"time"
	//	"fmt"
)

type DisplayProjection struct {
	perspective_ratio      float32
	perspective_increasing bool
	projection_matrix      [16]float32
	Dim                    float64
	AspectRatio            float64
	FieldOfView            float64
}

func NewPerspectiveProjection() DisplayProjection {
	var tmp [16]float32
	return DisplayProjection{1, true, tmp, 15, 1, 55.0}
}

func (self *DisplayProjection) updateProjection() {
	tmp := self.perspective_ratio
	if self.perspective_increasing && self.perspective_ratio < 1.0 {
		tmp += .05
	} else if !self.perspective_increasing && self.perspective_ratio > 0.0 {
		tmp -= .05
	}

	if tmp < 0 {
		tmp = 0
	} else if tmp > 1 {
		tmp = 1
	}
	self.perspective_ratio = tmp
}

func (self *DisplayProjection) SetProjectionMatrix() {
	//  Tell OpenGL we want to manipulate the projection matrix
	gl.MatrixMode(gl.PROJECTION)
	//  Undo previous transformations
	gl.LoadIdentity()

	persp_mtx := self.getPerspMatrix()
	ortho_mtx := self.getOrthoMatrix()

	var proj_mtx [16]float32

	ratio := self.perspective_ratio
	ratio = ratio * ratio * ratio
	for idx, _ := range persp_mtx {
		proj_mtx[idx] = (persp_mtx[idx] * ratio) +
			(ortho_mtx[idx] * (1 - ratio))
	}
	gl.LoadMatrixf(&proj_mtx)

	//  Switch to manipulating the model matrix
	gl.MatrixMode(gl.MODELVIEW)
	//  Undo previous transformations
	gl.LoadIdentity()
}

func (self *DisplayProjection) SwitchProjectionMode() {
	self.perspective_increasing = !self.perspective_increasing
}

func (self *DisplayProjection) getPerspMatrix() *[16]float32 {
	gl.MatrixMode(gl.PROJECTION)

	gl.LoadIdentity()
	glu.Perspective(self.FieldOfView, self.AspectRatio, self.Dim/16, 32*self.Dim)
	var tmp [16]float32
	output := tmp[:]
	gl.GetFloatv(gl.PROJECTION_MATRIX, output)
	return &tmp
}

func (self *DisplayProjection) getOrthoMatrix() *[16]float32 {
	gl.MatrixMode(gl.PROJECTION)

	gl.LoadIdentity()
	gl.Ortho(-self.Dim*self.AspectRatio, +self.Dim*self.AspectRatio,
		-self.Dim, +self.Dim,
		-16*self.Dim, 32*self.Dim)
	var tmp [16]float32
	output := tmp[:]
	gl.GetFloatv(gl.PROJECTION_MATRIX, output)
	return &tmp
}
