package main

import(
	"fmt"
	"io"
	"io/ioutil"	
	"os"
)

// scan a new folder for Git repositories
func scan(folder string) {
    fmt.Printf("Found folders:\n\n")
    repositories := recursiveScanFolder(folder)
    filePath := getDotFilePath()
    addNewSliceElementsToFile(filePath, repositories)
    fmt.Printf("\n\nSuccessfully added\n\n")
}