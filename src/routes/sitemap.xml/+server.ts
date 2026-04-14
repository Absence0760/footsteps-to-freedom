import type { RequestHandler } from "./$types";

const SITE_URL = "https://absence0760.github.io/footsteps-to-freedom";

const staticRoutes = [
	"",
	"/about",
	"/tours",
	"/contact",
	"/faqs",
	"/privacy",
	"/terms",
	"/reviews",
];

function getContentSlugs(modules: Record<string, unknown>): string[] {
	return Object.values(modules).map(
		(mod) => (mod as { metadata: { slug: string } }).metadata.slug
	);
}

const tourModules = import.meta.glob("/src/content/tours/*.md", { eager: true });
const teamModules = import.meta.glob("/src/content/team/*.md", { eager: true });

const tourSlugs = getContentSlugs(tourModules);
const teamSlugs = getContentSlugs(teamModules);

function urlEntry(path: string): string {
	return `<url><loc>${SITE_URL}${path}</loc></url>`;
}

export const prerender = true;

export const GET: RequestHandler = async () => {
	const urls = [
		...staticRoutes.map((route) => urlEntry(route)),
		...tourSlugs.map((slug) => urlEntry(`/tours/${slug}`)),
		...teamSlugs.map((slug) => urlEntry(`/about/${slug}`)),
	];

	return new Response(
		`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="https://www.sitemaps.org/schemas/sitemap/0.9">
${urls.join("\n")}
</urlset>`.trim(),
		{
			headers: {
				"Content-Type": "application/xml",
			},
		},
	);
};
