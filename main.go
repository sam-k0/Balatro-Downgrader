package main

import (
	"os"

	"github.com/icedream/go-bsdiff"
)

func makePatch(oldPath string, newPath string, patchPath string) {

	oldFile, err := os.Open(oldPath)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Open(newPath)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	patchFile, err := os.Create(patchPath)
	if err != nil {
		panic(err)
	}
	defer patchFile.Close()

	err = bsdiff.Diff(oldFile, newFile, patchFile)
	if err != nil {
		panic(err)
	}
}

func applyPatch(oldPath string, newPath string, patchPath string) {

	oldFile, err := os.Open(oldPath)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	patchFile, err := os.Open(patchPath)
	if err != nil {
		panic(err)
	}
	defer patchFile.Close()

	newFile, err := os.Create(newPath)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	err = bsdiff.Patch(oldFile, newFile, patchFile)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Get the command line arguments: oldFile, newFile, patchFile, and mode: make or apply
	args := os.Args[1:]

	if len(args) < 4 {
		panic("Invalid number of arguments: 4 arguments required")
	}

	oldFile := args[0]
	newFile := args[1]
	patchFile := args[2]
	mode := args[3]

	if mode == "make" {
		// For making: newestGameFile, oldGameFile, outputPatchFile make
		makePatch(oldFile, newFile, patchFile)
	}

	if mode == "apply" {
		// For applying: newestGameFile, outputOldFile, patchFile apply
		applyPatch(oldFile, newFile, patchFile)
	}

}
