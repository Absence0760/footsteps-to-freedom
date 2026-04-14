import { describe, it, expect, vi } from 'vitest';

vi.mock('$app/paths', () => ({
	base: '/footsteps-to-freedom',
}));

import { asset } from './assets';

describe('asset', () => {
	it('prepends base path to asset URL', () => {
		expect(asset('/hero-image.png')).toBe('/footsteps-to-freedom/hero-image.png');
	});

	it('prepends base path to nested paths', () => {
		expect(asset('/images/photo.jpg')).toBe('/footsteps-to-freedom/images/photo.jpg');
	});

	it('handles empty path', () => {
		expect(asset('')).toBe('/footsteps-to-freedom');
	});
});
