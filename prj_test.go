package prj

import (
    "github.com/franela/goblin"
    "testing"
    "math"
)

func TestCart(t *testing.T) {
    g := goblin.Goblin(t)

	g.Describe("geom.point", func() {
		g.It("x, y access & null", func() {
			coords :=[][]float64{{28.00372, 40.81747}}

			res := Transform(4326,3857, coords, true)
			x, y := res[0][0], res[0][1]
			g.Assert(math.Abs(x - 3117359.85071741) < 1.0e-8).IsTrue()
			g.Assert(math.Abs(y - 4985455.69596696) < 1.0e-8).IsTrue()
		})

    })
}