package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RecurseDownFS(path string) int {

	movie_count := 0

	dirEntries, err1 := os.ReadDir(path)
	checkError(err1)

	if dirEntries == nil || len(dirEntries) == 0 {
		return 0
	}

	for i := range dirEntries {
		fmt.Println(dirEntries[i].Name(), "TYPE= ", dirEntries[i].Type())
		fmt.Println("CURRENT DIR== " + path + "/" + dirEntries[i].Name())

		info, err2 := dirEntries[i].Info()
		checkError(err2)

		if info.IsDir() {
			movie_count = movie_count + RecurseDownFS(path+"/"+dirEntries[i].Name())
		}

		if strings.Contains(info.Name(), string(".mp4")) ||
			strings.Contains(info.Name(), string(".mkv")) {
			movie_count++
		}
	}

	return movie_count
}

func main() {
	fmt.Println("hello world")

	fs := os.DirFS("/")
	fmt.Println(fs)

	dirEntries, err1 := os.ReadDir("E:/movies")
	checkError(err1)

	for i := range dirEntries {
		fmt.Println(dirEntries[i].Name(), "TYPE= ", dirEntries[i].Type())
	}

	fmt.Println("===========================================START=============================================")
	fmt.Println("===========================================START=============================================")
	fmt.Println("===========================================START=============================================")

	count := RecurseDownFS("E:/movies")

	fmt.Println("movie count = ", count)

	fmt.Println("============================================END===============================================")
	fmt.Println("============================================END===============================================")
	fmt.Println("============================================END===============================================")
}
