// Tests for the statusline package.

package status

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type captureWriter struct{ text string }

var text string

func (c captureWriter) Write(p []byte) (n int, err error) {
	text += string(p)
	return len(p), nil
}

func TestMain(m *testing.M) {
	SetOutput(captureWriter{}) // capture the output
	os.Exit(m.Run())
}

func expectGenerator(color Color, status string, msgLine, msgStatus string) string {
	reps := StatusLineIndent - len(msgLine)
	if reps < 0 {
		reps = 0
	}
	return fmt.Sprintf("- %s%s[%s] %s\n", msgLine, strings.Repeat(".", reps), colors[color](status), msgStatus)
}

func testFunction(t *testing.T, msgLine, msgStatus, status string, funcToTest func(msg string), color Color) {

	expect := expectGenerator(color, status, msgLine, msgStatus)

	Line(msgLine)
	funcToTest(msgStatus)

	// read captured text
	if text != expect {
		t.Fatalf("have: %s; expect: %s", text, expect)
	}

	text = ""
}

func TestLineOk(t *testing.T) {
	testFunction(t, "This will work", "told you", ok, Ok, Green)
}

func TestLineOkLong(t *testing.T) {
	testFunction(t, "This will work and will be a very long status message that doesn't fit the standard length anymore", "told you", ok, Ok, Green)
}

func TestLineAttn(t *testing.T) {
	testFunction(t, "This will be a warn", "told you", warn, Attn, Yellow)
}

func TestLineWarn(t *testing.T) {
	testFunction(t, "This will be a alias for warn", "told you", warn, Warn, Yellow)
}

func TestLineFail(t *testing.T) {
	testFunction(t, "This will fail", "told you", fail, Fail, Red)
}

func TestCustom(t *testing.T) {
	msg := "This is custom"
	expect := fmt.Sprintf("- %s%s[ %s ] %s\n", msg, strings.Repeat(".", StatusLineIndent-len(msg)), colors[Blue]("DONE"), "World")

	Line(msg)
	Custom(Blue, "DONE", "World")

	// read captured text
	if text != expect {
		t.Fatalf("have: %s; expect: %s", text, expect)
	}
	text = ""

	expect = fmt.Sprintf("- %s%s[ %s ] %s\n", msg, strings.Repeat(".", StatusLineIndent-len(msg)), "DONE", "World")
	Line(msg)
	Custom(10, "DONEthiswillberemoved", "World")

	// read captured text
	if text != expect {
		t.Fatalf("have: %s; expect: %s", text, expect)
	}

	text = ""
}

func TestAnyStatus(t *testing.T) {
	msg := "Any status"
	expect := fmt.Sprintf("- %s%s[ %s ] %s\n", msg, strings.Repeat(".", StatusLineIndent-len(msg)), "THIS IS WAY TOO LONG", "World")

	Line("Any status")

	// choose an out of bounds color
	AnyStatus(256, "THIS IS WAY TOO LONG", "World")

	// read captured text
	if text != expect {
		t.Fatalf("have: %s; expect: %s", text, expect)
	}
	text = ""
}
