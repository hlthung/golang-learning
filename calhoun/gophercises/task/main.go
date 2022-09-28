package main

import (
	"github.com/hlthung/golang-learning/calhoun/gophercises/task/cmd"
)

// According to Calhoun, Go is awesome for building CLI that distribute to different OS,
// As Go simply converts / compiles the code into a binary file for any given platform
// while Java needs the JVM to interpret compiled code. However, below articles does mention
// the limitation of Go and did a comparison between Java and Go.
// https://spiralscout.com/blog/when-to-use-go-vs-java-one-programmers-take-on-two-top-languages
// https://www.turing.com/blog/golang-vs-java-which-language-is-best/
// Note: This behavior actually made Go faster than Java on almost every benchmark.
// This is due to how it is compiled: Go doesn’t rely on a virtual machine to compile its code.
// It gets compiled directly into a binary file. Because Go does not have the VM, it is faster.
// But that that VM also helps Java work on more platforms

// cobra command is now cobra-cli (sept 2022)
func main() {
	cmd.RootCmd.Execute()
}
