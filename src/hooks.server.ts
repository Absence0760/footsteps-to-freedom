import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	const response = await resolve(event);

	// Add security headers to protect against common attacks
	response.headers.set('X-Frame-Options', 'DENY'); // Prevent clickjacking
	response.headers.set('X-Content-Type-Options', 'nosniff'); // Prevent MIME type sniffing
	response.headers.set('X-XSS-Protection', '1; mode=block'); // Legacy XSS protection
	response.headers.set('Referrer-Policy', 'strict-origin-when-cross-origin'); // Control referrer information

	// Content Security Policy - allows inline styles for Svelte but blocks unsafe scripts
	response.headers.set(
		'Content-Security-Policy',
		"default-src 'self'; " +
		"script-src 'self' 'unsafe-inline'; " + // unsafe-inline needed for Svelte hydration
		"style-src 'self' 'unsafe-inline'; " + // unsafe-inline needed for Svelte component styles
		"img-src 'self' data: https:; " +
		"font-src 'self' data:; " +
		"connect-src 'self'; " +
		"frame-ancestors 'none'; " +
		"base-uri 'self'; " +
		"form-action 'self'"
	);

	// Permissions Policy - restrict browser features
	response.headers.set(
		'Permissions-Policy',
		'camera=(), microphone=(), geolocation=(), interest-cohort=()'
	);

	return response;
};
