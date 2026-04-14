# Project Overview

**Footsteps to Freedom** is a marketing website for a boutique Cape Town tour company. It showcases guided tours focused on South African history (Mandela, townships, Cape Peninsula, etc.) with team bios, a contact form, and testimonials. Deployed as a static site to GitHub Pages.

## Stack

- **Framework**: SvelteKit 2 with Svelte 5 (runes)
- **Language**: TypeScript
- **Package manager**: pnpm
- **Adapters**: `@sveltejs/adapter-static` (GitHub Pages), `@sveltejs/adapter-vercel` (Vercel)
- **Styling**: normalize.css + custom CSS variables in `src/theme.css`, global styles in `src/app.css`
- **Icons**: unplugin-icons with `@iconify-json/material-symbols`
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
pnpm test             # Run tests once
pnpm test:watch       # Run tests in watch mode
```

## Verification

**Always run these checks before considering work complete:**

```bash
pnpm check            # Must pass — catches type errors and Svelte compiler issues
pnpm test             # Must pass — catches content, utility, and data loading regressions
pnpm build            # Must succeed — verifies the static build works end-to-end
```

If any step fails, fix the errors before committing. `pnpm check` and `pnpm test` are fast; run them after every non-trivial change.

## Testing

Tests use **vitest** (configured in `vite.config.ts`). Test files live next to the code they test as `*.test.ts`.

Current test coverage:
- **Content integrity** (`src/content/content.test.ts`): validates all tour and team markdown frontmatter, checks images exist in `static/`, ensures slugs are unique and match filenames
- **Asset utility** (`src/lib/utils/assets.test.ts`): tests base path prepending for GitHub Pages
- **Tour data loading** (`src/routes/tours/[slug]/page.test.ts`): tests `entries()` and `load()` functions, including 404 handling
- **Team data loading** (`src/routes/about/[slug]/page.test.ts`): tests `entries()` and `load()` functions, including 404 handling

When adding new content or utilities, add corresponding tests.

## Conventions

- Use Svelte 5 runes syntax (`$state`, `$derived`, `$effect`, `$props`) — not the legacy options API
- TypeScript throughout; `lang="ts"` on all `<script>` blocks
- All routes are prerendered (`prerender.default: true` in svelte.config.js)
- Use `src/lib/utils/assets.ts` to resolve static asset paths — never hardcode `/static/...`; the `BASE_PATH` env var adjusts paths for GitHub Pages subdirectory hosting
- CSS design tokens live in `src/theme.css`; use them via CSS variables (e.g., `var(--primary)`, `var(--text-body)`) rather than hardcoded values
- Contact page sets `ssr: false` due to Web3Forms client-only requirements
- New components go in the appropriate atomic level: `atoms/` for primitives, `molecules/` for composites, `organisms/` for full sections
- Page-level content components go in `src/lib/pages/` and are named `*PageContent.svelte`

## Common Tasks

### Adding a new tour
1. Create `src/content/tours/<slug>.md` with frontmatter: `id`, `slug`, `name`, `description`, `image`, `category`
2. Add the tour image to `static/`
3. The tour auto-renders at `/tours/<slug>` — no route changes needed

### Adding a team member
1. Create `src/content/team/<slug>.md` with frontmatter: `id`, `slug`, `name`, `role`, `bio`, `image`
2. Add the team member photo to `static/`
3. The bio auto-renders at `/about/<slug>` — no route changes needed

### Adding a new page
1. Create `src/routes/<page>/+page.svelte`
2. Create the content component in `src/lib/pages/<Page>PageContent.svelte`
3. Import and render the content component from the route file

### Modifying styles
- Design tokens (colors, spacing, typography) are in `src/theme.css`
- Global styles in `src/app.css`
- Component-scoped styles use `<style>` blocks within `.svelte` files

## Deployment

- **GitHub Pages**: push to `main` triggers `deploy.yml`, which builds with `BASE_PATH='/<repo-name>'` and deploys `build/`
- The `build/.nojekyll` file is created at build time to bypass Jekyll processing
- Vercel adapter is installed as an alternative deployment target
- **Workflow**: develop on `dev`, merge to `main` when ready to deploy. No PRs — direct merge.

## CI

- `test.yml` runs on push to `main`/`dev`: type check → tests → build
- `claude.yml` runs when `@claude` is mentioned in issues/comments (requires `ANTHROPIC_API_KEY` secret and `CLAUDE_ALLOWED_USERS` variable in GitHub repo settings)

## Custom Claude Commands

- `/add-tour <name>` — scaffold a new tour markdown file with frontmatter and content
- `/add-team-member <name>` — scaffold a new team member bio
- `/verify` — run the full check → test → build pipeline and report results
