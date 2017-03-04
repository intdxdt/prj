package prj


const (
	p = iota
	g
)

type SRS struct {
	prj4  string
	_type int
	to    *SRS
}

