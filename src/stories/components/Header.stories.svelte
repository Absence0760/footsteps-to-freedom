<script lang="ts" context="module">
import { Header } from "$lib/components";
import { defineMeta } from "@storybook/addon-svelte-csf";
import { expect, userEvent, within } from "@storybook/test";

const { Story } = defineMeta({
	title: "Example/Header",
	component: Header,
	tags: ["autodocs"],
	args: {
		siteName: "Footsteps To Freedom Tours",
		navLinks: [
			{ name: "Home", href: "/" },
			{ name: "Tours", href: "/tours" },
			{ name: "About Us", href: "/about" },
			{ name: "Contact", href: "/contact" },
		],
		isMenuOpen: false,
	},
	argTypes: {
		siteName: {
			control: "text",
			description: "The name of the site displayed in the header.",
		},
		navLinks: {
			control: "object",
			description: "An array of navigation links (name and href).",
		},
		isMenuOpen: {
			control: "boolean",
			description: "Controls the visibility of the mobile navigation menu.",
		},
	},
});
</script>

<Story name="Default" />
<Story
  name="Custom Content"
  args={{
    siteName: "Global Expeditions",
    navLinks: [
      { name: "Destinations", href: "/destinations" },
      { name: "Guides", href: "/guides" },
      { name: "Contact Us", href: "/contact-us" },
    ],
  }}
/>

