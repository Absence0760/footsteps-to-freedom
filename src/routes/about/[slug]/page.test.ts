import { describe, it, expect } from 'vitest';
import { entries, load } from './+page.js';

describe('about/[slug] data loading', () => {
	describe('entries', () => {
		it('returns an array of slug objects', () => {
			const result = entries() as Array<{ slug: string }>;
			expect(Array.isArray(result)).toBe(true);
			expect(result.length).toBeGreaterThan(0);
			for (const entry of result) {
				expect(entry).toHaveProperty('slug');
				expect(typeof entry.slug).toBe('string');
				expect(entry.slug.length).toBeGreaterThan(0);
			}
		});

		it('includes known team member slugs', () => {
			const slugs = (entries() as Array<{ slug: string }>).map((e) => e.slug);
			expect(slugs).toContain('garth-angus');
			expect(slugs).toContain('mary-ann-wenham');
		});
	});

	describe('load', () => {
		it('returns member data for a valid slug', async () => {
			const result = await load({ params: { slug: 'garth-angus' } } as any);
			expect(result.member).toBeDefined();
			expect(result.member.name).toBe('Garth Angus');
			expect(result.member.slug).toBe('garth-angus');
			expect(result.member.component).toBeDefined();
		});

		it('returns all expected metadata fields', async () => {
			const result = await load({ params: { slug: 'garth-angus' } } as any);
			const { member } = result;
			expect(member.slug).toBeTruthy();
			expect(member.name).toBeTruthy();
			expect(member.role).toBeTruthy();
			expect(member.bio).toBeTruthy();
			expect(member.image).toBeTruthy();
		});

		it('throws 404 for an invalid slug', async () => {
			await expect(load({ params: { slug: 'nonexistent-member' } } as any)).rejects.toThrow();
		});
	});
});
