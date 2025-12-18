package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Point struct {
	X float64
	Y float64
}

func bezier(t float64, p0, p1, p2, p3 Point) Point {
	u := 1 - t
	tt := t * t
	uu := u * u
	uuu := uu * u
	ttt := tt * t

	x := uuu*p0.X + 3*uu*t*p1.X + 3*u*tt*p2.X + ttt*p3.X
	y := uuu*p0.Y + 3*uu*t*p1.Y + 3*u*tt*p2.Y + ttt*p3.Y

	return Point{X: x, Y: y}
}

func MoveMouseHuman(page *rod.Page, from, to Point) error {
	rand.Seed(time.Now().UnixNano())

	p1 := Point{
		X: from.X + rand.Float64()*120 - 60,
		Y: from.Y + rand.Float64()*120 - 60,
	}
	p2 := Point{
		X: to.X + rand.Float64()*120 - 60,
		Y: to.Y + rand.Float64()*120 - 60,
	}

	steps := rand.Intn(25) + 35

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		p := bezier(t, from, p1, p2, to)

		page.Mouse.MoveTo(proto.Point{X: p.X, Y: p.Y})
		time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond)
	}

	page.Mouse.MoveTo(proto.Point{
		X: to.X + rand.Float64()*2 - 1,
		Y: to.Y + rand.Float64()*2 - 1,
	})

	return nil
}
