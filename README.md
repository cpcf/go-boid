# go-boids

A go implementation of the boids simulation described by Craig Reynolds in 1986.

https://dl.acm.org/doi/10.1145/37402.37406

## Output
Runs in the terminal, using the https://github.com/charmbracelet/bubbletea framework to handle terminal output.

Scales the number of boids to the terminal size. Works best on larger terminals. Very large terminal sizes may cause performance issues.

![](https://github.com/cpcf/go-boids/blob/main/go-boids.gif)

## Installation

```bash
go get github.com/cpcf/go-boids
```

## Usage

```bash
go-boids
```

## Parameters
Currently hard coded in main.go and boid.go
