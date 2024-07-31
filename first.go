// you require atleast 1 package per app and main package is required to run the app
// this helps export import packages from other files
// Go has collection of packages called standard library (build-in packages)

// package main
// // main = entry point of the app
// import "fmt"

// func main() {
// 	fmt.Print("Hello, World!")
// }

// ####################################
// Not use Go Run - they may not have installed Go in their system
// Go module = multiple packages 
//go: go.mod file not found in current directory or any parent directory; see 'go help modules'
// go mod init exmaple/name = create go.mod file
// after this go build => no error
// above gets you a executable file in the same directory
// exe - in case of windows
// .\first.exe
// Hello, World!

// #########################################
// which code to execute first 
// func main() ==> only one main function in the entire app
// one folder = 1 main 
// thats why utility packages like fmt don't have func main in them
