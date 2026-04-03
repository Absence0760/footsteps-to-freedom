<script lang="ts">
import { base } from '$app/paths';
import { Link, Text } from "$lib/components/atoms";
import { CallToActionSection, HeroSection, TourModal } from "$lib/components/organisms";

const tourModules = import.meta.glob('/src/content/tours/*.md', { eager: true });
const allTours = Object.values(tourModules).map((mod: any) => ({
  ...mod.metadata,
  component: mod.default,
}));

const halfDayTours = allTours.filter(t => t.duration === 'half-day');
const fullDayTours = allTours.filter(t => t.duration === 'full-day');
const curatedTours = allTours.filter(t => t.duration === 'curated');

let selectedTour: typeof allTours[number] | null = $state(null);

function openModal(tour: typeof allTours[number]) {
  selectedTour = tour;
  if (typeof document !== 'undefined') document.body.style.overflow = 'hidden';
}

function closeModal() {
  selectedTour = null;
  if (typeof document !== 'undefined') document.body.style.overflow = '';
}

const resolve = (src: string) =>
  src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;
</script>

<div class="page-container">
  <!-- Hero -->
  <HeroSection image="/african-buffalo.jpg" mobileImage="/vertical-lioness.jpg" alt="South Africa's iconic wildlife and landscapes">
    <Text variant="h1">Our Inspiring Tours</Text>
    <Text variant="h2">Step back in time and walk the paths of resilience, struggle, and triumph</Text>
    <div class="hero-buttons">
      <Link variant="button" href="#half-day">Browse Tours</Link>
      <Link
        variant="button"
        to="/contact"
        style="--button-primary-bg: var(--secondary); --button-primary-hover-bg: var(--secondary-hover);"
      >Get in Touch</Link>
    </div>
  </HeroSection>

  <!-- Tour Type Strip -->
  <div class="type-strip">
    <a href="#half-day" class="type-item">
      <p class="type-time">3–4 hours</p>
      <h3 class="type-name">Half-Day Tours</h3>
      <p class="type-count">{halfDayTours.length} tours available</p>
    </a>
    <div class="type-divider" aria-hidden="true"></div>
    <a href="#full-day" class="type-item">
      <p class="type-time">Full day</p>
      <h3 class="type-name">Full-Day Tours</h3>
      <p class="type-count">{fullDayTours.length} tours available</p>
    </a>
    <div class="type-divider" aria-hidden="true"></div>
    <a href="#curated" class="type-item">
      <p class="type-time">Your schedule</p>
      <h3 class="type-name">Curated Tours</h3>
      <p class="type-count">{curatedTours.length} tours available</p>
    </a>
  </div>

  <!-- Half-Day Tours -->
  <section class="tours-section" id="half-day">
    <div class="section-header">
      <p class="section-eyebrow">3–4 hours</p>
      <h2 class="section-heading">Half-Day Tours</h2>
      <p class="section-subheading">Perfect for combining with other activities — maximum discovery in minimum time</p>
    </div>
    <div class="tour-grid">
      {#each halfDayTours as tour}
        <button class="tour-card" onclick={() => openModal(tour)}>
          <div class="tour-card__image">
            <img src={resolve(tour.image)} alt={tour.name} loading="lazy" />
            {#if tour.category}
              <span class="tour-card__badge">{tour.category}</span>
            {/if}
          </div>
          <div class="tour-card__body">
            <h3 class="tour-card__title">{tour.name}</h3>
            <p class="tour-card__desc">{tour.description}</p>
            <span class="tour-card__cta">Discover more &rarr;</span>
          </div>
        </button>
      {/each}
    </div>
  </section>

  <!-- Full-Day Tours -->
  <section class="tours-section tours-section--dark" id="full-day">
    <div class="section-header">
      <p class="section-eyebrow">Full day</p>
      <h2 class="section-heading section-heading--light">Full-Day Tours</h2>
      <p class="section-subheading section-subheading--light">Immerse yourself completely — we handle every detail so you can focus on discovery</p>
    </div>
    <div class="tour-grid">
      {#each fullDayTours as tour}
        <button class="tour-card tour-card--on-dark" onclick={() => openModal(tour)}>
          <div class="tour-card__image">
            <img src={resolve(tour.image)} alt={tour.name} loading="lazy" />
            {#if tour.category}
              <span class="tour-card__badge">{tour.category}</span>
            {/if}
          </div>
          <div class="tour-card__body">
            <h3 class="tour-card__title">{tour.name}</h3>
            <p class="tour-card__desc">{tour.description}</p>
            <span class="tour-card__cta">Discover more &rarr;</span>
          </div>
        </button>
      {/each}
    </div>
  </section>

  <!-- Curated Tours -->
  <section class="tours-section tours-section--curated" id="curated">
    <div class="section-header">
      <p class="section-eyebrow section-eyebrow--gold">Bespoke</p>
      <h2 class="section-heading">Curated Tours</h2>
      <p class="section-subheading">Tailor-made experiences designed around your specific interests and pace</p>
    </div>
    <div class="tour-grid">
      {#each curatedTours as tour}
        <button class="tour-card" onclick={() => openModal(tour)}>
          <div class="tour-card__image">
            <img src={resolve(tour.image)} alt={tour.name} loading="lazy" />
            {#if tour.category}
              <span class="tour-card__badge tour-card__badge--gold">{tour.category}</span>
            {/if}
          </div>
          <div class="tour-card__body">
            <h3 class="tour-card__title">{tour.name}</h3>
            <p class="tour-card__desc">{tour.description}</p>
            <span class="tour-card__cta tour-card__cta--gold">Discover more &rarr;</span>
          </div>
        </button>
      {/each}
    </div>
  </section>

  <!-- CTA -->
  <div
    class="cta-wrapper"
    style="background-image: url({resolve('/runner-with-table-mountain.jpg')})"
  >
    <CallToActionSection
      title="Ready for Your Journey?"
      description="Contact us with any questions or to book your adventure. We'd love to hear from you."
      buttonLabel="Book Your Tour"
      buttonHref="/contact"
    />
  </div>
</div>

<!-- Tour Detail Modal -->
<TourModal tour={selectedTour} onClose={closeModal} />

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

  /* ── Tour type strip ────────────────────────────────────────────── */
  .type-strip {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
    display: flex;
    align-items: stretch;
    justify-content: center;
    margin-top: calc(-1 * clamp(2rem, 5vw, 3rem));
  }

  .type-item {
    flex: 1 1 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 0.25rem;
    padding: clamp(1.25rem, 3vw, 2rem) clamp(1rem, 3vw, 2rem);
    text-decoration: none;
    text-align: center;
    transition: background 0.2s ease;
    max-width: 360px;
  }

  .type-item:hover {
    background: rgba(255, 255, 255, 0.06);
  }

  .type-time {
    font-size: 0.75rem;
    font-weight: 600;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0;
  }

  .type-name {
    font-size: clamp(1rem, 2.5vw, 1.25rem);
    font-weight: 700;
    color: #ffffff;
    margin: 0.25rem 0;
    line-height: 1.2;
  }

  .type-count {
    font-size: 0.8rem;
    color: rgba(255, 255, 255, 0.55);
    margin: 0;
  }

  .type-divider {
    width: 1px;
    background: rgba(255, 255, 255, 0.1);
    align-self: stretch;
  }

  /* ── Tour sections ──────────────────────────────────────────────── */
  .tours-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    padding: clamp(3rem, 6vw, 5rem) clamp(1.5rem, 4vw, 3rem);
    background: var(--body-bg);
  }

  .tours-section--dark {
    background: var(--text-heading);
  }

  .tours-section--curated {
    background: var(--secondary-bg);
  }

  /* ── Section headers ────────────────────────────────────────────── */
  .section-header {
    text-align: center;
    max-width: 700px;
    margin: 0 auto clamp(2rem, 4vw, 3rem);
  }

  .section-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0 0 0.5rem;
  }

  .section-eyebrow--gold {
    color: var(--accent-gold);
  }

  .section-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: var(--text-heading);
    margin: 0 0 0.75rem;
    letter-spacing: -0.02em;
    line-height: 1.15;
  }

  .section-heading--light {
    color: #ffffff;
  }

  .section-subheading {
    font-size: clamp(0.95rem, 2vw, 1.05rem);
    color: var(--text-muted);
    margin: 0;
    line-height: 1.6;
  }

  .section-subheading--light {
    color: rgba(255, 255, 255, 0.65);
  }

  /* ── Tour grid ──────────────────────────────────────────────────── */
  .tour-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: clamp(1.25rem, 2.5vw, 2rem);
    max-width: 1200px;
    margin: 0 auto;
  }

  /* ── Tour cards ─────────────────────────────────────────────────── */
  .tour-card {
    display: flex;
    flex-direction: column;
    background: #ffffff;
    border-radius: 0.75rem;
    overflow: hidden;
    cursor: pointer;
    border: none;
    padding: 0;
    text-align: left;
    box-shadow: var(--shadow-md);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    width: 100%;
  }

  .tour-card:hover {
    transform: translateY(-6px);
    box-shadow: var(--shadow-xl);
  }

  .tour-card:focus-visible {
    outline: 3px solid var(--primary);
    outline-offset: 3px;
  }

  .tour-card--on-dark {
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
  }

  .tour-card--on-dark:hover {
    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.55);
  }

  .tour-card__image {
    position: relative;
    aspect-ratio: 16 / 10;
    overflow: hidden;
    flex-shrink: 0;
  }

  .tour-card__image img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
    transition: transform 0.5s ease;
  }

  .tour-card:hover .tour-card__image img {
    transform: scale(1.06);
  }

  .tour-card__badge {
    position: absolute;
    top: 0.75rem;
    left: 0.75rem;
    background: var(--primary);
    color: #ffffff;
    font-size: 0.65rem;
    font-weight: 700;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    padding: 0.25rem 0.6rem;
    border-radius: 999px;
    white-space: nowrap;
  }

  .tour-card__badge--gold {
    background: var(--accent-gold);
    color: var(--text-heading);
  }

  .tour-card__body {
    padding: clamp(1rem, 2.5vw, 1.5rem);
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    flex: 1;
  }

  .tour-card__title {
    font-size: clamp(1rem, 2vw, 1.2rem);
    font-weight: 700;
    color: var(--text-heading);
    margin: 0;
    line-height: 1.3;
  }

  .tour-card__desc {
    font-size: clamp(0.85rem, 1.5vw, 0.95rem);
    color: var(--text-muted);
    line-height: 1.65;
    margin: 0;
    flex: 1;
  }

  .tour-card__cta {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--primary);
    margin-top: 0.25rem;
    display: block;
  }

  .tour-card__cta--gold {
    color: #a07c00;
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
  @media (max-width: 900px) {
    .tour-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 640px) {
    .type-strip {
      flex-wrap: wrap;
    }

    .type-item {
      flex: 1 1 calc(50% - 1px);
    }

    .type-divider:last-of-type {
      display: none;
    }

    .tour-grid {
      grid-template-columns: 1fr;
      gap: 1.25rem;
    }

    .hero-buttons {
      flex-direction: column;
      width: 100%;
      max-width: 300px;
      margin-left: auto;
      margin-right: auto;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .tour-card__image img {
      transition: none;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }
</style>
