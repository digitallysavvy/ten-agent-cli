# TEN Agent CLI - Build from Source

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
go build -ldflags="-X github.com/digitallysavvy/ten-agent-cli/cmd.anthropicAPIKey=${ANTHROPIC_API_KEY} -X github.com/digitallysavvy/ten-agent-cli/cmd.version=dev -X github.com/digitallysavvy/ten-agent-cli/cmd.commit=$(git rev-parse --short HEAD) -X github.com/digitallysavvy/ten-agent-cli/cmd.date=$(date -u +%Y-%m-%dT%H:%M:%SZ) -s -w" -o ten-agent
```

This command will:

- Set the Anthropic API key from the environment variable
- Set the version to "dev"
- Include the short commit hash
- Add the current date and time
- Strip debugging information to reduce binary size

Note: Make sure the `ANTHROPIC_API_KEY` environment variable is set before running this command.

5. (Optional) Add the CLI to your PATH for easier access:

```
sudo mv ten-agent /usr/local/bin/
```
