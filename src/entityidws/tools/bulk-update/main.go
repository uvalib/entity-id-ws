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
   "entityidws/api"
   "flag"
)

type testConfig struct {
   Endpoint string
   Token    string
}

var cfg = loadConfig()

func main() {

   if len( os.Args ) == 1 {
      fmt.Printf( "Update metadata for a set of DOI's\n" )
      fmt.Printf( "use: %s <file> [-ignore]\n", os.Args[ 0 ] )
      os.Exit( 0 )
   }

   ignoreError := false
   flag.BoolVar( &ignoreError, "ignore", true, "Ignore errors")
   flag.Parse( )

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
      entity := api.Request{ ID: doi, Schema: "datacite", DataCite: api.DataCiteSchema{ Title: "The title" } }
      status := client.Update( cfg.Endpoint, entity, cfg.Token )
      if status == expected {
         fmt.Printf( "Updated %s\n", doi )
      } else {
         fmt.Printf("ERROR: updating %s. Expected %v, got %v\n", doi, expected, status)
         if ignoreError == false {
            os.Exit(status)
         }
      }
   }
   os.Exit( 0 )
}

func loadConfig() testConfig {

   data, err := ioutil.ReadFile("src/entityidws/tools/bulk-update/config.yml")
   if err != nil {
      log.Fatal(err)
   }

   var c testConfig
   if err := yaml.Unmarshal(data, &c); err != nil {
      log.Fatal(err)
   }

   fmt.Printf("endpoint [%s]\n", c.Endpoint )
   //fmt.Printf("token    [%s]\n", c.Token )

   return c
}

//
// end of file
//

