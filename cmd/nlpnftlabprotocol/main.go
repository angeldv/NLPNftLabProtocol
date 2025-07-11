// cmd/nlpnftlabprotocol/main.go
package main

import (
"flag"
"log"
"os"

"nlpnftlabprotocol/internal/nlpnftlabprotocol"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nlpnftlabprotocol.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
