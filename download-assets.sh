#!/bin/bash

# Script to download frontend assets for local use
# This allows the application to work offline and improves performance

set -e

ASSETS_DIR="frontend/assets"
CSS_DIR="$ASSETS_DIR/css"
JS_DIR="$ASSETS_DIR/js"

echo "📦 Downloading frontend assets..."

# Create directories if they don't exist
mkdir -p "$CSS_DIR"
mkdir -p "$JS_DIR"

# Download Tailwind CSS (standalone build with all features)
echo "⬇️  Downloading Tailwind CSS..."
curl -sL https://cdn.tailwindcss.com -o "$CSS_DIR/tailwind.js"
echo "✅ Tailwind CSS downloaded"

# Download HTMX
echo "⬇️  Downloading HTMX..."
curl -sL https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js -o "$JS_DIR/htmx.min.js"
echo "✅ HTMX downloaded"

# Download Alpine.js
echo "⬇️  Downloading Alpine.js..."
curl -sL https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js -o "$JS_DIR/alpine.min.js"
echo "✅ Alpine.js downloaded"

# Download specific Alpine.js version for consistency (latest 3.x)
echo "⬇️  Downloading Alpine.js (versioned)..."
ALPINE_VERSION=$(curl -s https://registry.npmjs.org/alpinejs/latest | grep -o '"version":"[^"]*' | cut -d'"' -f4)
curl -sL "https://cdn.jsdelivr.net/npm/alpinejs@${ALPINE_VERSION}/dist/cdn.min.js" -o "$JS_DIR/alpine-${ALPINE_VERSION}.min.js"
echo "✅ Alpine.js v${ALPINE_VERSION} downloaded"

# Create a .gitignore for assets if it doesn't exist
if [ ! -f "$ASSETS_DIR/.gitignore" ]; then
    echo "# Downloaded assets" > "$ASSETS_DIR/.gitignore"
    echo "# Uncomment to ignore downloaded files (force CDN use)" >> "$ASSETS_DIR/.gitignore"
    echo "# css/" >> "$ASSETS_DIR/.gitignore"
    echo "# js/" >> "$ASSETS_DIR/.gitignore"
fi

echo ""
echo "✅ All assets downloaded successfully!"
echo ""
echo "📊 Asset sizes:"
du -h "$CSS_DIR"/* "$JS_DIR"/* 2>/dev/null | awk '{print "   "$2": "$1}'
echo ""
echo "💡 To use CDN instead of local assets, delete the assets directory:"
echo "   rm -rf $ASSETS_DIR/{css,js}"
echo ""
