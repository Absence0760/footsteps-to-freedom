<script lang="ts">
import { base } from '$app/paths';
import { Link, Text } from "$lib/components/atoms";
import {
	CallToActionSection,
	HeroSlideshow,
	TestimonialsSection,
} from "$lib/components/organisms";

const heroSlides = [
	{
		image: '/cape-town-from-table-mountain.jpg',
		mobileImage: '/vertical-distant-table-mountain.jpg',
		alt: 'Cape Town skyline viewed from Table Mountain',
		caption: 'Cape Town, Western Cape',
	},
	{
		image: '/zebra-giraffe-standing-rear-green-trees.jpg',
		mobileImage: '/zebra-giraffe-standing-rear-green-trees.jpg',
		alt: 'Zebra and giraffe among green trees',
		caption: 'South African wildlife',
	},
	{
		image: '/sandy-beach-with-table-mountain.jpg',
		mobileImage: '/vertical-man-surfing.jpg',
		alt: 'Sandy beach with Table Mountain backdrop',
		caption: 'Atlantic Seaboard',
	},
	{
		image: '/downtown-with-table-mountain.jpg',
		mobileImage: '/vertical-woman-with-lions-head.jpg',
		alt: 'Downtown Cape Town with Table Mountain',
		caption: "Lion's Head, Cape Town",
	},
];

const resolve = (src: string) =>
	src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;
</script>

<div class="page-container">
  <!-- Hero Slideshow -->
  <HeroSlideshow slides={heroSlides} interval={5000}>
    <Text variant="h1">Walk Through History. Discover Freedom.</Text>
    <Text variant="h2">Experience South Africa's rich heritage with authentic, guided tours</Text>
    <div class="hero-buttons">
      <Link variant="button" to="/tours">Explore Tours</Link>
      <Link
        variant="button"
        to="/contact"
        style="--button-primary-bg: var(--secondary); --button-primary-hover-bg: var(--secondary-hover);"
      >Contact Us</Link>
    </div>
  </HeroSlideshow>

  <!-- Trust Bar -->
  <div class="trust-bar">
    <div class="trust-bar__inner">
      <div class="trust-item">
        <span class="trust-item__label">Expert Local Guides</span>
      </div>
      <div class="trust-divider" aria-hidden="true"></div>
      <div class="trust-item">
        <span class="trust-item__label">5-Star TripAdvisor</span>
      </div>
      <div class="trust-divider" aria-hidden="true"></div>
      <div class="trust-item">
        <span class="trust-item__label">Small Group Tours</span>
      </div>
      <div class="trust-divider" aria-hidden="true"></div>
      <div class="trust-item">
        <span class="trust-item__label">Award-Winning</span>
      </div>
    </div>
  </div>

  <!-- Mission Split Section -->
  <section class="mission-split">
    <div class="mission-split__image-col">
      <img
        src={resolve('/scenic-view-of-man-sitting-on-rock-table-mountain.jpg')}
        alt="Traveller seated on Table Mountain rock, overlooking Cape Town"
        class="mission-split__img"
      />
    </div>
    <div class="mission-split__text-col">
      <p class="mission-split__eyebrow">Our Mission</p>
      <h2 class="mission-split__heading">Bringing History to Life</h2>
      <p class="mission-split__body">
        The natural beauty of the Cape is self-evident and widely recorded.
        Less evident are the human stories behind its scenic splendour.
        Against this backdrop, Footsteps to Freedom brings together curious
        travellers and well-informed guides to create enduring memories and
        firm friendships.
      </p>
      <Link variant="button" to="/about">Discover Our Story</Link>
    </div>
  </section>

  <!-- Photo Mosaic Gallery -->
  <section class="gallery-section">
    <div class="gallery-header">
      <p class="gallery-eyebrow">South Africa Awaits</p>
      <h2 class="gallery-heading">Experience the Cape</h2>
    </div>

    <div class="gallery-mosaic">
      <a href="{base}/tours" class="gallery-item gallery-item--tall">
        <img
          src={resolve('/lion.jpg')}
          alt="Lion in the wild, South Africa"
          class="gallery-img"
          loading="lazy"
        />
        <div class="gallery-overlay">
          <span class="gallery-label">Wildlife Safari</span>
        </div>
      </a>

      <a href="{base}/tours" class="gallery-item">
        <img
          src={resolve('/waterfall-mpumalanga.jpg')}
          alt="Waterfall in Mpumalanga"
          class="gallery-img"
          loading="lazy"
        />
        <div class="gallery-overlay">
          <span class="gallery-label">Natural Wonders</span>
        </div>
      </a>

      <a href="{base}/tours" class="gallery-item">
        <img
          src={resolve('/spring-buck-standing-in-yellow-flowers.jpg')}
          alt="Springbok standing in yellow wildflowers"
          class="gallery-img"
          loading="lazy"
        />
        <div class="gallery-overlay">
          <span class="gallery-label">Namaqualand</span>
        </div>
      </a>

      <a href="{base}/tours" class="gallery-item">
        <img
          src={resolve('/downtown-cape-town.jpg')}
          alt="Downtown Cape Town streetscape"
          class="gallery-img"
          loading="lazy"
        />
        <div class="gallery-overlay">
          <span class="gallery-label">Cape Town City</span>
        </div>
      </a>

      <a href="{base}/tours" class="gallery-item">
        <img
          src={resolve('/protea.jpg')}
          alt="King Protea, South Africa's national flower"
          class="gallery-img"
          loading="lazy"
        />
        <div class="gallery-overlay">
          <span class="gallery-label">Fynbos, Western Cape</span>
        </div>
      </a>
    </div>

    <div class="gallery-cta">
      <Link variant="button" to="/tours">View All Tours</Link>
    </div>
  </section>

  <!-- Core Values Section -->
  <section class="values-section">
    <div class="values-inner">
      <div class="values-header">
        <p class="values-eyebrow">What We Stand For</p>
        <h2 class="values-heading">Our Core Values</h2>
      </div>

      <div class="values-grid">
        <div class="value-item">
          <span class="value-number" aria-hidden="true">01</span>
          <h3 class="value-title">Authenticity</h3>
          <p class="value-body">We strive for historical accuracy and genuine cultural immersion in every tour.</p>
        </div>
        <div class="value-item">
          <span class="value-number" aria-hidden="true">02</span>
          <h3 class="value-title">Respect</h3>
          <p class="value-body">Respect for the past, the present, and the communities we visit is paramount.</p>
        </div>
        <div class="value-item">
          <span class="value-number" aria-hidden="true">03</span>
          <h3 class="value-title">Education</h3>
          <p class="value-body">Every journey is designed to be a profound learning experience.</p>
        </div>
        <div class="value-item">
          <span class="value-number" aria-hidden="true">04</span>
          <h3 class="value-title">Empathy</h3>
          <p class="value-body">We encourage understanding and connection with the human stories of history.</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Testimonials Section -->
  <section class="testimonials-wrapper">
    <TestimonialsSection />
  </section>

  <!-- Call to Action -->
  <div
    class="cta-wrapper"
    style="background-image: url({resolve('/runner-with-table-mountain.jpg')})"
  >
    <CallToActionSection
      title="Ready to Walk Through History?"
      description="Join a small-group tour led by expert local guides. Cape Town's stories are waiting for you."
      buttonLabel="Book Your Tour"
      buttonHref="/contact"
    />
  </div>
</div>

<style>
  /* ── Hero buttons ───────────────────────────────────────────────── */
  .hero-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    align-items: center;
    margin-top: 1.5rem;
    flex-wrap: wrap;
  }

  /* ── Trust bar ──────────────────────────────────────────────────── */
  .trust-bar {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--primary);
    /* Pull up to sit flush against the hero */
    margin-top: calc(-1 * clamp(2rem, 5vw, 3rem));
  }

  .trust-bar__inner {
    max-width: 1100px;
    margin: 0 auto;
    padding: 1rem 2rem;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
  }

  .trust-item {
    flex: 1 1 auto;
    text-align: center;
    padding: 0.75rem 1.5rem;
  }

  .trust-item__label {
    font-size: clamp(0.75rem, 2vw, 0.9rem);
    font-weight: 700;
    color: #ffffff;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    white-space: nowrap;
  }

  .trust-divider {
    width: 1px;
    height: 1.5rem;
    background: rgba(255, 255, 255, 0.35);
    flex-shrink: 0;
  }

  /* ── Mission split ──────────────────────────────────────────────── */
  .mission-split {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    display: grid;
    grid-template-columns: 1fr 1fr;
    min-height: 520px;
    overflow: hidden;
  }

  .mission-split__image-col {
    position: relative;
    overflow: hidden;
  }

  .mission-split__img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    object-position: center 30%;
    display: block;
    transition: transform 0.6s ease;
  }

  .mission-split:hover .mission-split__img {
    transform: scale(1.03);
  }

  .mission-split__text-col {
    background: var(--body-bg);
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: clamp(2.5rem, 6vw, 5rem) clamp(2rem, 5vw, 4rem);
    gap: 1.25rem;
  }

  .mission-split__eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0;
  }

  .mission-split__heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: var(--text-heading);
    line-height: 1.15;
    margin: 0;
    letter-spacing: -0.02em;
  }

  .mission-split__body {
    font-size: clamp(1rem, 2vw, 1.125rem);
    color: var(--text-body);
    line-height: 1.8;
    margin: 0;
    max-width: 50ch;
  }

  /* ── Gallery mosaic ─────────────────────────────────────────────── */
  .gallery-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
    padding: clamp(3rem, 6vw, 5rem) clamp(1.5rem, 4vw, 3rem);
  }

  .gallery-header {
    text-align: center;
    margin-bottom: clamp(1.5rem, 3vw, 2.5rem);
  }

  .gallery-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--accent-gold);
    margin: 0 0 0.5rem;
  }

  .gallery-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0;
    letter-spacing: -0.02em;
  }

  .gallery-mosaic {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 280px 280px;
    gap: 0.5rem;
    max-width: 1100px;
    margin: 0 auto;
  }

  .gallery-item--tall {
    grid-row: 1 / 3;
  }

  .gallery-item {
    position: relative;
    overflow: hidden;
    border-radius: 0.25rem;
    display: block;
    text-decoration: none;
  }

  .gallery-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
    transition: transform 0.5s ease;
  }

  .gallery-item:hover .gallery-img,
  .gallery-item:focus-visible .gallery-img {
    transform: scale(1.07);
  }

  .gallery-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(
      to top,
      rgba(26, 26, 46, 0.85) 0%,
      rgba(26, 26, 46, 0) 55%
    );
    display: flex;
    align-items: flex-end;
    padding: 1rem 1.25rem;
    opacity: 0;
    transition: opacity 0.35s ease;
  }

  .gallery-item:hover .gallery-overlay,
  .gallery-item:focus-visible .gallery-overlay {
    opacity: 1;
  }

  .gallery-label {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: #ffffff;
  }

  .gallery-cta {
    text-align: center;
    margin-top: clamp(1.5rem, 3vw, 2.5rem);
  }

  /* ── Values section ─────────────────────────────────────────────── */
  .values-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
    padding: clamp(3rem, 7vw, 6rem) clamp(1.5rem, 4vw, 3rem);
  }

  .values-inner {
    max-width: 1100px;
    margin: 0 auto;
  }

  .values-header {
    text-align: center;
    margin-bottom: clamp(2rem, 4vw, 3.5rem);
  }

  .values-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--accent-gold);
    margin: 0 0 0.5rem;
  }

  .values-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0;
    letter-spacing: -0.02em;
  }

  .values-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: clamp(1.5rem, 3vw, 2.5rem);
  }

  .value-item {
    padding: clamp(1.5rem, 3vw, 2rem);
    border-top: 2px solid rgba(255, 255, 255, 0.12);
    transition: border-color 0.3s ease;
  }

  .value-item:hover {
    border-color: var(--primary);
  }

  .value-number {
    display: block;
    font-size: clamp(3rem, 6vw, 4.5rem);
    font-weight: 900;
    line-height: 1;
    color: rgba(255, 255, 255, 0.06);
    letter-spacing: -0.04em;
    margin-bottom: 0.75rem;
    user-select: none;
  }

  .value-title {
    font-size: clamp(1.1rem, 2.5vw, 1.35rem);
    font-weight: 700;
    color: #ffffff;
    margin: 0 0 0.75rem;
    line-height: 1.3;
  }

  .value-body {
    font-size: clamp(0.875rem, 2vw, 1rem);
    color: rgba(255, 255, 255, 0.7);
    line-height: 1.75;
    margin: 0;
  }

  /* ── Testimonials wrapper ───────────────────────────────────────── */
  .testimonials-wrapper {
    width: 100%;
  }

  /* ── CTA wrapper ────────────────────────────────────────────────── */
  .cta-wrapper {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background-size: cover;
    background-position: center 40%;
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

  .cta-wrapper :global(.text--h2) {
    color: #ffffff;
  }

  .cta-wrapper :global(.text--p) {
    color: rgba(255, 255, 255, 0.88);
  }

  /* ── Responsive ─────────────────────────────────────────────────── */
  @media (max-width: 768px) {
    .mission-split {
      grid-template-columns: 1fr;
      min-height: auto;
    }

    .mission-split__image-col {
      height: 320px;
    }

    .values-grid {
      grid-template-columns: repeat(2, 1fr);
    }

    .trust-divider {
      display: none;
    }

    .trust-item {
      padding: 0.6rem 1rem;
    }
  }

  @media (max-width: 640px) {
    .gallery-mosaic {
      grid-template-columns: 1fr 1fr;
      grid-template-rows: 160px 160px;
      gap: 0.35rem;
    }

    .gallery-item--tall {
      grid-row: 1 / 2;
    }

    .gallery-item:nth-child(5) {
      display: none;
    }

    .gallery-overlay {
      opacity: 1;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }

  @media (max-width: 540px) {
    .values-grid {
      grid-template-columns: 1fr;
    }

    .hero-buttons {
      flex-direction: column;
      width: 100%;
      max-width: 300px;
      margin-left: auto;
      margin-right: auto;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .mission-split__img,
    .gallery-img {
      transition: none;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }
</style>
