package handlers

import (
    "path/filepath"
    "strings"
)

func Version( ) string {
    files, _ := filepath.Glob( "buildtag.*" )
    if len( files ) == 1 {
       return strings.Replace( files[ 0 ], "buildtag.", "", 1 )
    }

    return "unknown"
}
