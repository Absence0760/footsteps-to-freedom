<script lang="ts">
import { Image, Link, Text } from "$lib/components/atoms";
import type { Snippet } from "svelte";

interface Story {
	title: string;
	snippet: string;
	image: string;
	link: string;
}

interface Props {
	stories?: Story[];
	children?: Snippet;
}

const {
	stories = [
		{
			title: "The Silent Code of the Underground Railroad",
			snippet:
				"Discover the ingenious methods and hidden networks that guided countless individuals to freedom.",
			image: "",
			link: "/blog/underground-railroad-code",
		},
		{
			title: "Unsung Heroes of the Civil Rights Movement",
			snippet:
				"Beyond the famous names, explore the courage of everyday people who shaped history.",
			image: "",
			link: "/blog/unsung-civil-rights-heroes",
		},
		{
			title: "Voices of Suffrage: A Century of Struggle",
			snippet:
				"Learn about the fierce determination and strategic brilliance behind women's fight for the vote.",
			image: "",
			link: "/blog/womens-suffrage-voices",
		},
		{
			title: "Echoes of Emancipation: A New Beginning",
			snippet:
				"Delve into the complex aftermath of freedom and the journey towards true equality.",
			image: "",
			link: "/blog/emancipation-echoes",
		},
	],
	children,
}: Props = $props();
</script>

<section>
  <Text
    level={2}
    color="var(--text-heading, #1f2937)"
    textAlign="center"
    marginBottom="2rem"
  >
    Inspiring Stories from History
  </Text>
  <div class="stories-grid">
    {#each stories as story}
      <div class="story-card">
        <Image
          src={story.image}
          alt={story.title}
          width="100%"
          height="200px"
          borderRadius="0.5rem"
          shadow="var(--shadow-md, 0 4px 6px -1px rgba(0, 0, 0, 0.1))"
        />
        <Text
          level={3}
          color="var(--text-heading, #1f2937)"
          fontSize="1.25rem"
          marginBottom="0.5rem"
        >
          {story.title}
        </Text>
        <Text color="var(--text-body, #4b5563)" fontSize="0.875rem" marginBottom="1rem">
          {story.snippet}
        </Text>
        <Link
          href={story.link}
          color="var(--button-primary-bg, #4f46e5)"
        >
          Read More
        </Link>
      </div>
    {/each}
  </div>
  {#if children}
    {@render children()}
  {/if}
</section>

<style>
  .stories-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }

  .story-card {
    display: flex;
    flex-direction: column;
    padding: 1rem;
    background-color: var(--card-bg, #ffffff);
    border-radius: 0.5rem;
    box-shadow: var(--shadow-md, 0 4px 6px -1px rgba(0, 0, 0, 0.1));
  }
</style>