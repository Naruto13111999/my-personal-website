# gyanankur.dev

Personal portfolio website for **Gyanankur Dey** — built with Go backend and a modern responsive frontend.

## Quick Start (Local)

```bash
cd gyanankur.dev
go run ./cmd/server
```

Open [http://localhost:8080](http://localhost:8080)

### API Endpoints

| Endpoint         | Description              |
|------------------|--------------------------|
| `GET /`          | Portfolio homepage       |
| `GET /api/profile` | Profile JSON (API)     |
| `GET /health`    | Health check             |
| `GET /static/*`  | CSS, JS assets           |

## Project Structure

```
gyanankur.dev/
├── cmd/server/main.go       # Entry point
├── internal/
│   ├── data/profile.go      # Resume data & projects
│   └── handlers/handlers.go # HTTP handlers
├── web/
│   ├── templates/index.html # Go HTML template
│   └── static/css, js       # Frontend assets
├── Dockerfile
├── fly.toml                 # Fly.io config
└── README.md
```

---

## Hosting Publicly

Once you're happy with the site locally, pick one of these options to deploy **gyanankur.dev** publicly.

### Prerequisites (All Options)

1. **Buy the domain** `gyanankur.dev` from a registrar (Namecheap, Google Domains, Cloudflare, etc.)
2. **Push code to GitHub** (recommended for CI/CD deploys)

---

## Option 1: Fly.io (Recommended — Simple & Free Tier)

Fly.io runs your Docker container close to users with a generous free tier.

### Steps

```bash
# Install flyctl
brew install flyctl   # macOS
# or: curl -L https://fly.io/install.sh | sh

# Login
fly auth login

# From project root
cd gyanankur.dev
fly launch          # Accept defaults, say NO to Postgres
fly deploy
```

Your app will be live at `https://gyanankur-dev.fly.dev` (or similar).

### Custom Domain (gyanankur.dev)

```bash
# Add your domain
fly certs add gyanankur.dev
fly certs add www.gyanankur.dev
```

Fly will show DNS records. At your domain registrar, add:

| Type  | Name | Value                    |
|-------|------|--------------------------|
| A     | @    | (IP from `fly ips list`) |
| AAAA  | @    | (IPv6 from fly)          |
| CNAME | www  | gyanankur-dev.fly.dev    |

Verify: `fly certs check gyanankur.dev`

**Cost:** Free tier covers small personal sites; ~$0–5/mo if you scale.

---

## Option 2: Railway

Great for GitHub-connected deploys with minimal config.

### Steps

1. Push repo to GitHub
2. Go to [railway.app](https://railway.app) → New Project → Deploy from GitHub
3. Select your repo; Railway auto-detects the Dockerfile
4. Set **PORT** env var to `8080` (usually auto-set)
5. Generate a public domain in Settings → Networking

### Custom Domain

In Railway → Settings → Domains → Add `gyanankur.dev`  
Add the CNAME record Railway provides at your registrar.

**Cost:** $5/mo hobby plan (includes $5 credit).

---

## Option 3: Render

Similar to Railway with a free tier (spins down after inactivity).

### Steps

1. Push to GitHub
2. [render.com](https://render.com) → New → Web Service
3. Connect repo, set:
   - **Runtime:** Docker
   - **Port:** 8080
   - **Health check path:** `/health`
4. Deploy

### Custom Domain

Render Dashboard → Settings → Custom Domains → Add `gyanankur.dev`  
Follow DNS instructions (CNAME to `your-app.onrender.com`).

**Cost:** Free tier available (cold starts after 15 min idle).

---

## Option 4: Google Cloud Run

Pay-per-request, scales to zero — good for low-traffic portfolios.

### Steps

```bash
# Install gcloud CLI, authenticate
gcloud auth login
gcloud config set project YOUR_PROJECT_ID

# Build & deploy
cd gyanankur.dev
gcloud run deploy gyanankur-dev \
  --source . \
  --region asia-south1 \
  --allow-unauthenticated \
  --port 8080
```

### Custom Domain

Cloud Run → Manage Custom Domains → Map `gyanankur.dev`  
Update DNS at registrar per Google's instructions.

**Cost:** Free tier: 2M requests/month.

---

## Option 5: AWS (ECS / App Runner / EC2)

For AWS-native hosting:

### App Runner (Easiest on AWS)

1. Push Docker image to ECR
2. App Runner → Create service from ECR image, port 8080
3. Route 53 or external DNS → CNAME to App Runner URL

### EC2 (Manual)

```bash
# On a t3.micro instance
docker build -t gyanankur-dev .
docker run -d -p 80:8080 --restart always gyanankur-dev
```

Point `gyanankur.dev` A record to EC2 elastic IP. Use Nginx + Let's Encrypt for HTTPS.

**Cost:** t3.micro free tier eligible; App Runner ~$5+/mo.

---

## Option 6: DigitalOcean App Platform

1. [cloud.digitalocean.com](https://cloud.digitalocean.com) → Apps → Create App
2. Connect GitHub repo, detect Dockerfile
3. Set HTTP port to 8080
4. Add custom domain in app settings

**Cost:** ~$5/mo for basic tier.

---

## Option 7: VPS (Hetzner / DigitalOcean Droplet)

Most control, cheapest at scale.

```bash
# On Ubuntu VPS
sudo apt update && sudo apt install -y docker.io
git clone YOUR_REPO_URL
cd gyanankur.dev
docker build -t gyanankur-dev .
docker run -d -p 80:8080 --name gyanankur --restart always gyanankur-dev

# HTTPS with Caddy (recommended)
sudo apt install -y caddy
# /etc/caddy/Caddyfile:
# gyanankur.dev {
#   reverse_proxy localhost:8080
# }
sudo systemctl reload caddy
```

Point domain A record to VPS IP.

**Cost:** ~$4–6/mo (Hetzner CX22).

---

## DNS Setup Summary for gyanankur.dev

Regardless of host, you'll typically configure:

| Record | Purpose                          |
|--------|----------------------------------|
| A      | Points root domain to server IP  |
| AAAA   | IPv6 (if provider supports)      |
| CNAME  | `www` → your hosting URL         |

**Tip:** Use [Cloudflare](https://cloudflare.com) as DNS proxy for free SSL, DDoS protection, and caching.

---

## Docker (Manual)

```bash
docker build -t gyanankur-dev .
docker run -p 8080:8080 gyanankur-dev
```

---

## Environment Variables

| Variable | Default | Description     |
|----------|---------|-----------------|
| `PORT`   | `8080`  | HTTP listen port |

---

## Recommended Path

For a personal portfolio like this:

1. **Develop locally** → `go run ./cmd/server`
2. **Deploy to Fly.io** (fastest path to production)
3. **Point gyanankur.dev** DNS to Fly
4. **Enable HTTPS** (automatic on Fly/Railway/Render)

Total time to go live: ~30 minutes after you're satisfied locally.

---

## License

Personal portfolio — all rights reserved © Gyanankur Dey
