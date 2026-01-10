<script lang="ts">
import { goto } from "$app/navigation";
import type { IconName } from "$lib/components/atoms";
import type { Snippet } from "svelte";

export type Variant = "primary" | "secondary" | "contingent";

interface Props {
	variant?: Variant;
	iconType?: IconName;
	imageSrc?: string;
	slug?: string;
	children: Snippet;
}

let {
	variant = "primary",
	iconType,
	imageSrc,
	slug = "",
	children,
	...restProps
}: Props = $props();

let imageAlt = "/image.png";
let isHovered = $state(false);

function handleClick() {
	if (slug) goto(`/${slug}`);
}
</script>

<div
  class="card {variant}"
  class:hovered={isHovered}
  class:clickable={slug}
  role={slug ? "button" : undefined}
  tabindex={slug ? 0 : undefined}
  aria-labelledby="card-title-{variant}"
  onclick={handleClick}
  onmouseover={() => (isHovered = true)}
  onmouseout={() => (isHovered = false)}
  onfocus={() => (isHovered = true)}
  onblur={() => (isHovered = false)}
  {...restProps}
>
  {#if imageSrc}
    <div class="image-container">
      <img src={imageSrc} alt={imageAlt} class="card-image" />
    </div>
  {/if}

  {#if children}
    <div class="content">
      {@render children()}
    </div>
  {/if}
</div>

<style>
  .card {
    position: relative;
    display: flex;
    flex-direction: column;
    padding: var(--card-padding, clamp(1rem, 4vw, 1.5rem));
    margin: var(--card-margin, 1rem);
    border-radius: var(--card-radius, 0.75rem);
    box-shadow: var(--shadow-md, 0 4px 6px -1px rgba(0, 0, 0, 0.1));
    background: var(--card-bg, #ffffff);
    transition:
      transform 0.3s ease,
      box-shadow 0.3s ease;
    box-sizing: border-box;
    max-width: 100%;
  }

  .card.clickable {
    cursor: pointer;
  }

  .card.clickable:hover,
  .card.clickable.hovered:not(.disabled) {
    transform: translateY(-6px);
    box-shadow: var(--shadow-xl, 0 20px 30px -8px rgba(0, 0, 0, 0.2));
  }

  .card.disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .card.primary {
    width: 100%;
    min-width: 280px;
    max-width: 450px;
    min-height: 280px;
    --card-bg: #ffffff;
    --icon-color: var(--primary);
    box-shadow: var(--shadow-vibrant);
  }

  .card.primary.clickable:hover {
    box-shadow: var(--shadow-vibrant), var(--shadow-xl);
  }

  .card.secondary {
    width: 100%;
    max-width: 900px;
    min-height: 300px;
    --card-padding: clamp(2rem, 6vw, 3rem);
    --card-bg: var(--secondary-bg, #FFF5E6);
    --icon-color: var(--secondary);
  }

  .card.contingent {
    width: 100%;
    max-width: 1200px;
    min-height: 350px;
    --card-bg: var(--card-bg, #ffffff);
    --icon-color: var(--accent-purple);
  }

  .image-container {
    flex: 0 0 auto;
    width: 100%;
  }

  .card.horizontal .image-container {
    width: var(--image-width, 40%);
    max-width: 300px;
  }

  .card-image {
    width: 100%;
    height: auto;
    border-radius: var(--image-radius, 0.5rem);
    object-fit: cover;
    aspect-ratio: var(--image-aspect-ratio, 16 / 9);
  }

  .content {
    align-items: center;
    display: flex;
    flex-direction: column;
    gap: var(--content-gap, 1rem);
    flex: 1;
  }

  .action-icon {
    align-self: center;
    transition: transform 0.2s ease;
  }
  @media (max-width: 768px) {
    .card {
      --card-padding: clamp(1rem, 4vw, 1.5rem);
      --card-margin: 1rem;
    }

    .image-container {
      width: 100%;
      max-width: none;
    }
  }

  @media (max-width: 600px) {
    .card.primary {
      max-width: 100%;
    }
    .card.secondary {
      max-width: 100%;
    }
    .card.contingent {
      max-width: 100%;
    }
  }
</style>
