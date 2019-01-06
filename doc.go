/*
Package status is a go implementation of a debian-like statusline.

The idea of the package is that no prior configuration needs to be applied before
using it so that functions can be called "off-the-shelf"

Example:

    package main
    import (
        "os"
        "time"

        "github.com/els0r/status"
    )

    func main() {
        // print a status line
        status.Linef("Waiting %d seconds", 5)
        time.Sleep(5*time.Second)
        status.Ok("")

        // now write to stderr
        status.SetOutput(os.Stderr)

        // print another status line
        status.Linef("Alerting you in %d seconds", 5)
        time.Sleep(5*time.Second)
        status.Warn("this went to stderr")
    }
*/
package status
