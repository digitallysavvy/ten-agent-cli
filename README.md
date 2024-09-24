# TEN Agent CLI

TEN Agent CLI is a command-line tool designed to help developers initialize, start, and manage AI Agents using the TEN framework.

## Installation

### Option 1: Install with Homebrew (macOS and Linux)

You can install TEN Agent CLI using Homebrew:

```
brew install digitallysavvy/tap/ten-agent-cli
```

### Option 2: Install from Releases

#### Mac and Linux

1. Download the latest release for your operating system from the [releases page](https://github.com/digitallysavvy/ten-agent-cli/releases).
2. Make the downloaded file executable:
   ```
   chmod +x ten-agent
   ```
3. Move the executable to a directory in your PATH:
   ```
   sudo mv ten-agent /usr/local/bin/
   ```

#### Windows

1. Download the latest Windows release from the [releases page](https://github.com/digitallysavvy/ten-agent-cli/releases).
2. Add the directory containing the `ten-agent.exe` file to your system's PATH.

### Option 3: Build from Source

#### Prerequisites

- Go 1.16 or later
- Docker
- Docker Compose

#### Building the CLI

1. Clone this repository:

```
git clone https://github.com/digitallysavvy/ten-agent-cli.git
cd ten-agent-cli
```

2. Copy the example environment file:

```
cp .env.example .env
```

3. Edit the `.env` file and set the necessary environment variables.

4. Build the CLI:

```
go build -o ten-agent
```

5. (Optional) Add the CLI to your PATH for easier access:

```
sudo mv ten-agent /usr/local/bin/
```

## Usage

### Initializing a new TEN Agent project

```
ten-agent init [project-name]
```

This command will:

- Create a new directory with the given project name
- Clone the TEN-Agent template into this directory
- Prompt you for necessary environment variables
- Set up the initial project structure

### Starting the TEN Agent services

```
ten-agent start
```

This command will:

- Start the Docker containers for the TEN Agent services
- Enter the main container, giving you a shell to work with

### Stopping the TEN Agent services

```
ten-agent stop
```

This command will stop and remove the Docker containers for the TEN Agent services.

## Development and Testing

### Running the CLI locally

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

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Apache 2.0[./LICENSE]
