<script lang="ts">
import type { Snippet } from "svelte";

interface Props {
	type?: "success" | "error" | "warning" | "info";
	children?: Snippet;
	duration?: number; // New prop for dismissal duration
}

const { type = "info", children, duration = 3000 }: Props = $props();
let isVisible = $state(true); // Track visibility

// Auto-dismiss after duration
if (duration > 0) {
	setTimeout(() => {
		isVisible = false;
	}, duration);
}

const colors = $derived(
	{
		success: "var(--button-primary-bg, #4f46e5)",
		error: "#ef4444",
		warning: "#f59e0b",
		info: "var(--text-body, #4b5563)",
	}[type],
);
</script>

{#if isVisible}
  <div
    class="alert"
    style:background-color={colors}
    style:color="var(--button-primary-text, #ffffff)"
  >
    {#if children}
      {@render children()}
    {/if}
  </div>
{/if}

<style>
  .alert {
    position: fixed;
    top: 1rem;
    right: 1rem;
    padding: 1rem;
    border-radius: 0.5rem;
    box-shadow: var(--shadow-sm, 0 1px 2px 0 rgba(0, 0, 0, 0.05));
    font-size: 0.875rem;
    max-width: 200px;
    z-index: 1000; /* Ensure it appears above other content */
  }
</style>