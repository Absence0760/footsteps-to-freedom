<script lang="ts">
import { Text } from "$lib/components/atoms";
import { Review } from "$lib/components/molecules";
import type { Snippet } from "svelte";

interface Testimonial {
	quote: string;
	author: string;
	location: string;
}

interface Props {
	testimonials?: Testimonial[];
	children?: Snippet;
}

const {
	testimonials = [
		{
			quote:
				"Footsteps To Freedom provided an incredibly moving and educational experience. Our guide was exceptional!",
			author: "Sarah M.",
			location: "New York, NY",
		},
		{
			quote:
				"I learned so much more than I ever could from a textbook. Truly an unforgettable journey.",
			author: "James T.",
			location: "Atlanta, GA",
		},
		{
			quote:
				"The attention to detail and respect for history on this tour was unparalleled. Highly recommend!",
			author: "Emily R.",
			location: "Chicago, IL",
		},
	],
	children,
}: Props = $props();
</script>

<section>
  <Text variant="h2">
    What Our Travelers Say
  </Text>
  <div class="testimonials-grid">
    {#each testimonials as testimonial}
      <Review
        rating={5}
        comment={testimonial.quote}
        author={`${testimonial.author}, ${testimonial.location}`}
      />
    {/each}
  </div>
  {#if children}
    {@render children()}
  {/if}
</section>

<style>
  .testimonials-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }
</style>