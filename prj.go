package prj

import (
	"github.com/pebbe/go-proj-4/proj"
	"log"
)

const (
	p = iota
	g
)

type SRS struct {
	prj4  string
	_type int
	to    *SRS
}

func (self *SRS) AsProjected() *SRS {
	self._type = p
	return self
}

func (self *SRS) AsGeographic() *SRS {
	self._type = g
	return self
}

func (self *SRS) isGeog() bool {
	return self._type == g
}

func (self *SRS) isProj() bool {
	return self._type == p
}


func (self *SRS) To(dest *SRS) *SRS {
	self.to = dest
	return self
}

func (self *SRS) Trans(x, y float64) (float64, float64, error) {
	if self.to == nil {
		log.Fatalln("destination srs not set ")
	}
	srcPrj, err := proj.NewProj(self.prj4)
	defer srcPrj.Close()
	if err != nil {
		log.Fatal(err)
	}

	destPrj, err := proj.NewProj(self.to.prj4)
	defer destPrj.Close()
	if err != nil {
		log.Fatal(err)
	}

	if self.isGeog() {
		return proj.Transform2(srcPrj, destPrj, proj.DegToRad(x), proj.DegToRad(y))
	}
	return proj.Transform2(srcPrj, destPrj, x, y)
}


func NewSRS(srid int) *SRS {
	var prj4 = prjDB[srid]
	if prj4 == "" {
		log.Fatalln("srid not found")
	}
	srs := &SRS{prj4:prj4}
	return srs.AsGeographic()
}

