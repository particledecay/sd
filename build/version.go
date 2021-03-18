package build

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/rs/zerolog/log"
)

// NOTE: build.Version is set during the build process (on the command line)

var (
	// Version holds the tag of the build
	Version = ""
	// Commit holds the git commit sha
	Commit = ""
	// Date is the build date
	Date = ""
)

const versionTpl = `
Version:	v{{ .Version }}
SHA:		{{ .Commit }}
Built On:	{{ .Date }}`

// PrintVersion outputs the version info
func PrintVersion(verbose bool) string {
	if !verbose {
		if Version != "" {
			return fmt.Sprintf("v%s", Version)
		}
		return ""
	}

	if Version == "" {
		return ""
	}

	data := struct {
		Version string
		Commit  string
		Date    string
	}{
		Version: Version,
		Commit:  Commit,
		Date:    Date,
	}

	var tpl bytes.Buffer

	t, err := template.New("build").Parse(versionTpl)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	if err := t.Execute(&tpl, data); err != nil {
		log.Fatal().Msgf("%v", err)
	}

	return fmt.Sprintf(tpl.String())
}
