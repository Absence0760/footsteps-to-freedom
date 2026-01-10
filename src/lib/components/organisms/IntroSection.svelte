<script lang="ts">
import { Link, Text } from "$lib/components/atoms/";
import type { Snippet } from "svelte";

interface Props {
	headline?: string;
	description?: string;
	ctaText?: string;
	ctaHref?: string;
	children?: Snippet;
}

const {
	headline = "Walk Through History. Discover Freedom.",
	description = `
    The natural beauty of the Cape is self evident and widely recorded. 
    Less evident are the human stories behind its scenic splendour.
    Against this backdrop, Footsteps to Freedom brings together curious travellers and well-informed and personable tourist guides, 
    to create enduring memories and firm friendships. Individually we are all unique, with our own history and our own interests. 
    Likewise, our tours are also unique. Select a tour that suits you (Cultural and Historical, Leisure, Active), 
    or tell us about yourself and your interests and we will create an unforgettable experience for you.
    `,
	ctaText = "Learn More About Us",
	ctaHref = "/about",
	children,
}: Props = $props();
</script>

<section class="intro-section">
  <Text variant="h2">
    {@html headline}
  </Text>
  <Text variant="p">
    {@html description}
  </Text>
  <Link variant="inline-text" to={ctaHref}>
    {ctaText}
  </Link>
  {#if children}
    {@render children()}
  {/if}
</section>

<style>
  .intro-section {
    padding: 3rem 2rem;
    background-color: var(--card-bg, #ffffff);
    border-radius: 0.5rem;
    box-shadow: var(--shadow-lg, 0 10px 15px -3px rgba(0, 0, 0, 0.1));
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1rem;
    max-width: 1200px;
    margin: 2rem auto;
  }
</style>