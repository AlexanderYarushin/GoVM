package utils

var Reset = "\033[0m"
var Red = "\033[31m"

func Error(error string) {
	println(Red + "Error: " + Reset + error)
}

func CompilationError(error string) {
	println(Red + "Compilation error: " + Reset + error)
}

var Errors = map[string]func(args ...string) string{
	"SourceNotFound": func(args ...string) string { return "not found .vms source code file" },
	"InvalidArgNum": func(args ...string) string {
		return "invalid number of arguments in command '" + args[0] + "'"
	},
	"InvalidOperand":  func(args ...string) string { return "invalid operand '" + args[0] + "'" },
	"DuplicateLabels": func(args ...string) string { return "duplicate labels" },
	"InvalidCommand":  func(args ...string) string { return "invalid command '" + args[0] + "'" },
}
