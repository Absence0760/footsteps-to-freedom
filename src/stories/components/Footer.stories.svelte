<script lang="ts" context="module">
import { Footer } from "$lib/components";
import { defineMeta } from "@storybook/addon-svelte-csf";

const { Story } = defineMeta({
	title: "Example/Footer",
	component: Footer,
	tags: ["autodocs"],
	args: {
		year: new Date().getFullYear(),
		siteName: "Footsteps to Freedom",
		footerLinks: [
			{ name: "Privacy Policy", href: "/privacy" },
			{ name: "Terms of Service", href: "/terms" },
			{ name: "FAQs", href: "/faqs" },
		],
	},
	// Define controls for arguments in the Storybook UI
	argTypes: {
		year: { control: "number", description: "The current year for copyright." },
		siteName: { control: "text", description: "The name of the website." },
		footerLinks: {
			control: "object",
			description: "An array of navigation links for the footer.",
		},
	},
});
</script>

<Story name="Default" />
<Story name="Fewer Links" args={{
	footerLinks: [
		{ name: "Support", href: "/support" },
		{ name: "Blog", href: "/blog" },
	]
}} />
<Story name="Custom Site Name" args={{
	siteName: "My Awesome Site"
}} />
<Story name="Specific Year" args={{
	year: 2023
}} />
