import { error } from "@sveltejs/kit";

export const prerender = true;

/** @type {Record<string, any>} */
const tourModules = import.meta.glob('/src/content/tours/*.md', { eager: true });
const allTours = Object.values(tourModules).map((mod) => mod.metadata);

export function entries() {
	return allTours.map((tour) => ({
		slug: tour.slug,
	}));
}

export async function load({ params }) {
	const entry = Object.entries(tourModules).find(([, mod]) => mod.metadata.slug === params.slug);
	if (!entry) error(404, "Tour not found");
	const [, mod] = entry;
	return {
		tour: {
			...mod.metadata,
			component: mod.default,
		},
	};
}
