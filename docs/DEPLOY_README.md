# TEN Agent CLI - Deployment

## Prerequisites

- Have a TEN Agent project initialized and working.

## Deploy using Railway and GitHub

1. Create a new Railway project using your GitHub repository
2. Add environment variables
3. Deploy

## Deploy using Railway CLI

1. Install the Railway CLI using `brew install railway`
2. Run `railway login` and login using your Railway credentials
3. Run `railway init` to create a new project from the dockerfile
4. Run `railway up` to deploy
5. Go into the Railway project and add the environment variables

Now whenever you make any changes run `railway up` to deploy.
