package utils

func CmdArgs(args []string) {
	if len(args) == 1 {
		ShowUsage()
	}

	for _, arg := range args[1:] {
		switch arg {
		case "-v", "--version":
			ShowVersion()
		default:
			ShowUsage()
		}
	}
}
