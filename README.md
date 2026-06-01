# procenv

Read one environment variable from another process by pid.

## Install

```bash
go get github.com/brandonkramer/procenv
```

## Usage

```go
value, ok := procenv.ProcEnv(pid, "MYAPP_RUN_ID")
```

Linux reads `/proc/<pid>/environ`. macOS uses `ps eww`. Other platforms return `("", false)`.
