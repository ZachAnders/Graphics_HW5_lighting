package main

/*
* CSCI 4229 Final Project
*
* Author: Zach Anders
*
* Some code derived from CSCI 4229 Examples
*
* Code for 2D text display (due to missing windowpos2i) based off some example code
* at http://programmingexamples.net/wiki/OpenGL/Text
 */

import (
	"actor"
	"entity"
	"math/rand"
	"os"
	"runtime"
	"time"

	//	. "entity"
	"glutil"
	"util"
	"world"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/ianremmler/ode"
	"github.com/rhencke/glut"
)

var (
	light    actor.OrbitActor
	crateTex *glh.Texture
	skybox   *entity.Skybox
)

var renderQueue util.RenderQueue = util.NewEmptyRenderQueue()

//var currentCamera = CreateDefaultViewport()
var my_world = world.NewWorld()
var player = actor.NewPlayer(&my_world, glutil.Point3D{0, 5, 0}, 5)
var currentMouse = glutil.CreateMouseState()
var cursor = []int{0, 0}

// Normal order:
//Translate -> Rotate -> Scale

// Creates the display function
func DisplayFunc() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.Enable(gl.DEPTH_TEST)
	gl.LoadIdentity()

	player.PositionSelf()
	skybox.Position = player.GetPosition()

	renderQueue.RenderAll()
	my_world.Tick()

	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.Color3f(.9, .9, .9)
	glutil.Print2d(5, 35, "Keys: W,S,A,D : Move Around, Space : Jump | Q / E : Move Light/Sun Up/Down | L : Pause light")
	glutil.Print2d(5, 20, "Click and Drag: Rotate | Arrow keys: Rotate | B : Spawn New Creates | R : Reset Location | G : Grab / throw crate")
	glutil.Print2d(5, 5, "%s", player.ToString())
	glutil.Print2d(cursor[0], cursor[1], "+")

	gl.Flush()
	glut.SwapBuffers()
}

// Creates the key handler function
func KeyDownFunc(ch byte, x int, y int) {
	switch ch {
	case 27:
		os.Exit(0)
		break
	case 's':
		player.Translate(-2.5, 0.0, 0.0)
		break
	case 'w':
		player.Translate(2.5, 0.0, 0.0)
		break
	case 'd':
		player.Translate(0.0, 0.0, -2.5)
		break
	case 'a':
		player.Translate(0.0, 0.0, 2.5)
		break
	case 'l':
		light.Toggle()
		break
	case 'q':
		light.AdjustAngle(-5)
		break
	case 'e':
		light.AdjustAngle(5)
		break
	case ' ':
		player.Translate(0, 10, 0)
		break
	case 'g':
		player.GrabThrow()
		break
	case 'b':
		x := (rand.Float64() * 200) - (100)
		y := (rand.Float64() * 20) - (10) + 25
		z := (rand.Float64() * 200) - (100)
		boxSpawn(x, y, z)
		break
	case 'r':
		player.SetPosition(glutil.Point3D{-20, 5, 10})
		player.Model.Body().SetForce(ode.V3(0, 0, 0))
		player.Model.Body().SetLinearVelocity(ode.V3(0, 0, 0))
	}
	glut.PostRedisplay()
}

// Creates the special key handler function
func SpecialFunc(key int, x int, y int) {
	// Phi: Elevation, Theta: Azimuth
	if key == glut.KEY_RIGHT {
		player.Rotate(10, 0)
	} else if key == glut.KEY_LEFT {
		player.Rotate(-10, 0)
	} else if key == glut.KEY_UP {
		player.Rotate(0, 10)
	} else if key == glut.KEY_DOWN {
		player.Rotate(0, -10)
	}

	//  Tell GLUT it is necessary to redisplay the scene
	glut.PostRedisplay()
}

func MouseMotion(x, y int) {
	if currentMouse.LeftDown {
		horiz_delta := currentMouse.X - x
		vert_delta := currentMouse.Y - y
		player.Rotate(int32(horiz_delta)/-5, int32(vert_delta)/5)
		player.ImmediateLook()
	}
	currentMouse.X, currentMouse.Y = x, y
}

func MouseDown(button, state, x, y int) {
	currentMouse.X, currentMouse.Y = x, y
	switch button {
	case glut.LEFT_BUTTON:
		currentMouse.LeftDown = (state == glut.DOWN)
		break
	case glut.MIDDLE_BUTTON:
		currentMouse.MiddleDown = (state == glut.DOWN)
		break
	case glut.RIGHT_BUTTON:
		currentMouse.RightDown = (state == glut.DOWN)
		player.GrabThrow()
		break
	}
}

/*
 *  GLUT calls this routine when the window is resized
 */
func Reshape(width int, height int) {
	asp := float64(1)
	//  Ratio of the width to the height of the window
	if height > 0 {
		asp = float64(width) / float64(height)
	}

	cursor[0], cursor[1] = width/2, height/2

	//  Set the viewport to the entire window
	gl.Viewport(0, 0, width, height)

	player.GetProjection().AspectRatio = asp
	player.GetProjection().SetProjectionMatrix()
}

// Idler function. Called whenever GLUT is idle.
func IdleFunc() {
	s_time := glut.Get(glut.ELAPSED_TIME)
	glut.PostRedisplay()
	e_time := glut.Get(glut.ELAPSED_TIME) - s_time
	if e_time < 16 {
		time.Sleep(time.Duration(16-e_time) * time.Millisecond)
	}
}

func boxSpawn(x, y, z float64) {
	testblock := actor.NewBlock(&my_world, glutil.Point3D{x, y, z}, glutil.Point3D{3, 3, 3}, glutil.Color4D{.5, .5, .5, 1}, true)
	testblock.Texture = crateTex
	testblock.TexScale = 3
	my_world.AddActor(&testblock)
}

func setup() {
	crateTex = glh.NewTexture(256, 256)
	file, _ := os.Open("crate.png")
	crateTex.FromPngReader(file, 0)
	file.Close()

	lightsource := actor.NewBallLight(&my_world, glutil.Point3D{15, 15, 15})
	light = actor.NewOrbitActor(&my_world, &lightsource, glutil.Point3D{0, 0, 0}, 450)
	light.SetAxis(glutil.Point3D{1, 0, 0})
	my_world.AddActor(&light)

	actor.NewPlane(&my_world, glutil.Point3D{0, 0, 0})

	skybox = entity.NewSkybox()
	renderQueue.Add(skybox)

	rand.Seed(0x12345 + 8)
	my_world.AddActor(player)

	//stairs := actor.NewStairs(&my_world, glutil.Point3D{0, 0, 0})
	//renderQueue.Add(&stairs)

	trunkTex := glh.NewTexture(256, 256)
	file, _ = os.Open("trunk.png")
	trunkTex.FromPngReader(file, 0)
	file.Close()

	treeTex := glh.NewTexture(256, 256)
	file, _ = os.Open("tree.png")
	treeTex.FromPngReader(file, 0)
	file.Close()

	scale := float64(300)
	for i := 0; i < 200; i++ {
		x := (rand.Float64() * scale) - (scale / 2)
		z := (rand.Float64() * scale) - (scale / 2)
		x_scale, z_scale := (rand.Float64()/2)+.75, (rand.Float64()/2)+.75
		y_scale := (rand.Float64() * 1) + .75
		//	if (x < -5 || x > 20) && (z < -5 || z > 20) {
		tree1 := actor.NewTree(&my_world, glutil.Point3D{x, 0, z}, x_scale, y_scale, z_scale)
		tree1.TrunkTex = trunkTex
		tree1.TreeTex = treeTex
		my_world.AddActor(&tree1)
		//	}
	}
	for i := 0; i < 50; i++ {
		x := (rand.Float64() * scale) - (scale / 2)
		y := (rand.Float64() * 10) + 15
		z := (rand.Float64() * scale) - (scale / 2)
		boxSpawn(x, y, z)

	}

	// Put camera somewhere
	//currentCamera.Translate(5, -5, 28)
	//player.Translate(10, 0, -20)
	player.Model.SetPosition(ode.V3(10, 5, -20))
	player.ImmediateJump() // Skip interpolation
	//player.Rotate(-120, 0)
	player.ImmediateLook() // Skip interpolation
	my_world.Start()
}

func main() {
	// Let go use 2 threads
	runtime.GOMAXPROCS(2)

	// Init
	glut.InitDisplayMode(glut.RGB | glut.DOUBLE | glut.DEPTH)
	glut.InitWindowSize(750, 750)
	glut.CreateWindow("Zach Anders - Assignment 7")
	setup()

	// Display Callbacks
	glut.DisplayFunc(DisplayFunc)
	glut.IdleFunc(IdleFunc)
	glut.ReshapeFunc(Reshape)

	// Input Callbacks
	glut.SpecialFunc(SpecialFunc)
	glut.KeyboardFunc(KeyDownFunc)
	glut.MotionFunc(MouseMotion)
	glut.MouseFunc(MouseDown)

	glut.MainLoop()
}
