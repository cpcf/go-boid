package main

// A simple example demonstrating how to draw and animate on a cellular grid.
// Note that the cellbuffer implementation in this example does not support
// double-width runes.

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	fps           = 120
	boidCount     = 250
	asterisk      = "*"
	bouncey       = true
	clampMinSpeed = true
)

type cellbuffer struct {
	cells  []string
	stride int
}

func (c *cellbuffer) init(w, h int) {
	if w == 0 {
		return
	}
	c.stride = w
	c.cells = make([]string, w*h)
	c.wipe()
}

func (c cellbuffer) set(x, y int) {
	i := y*c.stride + x
	if i > len(c.cells)-1 || x < 0 || y < 0 || x >= c.width() || y >= c.height() {
		return
	}
	c.cells[i] = asterisk
}

func (c *cellbuffer) wipe() {
	for i := range c.cells {
		c.cells[i] = " "
	}
}

func (c cellbuffer) width() int {
	return c.stride
}

func (c cellbuffer) height() int {
	h := len(c.cells) / c.stride
	if len(c.cells)%c.stride != 0 {
		h++
	}
	return h
}

func (c cellbuffer) ready() bool {
	return len(c.cells) > 0
}

func (c cellbuffer) String() string {
	var b strings.Builder
	for i := 0; i < len(c.cells); i++ {
		if i > 0 && i%c.stride == 0 && i < len(c.cells)-1 {
			b.WriteRune('\n')
		}
		b.WriteString(c.cells[i])
	}
	return b.String()
}

type frameMsg struct{}

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(_ time.Time) tea.Msg {
		return frameMsg{}
	})
}

type model struct {
	cells cellbuffer
	boids []boid
}

func (m model) Init() tea.Cmd {
	return animate()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		// if !m.cells.ready() {
		// 	m.targetX, m.targetY = float64(msg.Width)/2, float64(msg.Height)/2
		// }
		m.cells.init(msg.Width, msg.Height)
		m.boids = initRandomBoids(boidCount, msg.Width, msg.Height)
		return m, nil
	// case tea.MouseMsg:
	// 	if !m.cells.ready() {
	// 		return m, nil
	// 	}
	// 	m.targetX, m.targetY = float64(msg.X), float64(msg.Y)
	// 	return m, nil

	case frameMsg:
		if !m.cells.ready() {
			return m, nil
		}

		m.cells.wipe()
		m.updateBoids()
		return m, animate()
	default:
		return m, nil
	}
}

func (m model) View() string {
	return m.cells.String()
}

func main() {
	m := model{}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}

func initRandomBoids(count int, screenWidth, screenHeight int) []boid {
	maxSpeed := 2.0
	boids := make([]boid, count)
	for i := range boids {
		boids[i] = boid{
			pos: Point{
				x: rand.Float64() * float64(screenWidth),
				y: rand.Float64() * float64(screenHeight),
			},
			vel: Point{
				x: rand.Float64()*maxSpeed*2 - maxSpeed,
				y: rand.Float64()*maxSpeed*2 - maxSpeed,
			},
			maxX:          float64(screenWidth),
			maxY:          float64(screenHeight),
			bounce:        bouncey,
			clampMinSpeed: clampMinSpeed,
		}
	}
	return boids
}

func (m model) updateBoids() {
	var wg sync.WaitGroup
	for i := range m.boids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			m.boids[i].update(m.boids)
		}(i)
		wg.Wait()
		m.boids[i].move()
		drawEllipse(&m.cells, m.boids[i].pos.x, m.boids[i].pos.y, 1, 1)
	}
}

func drawEllipse(cb *cellbuffer, xc, yc, rx, ry float64) {
	var (
		dx, dy, d1, d2 float64
		x              float64
		y              = ry
	)

	d1 = ry*ry - rx*rx*ry + 0.25*rx*rx
	dx = 2 * ry * ry * x
	dy = 2 * rx * rx * y

	for dx < dy {
		cb.set(int(x+xc), int(y+yc))
		cb.set(int(-x+xc), int(y+yc))
		cb.set(int(x+xc), int(-y+yc))
		cb.set(int(-x+xc), int(-y+yc))
		if d1 < 0 {
			x++
			dx = dx + (2 * ry * ry)
			d1 = d1 + dx + (ry * ry)
		} else {
			x++
			y--
			dx = dx + (2 * ry * ry)
			dy = dy - (2 * rx * rx)
			d1 = d1 + dx - dy + (ry * ry)
		}
	}

	d2 = ((ry * ry) * ((x + 0.5) * (x + 0.5))) + ((rx * rx) * ((y - 1) * (y - 1))) - (rx * rx * ry * ry)

	for y >= 0 {
		cb.set(int(x+xc), int(y+yc))
		cb.set(int(-x+xc), int(y+yc))
		cb.set(int(x+xc), int(-y+yc))
		cb.set(int(-x+xc), int(-y+yc))
		if d2 > 0 {
			y--
			dy = dy - (2 * rx * rx)
			d2 = d2 + (rx * rx) - dy
		} else {
			y--
			x++
			dx = dx + (2 * ry * ry)
			dy = dy - (2 * rx * rx)
			d2 = d2 + dx - dy + (rx * rx)
		}
	}
}
