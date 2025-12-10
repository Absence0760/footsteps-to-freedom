# Frontend Assets Management

This document explains how frontend assets (Tailwind CSS, HTMX, Alpine.js) are managed in this project.

## Overview

The application supports **two modes** for loading frontend dependencies:

1. **Local Assets** (Production recommended) - Assets are served from the application
2. **CDN Fallback** - If local assets aren't available, falls back to CDN

## Why Use Local Assets?

### Benefits:
- **Performance**: Faster load times (no external DNS lookups)
- **Reliability**: Works offline and isn't affected by CDN outages
- **Privacy**: No external requests that could track users
- **Control**: Specific versions guaranteed (no unexpected updates)
- **Cost**: No CDN bandwidth costs

### Trade-offs:
- Slightly larger Docker image (~500KB)
- Need to manually update dependencies
- Your bandwidth usage increases slightly

## Quick Start

### Download Assets

Run the download script to fetch the latest versions:

```bash
./download-assets.sh
```

This downloads:
- **Tailwind CSS** (~400KB) - Full standalone build
- **HTMX** (~48KB) - Minimized version 1.9.10
- **Alpine.js** (~44KB) - Latest 3.x version

### Verify Assets

Check that assets were downloaded:

```bash
ls -lh frontend/assets/css/
ls -lh frontend/assets/js/
```

## How It Works

### Automatic Fallback

All HTML files use a smart loading pattern:

```html
<!-- Tries local first, falls back to CDN if unavailable -->
<script src="/assets/js/htmx.min.js"
        onerror="this.onerror=null; this.src='https://unpkg.com/htmx.org@1.9.10'">
</script>
```

### File Locations

```
frontend/
├── assets/
│   ├── css/
│   │   └── tailwind.js          # Tailwind CSS standalone
│   └── js/
│       ├── htmx.min.js          # HTMX library
│       ├── alpine.min.js        # Alpine.js (latest)
│       └── alpine-3.15.2.min.js # Alpine.js (versioned)
├── index.html
├── about.html
└── contact.html
```

## Using CDN Only

If you prefer to always use CDN versions:

### Option 1: Delete Local Assets

```bash
rm -rf frontend/assets/css frontend/assets/js
```

The HTML files will automatically fall back to CDN.

### Option 2: Ignore in Git

Uncomment these lines in `.gitignore`:

```gitignore
/frontend/assets/css/
/frontend/assets/js/
```

Then remove from repository:

```bash
git rm -r --cached frontend/assets/
```

## Updating Dependencies

### Update All Assets

Simply re-run the download script:

```bash
./download-assets.sh
```

### Update Individual Libraries

**Update Tailwind CSS:**
```bash
curl -sL https://cdn.tailwindcss.com -o frontend/assets/css/tailwind.js
```

**Update HTMX:**
```bash
curl -sL https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js -o frontend/assets/js/htmx.min.js
```

**Update Alpine.js:**
```bash
curl -sL https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js -o frontend/assets/js/alpine.min.js
```

## Docker Deployment

### Production Build

The Dockerfile automatically includes assets:

```dockerfile
# Copy frontend files (includes assets/)
COPY frontend/ ./static/
```

Assets are served by the Go Fiber backend at `/assets/*`.

### Build & Deploy

```bash
# Build with local assets
docker-compose build

# Or deploy to AWS
./deploy.sh
```

## Troubleshooting

### Assets Not Loading

**Problem**: Page shows unstyled content or errors

**Solutions**:

1. Check assets exist:
   ```bash
   ls frontend/assets/css/
   ls frontend/assets/js/
   ```

2. Download assets:
   ```bash
   ./download-assets.sh
   ```

3. Verify path in browser DevTools (Network tab)

4. Check server logs for 404 errors

### Version Mismatch

**Problem**: Different versions in dev vs production

**Solution**: Commit assets to git:

```bash
git add frontend/assets/
git commit -m "Add frontend assets"
```

### CDN Fallback Not Working

**Problem**: Page fails when local assets missing

**Solution**: Check `onerror` handler in HTML:

```html
<script src="/assets/js/htmx.min.js"
        onerror="this.onerror=null; this.src='https://unpkg.com/htmx.org@1.9.10'">
</script>
```

## Performance Comparison

### With Local Assets (Recommended)

```
- First Load: ~500KB (all assets)
- Subsequent: 0KB (cached)
- DNS Lookups: 0
- External Requests: 0
```

### With CDN Only

```
- First Load: ~500KB (from CDN)
- Subsequent: 0KB (cached)
- DNS Lookups: 3 (cdn.tailwindcss.com, unpkg.com, cdn.jsdelivr.net)
- External Requests: 3
- Privacy Impact: User IP visible to 3rd parties
```

## Best Practices

### For Production
✅ Use local assets (better performance & privacy)
✅ Commit assets to repository (version consistency)
✅ Set cache headers for assets (handled by Fiber)

### For Development
✅ Either approach works fine
✅ CDN can be faster for rapid prototyping
✅ Local assets recommended for offline development

### For Client Projects
✅ Use local assets for enterprise clients (compliance)
✅ Document asset versions in commits
✅ Test asset updates in staging first

## Related Files

- `download-assets.sh` - Asset download script
- `.gitignore` - Asset ignore rules (commented out by default)
- `Dockerfile` - Copies assets to production image
- `backend/main.go` - Serves static files including assets

## Version Information

Current versions (as of last download):

| Library | Version | Size |
|---------|---------|------|
| Tailwind CSS | Latest | ~400KB |
| HTMX | 1.9.10 | ~48KB |
| Alpine.js | 3.15.2 | ~44KB |

**Total**: ~492KB (minified)

Run `./download-assets.sh` to see current versions.
