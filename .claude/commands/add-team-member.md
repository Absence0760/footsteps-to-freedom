Add a new team member to the website.

Name: $ARGUMENTS

Steps:
1. Create a new markdown file at `src/content/team/<slug>.md` where `<slug>` is the kebab-case version of the name
2. Include YAML frontmatter with these required fields: `name`, `role`, `bio`, `image`, `slug`
3. Ask me which image to use from `static/` or if I want to add a new one
4. Write a markdown biography for the team member
5. Run `pnpm test` to verify the new content passes integrity checks
6. Run `pnpm check` to verify no type errors
