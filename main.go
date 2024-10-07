package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	updateVars()

	m := model{}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}

var fps = 120
var bouncey = true
var clampMinSpeed = true

var radius = 7.0
var maxSpeed = 0.5
var adjustRate = 0.025
var alignmentRate = 1.0
var cohesionRate = 1.0
var separationRate = 1.0
var targetMinSpeed = 0.01

func updateVars() {
	if err := godotenv.Overload(); err != nil {
		log.Fatal("Error loading .env file")
	}
	if v := os.Getenv("FPS"); v != "" {
		fps, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("BOUNCE"); v != "" {
		bouncey, _ = strconv.ParseBool(v)
	}
	if v := os.Getenv("CLAMP_MIN_SPEED"); v != "" {
		clampMinSpeed, _ = strconv.ParseBool(v)
	}
	if v := os.Getenv("RADIUS"); v != "" {
		radius, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("MAX_SPEED"); v != "" {
		maxSpeed, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("ADJUST_RATE"); v != "" {
		adjustRate, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("ALIGNMENT_RATE"); v != "" {
		alignmentRate, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("COHESION_RATE"); v != "" {
		cohesionRate, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("SEPARATION_RATE"); v != "" {
		separationRate, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("TARGET_MIN_SPEED"); v != "" {
		targetMinSpeed, _ = strconv.ParseFloat(v, 64)
	}
}
