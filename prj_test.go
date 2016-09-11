package prj

import (
	. "github.com/franela/goblin"
	"testing"
	"math"
)

func TestCart(t *testing.T) {
	g := Goblin(t)

	g.Describe("geom.point", func() {
		g.It("x, y access & null", func() {
			lng, lat := 28.00372, 40.81747

			x, y, err := NewSRS(4326).AsGeographic().To(
				NewSRS(3857).AsProjected(),
			).Trans(lng, lat)

			g.Assert(err == nil).IsTrue()
			g.Assert(math.Abs(x - 3117359.85071741) < 1.0e-8).IsTrue()
			g.Assert(math.Abs(y - 4985455.69596696) < 1.0e-8).IsTrue()
		})

	})
}