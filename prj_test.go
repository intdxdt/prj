package prj

import (
    . "github.com/franela/goblin"
    "testing"
    "math"
)

func TestCart(t *testing.T) {
    g := Goblin(t)

    g.Describe("project point", func() {
        g.It("srs", func() {
            var srs = NewSRS(4326).AsGeographic()
            var srs_prjs = NewSRS(3857).AsProjected()

            g.Assert(srs.isGeog()).IsTrue()
            g.Assert(srs.isProj()).IsFalse()

            g.Assert(srs_prjs.isProj()).IsTrue()
            g.Assert(srs_prjs.isGeog()).IsFalse()

        })

        g.It("transformation x,y", func() {
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