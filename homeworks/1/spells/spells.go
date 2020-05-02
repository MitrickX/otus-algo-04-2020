package spells

import (
	"fmt"
	"io"
)

// Spell says draw # or ., where # is TRUE, . is FALSE
// Notice that y is column, x is row (usually the opposite, but not for this task)
type Spell func(x, y int) bool

// Draw draws 25x25 picture corresponds to spell function
// Writer is where to write
func Draw(spell Spell, writer io.Writer) {
	rows := 25
	cols := 25
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if spell(x, y) {
				print(writer, "#")
			} else {
				print(writer, ".")
			}
			if y < cols-1 {
				print(writer, " ")
			}
		}
		println(writer)
	}
}

// print prints into writer ignoring error
func print(w io.Writer, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

// println print new line into writer ignoring error
func println(w io.Writer, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}
