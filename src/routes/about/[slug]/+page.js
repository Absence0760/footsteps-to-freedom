import { teamMembers } from "$lib/data/team";
import { error } from "@sveltejs/kit";

export const prerender = true;

// Tell SvelteKit which slugs to prerender
export function entries() {
	return teamMembers.map((member) => ({
		slug: member.slug,
	}));
}

export async function load({ params }) {
	const member = teamMembers.find((t) => t.slug === params.slug);

	if (!member) {
		error(404, "member not found");
	}

	return {
		member,
	};
}
