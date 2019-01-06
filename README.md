status [![GoDoc](https://godoc.org/github.com/els0r/status?status.svg)](https://godoc.org/github.com/els0r/status) [![Go Report Card](https://goreportcard.com/badge/github.com/els0r/status)](https://goreportcard.com/report/github.com/els0r/status)
============

Print status lines debian init-style.

### Usage

In any package where status output should be printed to the command line, simply ``import "github.com/els0r/status"`` and access the familiar functions with

```golang
status.Line("Updating database")
// do something awesomely magical
// ...
status.Okf("%d/10 entries updated", 10)
```

The two functions `Custom` and `AnyStatus` allow you to customize what is written
into the enclosed status. Say you want to print DONE instead of OK, then use `Custom`:

```golang
status.Line("Waiting for database update")
status.Custom(status.Blue, "DONE", "10/10 entries updated")
```
By default `Custom` truncates whatever you supply as the status argument to four characters such that the width is not altered with regard to the standard status types.

If you don't care and want it to be of arbitrary width, use `AnyStatus`:

```golang
status.Line("Status of database")
status.AnyStatus(status.Green, "UPDATED", "10/10 entries updated")
```

If you want to control where the output is sent, you can explicitly set an `io.Writer` object with `SetOutput`. The default writer will be `os.Stdout`. Example:
```golang
type myWriter struct{} // implements Write method of io.Writer

mw := myWriter{}
status.SetOutput(mw) // statusline output now handled by your own Write() implementation
```
