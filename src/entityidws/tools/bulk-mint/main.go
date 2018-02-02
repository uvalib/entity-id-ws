package main

import (
   "os"
   "log"
   "gopkg.in/yaml.v2"
   "io/ioutil"
   "net/http"
   "entityidws/client"
   "fmt"
   "entityidws/api"
   "strconv"
)

type testConfig struct {
   Endpoint string
   Token    string
   Shoulder string
}

var cfg = loadConfig()

func main() {

   if len( os.Args ) == 1 {
      fmt.Printf( "Mint a new block of DOI's\n" )
      fmt.Printf( "use: %s <count>\n", os.Args[ 0 ] )
      os.Exit( 0 )
   }

   count, _ := strconv.Atoi( os.Args[ 1 ] )
   for current := 0; current < count; current ++ {

      expected := http.StatusOK

      status, entity := client.Create( cfg.Endpoint, cfg.Shoulder, api.Request{ Schema: "datacite" }, cfg.Token )
      if status != expected {
         fmt.Printf("ERROR minting. Expected %v, got %v\n", expected, status)
         os.Exit( status )
      }

      fmt.Printf( "%03d -> %s\n", current + 1, entity.ID )

   }
   os.Exit( 0 )
}

func loadConfig() testConfig {

   data, err := ioutil.ReadFile("src/entityidws/tools/mint/config.yml")
   if err != nil {
      log.Fatal(err)
   }

   var c testConfig
   if err := yaml.Unmarshal(data, &c); err != nil {
      log.Fatal(err)
   }

   fmt.Printf("endpoint [%s]\n", c.Endpoint )
   fmt.Printf("token    [%s]\n", c.Token )
   fmt.Printf("shoulder [%s]\n", c.Shoulder )

   return c
}

//
// end of file
//

