# Deploying DeepLX to Vercel

This guide explains how to deploy DeepLX to Vercel's serverless platform.

## Prerequisites

- A [Vercel](https://vercel.com) account
- [Vercel CLI](https://vercel.com/docs/cli) (optional)

## Deployment Steps

### Method 1: Deploy via Vercel Dashboard

1. Fork this repository to your GitHub account
2. Log in to your Vercel account and click "Add New Project"
3. Import your forked repository
4. Configure the following environment variables (if needed):
   - `TOKEN`: Optional authentication token
   - `DL_SESSION`: Optional DeepL session for Pro API features
   - `PROXY`: Optional HTTP proxy URL

5. Click "Deploy"

### Method 2: Deploy via Vercel CLI

1. Install Vercel CLI: `npm i -g vercel`
2. Clone this repository: `git clone https://github.com/OwO-Network/DeepLX.git`
3. Navigate to the project directory: `cd DeepLX`
4. Run: `vercel` and follow the prompts
5. Set environment variables if needed:
   ```
   vercel env add TOKEN your_token_here
   vercel env add DL_SESSION your_dl_session_here
   vercel env add PROXY your_proxy_url_here
   ```

## Using the Deployed API

Once deployed, you can use the API endpoints just like with the self-hosted version:

- `/translate`: Main translation endpoint
- `/v1/translate`: Pro API endpoint (requires DL_SESSION)
- `/v2/translate`: API endpoint compatible with official API format

## Notes

- Vercel has a timeout limit of 10 seconds for serverless functions on free tier
- Memory is limited to 1024MB per serverless function
- There may be cold start latency when the function hasn't been used recently

For more information on using DeepLX, see the [main documentation](https://deeplx.owo.network). 