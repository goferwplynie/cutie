# Cutie
A cute CLI project manager to keep your projects organized and encourage you to keep coding :33

---

## Installation

### 1. Install with Go

Make sure you have Go installed and your `$GOPATH/bin` is in your PATH. Then run:

```bash
go install github.com/goferwplynie/cutie@latest
```

### 2. Build from source
```bash
git clone https://github.com/goferwplynie/cutie.git
cd cutie
makepkg -si

```

---

## Usage

### Starting a new project
```bash
cutie init <path> <name> --dl 2025-11-25 --template go.json --reminder 3 -v
```
| Flag         | Description |
|--------------|-------------|
| `-v`         | Enable verbose mode |
| `--dl`       | Deadline in `YYYY-MM-DD` format |
| `--template` | Template name stored in the config directory |
| `--reminder` | After how many days `cutie remind` should give a reminder message for this project |

`cutie projects`
shows all projects in a pretty table :33

`cutie remind [--nc]`
Checks the reminders file and displays encouraging messages for your projects.

A project is added to reminders if:

- It hasn’t been modified for the number of days set with --reminder, or

- There is one week left until its deadline

--nc — turn off colored output (color is on by default)

💡 You can run this command automatically when your shell starts. It only checks for reminders in the file and doesn’t add new ones unless the date changed. Reminders are always refreshed after adding a new project or when the date changes.
checks the reminders file and displays some encouraging texts for your project.
project gets into reminders file after it wasnt modified for the number of days set with `--reminder` flag or if there is one week left till deadline

example output of `cutie remind` :33 :
```
Projects that need some love:
   Your project librusApi is lonely~ a little attention will make it happy!
Upcoming Deadlines:
   Don’t forget! librusApi deadline is getting close (3 days left)~ cheer up!
   cutie is due in 2 days! Keep going, you’re doing great~
```

---

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
templates are stored in config directory `cutie/templates`

available variables in templates:
- `$NAME` - project name

---

## Optional: auto-run cutie remind on terminal startup

To always check your project reminders, you can add this line to your shell profile (~/.bashrc, ~/.zshrc, etc.):
```bash
cutie remind

```

---

## Future updates

for future i plan making better project template functionality by defining them with just files and folders. Also i plan on adding more variables for templates including env variables :3c
