package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	// delete all files from slice
	if !printFiles {
		for i, rcount, rlen := 0, 0, len(files); i < rlen; i++ {
			j := i - rcount
			if files[j].IsDir() == false {
				files = append(files[:j], files[j+1:]...)
				rcount++
			}
		}
	}

	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	lenfiles := len(files)
	var precond, subprecond string
	for _, f := range files {
		lenfiles--
		if lenfiles == 0 {
			precond = "└───"
			subprecond = "	"
		} else {
			precond = "├───"
			subprecond = "│	"
		}

		if f.IsDir() {
			fmt.Fprintf(out, precond+f.Name()+"\n")
			subout := new(bytes.Buffer)

			subpath := path + "/" + f.Name()
			suberr := dirTree(subout, subpath, printFiles)
			if suberr != nil {
				panic(suberr.Error())
			}

			s := strings.Split(subout.String(), "\n")
			for si := range s {
				if si < len(s)-1 {
					fmt.Fprintf(out, subprecond+s[si]+"\n")
				}
			}

		} else {
			if printFiles {
				var mysize string
				if f.Size() == 0 {
					mysize = " (empty)"
				} else {
					mysize = " (" + strconv.FormatInt(f.Size(), 10) + "b)"
				}
				fmt.Fprintf(out, precond+f.Name()+mysize+"\n")
			}
		}
	}
	return nil
}

// $ go run main.go ./testdata
// $ go run main.go ./testdata/zline/lorem/ -f
// $ go run main.go ./testdata -f

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(out)
}
