package project

import (
	"fmt"
	"os"
	"os/exec"
)

func Initialize(projectName string, envVars map[string]string) error {
	// Create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Change to project directory
	err = os.Chdir(projectName)
	if err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Clone the TEN-Agent template directly into the current directory
	cmd := exec.Command("git", "clone", "--depth", "1", "https://github.com/digitallysavvy/TEN-Agent", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone TEN-Agent template: %w\n%s", err, output)
	}

	// Remove the .git directory to disassociate from the template repository
	err = os.RemoveAll(".git")
	if err != nil {
		return fmt.Errorf("failed to remove .git directory: %w", err)
	}

	// Initialize a new git repository
	cmd = exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize new git repository: %w", err)
	}

	// Create .env file with provided environment variables
	err = createEnvFile(envVars)
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}

	// Install dependencies
	cmd = exec.Command("docker", "compose", "pull")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to pull Docker images: %w", err)
	}

	// Build the project
	cmd = exec.Command("docker", "compose", "run", "--rm", "astra_agents_dev", "make", "build")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to build the project: %w", err)
	}

	return nil
}

func createEnvFile(envVars map[string]string) error {
	file, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content
	content := `# ------------------------------
# Environment Variables for server & worker
# ------------------------------

# ------------------------------
# Server Configuration
# ------------------------------

# Log path
LOG_PATH=%s
LOG_STDOUT=%s
# Graph designer server port
GRAPH_DESIGNER_SERVER_PORT=%s
# Server port
SERVER_PORT=%s
# Maximum number of workers
WORKERS_MAX=%s
# Worker quit timeout in seconds
WORKER_QUIT_TIMEOUT_SECONDES=%s

# Agora App ID and Agora App Certificate
AGORA_APP_ID=%s
AGORA_APP_CERTIFICATE=%s

# ------------------------------
# Worker Configuration
# ------------------------------

# Extension: aliyun_analyticdb_vector_storage
ALIBABA_CLOUD_ACCESS_KEY_ID=
ALIBABA_CLOUD_ACCESS_KEY_SECRET=
ALIYUN_ANALYTICDB_ACCOUNT=
ALIYUN_ANALYTICDB_ACCOUNT_PASSWORD=
ALIYUN_ANALYTICDB_INSTANCE_ID=
ALIYUN_ANALYTICDB_INSTANCE_REGION=cn-shanghai
ALIYUN_ANALYTICDB_NAMESPACE=
ALIYUN_ANALYTICDB_NAMESPACE_PASSWORD=

# Extension: aliyun_text_embedding
ALIYUN_TEXT_EMBEDDING_API_KEY=

# Extension: bedrock_llm
# Extension: polly_tts
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=

# Extension: agora_rtc
# Azure STT key and region
AZURE_STT_KEY=%s
AZURE_STT_REGION=%s

# Extension: azure_tts
# Azure TTS key and region
AZURE_TTS_KEY=%s
AZURE_TTS_REGION=%s

# Extension: cosy_tts
# Cosy TTS key
COSY_TTS_KEY=

# Extension: elevenlabs_tts
# ElevenLabs TTS key
ELEVENLABS_TTS_KEY=

# Extension: gemini_llm
# Gemini API key
GEMINI_API_KEY=

# Extension: litellm
# Using Environment Variables, refer to https://docs.litellm.ai/docs/providers
# For example:
#     OpenAI
#         OPENAI_API_KEY=<your-api-key>
#         OPENAI_API_BASE=<openai-api-base>
#     AWS Bedrock
#         AWS_ACCESS_KEY_ID=<your-aws-access-key-id>
#         AWS_SECRET_ACCESS_KEY=<your-aws-secret-access-key>
#         AWS_REGION_NAME=<aws-region-name>
LITELLM_MODEL=gpt-4o-mini

# Extension: openai_chatgpt
# OpenAI API key
OPENAI_API_KEY=%s
# OpenAI proxy URL
OPENAI_PROXY_URL=

# Extension: qwen_llm
# Qwen API key
QWEN_API_KEY=
`

	_, err = fmt.Fprintf(file, content,
		envVars["LOG_PATH"],
		envVars["LOG_STDOUT"],
		envVars["GRAPH_DESIGNER_SERVER_PORT"],
		envVars["SERVER_PORT"],
		envVars["WORKERS_MAX"],
		envVars["WORKER_QUIT_TIMEOUT_SECONDES"],
		envVars["AGORA_APP_ID"],
		envVars["AGORA_APP_CERTIFICATE"],
		envVars["AZURE_STT_KEY"],
		envVars["AZURE_STT_REGION"],
		envVars["AZURE_TTS_KEY"],
		envVars["AZURE_TTS_REGION"],
		envVars["OPENAI_API_KEY"],
	)

	return err
}
