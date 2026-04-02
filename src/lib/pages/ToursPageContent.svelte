<script lang="ts">
import { Text } from "$lib/components/atoms";
import { Card } from "$lib/components/molecules";
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
</script>

<div class="page-container">
  <!-- Hero Section -->
  <HeroSection image="/rhino.jpg" mobileImage="/vertical-man-surfing.jpg" alt="South Africa tour destinations">
    <Text variant="h1">Our Inspiring Tours</Text>
    <Text variant="h2">Step back in time and walk the paths of resilience, struggle, and triumph</Text>
  </HeroSection>

  <!-- Introduction Section -->
  <section class="intro-section">
    <Text variant="h2">Discover South Africa's Rich Heritage</Text>
    <Text variant="p">
      Each Footsteps To Freedom tour offers a unique and immersive historical experience. Our expert guides bring history to life, sharing the vibrant human stories that shaped this remarkable nation.
    </Text>
  </section>

  <!-- Half-Day Tours -->
  <section class="tours-section">
    <div class="section-header">
      <Text variant="h2">Half-Day Tours</Text>
      <p class="section-subheading">3–4 hours of discovery — perfect for combining with other activities</p>
    </div>
    <div class="tour-grid">
      {#each halfDayTours as tour}
        <Card
          variant="primary"
          imageSrc={tour.image}
          onCardClick={() => openModal(tour)}
        >
          <Text variant="h3">{tour.name}</Text>
          <Text variant="p">{tour.description}</Text>
        </Card>
      {/each}
    </div>
  </section>

  <!-- Full-Day Tours -->
  <section class="tours-section tours-section--alt">
    <div class="section-header">
      <Text variant="h2">Full-Day Tours</Text>
      <p class="section-subheading">A full day of exploration — we handle every detail</p>
    </div>
    <div class="tour-grid">
      {#each fullDayTours as tour}
        <Card
          variant="primary"
          imageSrc={tour.image}
          onCardClick={() => openModal(tour)}
        >
          <Text variant="h3">{tour.name}</Text>
          <Text variant="p">{tour.description}</Text>
        </Card>
      {/each}
    </div>
  </section>

  <!-- Curated Tours -->
  <section class="tours-section">
    <div class="section-header">
      <Text variant="h2">Curated Tours</Text>
      <p class="section-subheading">Tailor-made experiences designed around your specific interests and pace</p>
    </div>
    <div class="tour-grid">
      {#each curatedTours as tour}
        <Card
          variant="primary"
          imageSrc={tour.image}
          onCardClick={() => openModal(tour)}
        >
          <Text variant="h3">{tour.name}</Text>
          <Text variant="p">{tour.description}</Text>
        </Card>
      {/each}
    </div>
  </section>

  <!-- CTA Section -->
  <CallToActionSection
    title="Ready for Your Journey?"
    description="Explore the tours above and contact us if you have any questions or to book your adventure."
    buttonLabel="Contact Us"
    buttonHref="/contact"
  />
</div>

<!-- Tour Detail Modal -->
<TourModal tour={selectedTour} onClose={closeModal} />

<style>
  .intro-section {
    padding: clamp(2rem, 5vw, 4rem) 0;
    text-align: center;
    max-width: 900px;
    margin: 0 auto;
  }

  .tours-section {
    padding: clamp(2rem, 5vw, 4rem) clamp(1rem, 4vw, 3rem);
  }

  .tours-section--alt {
    background: var(--gradient-section);
    border-radius: 1rem;
  }

  .section-header {
    text-align: center;
    margin-bottom: clamp(1.5rem, 3vw, 2.5rem);
  }

  .section-subheading {
    margin: 0;
    color: var(--text-muted, #4A5568);
    font-size: 1rem;
    line-height: 1.5;
  }

  .tour-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: clamp(1.5rem, 3vw, 2.5rem);
    justify-items: center;
  }

  @media (max-width: 900px) {
    .tour-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 600px) {
    .tour-grid {
      grid-template-columns: 1fr;
      gap: 1.5rem;
    }

    .tours-section {
      padding: clamp(1.5rem, 4vw, 2rem) 0.75rem;
    }
  }
</style>
