<script lang="ts">
import type { Snippet } from "svelte";

interface Props {
	variant?: "h1" | "h2" | "h3" | "h4" | "p" | "caption" | "label";
	tag?: string;
	id?: string;
	children?: Snippet;
}

const {
	variant = "h1",
	tag = variant === "caption" || variant === "label" ? "span" : variant,
	id = "",
	children,
	...restProps
}: Props = $props();
</script>

<svelte:element this={tag} {id} class="text text--{variant}" {...restProps}>
  {#if children}
    {@render children()}
  {/if}
</svelte:element>

<style>
  .text {
    all: unset;
    display: block;
  }

  .text--h1 {
    color: var(--text-heading, #1A1A2E);
    font-size: clamp(2.5rem, 6vw, 3.5rem);
    font-weight: 800;
    line-height: 1.1;
    letter-spacing: -0.02em;
    text-align: center;
    margin-bottom: 1rem;
    max-width: 800px;
    margin-left: auto;
    margin-right: auto;
  }

  .text--h2 {
    color: var(--text-heading, #1A1A2E);
    font-size: clamp(1.75rem, 4.5vw, 2.5rem);
    font-weight: 700;
    line-height: 1.2;
    text-align: center;
    margin-bottom: 1rem;
    max-width: 600px;
    margin-left: auto;
    margin-right: auto;
  }

  .text--h3 {
    color: var(--text-heading, #1A1A2E);
    font-size: clamp(1.25rem, 3vw, 1.75rem);
    font-weight: 600;
    line-height: 1.4;
    text-align: center;
    margin-bottom: 0.75rem;
    max-width: 600px;
    margin-left: auto;
    margin-right: auto;
  }

  .text--h4 {
    color: var(--text-heading, #1A1A2E);
    font-size: clamp(1rem, 2vw, 1.25rem);
    font-weight: 600;
    line-height: 1.4;
    text-align: left;
    margin-bottom: 0.75rem;
    max-width: 600px;
  }

  .text--p {
    color: var(--text-body, #2D3748);
    font-size: clamp(1rem, 2.5vw, 1.125rem);
    font-weight: 400;
    line-height: 1.75;
    text-align: left;
    margin-bottom: 1rem;
    max-width: 70ch;
  }

  .text--caption {
    color: var(--text-muted, #4A5568);
    font-size: clamp(0.75rem, 2vw, 0.875rem);
    font-weight: 400;
    line-height: 1.5;
    text-align: left;
    margin-bottom: 0.5rem;
    max-width: 600px;
  }

  .text--label {
    color: var(--text-label, #2D3748);
    font-size: clamp(0.75rem, 2vw, 0.875rem);
    font-weight: 500;
    line-height: 1.5;
    text-align: left;
    margin-bottom: 0.5rem;
    max-width: 600px;
  }

  @media (max-width: 768px) {
    .text--h1 {
      font-size: clamp(1.75rem, 4.5vw, 2rem);
      margin-bottom: 0.75rem;
    }

    .text--h2 {
      font-size: clamp(1.25rem, 3.5vw, 1.5rem);
      margin-bottom: 0.75rem;
    }

    .text--h3 {
      font-size: clamp(1rem, 2.5vw, 1.25rem);
      margin-bottom: 0.5rem;
    }

    .text--p {
      font-size: clamp(0.875rem, 2.25vw, 0.9375rem);
    }

    .text--caption {
      font-size: clamp(0.6875rem, 1.75vw, 0.8125rem);
    }
  }
</style>