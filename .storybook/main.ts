import { dirname, join } from "node:path";
import type { StorybookConfig } from "@storybook/sveltekit";
import Icons from "unplugin-icons/vite";
import { mergeConfig } from "vite";
import { optimizeDeps } from "vite";
import type { InlineConfig } from "vite";

const config: StorybookConfig = {
	stories: ["../src/**/*.stories.@(ts|svelte)"],
	addons: [
		getAbsolutePath("@storybook/addon-links"),
		getAbsolutePath("@storybook/addon-essentials"),
		getAbsolutePath("@storybook/addon-interactions"),
		"@storybook/addon-svelte-csf",
		getAbsolutePath("@storybook/addon-a11y"),
	],

	framework: {
		name: getAbsolutePath("@storybook/sveltekit"),
		options: {},
	},

	async viteFinal(config) {
		return mergeConfig(svelteDocgenUnpluginIconsWorkaround(config), {
			plugins: [
				Icons({
					autoInstall: true,
					compiler: "svelte",
				}),
			],

			// Add dependencies to pre-optimization
			optimizeDeps: {
				// include: ["storybook-dark-mode"],
			},
		});
	},
};

export default config;

/**
 * This function is used to resolve the absolute path of a package.
 * It is needed in projects that use Yarn PnP or are set up within a monorepo.
 */
function getAbsolutePath(value: string) {
	return dirname(require.resolve(join(value, "package.json")));
}

// https://github.com/storybookjs/storybook/issues/20562
/**
 * Workaround to allow svelte docgen plugin to work with unplugin-icons.
 */
function svelteDocgenUnpluginIconsWorkaround(config: InlineConfig) {
	if (!config.plugins) return config;

	const [_internalPlugins, ...userPlugins] = config.plugins;
	const docgenPlugin = userPlugins.find(
		(plugin) => plugin.name === "storybook:svelte-docgen-plugin",
	);
	if (docgenPlugin) {
		const origTransform = docgenPlugin.transform;
		const newTransform = (code, id, options) => {
			if (id.startsWith("~icons/")) {
				return;
			}
			return origTransform?.call(docgenPlugin, code, id, options);
		};
		docgenPlugin.transform = newTransform;
		docgenPlugin.enforce = "post";
	}
	return config;
}
