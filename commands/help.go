package commands

import (
	"fmt"
)

func Help() {
	fmt.Println("Nestor, your personal butler!")
	fmt.Println("Version: 0.1.0")
	fmt.Println("\nUSAGE:")
	fmt.Println("    nestor <command> <args>")
	fmt.Println("\nCOMMANDS:")
	fmt.Println("    - help    : Print this help")
	fmt.Println("    - init    : Initialize a local project. Args: `python` or `typescript`")
	fmt.Println("    - git     : Git helpers. Args: `clean`")
	fmt.Println("    - docker  : Docker helpers. Args: `clean`")
}
