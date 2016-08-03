package logger

import (
    "log"
)

func Log( msg string ) {
    log.Printf( "ENTITYID: %s", msg )
}