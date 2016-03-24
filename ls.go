// Copyright 2016 Lennart Espe. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
