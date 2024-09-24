# TEN Agent CLI

TEN Agent CLI is a command-line tool designed to help developers initialize, start, and manage AI Agents using the TEN framework.

## Installation

### Option 1: Install with Homebrew (macOS and Linux)

You can install TEN Agent CLI using Homebrew:

```
brew install digitallysavvy/tap/ten-agent
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

## Build from Source

See: [BUILD_README.md](./docs/BUILD_README.md)

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

### Running the TEN Agent Development Environment

There are two options for starting the TEN Agent development environment:

1. VSCode - Start the TEN Agent services with VSCode Dev Containers. Open the project in VSCode, and use the `Dev Containers: Reopen in Container` command from the Command Palette (⇧⌘P).

2. Ten-Agent - Start the TEN Agent services with Docker.

```
ten-agent start
```

This command will:

- Start the Docker containers for the TEN Agent services
- Enter the main container, giving you a shell to work with

### Stopping the TEN Agent services

When starting the TEN Agent with `ten-agent start`, the containers are left running in detached mode. To stop the TEN Agent services, use the following command:

```
ten-agent stop
```

This command will stop and remove the Docker containers for the TEN Agent services.

### Generating an Extension

```
ten-agent generate [extension-name] --verbose
```

This command will generate an extension using Claude 3.5 Sonnet. The `--verbose` flag is optional and will output the generated code to the console. Feel free to omit it to generate the files silently.

### Help

```
ten-agent help
```

This command will display the help information for the TEN Agent CLI.

## Development and Testing

See: [LOCAL_DEV_README.md](./docs/LOCAL_DEV_README.md)

### Deploying to Railway

See: [DEPLOY_README.md](./docs/DEPLOY_README.md)

## Contributing

Contributions are welcome! Please feel free to fork the repository and submit a Pull Request.

## License

Apache 2.0[./LICENSE]
