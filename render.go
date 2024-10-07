package main

import (
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	fps           = 120
	bouncey       = true
	clampMinSpeed = true
)

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
		m.cells.init(msg.Width, msg.Height)
		m.boids = initBoidsOnScreenSize(msg.Width, msg.Height)
		return m, nil
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
		drawTriangle(&m.cells, m.boids[i].pos, m.boids[i].forward)
	}
}

func drawTriangle(cb *cellbuffer, centre, dir Point) {
	cb.set(int(centre.x), int(centre.y), triangleRuneTable[dir])
}

var triangleRuneTable = map[Point]string{
	{-1, -1}: "◤",
	{-1, 0}:  "◀",
	{-1, 1}:  "◣",
	{0, 1}:   "▼",
	{1, 1}:   "◢",
	{1, 0}:   "▶",
	{1, -1}:  "◥",
	{0, -1}:  "▲",
}
