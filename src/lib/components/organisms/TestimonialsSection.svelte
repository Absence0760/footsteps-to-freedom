<script lang="ts">
import { Text } from "$lib/components/atoms";
import { Review } from "$lib/components/molecules";
import { onDestroy, onMount } from "svelte";
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
				"Whenever people ask me what moved me the most in South Africa, I hear myself say 'the city walks with Footsteps to Freedom'. Garth talks with so much passion that it feels like being absorbed in a film. Footsteps to Freedom leaves footprints on your soul.",
			author: "B. Bruyninckx",
			location: "Sint-Katelijne-Waver, Belgium",
		},
		{
			quote:
				"We had a couple of great days with Garth and highly recommend Footsteps to Freedom Tours when visiting Cape Town.",
			author: "Ricky W.",
			location: "TripAdvisor",
		},
		{
			quote:
				"Our lovely guide Zuzeka gave us an informative, authentic insight into the history of Langa where she lives. She answered all our questions honestly and showed us around. It was a great tour — highly recommended.",
			author: "Anne Marie C.",
			location: "London, United Kingdom",
		},
		{
			quote:
				"Our guide Ivor was superb. Very knowledgeable and patient to explain both what and why things happened the way they did. I found a new appreciation and respect for the South Africans. One of the best two hours I've invested and I'll come back for more.",
			author: "Wee T.",
			location: "TripAdvisor",
		},
		{
			quote:
				"Ivor's huge knowledge and enthusiasm of Cape Town contaminated us. He is an educated personality with a rich life experience. We appreciated his knowledge on history and politics in a broader world perspective. We walked for around 3.5 hours and it was the perfect start of our 4 days in Cape Town. We can highly recommend this tour with Ivor.",
			author: "Denbaar",
			location: "TripAdvisor",
		},
		{
			quote:
				"Garth is a fountain of knowledge about not only SA history, but US and world history as well, easily weaving together parallels and contrasts between the two. He is a natural teacher and a delightful tour guide. I highly recommend both the tour company and Garth personally.",
			author: "Leslye F.",
			location: "Boca Raton, Florida",
		},
		{
			quote:
				"Footsteps to Freedom went above and beyond to ensure I could experience a Township tour before I left the country. Our guide Nela provided great information, having lived his whole life in The White City. He had great insights into the history and present status of the townships. I highly recommend taking a tour with Footsteps to Freedom.",
			author: "Mark C.",
			location: "TripAdvisor",
		},
		{
			quote:
				"Garth is incredibly knowledgeable, extremely personable, and just all around easy to be with. Everywhere we went, he pointed out things we would've missed on our own, giving us in-depth history and background. Conversation always felt easy, like wandering around Cape Town with a family friend. We've hired guides in cities all over the world, and Garth was definitely one of the best!",
			author: "jbsf79",
			location: "Los Angeles, California",
		},
	],
	children,
}: Props = $props();

let current = $state(0);
let transitioning = $state(false);
let interval: ReturnType<typeof setInterval>;

function goTo(index: number) {
	if (transitioning) return;
	transitioning = true;
	current = (index + testimonials.length) % testimonials.length;
	setTimeout(() => (transitioning = false), 400);
}

function prev() { goTo(current - 1); resetTimer(); }
function next() { goTo(current + 1); resetTimer(); }

function resetTimer() {
	clearInterval(interval);
	interval = setInterval(() => goTo(current + 1), 5000);
}

onMount(() => { interval = setInterval(() => goTo(current + 1), 5000); });
onDestroy(() => clearInterval(interval));
</script>

<section>
  <Text variant="h2">What Our Travelers Say</Text>

  <div class="carousel">
    <button class="arrow left" onclick={prev} aria-label="Previous review">&#8592;</button>

    <div class="slides" class:fade={transitioning}>
      {#each [0, 1, 2] as offset}
        {@const t = testimonials[(current + offset) % testimonials.length]}
        <div class="slide">
          <Review
            rating={5}
            comment={t.quote}
            author={`${t.author}, ${t.location}`}
          />
        </div>
      {/each}
    </div>

    <button class="arrow right" onclick={next} aria-label="Next review">&#8594;</button>
  </div>

  <div class="dots">
    {#each testimonials as _, i}
      <button
        class="dot"
        class:active={i === current}
        onclick={() => { goTo(i); resetTimer(); }}
        aria-label={`Go to review ${i + 1}`}
      ></button>
    {/each}
  </div>

  <div class="tripadvisor-link">
    <a
      href="https://www.tripadvisor.com/Attraction_Review-g312659-d2098098-Reviews-Footsteps_to_Freedom_Tours-Cape_Town_Central_Western_Cape.html"
      target="_blank"
      rel="noopener noreferrer"
    >
      Read more reviews on TripAdvisor ↗
    </a>
  </div>

  {#if children}
    {@render children()}
  {/if}
</section>

<style>
  .carousel {
    display: flex;
    align-items: center;
    gap: 1rem;
    max-width: 1100px;
    margin: 1.5rem auto 0;
  }

  .slides {
    flex: 1;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1.5rem;
    transition: opacity 0.4s ease;
    align-items: stretch;
  }

  .slides.fade {
    opacity: 0;
  }

  .slide {
    display: flex;
    flex-direction: column;
  }

  .slide :global(.review) {
    flex: 1;
    height: 100%;
  }

  @media (max-width: 768px) {
    .slides {
      grid-template-columns: 1fr;
    }

    .slide:not(:first-child) {
      display: none;
    }
  }

  .arrow {
    flex-shrink: 0;
    background: none;
    border: 2px solid var(--button-primary-bg, #4f46e5);
    color: var(--button-primary-bg, #4f46e5);
    border-radius: 50%;
    width: 2.5rem;
    height: 2.5rem;
    font-size: 1.1rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color 0.2s ease, color 0.2s ease;
  }

  .arrow:hover {
    background-color: var(--button-primary-bg, #4f46e5);
    color: #ffffff;
  }

  .dots {
    display: flex;
    justify-content: center;
    gap: 0.5rem;
    margin-top: 1rem;
  }

  .dot {
    width: 0.6rem;
    height: 0.6rem;
    border-radius: 50%;
    background-color: #d1d5db;
    border: none;
    cursor: pointer;
    transition: background-color 0.3s ease;
    padding: 0;
  }

  .dot.active {
    background-color: var(--button-primary-bg, #4f46e5);
  }

  .tripadvisor-link {
    text-align: center;
    margin-top: 1.5rem;
  }

  .tripadvisor-link a {
    color: #00aa6c;
    font-weight: 600;
    font-size: 1rem;
    text-decoration: none;
  }

  .tripadvisor-link a:hover {
    text-decoration: underline;
  }
</style>
