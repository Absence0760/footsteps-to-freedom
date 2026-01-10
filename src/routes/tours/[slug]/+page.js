import { allTours } from "$lib/data/tours";
import { error } from "@sveltejs/kit";

export const prerender = true;

// Tell SvelteKit which slugs to prerender
export function entries() {
	return allTours.map((tour) => ({
		slug: tour.slug,
	}));
}

export async function load({ params }) {
	const tour = allTours.find((t) => t.slug === params.slug);

	if (!tour) {
		error(404, "Tour not found");
	}

	return {
		tour,
	};
}
