<script lang="ts">
import { base } from '$app/paths';
import { Link, Text } from "$lib/components/atoms";
import { HeroSection } from "$lib/components/organisms";

const teamModules = import.meta.glob('/src/content/team/*.md', { eager: true });
const teamMembers = Object.values(teamModules).map((mod: any) => mod.metadata);

const { title = "Meet Our Dedicated Team", ...restProps } = $props();

const resolve = (src: string) =>
  src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;
</script>

<div class="page-container">
  <!-- Hero -->
  <HeroSection image="/runner-with-table-mountain.jpg" mobileImage="/vertical-distant-table-mountain.jpg" alt="Cape Town from Table Mountain">
    <Text variant="h1">Our Story</Text>
    <Text variant="h2">Passionate guides sharing the human stories of history</Text>
  </HeroSection>

  <!-- Mission / Founding Story -->
  <section class="mission-section">
    <div class="mission-inner">
      <span class="mission-year" aria-hidden="true">2002</span>
      <div class="mission-content">
        <p class="mission-eyebrow">How It Began</p>
        <h2 class="mission-heading">History Unveiled</h2>
        <p class="mission-body">
          At <strong>Footsteps To Freedom Tours</strong>, we believe history isn't just about dates and names —
          it's about the vibrant human stories that shaped our world. We specialize in journeys that explore
          pivotal moments and untold tales of freedom, resistance, and human resilience.
        </p>
        <p class="mission-body">
          In 2002, the Footsteps to Freedom city walking tour of Cape Town was developed and launched by
          Garth Angus. He met tourist guide and ex-teacher Mary-Ann Wenham, who instilled in him a love of
          history, our city, its stories, and its people. Since then, Garth has met many other talented
          guides who share a love of Cape Town — and a love of learning new things.
        </p>
      </div>
    </div>
  </section>

  <!-- Philosophy -->
  <section class="philosophy-section">
    <div class="philosophy-inner">
      <div class="philosophy-header">
        <p class="philosophy-eyebrow">How We Work</p>
        <h2 class="philosophy-heading">Our Touring Philosophy</h2>
        <p class="philosophy-subhead">Simple, and revolving around people.</p>
      </div>

      <div class="philosophy-grid">
        <div class="pillar">
          <span class="pillar-number" aria-hidden="true">01</span>
          <h3 class="pillar-title">Personalized Matching</h3>
          <p class="pillar-body">We match the unique interests of our curious travellers with one of our passionate, enquiring tourist guides.</p>
        </div>
        <div class="pillar">
          <span class="pillar-number" aria-hidden="true">02</span>
          <h3 class="pillar-title">Flexible Tours</h3>
          <p class="pillar-body">Straitjacket tours are for mass market operators. Our tours are individual and flexible — shaped around you.</p>
        </div>
        <div class="pillar">
          <span class="pillar-number" aria-hidden="true">03</span>
          <h3 class="pillar-title">Your Interests Matter</h3>
          <p class="pillar-body">Tell us about yourself — your work, hobbies, nationality — and we'll design an itinerary you'll genuinely love.</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Team -->
  <section class="team-section">
    <div class="team-header">
      <p class="team-eyebrow">The People Behind the Tours</p>
      <h2 class="team-heading">{title}</h2>
    </div>

    {#each teamMembers as member, i}
      <div class="team-member" class:team-member--reversed={i % 2 !== 0}>
        <div class="team-member__image-col">
          <img
            src={resolve(member.image)}
            alt={member.name}
            class="team-member__img"
            loading="lazy"
          />
        </div>
        <div class="team-member__text-col">
          <p class="team-member__role">{member.role}</p>
          <h3 class="team-member__name">{member.name}</h3>
          <p class="team-member__bio">{member.bio}</p>
          <Link variant="button" to="/about/{member.slug}">Read Full Bio</Link>
        </div>
      </div>
    {/each}
  </section>

  <!-- CTA -->
  <div
    class="cta-wrapper"
    style="background-image: url({resolve('/spring-buck-standing-in-yellow-flowers.jpg')})"
  >
    <div class="cta-inner">
      <p class="cta-eyebrow">Start Your Journey</p>
      <h2 class="cta-heading">Ready to Explore with Us?</h2>
      <p class="cta-body">Let's create an unforgettable journey through South Africa's rich history together.</p>
      <div class="cta-buttons">
        <Link variant="button" to="/tours">View Our Tours</Link>
        <Link
          variant="button"
          to="/contact"
          style="--button-primary-bg: var(--secondary); --button-primary-hover-bg: var(--secondary-hover);"
        >Contact Us</Link>
      </div>
    </div>
  </div>
</div>

<style>
  /* ── Mission / founding story ───────────────────────────────────── */
  .mission-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
    padding: clamp(3rem, 7vw, 6rem) clamp(1.5rem, 4vw, 3rem);
    overflow: hidden;
  }

  .mission-inner {
    position: relative;
    max-width: 860px;
    margin: 0 auto;
  }

  /* Ghost year watermark */
  .mission-year {
    position: absolute;
    top: 50%;
    right: -1rem;
    transform: translateY(-50%);
    font-size: clamp(8rem, 20vw, 16rem);
    font-weight: 900;
    line-height: 1;
    color: rgba(255, 255, 255, 0.04);
    letter-spacing: -0.05em;
    user-select: none;
    pointer-events: none;
  }

  .mission-content {
    position: relative;
    z-index: 1;
  }

  .mission-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0 0 0.75rem;
  }

  .mission-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0 0 1.5rem;
    letter-spacing: -0.02em;
    line-height: 1.15;
  }

  .mission-body {
    font-size: clamp(0.95rem, 1.75vw, 1.075rem);
    color: rgba(255, 255, 255, 0.75);
    line-height: 1.85;
    margin: 0 0 1rem;
    max-width: 65ch;
  }

  .mission-body strong {
    color: rgba(255, 255, 255, 0.95);
    font-weight: 600;
  }

  /* ── Philosophy ─────────────────────────────────────────────────── */
  .philosophy-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--body-bg);
    padding: clamp(3rem, 6vw, 5rem) clamp(1.5rem, 4vw, 3rem);
  }

  .philosophy-inner {
    max-width: 1100px;
    margin: 0 auto;
  }

  .philosophy-header {
    text-align: center;
    margin-bottom: clamp(2rem, 4vw, 3.5rem);
  }

  .philosophy-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0 0 0.5rem;
  }

  .philosophy-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: var(--text-heading);
    margin: 0 0 0.5rem;
    letter-spacing: -0.02em;
  }

  .philosophy-subhead {
    font-size: 1rem;
    color: var(--text-muted);
    margin: 0;
  }

  .philosophy-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: clamp(1.5rem, 3vw, 2.5rem);
  }

  .pillar {
    padding: clamp(1.5rem, 3vw, 2rem);
    border-top: 2px solid rgba(26, 26, 46, 0.12);
    transition: border-color 0.3s ease;
  }

  .pillar:hover {
    border-color: var(--primary);
  }

  .pillar-number {
    display: block;
    font-size: clamp(3rem, 6vw, 4.5rem);
    font-weight: 900;
    line-height: 1;
    color: rgba(26, 26, 46, 0.06);
    letter-spacing: -0.04em;
    margin-bottom: 0.75rem;
    user-select: none;
  }

  .pillar-title {
    font-size: clamp(1.05rem, 2vw, 1.3rem);
    font-weight: 700;
    color: var(--text-heading);
    margin: 0 0 0.75rem;
    line-height: 1.3;
  }

  .pillar-body {
    font-size: clamp(0.875rem, 1.5vw, 1rem);
    color: var(--text-muted);
    line-height: 1.75;
    margin: 0;
  }

  /* ── Team section ───────────────────────────────────────────────── */
  .team-section {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background: var(--text-heading);
  }

  .team-header {
    text-align: center;
    padding: clamp(3rem, 6vw, 5rem) clamp(1.5rem, 4vw, 3rem) clamp(1.5rem, 3vw, 2.5rem);
  }

  .team-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--accent-gold);
    margin: 0 0 0.5rem;
  }

  .team-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0;
    letter-spacing: -0.02em;
  }

  /* Alternating guide rows */
  .team-member {
    display: grid;
    grid-template-columns: 1fr 1fr;
    min-height: 520px;
  }

  .team-member--reversed {
    direction: rtl;
  }

  .team-member--reversed > * {
    direction: ltr;
  }

  .team-member__image-col {
    position: relative;
    overflow: hidden;
  }

  .team-member__img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    object-position: center top;
    display: block;
    transition: transform 0.6s ease;
  }

  .team-member:hover .team-member__img {
    transform: scale(1.04);
  }

  .team-member__text-col {
    background: rgba(255, 255, 255, 0.04);
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 1rem;
    padding: clamp(2.5rem, 5vw, 4rem) clamp(2rem, 4vw, 3.5rem);
  }

  .team-member__role {
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0;
  }

  .team-member__name {
    font-size: clamp(1.5rem, 3.5vw, 2.25rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0;
    letter-spacing: -0.02em;
    line-height: 1.2;
  }

  .team-member__bio {
    font-size: clamp(0.95rem, 1.75vw, 1.075rem);
    color: rgba(255, 255, 255, 0.7);
    line-height: 1.8;
    margin: 0;
    max-width: 50ch;
  }

  /* ── CTA ────────────────────────────────────────────────────────── */
  .cta-wrapper {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    background-size: cover;
    background-position: center 60%;
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

  .cta-inner {
    position: relative;
    z-index: 1;
    text-align: center;
    padding: clamp(4rem, 8vw, 6rem) clamp(1.5rem, 4vw, 3rem);
    max-width: 800px;
    margin: 0 auto;
  }

  .cta-eyebrow {
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--accent-gold);
    margin: 0 0 0.75rem;
  }

  .cta-heading {
    font-size: clamp(1.75rem, 4vw, 2.75rem);
    font-weight: 800;
    color: #ffffff;
    margin: 0 0 1rem;
    letter-spacing: -0.02em;
  }

  .cta-body {
    font-size: clamp(1rem, 2vw, 1.125rem);
    color: rgba(255, 255, 255, 0.85);
    line-height: 1.7;
    margin: 0 0 2rem;
  }

  .cta-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
  }

  /* ── Responsive ─────────────────────────────────────────────────── */
  @media (max-width: 860px) {
    .team-member,
    .team-member--reversed {
      grid-template-columns: 1fr;
      direction: ltr;
      min-height: auto;
    }

    .team-member__image-col {
      height: 320px;
    }

    .philosophy-grid {
      grid-template-columns: 1fr;
    }
  }

  @media (max-width: 540px) {
    .cta-buttons {
      flex-direction: column;
      max-width: 300px;
      margin-left: auto;
      margin-right: auto;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .team-member__img {
      transition: none;
    }

    .cta-wrapper {
      background-attachment: scroll;
    }
  }
</style>
