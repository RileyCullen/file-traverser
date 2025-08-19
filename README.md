# file-traverser

A command line tool for traversing files/folders

## Building

Run `go mod -o ft ./src` from root.

## OS/Terminal Support

MacOS + zsh

## Installation

Pull this repo and build.

In root, run `mkdir Library/Application\ Support/ft`. Then run `touch Library/Application\ Support/ft/lastdir`

Add the following to your .zshrc file

```
ft() {
    export FT_LAST_DIR="$HOME/Library/Application Support/ft/lastdir"

    command $HOME/path/to/ft "$@"

    [ ! -f "$FT_LAST_DIR" ] || {
        . "$FT_LAST_DIR"
    }
}
```

Finally, restart your terminal an run `ft`
