package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		iterations, _ := strconv.Atoi(c.QueryParam("iterations"))
		if iterations < 1 {
			iterations = 1
		}

		return c.String(http.StatusOK, fmt.Sprintf("%f", pi(iterations)))
	})

	e.Logger.Fatal(e.Start(":3000"))
}

// Taken from https://go.dev/doc/play/pi.go
// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k < n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
