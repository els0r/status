package status

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Color provides access to colors used by the status printer
type Color int
type colorSprintFunc func(string, ...interface{}) string

// Supported colors (bold by default)
const (
	None Color = iota
	Red
	Yellow
	Green
	Blue
	Magenta
	Cyan
	White
	Black
	_NumColors
)

// StatusLineIndent is exported in case people want to adapt other formatting to the status line width
const StatusLineIndent = 54

// status strings
const (
	empty = "      "
	ok    = "  OK  "
	warn  = " ATTN "
	fail  = " FAIL "
)

var colors = [_NumColors]colorSprintFunc{
	fmt.Sprintf,
	color.New(color.Bold, color.FgRed).SprintfFunc(),
	color.New(color.Bold, color.FgYellow).SprintfFunc(),
	color.New(color.Bold, color.FgGreen).SprintfFunc(),
	color.New(color.Bold, color.FgBlue).SprintfFunc(),
	color.New(color.Bold, color.FgMagenta).SprintfFunc(),
	color.New(color.Bold, color.FgCyan).SprintfFunc(),
	color.New(color.Bold, color.FgWhite).SprintfFunc(),
	color.New(color.Bold, color.FgBlack).SprintfFunc(),
}

var ioDest io.Writer

// by default, statusline writes to standard out
func init() {
	ioDest = os.Stdout
}

// SetOutput allows the user to control where the status should be written to
func SetOutput(w io.Writer) {
	if w != nil {
		ioDest = w
	}
}

// Linef prints the status on a new line allowing for the status text to be formatted
func Linef(fmtStr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtStr, args...)
	if len(msg) < StatusLineIndent {
		fmt.Fprintf(ioDest, "- %s%s", msg, strings.Repeat(".", StatusLineIndent-len(msg)))
		return
	}
	fmt.Fprintf(ioDest, "- %s", msg)
}

// Line prints msg on a new line.
func Line(msg string) {
	Linef("%s", msg)
}

// Okf prints an OK status event enclosed by brackets with formatted status explanation
func Okf(fmtStr string, args ...interface{}) {
	fmt.Fprintf(ioDest, "[%s] %s\n", colors[Green](ok), fmt.Sprintf(fmtStr, args...))
}

// Ok prints an OK status event enclosed by brackets and appends msg as the status explanation
func Ok(msg string) {
	Okf("%s", msg)
}

// Warnf prints an ATTN status enclosed by brackets with formatted status explanation
func Warnf(fmtStr string, args ...interface{}) {
	fmt.Fprintf(ioDest, "[%s] %s\n", colors[Yellow](warn), fmt.Sprintf(fmtStr, args...))
}

// Warn prints an ATTN status enclosed by brackets and appends msg as the status explanation
func Warn(msg string) {
	Warnf("%s", msg)
}

// Attnf is an alias for Warnf
func Attnf(fmtStr string, args ...interface{}) {
	Warnf(fmtStr, args...)
}

// Attn is an alias for Warn
func Attn(msg string) {
	Attnf("%s", msg)
}

// Failf prints a FAIL event enclosed by brackets with formatted status explanation
func Failf(fmtStr string, args ...interface{}) {
	fmt.Fprintf(ioDest, "[%s] %s\n", colors[Red](fail), fmt.Sprintf(fmtStr, args...))
}

// Fail prints a FAIL event enclosed by brackets
func Fail(msg string) {
	Failf("%s", msg)
}

// Customf is a function allowing the user to set the status to something other than OK, ATTN, or FAIL.
//
// Arguments:
//   - color:  any available color from the enumerated constants (e.g. Blue)
//   - status: the enclosed status. Will be trimmed to 4 characters to fit the
//   - fmtStr: the format string (message)
//   - args:   arguments to the message
func Customf(color Color, status, fmtStr string, args ...interface{}) {

	// trim string if it is wider than 4 characters
	if len(status) > 4 {
		status = status[:4]
	}
	AnyStatusf(color, status, fmtStr, args...)
}

// Custom is Customf with with fixed message string
func Custom(color Color, status, msg string) {
	Customf(color, status, "%s", msg)
}

// AnyStatusf is the same as Customf without the status length constraint
func AnyStatusf(color Color, status, fmtStr string, args ...interface{}) {

	// default to NONE if color is out of range
	if color > _NumColors {
		color = None
	}
	fmt.Fprintf(ioDest, "[ %s ] %s\n", colors[color](status), fmt.Sprintf(fmtStr, args...))
}

// AnyStatus is the same as AnyStatusf with a fixed message string
func AnyStatus(color Color, status, msg string) {
	AnyStatusf(color, status, "%s", msg)
}
