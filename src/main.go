package main

import (
    "go-api/src/core"
)


func main() {
    const PORT int = 8080

    var server *core.Server = core.CreateServer(PORT)

    core.StartServer(server)
}

