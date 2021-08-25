package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main_menu() (map[int]string, int) {
	var str_n string
	re, _ := regexp.Compile(".noted$")
	n := 1
	directorylist, _ := ioutil.ReadDir(".")
	note_array := make(map[int]string)

	println("CRYPTEDNOTE - Note taking for the Paranoid")
	println("<<< Main Menu >>>")

	for _, file := range directorylist {
		if re.Match([]byte(file.Name())) {
			note_array[n] = file.Name()
			println(str_n + " > " + file.Name())
			n = n + 1
		}

	}

	str_n = strconv.FormatInt(int64(n), 10)
	println(string(str_n) + " > " + "New Item")
	println("\n\n")
	return note_array, n
}

func create_new(filename string) {
	os.Create(filename)
}

func main() {
	var password string
	var input string
	for true {
		file_map, n := main_menu()
		fmt.Scanln(&input)
		int_input, _ := strconv.Atoi(input)
		if int_input != n {
			converted_n, _ := strconv.Atoi(input)
			fmt.Println("Enter password: ")
			fmt.Scanln(&password)
			launch_kilo(file_map[converted_n], password)
		} else {
			fmt.Println("Enter new file name: ")
			fmt.Scanln(&input)
			create_new(input + ".noted")

		}

	}
}
