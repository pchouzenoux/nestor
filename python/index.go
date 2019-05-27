package python

import (
	"log"
)
import (
	"github.com/Nehorim/nestor/commons"
)

func getPythonEnvFile() (string, []byte) {
	var filecontent []byte = []byte(
		`echo Switch to local virtual env
source ./__venv__/bin/activate
`)
	return ".env", filecontent
}

func InitPythonVirtualEnv() {
	commons.ExecShellCmd("virtualenv", "__venv__")

	commons.ToFile(getPythonEnvFile())

	log.Println("Execute command `cd .` to activate python virtualenv")
}
