// gogithttpserv project main.go
package main

import (
    "log"
    "net/http"
//	"os"
    "github.com/bobquest33/go-git-http"
)

func main() {
    // Get git handler to serve a directory of repos
	//argPort := os.Args[1]
	//argRootgit := os.Args[2]	
	argPort := ":8080"
	argRootgit := "C:\\Users\\IBM_ADMIN\\Documents\\golang\\tools\\src"	
    git := githttp.New(argRootgit,"C:\\Users\\IBM_ADMIN\\rcs\\Git\\bin\\git.exe")

    // Attach handler to http server
    http.Handle("/", git)

    // Start HTTP server
    err := http.ListenAndServe(argPort, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
