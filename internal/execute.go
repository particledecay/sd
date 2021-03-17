package internal

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func RunScript(scriptPath string, args []string) {
	scriptInfo, err := os.Stat(scriptPath)
	if os.IsNotExist(err) || scriptInfo.Mode()&0111 == 0 {
		log.Fatal().Msgf("'%s' does not exist or is not executable", scriptPath)
	}
	cmd := exec.Command(scriptPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
}

func EditScript(scriptPath string) {
	// default editor if nothing else
	defaultEditor := "vi"

	// VISUAL env var only as a fallback
	visual, ok := os.LookupEnv("VISUAL")
	if ok {
		defaultEditor = visual
	}

	// prefer the EDITOR env var
	editor, ok := os.LookupEnv("EDITOR")
	if ok {
		defaultEditor = editor
	}

	cmd := exec.Command(defaultEditor, scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		_, err := os.Stat(scriptPath)
		if !os.IsNotExist(err) {
			err = os.Chmod(scriptPath, 0755)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}
		}
	}
}
