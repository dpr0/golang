package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {panic("usage go run main.go . [-f]")}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {panic(err.Error())}
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) (err error) {
	i := 0
	out.WriteString(printDir(path, i, printFiles, false, false))
    return err
}

func printDir(path string, i int, printFiles bool, last2 bool, last3 bool) (str string) {
	files, _ := os.ReadDir(path)
	var array [] fs.DirEntry
	for _, n := range files {
		if n.IsDir() || printFiles && n.Name() != ".DS_Store" {array = append(array, n)}
	}

	for idx, n := range array {
		last := len(array)-1 == idx

		var pre string
		if last { pre = "└───" } else { pre = "├───" }
		for j := 0; j <= i; j++ {
			var z string
			if j == 1 && last3 || j == 0 && last2 { z = "\t" } else { z = "│\t" }
			if j < i { pre = z + pre }
		}
		size := ""
		if !n.IsDir() {
			s, _ := n.Info()
			if s.Size() == 0 { size = " (empty)" } else { size = fmt.Sprintf(" (%vb)", s.Size()) }
		}
		str = str + pre + n.Name() + size + "\n" + printDir(path+"/"+n.Name(), i+1, printFiles, last, last2)
	}
	return str
}
