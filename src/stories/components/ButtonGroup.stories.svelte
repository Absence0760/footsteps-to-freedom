<script lang="ts" context="module">
import ButtonGroup from "$lib/components/ButtonGroup.svelte";
import { defineMeta } from "@storybook/addon-svelte-csf";

const { Story } = defineMeta({
	title: "Example/ButtonGroup",
	component: ButtonGroup,
	// This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
	tags: ["autodocs"],
	parameters: {
		// More on how to position stories at: https://storybook.js.org/docs/configure/story-layout
		layout: "fullscreen",
	},
});
</script>

<Story name="Main" />
