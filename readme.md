# Balatro Downgrader

A simple tool to downgrade your Balatro install to a previous version.

# Usage

Arguments for creating a patchfile
`go run main.go NewGame.exe OldGame.exe patchfile make`

Arguments for applying a patchfile
`go run main.go NewGame.exe OutputOldGame.exe patchfile apply`

Included is a patchfile (`p`) for downgrading Balatro from 1.0.1g-FULL to 1.0.0n-FULL.
Afaik this is not illegal as the patchfile is not redistributing any of the original files,
only the changes between them, so a patchfile cannot be used to pirate the game.
More info can be found here: http://www.daemonology.net/papers/bsdiff.pdf

# Dependencies

- go-bsdiff: github.com/icedream/go-bsdiff