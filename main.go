package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"github.com/buger/goterm"
	"github.com/GeistInDerSH/clearscreen"
)

func center(s string, n int, fill string) string {
	div := n / 2

	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
}



func main_menu() (map[int]string, int) {
	clearscreen.ClearScreen()
	var str_n string
	var str_nn string
	var str_nnn string
	re, _ := regexp.Compile(".noted$")
	n := 1
	directorylist, _ := ioutil.ReadDir(".")
	note_array := make(map[int]string)
	title := "CRYPTEDNOTE - Note taking for the Paranoid"
	crawl := "<<< Main Menu >>>"
	crawl2 := "<< Current Directory >>"
	dir, _ := os.Getwd()
	println(center(title,goterm.Width()-len(title)-2/2, " " ))
	println(center(crawl,goterm.Width()-len(crawl)-2/2, " " ))
	println(center(crawl2,goterm.Width()-len(crawl2)-2/2, " " ))
	println(center(dir,goterm.Width()-len(dir)-2/2, " " ))

	for _, file := range directorylist {
		if re.Match([]byte(file.Name())) {
			note_array[n] = file.Name()
			str_n = strconv.FormatInt(int64(n), 10)
			println(string(str_n) + " > " + file.Name())
		n = n + 1
		}

	}

	str_n = strconv.FormatInt(int64(n), 10)
	str_nn = strconv.FormatInt(int64(n+1), 10)
	str_nnn = strconv.FormatInt(int64(n+2), 10)
	println("\n")
	println(string(str_n) + " > " + "New Item")
	println(string(str_nn) + " > " + "Change Directory")
	println(string(str_nnn) + " > " + "Exit")

	println("\n")
	return note_array, n
}

func create_new(filename string) {
	os.Create(filename) 
}

func main() {
	var password string
	var input string
	var err error

	goterm.Clear()
	for {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println(err)
			}
		}()
		file_map, n := main_menu()
		fmt.Scanln(&input)
		if _, err := strconv.ParseInt(input, 10, 64); err == nil {
			switch int_input, _ := strconv.Atoi(input); {
				case int_input < n:
					converted_n, _ := strconv.Atoi(input)
					fmt.Println("Enter password: ")
					fmt.Scanln(&password)
					launch_kilo(file_map[converted_n], password)
				case int_input == n:
					fmt.Println("Enter new file name: ")
					fmt.Scanln(&input)
					create_new(input + ".noted")
				case int_input == n+1:
					fmt.Println("Enter new directory: ")
					fmt.Scanln(&input)
					if _, err := os.Stat(input); !os.IsNotExist(err) {
						os.Chdir(input)
					}

				case int_input == n+2:
					os.Exit(0)
				}

		}

	}

}
