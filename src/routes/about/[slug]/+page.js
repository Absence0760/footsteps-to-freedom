import { error } from "@sveltejs/kit";

export const prerender = true;

const teamModules = import.meta.glob('/src/content/team/*.md', { eager: true });
const allMembers = Object.values(teamModules).map((mod) => mod.metadata);

export function entries() {
	return allMembers.map((member) => ({
		slug: member.slug,
	}));
}

export async function load({ params }) {
	const entry = Object.entries(teamModules).find(([, mod]) => mod.metadata.slug === params.slug);
	if (!entry) error(404, "Team member not found");
	const [, mod] = entry;
	return {
		member: {
			...mod.metadata,
			component: mod.default,
		},
	};
}
