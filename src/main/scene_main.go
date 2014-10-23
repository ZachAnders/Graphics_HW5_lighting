package main

/*
* CSCI 4229 Assignment 5: Lighting
*
* Author: Zach Anders
*
* Some code derived from CSCI 4229 Examples 9, 10, and 13 (ex9.c, ex10.c, ex13.c)
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

	"glutil"
	"util"
	"world"

	"github.com/go-gl/gl"
	"github.com/ianremmler/ode"
	"github.com/rhencke/glut"
)

var renderQueue util.RenderQueue = util.NewEmptyRenderQueue()

//var currentCamera = CreateDefaultViewport()
var my_world = world.NewWorld()
var player = actor.NewPlayer(&my_world, glutil.Point3D{0, 0, 0}, 5)
var light actor.OrbitActor
var currentMouse = glutil.CreateMouseState()

// Normal order:
//Translate -> Rotate -> Scale

// Creates the display function
func DisplayFunc() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.Enable(gl.DEPTH_TEST)
	gl.LoadIdentity()

	// currentCamera.PositionSelf()
	player.PositionSelf()

	my_world.Tick()
	renderQueue.RenderAll()

	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.Color3f(.9, .9, .9)
	glutil.Print2d(5, 35, "Keys: W,S,A,D : Move Around | Q / E : Move Light left/right | L : Toggle light rotation")
	glutil.Print2d(5, 20, "Click and Drag: Rotate | Arrow keys: Rotate")
	glutil.Print2d(5, 5, "%s", player.ToString())

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
	case 'q':
		light.AdjustAngle(-5)
		break
	case 'e':
		light.AdjustAngle(5)
		break
	case 'l':
		light.Toggle()
		break
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
	if e_time < 33 {
		time.Sleep(time.Duration(33-e_time) * time.Millisecond)
	}
}

func setup() {

	lightsource := actor.NewBallLight(&my_world, glutil.Point3D{15, 15, 15})
	renderQueue.AddNamed(&lightsource, "light")
	light = actor.NewOrbitActor(&my_world, &lightsource, glutil.Point3D{0, 15, 0}, 50)
	my_world.AddActor(&light)

	rand.Seed(0x12345 + 8)
	my_world.AddActor(&player)
	block := actor.NewBlock(&my_world, glutil.Point3D{22, 0, 0}, glutil.Point3D{10, 20.5, 20}, glutil.Color4D{.5, .3, .1, 1})
	my_world.AddActor(&block)
	renderQueue.Add(&block)

	tmpList := []actor.OrbitActor{}
	for i := 0; i < 250; i++ {
		speed := rand.Intn(3)
		//radius := rand.Intn(50) + 125
		radius := 100
		x, y, z := float64(rand.Intn(10)-5), float64(rand.Intn(10)-5), float64(rand.Intn(150)-75)
		angle := float64(rand.Intn(360))
		//cylinder := NewCylinder(&my_world, glutil.Point3D{x, y + 30, z}, 1, 4)
		cylinder := actor.NewBlock(&my_world, glutil.Point3D{x, y + 30, z}, glutil.Point3D{.1, .1, .1}, glutil.Color4D{.8, .8, .8, .05})
		tmpList = append(tmpList, actor.NewOrbitActor(&my_world, &cylinder, glutil.Point3D{x, y + 0, z}, float64(radius)))
		tmpList[i].SetSpeed(int32(speed))
		tmpList[i].SetAngle(int32(angle))
		tmpList[i].SetAxis(glutil.Point3D{0, 0, 1})
		renderQueue.Add(&cylinder)
		my_world.AddActor(&tmpList[i])
	}

	// Create castle
	//myCastle := NewCastle()
	//renderQueue.AddNamed(&myCastle, "castle")

	// Create ground
	ground := entity.NewXYPlane(glutil.Point3D{100, 0, 100}, glutil.Point3D{-100, 0, -100}, 0)
	ground.SetColor(glutil.Color4D{0.4313725490 * .7, 0.24705882 * .7, 0.098039215 * .7, 1})
	ground.SetPolygonOffset(1.0)
	renderQueue.Add(&ground)
	my_world.AddEntity(&ground.Box)

	stairs := actor.NewStairs(&my_world, glutil.Point3D{0, 0, 0})
	renderQueue.Add(&stairs)

	scale := float64(200)
	for i := 0; i < 100; i++ {
		x := (rand.Float64() * scale) - (scale / 2)
		z := (rand.Float64() * scale) - (scale / 2)
		x_scale, z_scale := (rand.Float64()/2)+.75, (rand.Float64()/2)+.75
		y_scale := (rand.Float64() * 1) + .75
		if (x < -5 || x > 20) && (z < -5 || z > 20) {
			tree1 := entity.NewTree(glutil.Point3D{x, 0, z}, x_scale, y_scale, z_scale)
			renderQueue.Add(&tree1)
		}
	}

	// Put camera somewhere
	//currentCamera.Translate(5, -5, 28)
	player.Translate(10, 0, -20)
	player.ImmediateJump() // Skip interpolation
	player.Rotate(-120, 0)
	player.ImmediateLook() // Skip interpolation
}

func main() {
	// Let go use 2 threads
	runtime.GOMAXPROCS(2)

	// Init
	glut.InitDisplayMode(glut.RGB | glut.DOUBLE | glut.DEPTH)
	glut.InitWindowSize(750, 750)
	glut.CreateWindow("Zach Anders - Assignment 5")
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
