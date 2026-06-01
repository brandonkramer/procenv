# procenv

Read one environment variable from another process by pid.

## Install

From [pkg.go.dev](https://pkg.go.dev/github.com/brandonkramer/procenv):

```bash
go get github.com/brandonkramer/procenv@v0.1.0
```

## Usage

```go
value, ok := procenv.ProcEnv(pid, "MYAPP_RUN_ID")
if !ok {
    // pid invalid, process gone, or key missing
}
```

| Platform | Mechanism |
|----------|-----------|
| Linux | `/proc/<pid>/environ` |
| macOS | `ps eww -p <pid>` |
| Other | always `("", false)` |

Invalid `pid` (≤ 0) or empty `key` returns `("", false)` without syscalls.

## Development

Lefthook and golangci-lint are pinned in `go.mod` as **tools** (dev-only). Install git hooks once per clone:

```bash
make install-hooks
```

```bash
make check
make test
make lint
```
