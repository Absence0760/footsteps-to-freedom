import { describe, it, expect } from 'vitest';
import fs from 'node:fs';
import path from 'node:path';

const CONTENT_DIR = path.resolve('src/content');
const STATIC_DIR = path.resolve('static');

function parseFrontmatter(filePath: string): Record<string, string> {
	const raw = fs.readFileSync(filePath, 'utf-8');
	const match = raw.match(/^---\n([\s\S]*?)\n---/);
	if (!match) throw new Error(`No frontmatter found in ${filePath}`);
	const fields: Record<string, string> = {};
	for (const line of match[1].split('\n')) {
		const colon = line.indexOf(':');
		if (colon === -1) continue;
		const key = line.slice(0, colon).trim();
		const value = line.slice(colon + 1).trim();
		fields[key] = value;
	}
	return fields;
}

function getMarkdownFiles(dir: string): string[] {
	return fs
		.readdirSync(dir)
		.filter((f) => f.endsWith('.md'))
		.map((f) => path.join(dir, f));
}

describe('tour content', () => {
	const tourDir = path.join(CONTENT_DIR, 'tours');
	const tourFiles = getMarkdownFiles(tourDir);

	it('has at least one tour', () => {
		expect(tourFiles.length).toBeGreaterThan(0);
	});

	describe.each(tourFiles.map((f) => [path.basename(f), f]))('%s', (_, filePath) => {
		const fm = parseFrontmatter(filePath as string);

		it('has all required frontmatter fields', () => {
			for (const field of ['slug', 'name', 'description', 'image', 'category']) {
				expect(fm[field], `missing "${field}"`).toBeTruthy();
			}
		});

		it('references an image that exists in static/', () => {
			const imagePath = fm.image.startsWith('/') ? fm.image.slice(1) : fm.image;
			const fullPath = path.join(STATIC_DIR, imagePath);
			expect(fs.existsSync(fullPath), `image not found: static/${imagePath}`).toBe(true);
		});

		it('slug matches filename', () => {
			const filename = path.basename(filePath as string, '.md');
			expect(fm.slug).toBe(filename);
		});
	});

	it('has unique slugs', () => {
		const slugs = tourFiles.map((f) => parseFrontmatter(f).slug);
		expect(new Set(slugs).size).toBe(slugs.length);
	});
});

describe('team content', () => {
	const teamDir = path.join(CONTENT_DIR, 'team');
	const teamFiles = getMarkdownFiles(teamDir);

	it('has at least one team member', () => {
		expect(teamFiles.length).toBeGreaterThan(0);
	});

	describe.each(teamFiles.map((f) => [path.basename(f), f]))('%s', (_, filePath) => {
		const fm = parseFrontmatter(filePath as string);

		it('has all required frontmatter fields', () => {
			for (const field of ['slug', 'name', 'role', 'bio', 'image']) {
				expect(fm[field], `missing "${field}"`).toBeTruthy();
			}
		});

		it('references an image that exists in static/', () => {
			const imagePath = fm.image.startsWith('/') ? fm.image.slice(1) : fm.image;
			const fullPath = path.join(STATIC_DIR, imagePath);
			expect(fs.existsSync(fullPath), `image not found: static/${imagePath}`).toBe(true);
		});

		it('slug matches filename', () => {
			const filename = path.basename(filePath as string, '.md');
			expect(fm.slug).toBe(filename);
		});
	});

	it('has unique slugs', () => {
		const slugs = teamFiles.map((f) => parseFrontmatter(f).slug);
		expect(new Set(slugs).size).toBe(slugs.length);
	});
});
