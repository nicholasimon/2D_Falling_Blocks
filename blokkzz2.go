package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var (
	menuon                                                                 bool
	introtextx                                                             int
	introfadeon                                                            bool
	introfade                                                              = float32(0.2)
	introblockcolors                                                       = make([]rl.Color, 3800)
	introscreenon                                                          bool
	backtype                                                               int
	backtimer                                                              int
	backpolyz                                                              = make([]poly, 100)
	backcircles                                                            = make([]circle, 100)
	linetextylist                                                          = make([]int, 12)
	linetexty                                                              = monh / 6
	linetextlist                                                           = make([]int, 12)
	linetextcount                                                          int
	diamond4count, line2count, line3count, line4count                      int
	color1, color2, color3, color4, color5, color6, color7, color8, color9 rl.Color
	bloklayout                                                             = make([]blok, bloka)
	blok9                                                                  = make([]blok, 9)
	blok9holder                                                            = make([]blok, 9)
	blokw                                                                  = 22
	blokh                                                                  = 45
	bloka                                                                  = blokw * blokh
	monw                                                                   = 1280
	monh                                                                   = 720
	fps                                                                    = 60
	framecount                                                             int
	nextblockon, pauseon, numberson, debugon, gridon, centerlineson        bool
	onoff2, onoff3, onoff6, onoff10, onoff15, onoff30                      bool
	imgs                                                                   rl.Texture2D
	camera, camera2x, camera4x, camera8x                                   rl.Camera2D
)

type circle struct {
	x, y    int
	r, fade float32
	color   rl.Color
}

type blok struct {
	movecount, position, y, x int
	color                     rl.Color
	moveon, filled            bool
}
type poly struct {
	v2                     rl.Vector2
	sides                  int
	radius, rotation, fade float32
	color                  rl.Color
}

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "blokkzz")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(fps)
	rl.HideCursor()
	//rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {

		framecount++

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawnocamerabackground()

		rl.BeginMode2D(camera)
		drawlayers()
		rl.EndMode2D()

		drawnocamera()
		if introscreenon {
			drawintro()
		}
		if menuon {
			drawmenu()
		}
		update()

		rl.EndDrawing()

	}
	rl.CloseWindow()
}
func drawlayers() { // MARK: drawlayers

	for a := 0; a < len(blok9); a++ {
		rl.DrawRectangle(blok9[a].x, blok9[a].y, 15, 15, rl.White)
		rl.DrawRectangle(blok9[a].x, blok9[a].y, 15, 15, blok9[a].color)
	}
	for a := 0; a < len(bloklayout); a++ {
		if bloklayout[a].filled {
			rl.DrawRectangle(bloklayout[a].x, bloklayout[a].y, 15, 15, bloklayout[a].color)
		}
	}
}
func drawnocamerabackground() { // MARK: drawnocamerabackground

	switch backtype {
	case 2:
		// back polyz
		for a := 0; a < len(backpolyz); a++ {

			rl.DrawPoly(backpolyz[a].v2, backpolyz[a].sides, backpolyz[a].radius, backpolyz[a].rotation, rl.Fade(backpolyz[a].color, backpolyz[a].fade))
			//rl.DrawCircle(backcircles[a].x, backcircles[a].y, backcircles[a].r, rl.Fade(backcircles[a].color, backcircles[a].fade))
			backpolyz[a].radius -= 0.1
		}
		backtimer++
		if backtimer == 900 {
			createpolyz()
			backtimer = 0
		}
	case 1:
		// back circles
		for a := 0; a < len(backcircles); a++ {
			rl.DrawCircle(backcircles[a].x, backcircles[a].y, backcircles[a].r, rl.Fade(backcircles[a].color, backcircles[a].fade))
			backcircles[a].r -= 0.1
		}
		backtimer++
		if backtimer == 300 {
			createcircles()
			backtimer = 0
		}
	}
	// borders
	rl.DrawRectangle(470, 0, 2, monh, rl.Fade(rl.Green, 0.2))
	rl.DrawRectangle(808, 0, 2, monh, rl.Fade(rl.Green, 0.2))
	// background rec
	rl.DrawRectangle(472, 0, 336, monh, rl.Fade(rl.Black, 0.8))

}
func createcircles() { // MARK: createcircles
	for a := 0; a < len(backcircles); a++ {
		backcircles[a].fade = rF32(0.2, 0.6)
		backcircles[a].r = rFloat32(8, 51)
		backcircles[a].x = rInt(-10, monw+10)
		backcircles[a].y = rInt(-10, monh+10)
		backcircles[a].color = randomcolor()
	}
}
func createpolyz() { // MARK: createpolyz
	for a := 0; a < len(backpolyz); a++ {
		backpolyz[a].color = randomcolor()
		backpolyz[a].fade = rF32(0.2, 0.6)
		backpolyz[a].rotation = rFloat32(0, 360)
		backpolyz[a].radius = rFloat32(8, 51)
		backpolyz[a].sides = rInt(3, 9)
		x := rFloat32(-10, monw+10)
		y := rFloat32(-10, monh+10)
		backpolyz[a].v2 = rl.NewVector2(x, y)
	}
}

func drawnocamera() { // MARK: drawnocamera

	line2counttext := strconv.Itoa(line2count)
	line3counttext := strconv.Itoa(line3count)
	line4counttext := strconv.Itoa(line4count)
	diamond4counttext := strconv.Itoa(diamond4count)

	rl.DrawText(line2counttext, 10, 20, 20, rl.White)
	rl.DrawText("line2counttext", 50, 20, 20, rl.White)
	rl.DrawText(line3counttext, 10, 40, 20, rl.White)
	rl.DrawText("line3counttext", 50, 40, 20, rl.White)
	rl.DrawText(line4counttext, 10, 60, 20, rl.White)
	rl.DrawText("line4counttext", 50, 60, 20, rl.White)
	rl.DrawText(diamond4counttext, 10, 80, 20, rl.White)
	rl.DrawText("diamond4counttext", 50, 80, 20, rl.White)

	// numbers
	if numberson {
		x := 474
		y := 4
		count := 0
		for a := 0; a < bloka; a++ {
			rl.DrawText(strconv.Itoa(a), x, y, 10, rl.Fade(rl.Green, 0.4))
			x += 16
			count++
			if count == blokw {
				count = 0
				x = 474
				y += 16
			}

		}

	}
	// grid
	if gridon {
		y := 0
		x := 472

		for {
			rl.DrawLine(x, 0, x, monh, rl.Fade(rl.Green, 0.2))
			x += 16

			if x >= 808 {
				break
			}
		}
		x = 472
		for {
			rl.DrawLine(x, y, x+336, y, rl.Fade(rl.Green, 0.2))
			y += 16

			if y >= monh {
				break
			}
		}
	}
	// centerlines
	if centerlineson {
		rl.DrawLine(monw/2, 0, monw/2, monh, rl.Fade(rl.Green, 0.2))
		rl.DrawLine(0, monh/2, monw, monh/2, rl.Fade(rl.Green, 0.2))
	}
}

func update() { // MARK: update

	if !pauseon {
		linetext()
		if nextblockon {
			createblok()
			nextblockon = false

		}
		gravity()
		timers()
		updatebloks()

	}
	input()

	if debugon {
		debug()
	}
}
func updatebloks() {

	// find similar
	for a := 0; a < len(bloklayout)-(blokw+1); a++ {
		if bloklayout[a].filled {
			//right
			if bloklayout[a].color == bloklayout[a+1].color {
				count := 1
				for {
					if bloklayout[a+count].color == bloklayout[a+count+1].color {
						count++
					} else {
						break
					}
				}
				if count == 1 {
					line2count++
				} else if count == 2 {
					line3count++
				} else if count == 3 {
					line4count++
				}
				count++

				linetextlist[linetextcount] = count
				linetextcount++
				if linetextcount == len(linetextlist) {
					linetextcount = 0
				}

				for b := 0; b < count; b++ {
					bloklayout[a+b] = blok{}
				}
			}
			// above
			if bloklayout[a].color == bloklayout[a-blokw].color {
				count := 1
				for {
					if bloklayout[a-(count*blokw)].color == bloklayout[a-((count+1)*blokw)].color {
						count++
					} else {
						break
					}
				}

				if count == 1 {
					line2count++
				} else if count == 2 {
					line3count++
				} else if count == 3 {
					line4count++
				}

				count++
				for b := 0; b < count; b++ {
					bloklayout[a-(b*blokw)] = blok{}
				}
			}
			// diamond 4 block
			if bloklayout[a].color == bloklayout[(a+blokw)+1].color && bloklayout[(a+blokw)+1].color == bloklayout[(a+blokw)-1].color && bloklayout[a].color == bloklayout[a+(blokw*2)].color {
				bloklayout[a] = blok{}
				bloklayout[(a+blokw)+1] = blok{}
				bloklayout[(a+blokw)-1] = blok{}
				bloklayout[a+(blokw*2)] = blok{}
				diamond4count++
			}

		}
	}
	// move down empty space
	for a := len(bloklayout) - (blokw + 1); a > 0; a-- {
		if bloklayout[a].filled {
			if !bloklayout[a+blokw].filled {
				bloklayout[a].moveon = true
			}
		}
	}
	for a := len(bloklayout) - (blokw + 1); a > 0; a-- {
		if bloklayout[a].moveon && bloklayout[a].movecount < 4 && bloklayout[a].y < monh-16 {
			bloklayout[a].y += 4
			bloklayout[a].movecount++
		} else if bloklayout[a].moveon && bloklayout[a].movecount >= 3 {
			bloklayout[a].moveon = false
			bloklayout[a].movecount = 0
			bloklayout[a+blokw] = bloklayout[a]
			bloklayout[a] = blok{}
		}
	}

	// next block
	nextblock := true
	for a := 0; a < len(blok9); a++ {
		if blok9[a].moveon == true {
			nextblock = false
		}
	}
	if nextblock {
		nextblockon = true
	}

}
func drawintro() { // MARK: drawintro

	rl.DrawRectangle(0, 0, monw, monh, rl.Black)
	count := 0
	for a := 0; a < monw; a++ {
		y := 0
		for b := 0; b < 45; b++ {
			rl.DrawRectangle(a, y, 15, 15, rl.Fade(introblockcolors[count], introfade))
			y += 16
			count++
		}
		a += 16
	}

	if introfadeon {
		introfade -= 0.01
	} else {
		introfade += 0.01
	}

	if introfade >= 1.0 {
		introfadeon = true
	} else if introfade <= 0.2 {
		introfadeon = false
	}

	rl.DrawRectangle(0, monh/2-50, monw, 100, rl.Black)

	rl.DrawText("press space to start...", introtextx-600, monh/2-16, 40, rl.White)
	rl.DrawText("blokkz", introtextx, monh/2-36, 80, rl.White)
	rl.DrawText("blokkz", introtextx+400, monh/2-36, 80, rl.White)
	rl.DrawText("blokkz", introtextx+800, monh/2-36, 80, rl.White)
	liney := monh/2 - 34
	for a := 0; a < 25; a++ {
		rl.DrawLine(0, liney, monw, liney, rl.Fade(rl.Black, 0.4))
		liney += 3
	}

	introtextx += 4
	if introtextx > monw+650 {
		introtextx = -1150
	}

	rl.DrawRectangle(0, monh-50, monw, 50, rl.Black)
	rl.DrawText("a game by nicholasimon © 2021", monw/2-152, monh-33, 20, randomcolor())
	rl.DrawText("a game by nicholasimon © 2021", monw/2-151, monh-34, 20, rl.Black)
	rl.DrawText("a game by nicholasimon © 2021", monw/2-150, monh-35, 20, rl.White)

	if rl.IsKeyPressed(rl.KeySpace) {
		introscreenon = false
		pauseon = false
	}

}
func drawmenu() { // MARK: drawmenu

	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Black, 0.7))

}
func linetext() { // MARK: main

	rl.BeginMode2D(camera4x)

	for a := 0; a < len(linetextlist); a++ {

		if linetextlist[a] != 0 {
			text := strconv.Itoa(linetextlist[a])
			rl.DrawText(text, 222, linetextylist[a]+2, 40, randomcolor())
			rl.DrawText(text, 221, linetextylist[a]+1, 40, rl.Black)
			rl.DrawText(text, 220, linetextylist[a], 40, rl.White)

		}

	}

	rl.EndMode2D()

	if onoff3 {
		for a := 0; a < len(linetextylist); a++ {
			linetextylist[a]--
		}
	}
	if linetextylist[0] <= -144 {
		y := monh / 5
		for a := 0; a < len(linetextylist); a++ {
			linetextylist[a] = y
			y -= 40

		}
	}

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "setscreen")
	setscreen()
	rl.CloseWindow()
	setinitialvalues()
	raylib()
}
func timers() { // MARK: timers

	if framecount%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if framecount%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if framecount%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if framecount%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if framecount%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if framecount%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}

}
func gravity() { // MARK: gravity

	for a := 0; a < len(blok9); a++ {
		if blok9[a].moveon {
			if bloklayout[blok9[a].position+blokw].filled {
				storeblok(blok9[a], a)
			} else {
				if blok9[a].y < (monh - 16) {
					blok9[a].y += 4
					blok9[a].position = getposition(blok9[a].x, blok9[a].y)
				} else {
					storeblok(blok9[a], a)
				}
			}
		}
	}
}
func rotateleft() { // MARK: rotateleft

	blok9holder = blok9

	blok9[0] = blok9holder[8]
	blok9[1] = blok9holder[0]
	blok9[2] = blok9holder[1]
	blok9[3] = blok9holder[2]
	blok9[4] = blok9holder[3]
	blok9[5] = blok9holder[5]
	blok9[6] = blok9holder[5]
	blok9[7] = blok9holder[6]
	blok9[8] = blok9holder[7]

}
func storeblok(blok2store blok, blok9position int) { // MARK: storeblok
	blok2store.filled = true
	blok2store.moveon = false
	if blok2store.y%16 != 0 {
		for {
			blok2store.y--
			if blok2store.y%16 == 0 {
				break
			}
		}

	}

	bloklayout[blok2store.position] = blok2store
	blok9[blok9position] = blok{}
}
func getposition(x, y int) int { // MARK: getposition

	xpos := 0
	ypos := 0
	for a := 0; a < blokw; a++ {
		if x >= (a*16)+472 && x < (a*16)+16+472 {
			xpos = a
		}
	}
	for a := 0; a < blokh; a++ {
		if y >= a*16 && y < (a*16)+16 {
			ypos = a
		}
	}

	position := (ypos * blokw) + xpos
	return position

}
func createblok() { // MARK: createblok

	// randomize blok colors
	colors := []rl.Color{color1, color2, color3, color4, color5, color6, color7, color8, color9}
	rand.Shuffle(len(colors), func(i, j int) { colors[i], colors[j] = colors[j], colors[i] })

	for a := 0; a < len(blok9); a++ {
		blok9[a].color = colors[rInt(0, 9)]
		blok9[a].moveon = true
	}

	xchange := rInt(0, 18) * 16

	blok9[0].y = 0
	blok9[0].x = 488 + xchange
	blok9[1].y = 0
	blok9[1].x = 504 + xchange
	blok9[2].y = 0
	blok9[2].x = 520 + xchange

	blok9[3].y = 16
	blok9[3].x = 488 + xchange
	blok9[4].y = 16
	blok9[4].x = 504 + xchange
	blok9[5].y = 16
	blok9[5].x = 520 + xchange

	blok9[6].y = 32
	blok9[6].x = 488 + xchange
	blok9[7].y = 32
	blok9[7].x = 504 + xchange
	blok9[8].y = 32
	blok9[8].x = 520 + xchange

}
func moveright() { // MARK: moveright

	for a := 0; a < len(blok9); a++ {
		if blok9[8].x < 792 && blok9[8].moveon && blok9[7].moveon && blok9[6].moveon {
			blok9[a].x += 16
			blok9[a].position = getposition(blok9[a].x, blok9[a].y)
		}
	}

}
func moveleft() { // MARK: moveleft
	for a := 0; a < len(blok9); a++ {
		if blok9[8].x > 504 && blok9[8].moveon && blok9[7].moveon && blok9[6].moveon {
			blok9[a].x -= 16
			blok9[a].position = getposition(blok9[a].x, blok9[a].y)
		}

	}
}
func input() { // MARK: input

	if rl.IsKeyPressed(rl.KeyEscape) {

		if menuon {
			menuon = false
			pauseon = false

		} else {
			menuon = true
			pauseon = true

		}
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		if introscreenon {
			introscreenon = false
			pauseon = false
		} else {
			introscreenon = true
			pauseon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		rotateleft()
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		moveleft()
	}

	if rl.IsKeyDown(rl.KeyRight) {
		moveright()

	}

	if rl.IsKeyPressed(rl.KeyKp0) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDivide) {
		if numberson {
			numberson = false
		} else {
			numberson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if centerlineson {
			centerlineson = false
		} else {
			centerlineson = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 4.0
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 4.0 {
			camera.Zoom = 3.0
		}
	}
	if rl.IsKeyPressed(rl.KeyPause) {
		if pauseon {
			pauseon = false
		} else {
			pauseon = true
		}
	}

}
func debug() { // MARK: debug
	rl.DrawRectangle(monw-300, 0, 500, monw, rl.Fade(rl.Black, 0.8))
	rl.DrawFPS(monw-290, monh-100)

	xtext := strconv.Itoa(blok9[0].x)
	positiontext := strconv.Itoa(blok9[0].position)

	rl.DrawText(xtext, monw-290, 10, 10, rl.White)
	rl.DrawText("blok9[0].x", monw-150, 10, 10, rl.White)
	rl.DrawText(positiontext, monw-290, 20, 10, rl.White)
	rl.DrawText("blok9[0].position", monw-150, 20, 10, rl.White)

}
func createcolors() { // MARK: createcolors

	color1 = randomcolor()
	color2 = randomcolor()
	color3 = randomcolor()
	color4 = randomcolor()
	color5 = randomcolor()
	color6 = randomcolor()
	color7 = randomcolor()
	color8 = randomcolor()
	color9 = randomcolor()

}
func setinitialvalues() { // MARK: setinitialvalues

	for a := 0; a < len(introblockcolors); a++ {
		introblockcolors[a] = randomcolor()
	}

	backtype = rInt(1, 3)
	createcircles()
	createpolyz()
	y := monh / 5
	for a := 0; a < len(linetextylist); a++ {
		linetextylist[a] = y
		y -= 30

	}

	createcolors()
	createblok()

	// fill all color slots
	for a := 0; a < len(bloklayout); a++ {
		bloklayout[a].color = rl.White
	}

	// bottom blok border
	for a := bloka - 1; a > (bloka-blokw)-1; a-- {
		bloklayout[a].filled = true
	}

}

func setscreen() { // MARK: setscreen

	rl.SetWindowSize(monw, monh)

	camera.Zoom = 1.0
	camera.Target.X = 0
	camera.Target.Y = 0

	camera2x.Zoom = 2.0
	camera4x.Zoom = 4.0
	camera8x.Zoom = 8.0

	camera4x.Target.X = 0
	camera4x.Target.Y = 0
}

// random colors https://www.rapidtables.com/web/color/RGB_Color.html
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(105, 192)), uint8(rInt(105, 192)), uint8(rInt(105, 192)), 255)
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}

// random numbers
func rF32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
