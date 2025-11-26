package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

var (
	colBg       = color.RGBA{0xFA, 0xFB, 0xFC, 0xFF} // Very faint cool grey background
	colText     = color.RGBA{0x37, 0x35, 0x2F, 0xFF} // Text color
	colBtn      = color.RGBA{0xE1, 0xE4, 0xE8, 0xFF} // Neutral border/fill
	colBtnHover = color.RGBA{0xD0, 0xD7, 0xDE, 0xFF} // Darker hover
	colAccent   = color.RGBA{0x09, 0x69, 0xDA, 0xFF} // Subtle blue
)

type Game struct {
	sw       *Stopwatch
	fontFace font.Face
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= 20 && x <= 110 && y >= 100 && y <= 140 {
			if g.sw.state == StateRunning {
				g.sw.Pause()
			} else {
				g.sw.Start()
			}
		}

		if x >= 130 && x <= 220 && y >= 100 && y <= 140 {
			g.sw.Reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colBg)

	elapsed := g.sw.Elapsed()
	ms := int(elapsed.Milliseconds()) % 1000 / 10
	timeStr := fmt.Sprintf("%02d:%02d:%02d,%02d",
		int(elapsed.Hours()),
		int(elapsed.Minutes())%60,
		int(elapsed.Seconds())%60,
		ms)

	text.Draw(screen, timeStr, g.fontFace, 94, 70, colText)

	label := "Start"
	if g.sw.state == StateRunning {
		label = "Stop"
	}

	DrawButton(screen, 20, 100, 90, 40, label, g.fontFace)
	DrawButton(screen, 130, 100, 90, 40, "Restart", g.fontFace)
}

func DrawButton(screen *ebiten.Image, x, y, w, h int, label string, face font.Face) {
	mx, my := ebiten.CursorPosition()

	isHovered := mx >= x && mx <= x+w && my >= y && my <= y+h

	c := colBtn
	if isHovered {
		c = colBtnHover
	}

	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(w), float64(h), c)

	textX := x + (w-(len(label)*10))/2
	text.Draw(screen, label, face, textX, y+28, colText)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 200
}

func main() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
		return
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	game := &Game{
		sw:       NewStopwatch(),
		fontFace: face,
	}

	ebiten.SetWindowSize(300, 200)
	ebiten.SetWindowTitle("SaintWatch")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
