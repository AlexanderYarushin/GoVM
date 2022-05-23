package utils

var Reset = "\033[0m"
var Red = "\033[31m"

func Error(error string) {
	println(Red + "Error: " + Reset + error)
}

func CompilationError(error string) {
	println(Red + "Compilation error: " + Reset + error)
}
