<script lang="ts">
import { base } from '$app/paths';
import { HeroSection, CallToActionSection } from "$lib/components/organisms";

const resolve = (src: string) =>
  src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;

const testimonials = [
	{
		quote: "Whenever people ask me what moved me the most in South Africa, I hear myself say 'the city walks with Footsteps to Freedom'. Garth talks with so much passion that it feels like being absorbed in a film. Footsteps to Freedom leaves footprints on your soul.",
		author: "B. Bruyninckx",
		location: "Sint-Katelijne-Waver, Belgium",
	},
	{
		quote: "We had a couple of great days with Garth and highly recommend Footsteps to Freedom Tours when visiting Cape Town.",
		author: "Ricky W.",
		location: "TripAdvisor",
	},
	{
		quote: "Our lovely guide Zuzeka gave us an informative, authentic insight into the history of Langa where she lives. She answered all our questions honestly and showed us around. It was a great tour — highly recommended.",
		author: "Anne Marie C.",
		location: "London, United Kingdom",
	},
	{
		quote: "Our guide Ivor was superb. Very knowledgeable and patient to explain both what and why things happened the way they did. I found a new appreciation and respect for the South Africans. One of the best two hours I've invested and I'll come back for more.",
		author: "Wee T.",
		location: "TripAdvisor",
	},
	{
		quote: "Ivor's huge knowledge and enthusiasm of Cape Town contaminated us. He is an educated personality with a rich life experience. We appreciated his knowledge on history and politics in a broader world perspective. We walked for around 3.5 hours and it was the perfect start of our 4 days in Cape Town. We can highly recommend this tour with Ivor.",
		author: "Denbaar",
		location: "TripAdvisor",
	},
	{
		quote: "Garth is a fountain of knowledge about not only SA history, but US and world history as well, easily weaving together parallels and contrasts between the two. He is a natural teacher and a delightful tour guide. I highly recommend both the tour company and Garth personally.",
		author: "Leslye F.",
		location: "Boca Raton, Florida",
	},
	{
		quote: "Footsteps to Freedom went above and beyond to ensure I could experience a Township tour before I left the country. Our guide Nela provided great information, having lived his whole life in The White City. He had great insights into the history and present status of the townships. I highly recommend taking a tour with Footsteps to Freedom.",
		author: "Mark C.",
		location: "TripAdvisor",
	},
	{
		quote: "Garth is incredibly knowledgeable, extremely personable, and just all around easy to be with. Everywhere we went, he pointed out things we would've missed on our own, giving us in-depth history and background. Conversation always felt easy, like wandering around Cape Town with a family friend. We've hired guides in cities all over the world, and Garth was definitely one of the best!",
		author: "jbsf79",
		location: "Los Angeles, California",
	},
];

const featured = testimonials[0];
const rest = testimonials.slice(1);
</script>

<div class="page-container">
  <!-- Hero -->
  <HeroSection image="/zebra-giraffe-standing-rear-green-trees.jpg" mobileImage="/vertical-penguin.jpg" alt="South African wildlife — zebra and giraffe">
    <h1 class="hero-h1">What Our Travelers Say</h1>
    <p class="hero-sub">Real experiences from guests who've walked with us</p>
  </HeroSection>

  <!-- Featured quote -->
  <section class="featured-section">
    <div class="featured-inner">
      <div class="featured-stars" aria-label="5 out of 5 stars">★★★★★</div>
      <blockquote class="featured-quote">
        <span class="featured-mark" aria-hidden="true">"</span>
        <p>{featured.quote}</p>
      </blockquote>
      <p class="featured-attribution">— {featured.author}, {featured.location}</p>
    </div>
  </section>

  <!-- Reviews grid -->
  <section class="reviews-section">
    <div class="reviews-grid">
      {#each rest as review}
        <div class="quote-card">
          <div class="quote-card__stars" aria-label="5 out of 5 stars">★★★★★</div>
          <p class="quote-card__text">{review.quote}</p>
          <div class="quote-card__footer">
            <span class="quote-card__author">{review.author}</span>
            <span class="quote-card__location">{review.location}</span>
          </div>
        </div>
      {/each}
    </div>
  </section>

  <!-- TripAdvisor band -->
  <div class="tripadvisor-band">
    <div class="tripadvisor-inner">
      <div class="tripadvisor-text">
        <p class="tripadvisor-eyebrow">Rated Excellent</p>
        <h2 class="tripadvisor-heading">Read More on TripAdvisor</h2>
        <p class="tripadvisor-body">Hundreds of guests have shared their experiences. See what they have to say.</p>
      </div>
      <a
        href="https://www.tripadvisor.com/Attraction_Review-g312659-d2098098-Reviews-Footsteps_to_Freedom_Tours-Cape_Town_Central_Western_Cape.html"
        target="_blank"
        rel="noopener noreferrer"
        class="tripadvisor-btn"
      >
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="22" height="22" fill="currentColor" aria-hidden="true">
          <path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10 10-4.477 10-10S17.523 2 12 2zm0 18c-4.418 0-8-3.582-8-8s3.582-8 8-8 8 3.582 8 8-3.582 8-8 8zm-3.5-9a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm7 0a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm-3.5-2c-1.2 0-2.3.4-3.2 1H7l-1.5-1.5C6.8 7.2 9.3 6 12 6s5.2 1.2 6.5 3L17 10.5c-.9-.6-2-1-3.1-1H13c-.3-.3-.9-.5-1-.5zm0 7c-1.1 0-2-.4-2.7-1h5.4c-.7.6-1.6 1-2.7 1z"/>
        </svg>
        View All Reviews ↗
      </a>
    </div>
  </div>

  <!-- CTA -->
  <div
    class="cta-wrapper"
    style="background-image: url({resolve('/cape-town-from-table-mountain.jpg')})"
  >
    <CallToActionSection
      title="Ready to Make Your Own Memories?"
      description="Join us on a tour and experience Cape Town's stories for yourself."
      buttonLabel="Explore Tours"
      buttonHref="/tours"
    />
  </div>
</div>

<style>
  /* ── Hero text (bypasses Text component to avoid forced centering issues) */
  .hero-h1 {
    font-size: clamp(2rem, 5vw, 3.5rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0 0 0.75rem;
    letter-spacing: -0.02em;
    line-height: 1.1;
    text-shadow: 0 2px 8px rgba(0,0,0,0.35);
  }

  .hero-sub {
    font-size: clamp(1rem, 2.5vw, 1.25rem);
    font-weight: 500;
    color: rgba(255, 255, 255, 0.9);
    margin: 0;
    line-height: 1.5;
  }

  /* ── Featured quote ─────────────────────────────────────────────── */
  .featured-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
    padding: clamp(3.5rem, 7vw, 6rem) clamp(1.5rem, 4vw, 3rem);
  }

  .featured-inner {
    max-width: 860px;
    margin: 0 auto;
    text-align: center;
  }

  .featured-stars {
    font-size: 1.5rem;
    color: var(--accent-gold);
    letter-spacing: 0.1em;
    margin-bottom: 1.5rem;
  }

  .featured-quote {
    position: relative;
    margin: 0 0 1.5rem;
  }

  .featured-mark {
    display: block;
    font-size: clamp(5rem, 12vw, 9rem);
    line-height: 0.6;
    color: var(--primary);
    font-family: Georgia, serif;
    margin-bottom: 0.5rem;
    opacity: 0.8;
  }

  .featured-quote p {
    font-size: clamp(1.2rem, 2.5vw, 1.6rem);
    font-style: italic;
    color: rgba(255, 255, 255, 0.92);
    line-height: 1.65;
    margin: 0;
    font-weight: 400;
  }

  .featured-attribution {
    font-size: 0.875rem;
    font-weight: 700;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: var(--accent-gold);
    margin: 0;
  }

  /* ── Reviews grid ───────────────────────────────────────────────── */
  .reviews-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
    padding: 0 clamp(1.5rem, 4vw, 3rem) clamp(3rem, 6vw, 5rem);
  }

  .reviews-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: clamp(1.25rem, 2.5vw, 2rem);
    max-width: 1200px;
    margin: 0 auto;
  }

  .quote-card {
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 0.75rem;
    padding: clamp(1.5rem, 3vw, 2rem);
    display: flex;
    flex-direction: column;
    gap: 1rem;
    transition: background 0.25s ease, border-color 0.25s ease;
  }

  .quote-card:hover {
    background: rgba(255, 255, 255, 0.09);
    border-color: rgba(255, 107, 53, 0.35);
  }

  .quote-card__stars {
    font-size: 1rem;
    color: var(--accent-gold);
    letter-spacing: 0.05em;
  }

  .quote-card__text {
    font-size: clamp(0.9rem, 1.5vw, 1rem);
    font-style: italic;
    color: rgba(255, 255, 255, 0.8);
    line-height: 1.75;
    margin: 0;
    flex: 1;
  }

  .quote-card__footer {
    display: flex;
    flex-direction: column;
    gap: 0.15rem;
    padding-top: 0.75rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
  }

  .quote-card__author {
    font-size: 0.875rem;
    font-weight: 700;
    color: #ffffff;
  }

  .quote-card__location {
    font-size: 0.775rem;
    font-weight: 500;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: rgba(255, 255, 255, 0.4);
  }

  /* ── TripAdvisor band ───────────────────────────────────────────── */
  .tripadvisor-band {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--secondary-bg);
    padding: clamp(2.5rem, 5vw, 4rem) clamp(1.5rem, 4vw, 3rem);
  }

  .tripadvisor-inner {
    max-width: 900px;
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 2rem;
    flex-wrap: wrap;
  }

  .tripadvisor-text {
    flex: 1 1 auto;
  }

  .tripadvisor-eyebrow {
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: #00aa6c;
    margin: 0 0 0.4rem;
  }

  .tripadvisor-heading {
    font-size: clamp(1.4rem, 3vw, 2rem);
    font-weight: 800;
    color: var(--text-heading);
    margin: 0 0 0.5rem;
    letter-spacing: -0.02em;
  }

  .tripadvisor-body {
    font-size: 1rem;
    color: var(--text-muted);
    margin: 0;
    line-height: 1.6;
  }

  .tripadvisor-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    background: #00aa6c;
    color: #ffffff;
    font-size: 1rem;
    font-weight: 700;
    text-decoration: none;
    padding: 0.85rem 1.75rem;
    border-radius: 0.5rem;
    white-space: nowrap;
    flex-shrink: 0;
    transition: background 0.2s ease, transform 0.2s ease;
  }

  .tripadvisor-btn:hover {
    background: #008f5a;
    transform: translateY(-2px);
  }

  /* ── CTA wrapper ────────────────────────────────────────────────── */
  .cta-wrapper {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background-size: cover;
    background-position: center;
    background-attachment: fixed;
  }

  .cta-wrapper::before {
    content: '';
    position: absolute;
    inset: 0;
    background: linear-gradient(
      135deg,
      rgba(26, 26, 46, 0.82) 0%,
      rgba(255, 107, 53, 0.6) 100%
    );
    z-index: 0;
  }

  .cta-wrapper :global(.call-to-action-section) {
    position: relative;
    z-index: 1;
    background: transparent;
    box-shadow: none;
    padding: clamp(4rem, 8vw, 6rem) 2rem;
    max-width: 800px;
  }

  .cta-wrapper :global(.call-to-action-section:hover) {
    transform: none;
    box-shadow: none;
  }

  .cta-wrapper :global(.text--h2) { color: #ffffff; }
  .cta-wrapper :global(.text--p) { color: rgba(255, 255, 255, 0.88); }

  /* ── Responsive ─────────────────────────────────────────────────── */
  @media (max-width: 900px) {
    .reviews-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 600px) {
    .reviews-grid {
      grid-template-columns: 1fr;
    }

    .tripadvisor-inner {
      flex-direction: column;
      text-align: center;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .cta-wrapper {
      background-attachment: scroll;
    }
  }
</style>
