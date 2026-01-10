<script lang="ts" context="module">
import { Button } from "$lib/components";
import { defineMeta } from "@storybook/addon-svelte-csf";
import { expect, fn, userEvent } from "@storybook/test";

const { Story } = defineMeta({
	title: "Example/Button",
	component: Button,
	tags: ["autodocs"],
	args: { label: "Button", onclick: fn() },
	argTypes: {
		backgroundColor: { control: "color" },
		size: {
			control: { type: "select" },
			options: ["small", "medium", "large"],
		},
	},
	async play({ canvas, step }) {
		const button = canvas.getByRole("button");

		await step("Check the button text", async () => {
			await expect(button.textContent).toBe("Button");
		});

		await step("Click the button", async () => {
			await userEvent.click(button);
			await userEvent.tab();
		});
	},
});
</script>

<Story name="Primary" args={{ primary: true }} />
<Story name="Secondary" />
<Story name="Large" args={{ size: "large" }} />
<Story name="Small" args={{ size: "small" }} />
