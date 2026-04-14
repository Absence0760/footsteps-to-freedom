import { sveltekit } from "@sveltejs/kit/vite";
import Icons from "unplugin-icons/vite";
import { defineConfig } from "vitest/config";

export default defineConfig({
	plugins: [
		sveltekit(),
		Icons({
			autoInstall: true,
			compiler: "svelte",
		}),
	],
	test: {
		include: ["src/**/*.test.ts"],
	},
});
