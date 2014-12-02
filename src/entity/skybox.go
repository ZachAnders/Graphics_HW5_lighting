package entity

import (
	//	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	//	"github.com/ianremmler/ode"
	"glutil"
	//	"world"
	"os"
)

type Skybox struct {
	color    glutil.Color4D
	Texture  *glh.Texture
	Position glutil.Point3D
	TexScale float64
}

func NewSkybox() *Skybox {
	skyTex := glh.NewTexture(4096, 3072)
	//skyTex := glh.NewTexture(512, 384)
	file, _ := os.Open("skybox.png")
	skyTex.FromPngReader(file, 0)
	file.Close()

	color := glutil.Color4D{.25, .5, .5, 1}
	sb := Skybox{color, skyTex, glutil.Point3D{0, 0, 0}, 1.0}
	return &sb
}
func (self *Skybox) Render() {
	//	bot_color, top_color := self.color, self.color
	gl.PushMatrix()

	gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	// 2. Put the prism in the right place
	gl.Translated(self.Position.Unpack())

	// 1. Scale our cube to look like the prism we desire
	gl.Scaled(100, 100, 100)

	gl.Translated(-.5, -.5, -.5) // Center the cube

	gl.Disable(gl.DEPTH_TEST)

	self.Texture.Enter()

	// Code below is based on the simpleCube code in ex9.c
	gl.Begin(gl.QUADS)

	//  Front
	// GOOD
	//bot_color.Set()
	gl.Normal3f(0, 0, -1)
	gl.TexCoord2f(1.0/4, 2.0/3)
	gl.Vertex3f(0, 0, 1)
	gl.TexCoord2f(2.0/4, 2.0/3)
	gl.Vertex3f(+1, 0, 1)
	//top_color.Set()
	gl.TexCoord2f(2.0/4, 1.0/3)
	gl.Vertex3f(+1, +1, 1)
	gl.TexCoord2f(1.0/4, 1.0/3)
	gl.Vertex3f(0, +1, 1)

	//  Back
	//	bot_color.Set()
	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(3.0/4, 2.0/3)
	gl.Vertex3f(+1, 0, 0)
	gl.TexCoord2f(4.0/4, 2.0/3)
	gl.Vertex3f(0, 0, 0)
	//	top_color.Set()
	gl.TexCoord2f(4.0/4, 1.0/3)
	gl.Vertex3f(0, +1, 0)
	gl.TexCoord2f(3.0/4, 1.0/3)
	gl.Vertex3f(+1, +1, 0)

	//  Right
	// GOOD
	//	bot_color.Set()
	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2f(2.0/4, 2.0/3)
	gl.Vertex3f(+1, 0, +1)
	gl.TexCoord2f(3.0/4, 2.0/3)
	gl.Vertex3f(+1, 0, 0)
	//	top_color.Set()
	gl.TexCoord2f(3.0/4, 1.0/3)
	gl.Vertex3f(+1, +1, 0)
	gl.TexCoord2f(2.0/4, 1.0/3)
	gl.Vertex3f(+1, +1, +1)

	//  Left
	//	bot_color.Set()
	gl.Normal3f(+1, 0, 0)
	gl.TexCoord2f(0.0/4, 2.0/3)
	gl.Vertex3f(0, 0, 0)
	gl.TexCoord2f(1.0/4, 2.0/3)
	gl.Vertex3f(0, 0, +1)
	//	top_color.Set()
	gl.TexCoord2f(1.0/4, 1.0/3)
	gl.Vertex3f(0, +1, +1)
	gl.TexCoord2f(0.0/4, 1.0/3)
	gl.Vertex3f(0, +1, 0)

	//  Top
	//	top_color.Set()
	gl.Normal3f(0, -1, 0)
	gl.TexCoord2f(1.0/4, 1.0/3)
	gl.Vertex3f(0, +1, +1)
	gl.TexCoord2f(2.0/4, 1.0/3)
	gl.Vertex3f(+1, +1, +1)
	gl.TexCoord2f(2.0/4, 0.0/3)
	gl.Vertex3f(+1, +1, 0)
	gl.TexCoord2f(1.0/4, 0.0/3)
	gl.Vertex3f(0, +1, 0)

	// Bottom
	// bot_color.Set()
	gl.Normal3f(0, +1, 0)
	gl.TexCoord2f(1.0/4, 3.0/3)
	gl.Vertex3f(0, 0, 0)
	gl.TexCoord2f(2.0/4, 3.0/3)
	gl.Vertex3f(+1, 0, 0)
	gl.TexCoord2f(2.0/4, 2.0/3)
	gl.Vertex3f(+1, 0, +1)
	gl.TexCoord2f(1.0/4, 2.0/3)
	gl.Vertex3f(0, 0, +1)

	gl.End()
	self.Texture.Exit()
	gl.Enable(gl.DEPTH_TEST)
	gl.PopMatrix()

}
