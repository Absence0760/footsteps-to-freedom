<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import { Review } from "$lib/components/molecules";
import { HeroSection, CallToActionSection } from "$lib/components/organisms";

const testimonials = [
	{
		quote:
			"Whenever people ask me what moved me the most in South Africa, I hear myself say 'the city walks with Footsteps to Freedom'. Garth talks with so much passion that it feels like being absorbed in a film. Footsteps to Freedom leaves footprints on your soul.",
		author: "B. Bruyninckx",
		location: "Sint-Katelijne-Waver, Belgium",
		rating: 5,
	},
	{
		quote:
			"We had a couple of great days with Garth and highly recommend Footsteps to Freedom Tours when visiting Cape Town.",
		author: "Ricky W.",
		location: "TripAdvisor",
		rating: 5,
	},
	{
		quote:
			"Our lovely guide Zuzeka gave us an informative, authentic insight into the history of Langa where she lives. She answered all our questions honestly and showed us around. It was a great tour — highly recommended.",
		author: "Anne Marie C.",
		location: "London, United Kingdom",
		rating: 5,
	},
	{
		quote:
			"Our guide Ivor was superb. Very knowledgeable and patient to explain both what and why things happened the way they did. I found a new appreciation and respect for the South Africans. One of the best two hours I've invested and I'll come back for more.",
		author: "Wee T.",
		location: "TripAdvisor",
		rating: 5,
	},
	{
		quote:
			"Ivor's huge knowledge and enthusiasm of Cape Town contaminated us. He is an educated personality with a rich life experience. We appreciated his knowledge on history and politics in a broader world perspective. We walked for around 3.5 hours and it was the perfect start of our 4 days in Cape Town. We can highly recommend this tour with Ivor.",
		author: "Denbaar",
		location: "TripAdvisor",
		rating: 5,
	},
	{
		quote:
			"Garth is a fountain of knowledge about not only SA history, but US and world history as well, easily weaving together parallels and contrasts between the two. He is a natural teacher and a delightful tour guide. I highly recommend both the tour company and Garth personally.",
		author: "Leslye F.",
		location: "Boca Raton, Florida",
		rating: 5,
	},
	{
		quote:
			"Footsteps to Freedom went above and beyond to ensure I could experience a Township tour before I left the country. Our guide Nela provided great information, having lived his whole life in The White City. He had great insights into the history and present status of the townships. I highly recommend taking a tour with Footsteps to Freedom.",
		author: "Mark C.",
		location: "TripAdvisor",
		rating: 5,
	},
	{
		quote:
			"Garth is incredibly knowledgeable, extremely personable, and just all around easy to be with. Everywhere we went, he pointed out things we would've missed on our own, giving us in-depth history and background. Conversation always felt easy, like wandering around Cape Town with a family friend. We've hired guides in cities all over the world, and Garth was definitely one of the best!",
		author: "jbsf79",
		location: "Los Angeles, California",
		rating: 5,
	},
];
</script>

<div class="page-container">
  <HeroSection image="/zebra-giraffe-standing-rear-green-trees.jpg" mobileImage="/vertical-distant-table-mountain.jpg" alt="Cape Town landscape">
    <Text variant="h1">What Our Travelers Say</Text>
    <Text variant="h2">Real experiences from guests who've walked with us</Text>
  </HeroSection>

  <section class="intro-section">
    <Text variant="p">
      We're proud of the memories and connections we've helped create. Here's what some of our guests have shared about their experiences with Footsteps to Freedom.
    </Text>
    <div class="tripadvisor-link">
      <a
        href="https://www.tripadvisor.com/Attraction_Review-g312659-d2098098-Reviews-Footsteps_to_Freedom_Tours-Cape_Town_Central_Western_Cape.html"
        target="_blank"
        rel="noopener noreferrer"
      >
        Read more reviews on TripAdvisor ↗
      </a>
    </div>
  </section>

  <section class="reviews-section">
    <div class="reviews-grid">
      {#each testimonials as testimonial}
        <Review
          rating={testimonial.rating}
          comment={testimonial.quote}
          author={`${testimonial.author}, ${testimonial.location}`}
        />
      {/each}
    </div>
  </section>

  <section class="cta-wrapper">
    <CallToActionSection />
  </section>
</div>

<style>
  .intro-section {
    padding: clamp(2rem, 5vw, 3rem) clamp(1rem, 3vw, 2rem);
    text-align: center;
    max-width: 700px;
    margin: 0 auto;
  }

  .tripadvisor-link {
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

  .reviews-section {
    padding: clamp(1rem, 3vw, 2rem) clamp(1rem, 3vw, 2rem) clamp(2rem, 5vw, 4rem);
  }

  .reviews-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: clamp(1.5rem, 3vw, 2rem);
    max-width: 1200px;
    margin: 0 auto;
  }

  .cta-wrapper {
    width: 100%;
  }

  @media (max-width: 640px) {
    .reviews-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
