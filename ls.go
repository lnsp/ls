/*
MIT License

Copyright (c) 2016 Lennart Espe

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"golang.org/x/crypto/ssh/terminal"
)

// Size of min distance between items
const minItemMargin = 3

func main() {
	// Get console width and height
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	// Read in directory
	files, _ := ioutil.ReadDir("./")

	// Find the fitting item width
	maxLength := 0
	for _, f := range files {
		l := len(f.Name())
		if l > maxLength {
			maxLength = l
		}
	}

	// Calculate max items per row
	itemWidth := maxLength + minItemMargin
	if itemWidth > width {
		// String gets truncated if its longer than terminal size, looks better
		itemWidth = width
	}
	itemsPerRow := width / itemWidth

	i := 1
	for _, f := range files {
		fmt.Printf("%-" + strconv.Itoa(itemWidth) + "s", f.Name())
		if i % itemsPerRow == 0 {
			fmt.Println()
		}
		i++
	}
}
