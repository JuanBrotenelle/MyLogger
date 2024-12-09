# MyLogger

My logger, which I intend to use all the time in future projects. Of interest:

1. The message is nicely logged in the format `[LEVEL] TimeStamp - Message`

That's literally all there is to it 😊

### How to use:

#### 1. Install the package

```bash
go get github.com/JuanBrotenelle/mylogger
```

#### 2. Create a logger

```go
log := NewLogger("INFO")
```

> [!IMPORTANT] The following levels can also be used:
> ```
>    "INFO":  1,
>    "ERROR": 2,
>    "WARN":  3,
>    "DEBUG": 4,
> ```

#### 3. Use

```go

l := NewLogger("DEBUG")

l.Debug("Debug message")

l.Info("Info message")

l.Warn("Warn message")

l.Error("Error message")

l.Fatal("Fatal message")
```
