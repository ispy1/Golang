package main

func main() {
	println("in main before calling greeting")
	greeting()
	println("in main after calling greeting")
}
func greeting() {
	println("in greeting :hi!!1")
}
