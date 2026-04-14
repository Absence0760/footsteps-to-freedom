import { describe, it, expect } from 'vitest';
import { entries, load } from './+page.js';

describe('tours/[slug] data loading', () => {
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

		it('includes known tour slugs', () => {
			const slugs = (entries() as Array<{ slug: string }>).map((e) => e.slug);
			expect(slugs).toContain('city-walking');
			expect(slugs).toContain('cape-peninsula');
		});
	});

	describe('load', () => {
		it('returns tour data for a valid slug', async () => {
			const result = await load({ params: { slug: 'city-walking' } } as any);
			expect(result.tour).toBeDefined();
			expect(result.tour.name).toBe('City Walking');
			expect(result.tour.slug).toBe('city-walking');
			expect(result.tour.component).toBeDefined();
		});

		it('returns all expected metadata fields', async () => {
			const result = await load({ params: { slug: 'city-walking' } } as any);
			const { tour } = result;
			expect(tour.slug).toBeTruthy();
			expect(tour.name).toBeTruthy();
			expect(tour.description).toBeTruthy();
			expect(tour.image).toBeTruthy();
			expect(tour.category).toBeTruthy();
		});

		it('throws 404 for an invalid slug', async () => {
			await expect(load({ params: { slug: 'nonexistent-tour' } } as any)).rejects.toThrow();
		});
	});
});
