package prj

import (
	"fmt"
	"os"
	"log"
	"path"
	"os/exec"
	"path/filepath"
	"encoding/json"
)

type WinProj struct {
	Success bool
	Out [][]float64
}

func Transform(frm, to int, coords [][]float64, fromGeog ...bool) [][]float64 {
	checkErr := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := getCmdPath()
	coordArgs, err := json.Marshal(coords)
	checkErr(err)

	args := []string{
		"-f", fmt.Sprintf("%v", frm),
		"-t", fmt.Sprintf("%v", to),
		"-i", string(coordArgs),
	}

	out, err := exec.Command(cmd, args...).Output()
	checkErr(err)

	res := &WinProj{}
	err = json.Unmarshal(out, res)
	checkErr(err)

	return  res.Out
}

func getCmdPath() string {
	dir := getExecDir()
	if !IsDir(dir) {
		log.Fatalln("missing proj directory with binaries in parent folder - os:windows")
	}
	cmd := filepath.Join(dir, "proj/main.exe")
	if !IsFile(cmd) {
		log.Fatalln("missing main.exe in proj directory - os:windows")
	}
	return cmd
}

func getExecDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	return path.Dir(ex)
}

func IsFile(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return info.Mode().IsRegular()
}

func IsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	bln := false
	if err == nil && info.Mode().IsDir() {
		bln = true
	}
	return bln
}
