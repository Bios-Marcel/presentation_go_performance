## Tips & Tricks

### pprof Graph

```
go tool pprof "-http=:8080" cpu.pprof
```

### Bound Check Hints

```
go build -gcflags=”-d=ssa/check_bce/debug=1" .
```

### Escape Analysis Hints

CLI:

```
go build -gcflags=-m .
```

Vim:

```
cexpr system('go build -gcflags=-m .')
copen
```
