<script lang="ts">
import { base } from '$app/paths';
import { Text } from "$lib/components/atoms";
import type { Snippet } from "svelte";

interface Props {
	image: string;
	alt: string;
	children?: Snippet;
}

const { image: rawImage, alt, children }: Props = $props();

// Prepend base path to relative URLs
const image = rawImage.startsWith('/') && !rawImage.startsWith('//') ? `${base}${rawImage}` : rawImage;
</script>

<section class="hero-section" style:background-image={`url(${image})`} aria-label={alt}>
  <div class="hero-overlay"></div>
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
    padding-top: 5rem;
    background-size: cover;
    background-position: center;
    background-attachment: fixed;
    background-repeat: no-repeat;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1;
  }

  .hero-overlay {
    position: absolute;
    inset: 0;
    background-color: rgba(0, 0, 0, 0.45);
    z-index: 0;
  }

  .hero-content {
    position: relative;
    z-index: 1;
    text-align: center;
    padding: clamp(1.5rem, 4vw, 2rem);
    max-width: 800px;
    margin: 0 auto;
  }

  .hero-content :global(.text) {
    color: #ffffff;
    font-size: 1.1rem;
    font-weight: 600;
    line-height: 1.5;
    margin-bottom: 0.5rem;
  }

  @media (max-width: 600px) {
    .hero-section {
      height: 50vh;
    }

    .hero-content {
      padding: clamp(1rem, 3vw, 1.5rem);
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .hero-section {
      background-attachment: scroll;
    }
  }
</style>
