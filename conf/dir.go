package conf

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

var scriptDir string

// getScriptDir returns the path to the directory containing the scripts
func getScriptDir(create bool) string {
	// memoized
	if scriptDir != "" {
		return scriptDir
	}

	// prefer the env var if set
	overrideDir := os.Getenv("SD_PATH")
	if overrideDir != "" {
		if !dirExists(overrideDir) {
			if !create {
				log.Fatal().Msgf("SD_PATH is set to an invalid directory: '%s'", overrideDir)
			}
			err := os.MkdirAll(overrideDir, 0755)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}
		}
		scriptDir = overrideDir
		return overrideDir
	}

	// fallback to the default
	u, err := user.Current()
	if err != nil {
		log.Error().Msgf("%v", err)
		return ""
	}

	binDir := filepath.Join(u.HomeDir, "bin")

	if !dirExists(binDir) {
		if !create {
			log.Fatal().Msgf("%v", err)
		}
		err = os.MkdirAll(binDir, 0755)
		if err != nil {
			log.Fatal().Msgf("%v", err)
		}
	}
	scriptDir = binDir

	return scriptDir
}

// GetScriptPath returns the absolute path of a script
func GetScriptPath(category, name string, create bool) string {
	scriptDir := getScriptDir(create)

	categoryDir := filepath.Join(scriptDir, category)

	catInfo, err := os.Stat(categoryDir)
	if os.IsNotExist(err) || !catInfo.IsDir() {
		if !create {
			log.Fatal().Msgf("could not find directory '%s'", categoryDir)
		}
		log.Debug().Msgf("creating directory '%s'", categoryDir)
		err := os.MkdirAll(categoryDir, 0755)
		if err != nil {
			log.Fatal().Msgf("%v", err)
		}
	}

	scriptPath := filepath.Join(categoryDir, name)

	return scriptPath
}

func dirExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}
