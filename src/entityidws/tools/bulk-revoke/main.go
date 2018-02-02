package main

import (
   "os"
   "log"
   "gopkg.in/yaml.v2"
   "io/ioutil"
   "net/http"
   "entityidws/client"
   "fmt"
   "bufio"
)

type testConfig struct {
   Endpoint string
   Token    string
}

var cfg = loadConfig()

func main() {

   if len( os.Args ) == 1 {
      fmt.Printf( "Revoke a set of DOI's\n" )
      fmt.Printf( "use: %s <file>\n", os.Args[ 0 ] )
      os.Exit( 0 )
   }

   file, err := os.Open( os.Args[ 1 ] )
   if err != nil {
      fmt.Printf("%s\n", err )
      os.Exit( 1 )
   }

   defer file.Close()
   scanner := bufio.NewScanner( file )

   for scanner.Scan( ) {
      expected := http.StatusOK

      doi := scanner.Text( )
      status := client.Revoke( cfg.Endpoint, doi, cfg.Token )
      if status != expected {
         fmt.Printf("ERROR: revoking %s. Expected %v, got %v\n", doi, expected, status)
         os.Exit( status )
      }

      fmt.Printf( "Revoked %s\n", doi )

   }
   os.Exit( 0 )
}

func loadConfig() testConfig {

   data, err := ioutil.ReadFile("src/entityidws/tools/revoke/config.yml")
   if err != nil {
      log.Fatal(err)
   }

   var c testConfig
   if err := yaml.Unmarshal(data, &c); err != nil {
      log.Fatal(err)
   }

   fmt.Printf("endpoint [%s]\n", c.Endpoint )
   fmt.Printf("token    [%s]\n", c.Token )

   return c
}

//
// end of file
//

