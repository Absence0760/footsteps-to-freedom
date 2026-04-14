Add a new tour to the website.

Tour name: $ARGUMENTS

Steps:
1. Create a new markdown file at `src/content/tours/<slug>.md` where `<slug>` is the kebab-case version of the tour name
2. Include YAML frontmatter with these required fields: `id`, `slug`, `name`, `description`, `image`, `category`, `duration`
3. Ask me which image to use from `static/` or if I want to add a new one
4. Write engaging markdown body content describing the tour — include highlights, what visitors will see, and practical details
5. Run `pnpm test` to verify the new content passes integrity checks
6. Run `pnpm check` to verify no type errors
