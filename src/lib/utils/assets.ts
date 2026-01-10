import { base } from '$app/paths';

/**
 * Prepends the base path to asset URLs for GitHub Pages deployment
 * @param path - The asset path (e.g., '/hero-image.jpg')
 * @returns The full path with base prepended
 */
export function asset(path: string): string {
	return `${base}${path}`;
}
