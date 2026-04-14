<script lang="ts">
import { Icon, Text } from "$lib/components/atoms";
import type { Snippet } from "svelte";

interface Feature {
	icon: string;
	title: string;
	description: string;
}

interface Props {
	features?: Feature[];
	children?: Snippet;
}

const {
	features = [
		{
			icon: "book",
			title: "Expert Historians",
			description:
				"Journey with passionate guides who bring history to life with deep knowledge.",
		},
		{
			icon: "star",
			title: "Immersive Experiences",
			description:
				"Step into the past through vivid storytelling and authentic site visits.",
		},
		{
			icon: "globe",
			title: "Ethical & Responsible Travel",
			description:
				"Our tours support local communities and promote respectful engagement with heritage.",
		},
		{
			icon: "book",
			title: "Unforgettable Learning",
			description:
				"Gain profound insights into pivotal moments of human struggle and triumph.",
		},
	],
	children,
}: Props = $props();
</script>

<section>
  <Text variant="h2" style="margin-bottom: 2rem;">
    Why Choose Footsteps To Freedom
  </Text>
  <div class="features-grid">
    {#each features as feature}
      <div class="feature-card">
        <Icon
          icon="contact"
          color="var(--button-primary-bg, #4f46e5)"
          ariaLabel={feature.title}
        />
        <Text variant="h3" style="font-size: 1.25rem; margin-bottom: 0.5rem;">
          {feature.title}
        </Text>
        <Text variant="caption" style="font-size: 0.875rem;">
          {feature.description}
        </Text>
      </div>
    {/each}
  </div>
  {#if children}
    {@render children()}
  {/if}
</section>

<style>
  .features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }

  .feature-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    padding: 1.5rem;
    background-color: var(--card-bg, #ffffff);
    border-radius: 0.5rem;
    box-shadow: var(--shadow-md, 0 4px 6px -1px rgba(0, 0, 0, 0.1));
  }
</style>