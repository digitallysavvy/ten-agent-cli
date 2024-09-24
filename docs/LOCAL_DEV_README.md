# TEN Agent CLI - Build and Run the CLI locally

While developing, you can run the CLI directly with Go:

```
go run main.go [command]
```

For example:

```
go run main.go init my-agent
go run main.go start
go run main.go stop

```

### Testing

To run the tests for the CLI, use the following command:

```

go test ./...

```

This will run all tests in the project.

### Adding new commands

1. Create a new file in the `cmd` directory for your command (e.g., `cmd/newcommanddodo.go`).
2. Implement your command logic.
3. In the `init()` function of your new command file, add the command to the root command:
   ```go
   func init() {
       rootCmd.AddCommand(newcommandCmd)
   }
   ```
