# go-boid

A go implementation of the boids simulation described by Craig Reynolds in 1986.

https://dl.acm.org/doi/10.1145/37402.37406

## Output
Runs in the terminal, using the https://github.com/charmbracelet/bubbletea framework to handle terminal output.

Scales the number of boids to the terminal size. Works best on larger terminals. Very large terminal sizes may cause performance issues.

![](https://github.com/cpcf/go-boid/blob/main/go-boid.gif)

## Installation

```bash
go get github.com/go-boid/go-boid
```

## Usage

```bash
go-boid
```

## Parameters
Currently hard coded in main.go and boid.go