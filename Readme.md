# Cutie
## some random project manager or something because im bored

## Usage
Starting new project :33
`cutie init <path> <name> --dl 2025-11-25 --template go.json --reminder 3 -v`
- -v - verbose mode
- --dl - deadline in `YYYY-MM-DD` format
- --template - template name stored in config directory
- --reminder - after how many days `cutie remind` should give reminder message for this project

## Templates
project templates for faster project creation
example go.json file:
```
{
    "files":[
    "file.go",
    "internal/",
    "models/",
    "cmd/main.go"
],
    "commands":[
    "go mod init github.com/goferwplynie/$NAME",
    "echo -e 'package main\n\nfunc main(){\n}' > cmd/main.go"
]
}

```
there are available variables for templates:
$NAME - project name
