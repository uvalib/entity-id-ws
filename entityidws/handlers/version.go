package handlers

import (
	"path/filepath"
	"strings"
)

//
// Version -- get the version information
//
func Version() string {
	files, _ := filepath.Glob("buildtag.*")
	if len(files) == 1 {
		return strings.Replace(files[0], "buildtag.", "", 1)
	}

	return "unknown"
}

//
// end of file
//
