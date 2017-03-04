package prj

import (
	"log"
	"github.com/pebbe/go-proj-4/proj"
)

func Transform(frm, to int, coords [][]float64, fromGeog ...bool) [][]float64 {
	checkErr := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	frmprj4, ok := srsDB[frm]
	if !ok {
		log.Fatalln("src srs not found")
	}

	srcPrj, err := proj.NewProj(frmprj4)
	defer srcPrj.Close()
	checkErr(err)

	toprj4, ok := srsDB[to]
	if !ok {
		log.Fatalln("dest srs not found")
	}

	destPrj, err := proj.NewProj(toprj4)
	defer destPrj.Close()
	checkErr(err)

	isgeog := false
	if len(fromGeog) > 0 {
		isgeog = fromGeog[0]
	}
	results := make([][]float64, 0)
	for _, xy := range coords {
		x, y := xy[0], xy[1]
		if isgeog {
			x, y = proj.DegToRad(x), proj.DegToRad(y)
		}
		ox, oy, err := proj.Transform2(srcPrj, destPrj, x, y)
		checkErr(err)
		results = append(results, []float64{ox, oy})
	}

	return results
}
