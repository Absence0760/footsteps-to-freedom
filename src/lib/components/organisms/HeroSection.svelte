<script lang="ts">
import { base } from '$app/paths';
import { Text } from "$lib/components/atoms";
import type { Snippet } from "svelte";

interface Props {
	image: string;
	mobileImage?: string;
	alt: string;
	children?: Snippet;
}

const { image: rawImage, mobileImage: rawMobileImage, alt, children }: Props = $props();

// Prepend base path to relative URLs
const resolve = (src: string) => src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;
const image = resolve(rawImage);
const mobileImage = rawMobileImage ? resolve(rawMobileImage) : null;
</script>

<section
  class="hero-section"
  style:--hero-bg={`url(${image})`}
  style:--hero-mobile-bg={mobileImage ? `url(${mobileImage})` : `url(${image})`}
  aria-label={alt}
>
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
    background-image: var(--hero-bg);
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
      background-image: var(--hero-mobile-bg);
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
