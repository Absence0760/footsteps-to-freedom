# Project Overview

**Footsteps to Freedom** is a marketing website for a boutique Cape Town tour company. It showcases guided tours focused on South African history (Mandela, townships, Cape Peninsula, etc.) with team bios, a contact form, and testimonials. Deployed as a static site to GitHub Pages.

## Stack

- **Framework**: SvelteKit 2 with Svelte 5 (runes)
- **Language**: TypeScript
- **Package manager**: pnpm
- **Adapters**: `@sveltejs/adapter-static` (GitHub Pages), `@sveltejs/adapter-vercel` (Vercel)
- **Styling**: normalize.css + custom CSS variables in `src/theme.css`, global styles in `src/app.css`
- **Icons**: unplugin-icons with `@iconify-json/material-symbols`
- **Component explorer**: Storybook 8
- **Markdown**: mdsvex (for tour and team member content)
- **Forms**: Web3Forms API for contact form submissions
- **XSS sanitization**: isomorphic-dompurify

## Folder Structure

```
src/
  routes/
    +layout.svelte           # Header + Footer wrapper
    +layout.ts
    +page.svelte             # Home page
    about/
      +page.svelte           # Team listing
      [slug]/+page.svelte    # Team member detail
      [slug]/+page.js        # Prerender entries
    tours/
      +page.svelte           # Tours listing
      [slug]/+page.svelte    # Tour detail
      [slug]/+page.js        # Prerender entries
    contact/+page.svelte
    faqs/+page.svelte
    privacy/+page.svelte
    terms/+page.svelte
    sitemap.xml/+server.ts   # XML sitemap
  lib/
    components/
      atoms/                 # Button, Text, Icon, Input, Link, Alert, Image, Divider
      molecules/             # Card, Header, Footer, ContactForm, Review, SearchBar, Hamburger
      organisms/             # HeroSection, TestimonialsSection, CallToActionSection, etc.
    pages/                   # Page-level content components (HomePageContent, etc.)
    utils/assets.ts          # Prepends base path to static asset URLs
    server/index.ts          # Server-side utilities
  content/
    tours/*.md               # Tour content with YAML frontmatter
    team/*.md                # Team member bios with YAML frontmatter
  stories/                   # Storybook stories
  app.css                    # Global styles
  theme.css                  # Design tokens (colors, typography, shadows)
  app.d.ts
static/                      # Images (jpg/png) served as static assets
.github/
  workflows/
    deploy.yml               # Build and deploy to GitHub Pages on push to main
    claude.yml               # Claude Code automation (@claude mentions in issues/PRs)
```

## Content System

Tours and team members are authored as Markdown files with YAML frontmatter in `src/content/`. They are loaded at build time using `import.meta.glob` and prerendered into static pages at `/tours/[slug]` and `/about/[slug]`.

Frontmatter fields:
- **Tours**: `id`, `slug`, `name`, `description`, `image`, `category`
- **Team**: `id`, `slug`, `name`, `role`, `bio`, `image`

## Component Architecture

Follows atomic design:
- **Atoms**: Primitive elements (Button, Text, Icon, Input, Link, Alert)
- **Molecules**: Composite UI (Card, Header, Footer, FormField, Review)
- **Organisms**: Full-page sections (HeroSection, TestimonialsSection, ContactForm)
- **Pages**: Page-specific wrappers (HomePageContent, ToursPageContent, etc.)

## Development

```bash
pnpm i                # Install dependencies
pnpm dev              # Dev server on :7777
pnpm build            # Production build → build/
pnpm preview          # Preview build on :8888
pnpm check            # Type-check
pnpm check:watch      # Type-check in watch mode
pnpm storybook        # Storybook on :9999
pnpm build-storybook  # Build static Storybook
```

## Root package.json scripts — estate format

Root `scripts` follow the estate-wide format canonized in the templates repo's `base` CLAUDE.md; the exemplar is `project-running/package.json` — read it before restructuring this repo's scripts:

- `"//-- <group> --": "<one-line description>"` comment-key dividers above each cluster; the description carries load-bearing facts (ports, prerequisites, doc pointers), not filler.
- Verb-first, colon-namespaced names: `setup[:*]`, `dev:*` (orchestrators, then `dev:db:*`, `dev:run:<app>`, per-service groups), `build:<surface>`, `check:<surface>`, `test:<surface>[:unit|:e2e]`, `gen:<what>`. Long-running services reuse the lifecycle verbs `up`/`down`/`status`/`logs`.
- JSON holds one-liners only — anything longer delegates to a script under `bin/` or `scripts/`; workspace delegation goes through `pnpm -C <workspace> <script>`.
- New scripts join an existing group (or add a new `//--` divider in the right place); never append ungrouped entries at the bottom.
- Keep a `test:scripts` guard validating the root script targets (project-running's `scripts/check_root_scripts.mjs` is the reference shape; write it against this repo's layout).

If the current scripts block predates this format, migrate it the next time a change touches it — as its own commit, and renaming a script must update every caller (CI workflows, docs, `bin/`) in the same change.

## Conventions

- Use Svelte 5 runes syntax (`$state`, `$derived`, `$effect`, `$props`) — not the legacy options API
- TypeScript throughout; `lang="ts"` on all `<script>` blocks
- All routes are prerendered (`prerender.default: true` in svelte.config.js)
- Use `src/lib/utils/assets.ts` to resolve static asset paths — never hardcode `/static/...`; the `BASE_PATH` env var adjusts paths for GitHub Pages subdirectory hosting
- Component stories live in `src/stories/` as `*.stories.svelte`
- CSS design tokens live in `src/theme.css`; use them via CSS variables rather than hardcoded values
- Contact page sets `ssr: false` due to Web3Forms client-only requirements

## Deployment

- **GitHub Pages**: push to `main` triggers `deploy.yml`, which builds with `BASE_PATH='/<repo-name>'` and deploys `build/`
- The `build/.nojekyll` file is created at build time to bypass Jekyll processing
- Vercel adapter is installed as an alternative deployment target

## Pull Request Guidelines

- Target branch: `main`
- Keep PRs focused; one feature or fix per PR
- Draft PRs are fine for work-in-progress

## Merging & branch protection

`main` follows the estate "sealed main + CI gate" standard: every change reaches `origin/main` through a PR — **no direct pushes** (enforced on admins, including the owner). There are **0 required approvals**. Force-pushes, branch deletion, and unresolved conversations are blocked; history is linear. This repo has no functional test/build CI on PRs yet, so there is no required **`CI gate`** status check — when CI lands, add an aggregator job named `CI gate` (that `needs:` the CI jobs) to each functional workflow to make it the merge gate. Commit locally per-piece, but land via a PR. (`deploy.yml` still fires on the merge commit — a merge to `main` is itself a push.)
