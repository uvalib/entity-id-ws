package main

import (
   //"fmt"
   //"log"
   //"time"
)

func GetTokenDetails( token string ) ( Token, error ) {

	//start := time.Now( )

    //log.Printf( "Token %s NOT FOUND\t%s", token, time.Since( start ) )

    // return empty token if not found
    return Token{ }, nil
}

func ActivityIsOk( details Token, whom string, what string ) bool {
    return true
}