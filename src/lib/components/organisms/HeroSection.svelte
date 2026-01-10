<script lang="ts">
import { Text } from "$lib/components/atoms";
import type { Snippet } from "svelte";

interface Props {
	image: string;
	alt: string;
	children?: Snippet;
}

const { image, alt, children }: Props = $props();
</script>

<section class="hero-section" style:background-image={`url(${image})`} aria-label={alt}>
    <div class="hero-content">
      {#if children}
        {@render children()}
      {/if}
  </div>
</section>

<style>
  .hero-section {
    min-height: 100vh;
    position: relative;
    height: 60vh;
    background-size: cover;
    background-position: center;
    background-attachment: fixed;
    background-repeat: no-repeat;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1;
  }


  .hero-content {
    text-align: center;
    color: var(--text-header-dark, #FFFFFF);
    padding: clamp(1.5rem, 4vw, 2rem);
    max-width: 800px;
    margin: 0 auto;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  }

  @media (prefers-color-scheme: dark) {
    .hero-overlay {
      background-color: rgba(0, 0, 0, 0.5); /* Slightly darker overlay for dark mode */
    }

    .hero-content {
      color: var(--text-header-dark, #e5e7eb);
    }
  }

  @media (max-width: 600px) {
    .hero-section {
      height: 50vh; /* Shorter height for mobile */
    }

    .hero-content {
      padding: clamp(1rem, 3vw, 1.5rem);
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .hero-section {
      background-attachment: scroll; /* Disable fixed background for accessibility */
    }
  }
</style>
