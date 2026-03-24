<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import { Card } from "$lib/components/molecules";
import { CallToActionSection, HeroSection } from "$lib/components/organisms";

const tourModules = import.meta.glob('/src/content/tours/*.md', { eager: true });
const allTours = Object.values(tourModules).map((mod: any) => mod.metadata);
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

  <!-- Tours Grid -->
  <section class="tours-section">
    <div class="tour-grid">
      {#each allTours as tour}
        <Card
          variant="primary"
          imageSrc={tour.image}
          slug={"tours/"+tour.slug}
        >
          <Text variant="h3">{tour.name}</Text>
          <Text variant="p">{tour.description}</Text>
          {#if tour.category}
            <span class="category-tag">{tour.category}</span>
          {/if}
        </Card>
      {/each}
    </div>
  </section>

  <!-- Flexibility Section -->
  <section class="flexibility-section">
    <Text variant="h2">Tailored to Your Interests</Text>
    <Text variant="p">
      Whilst the above are some of our standard tours, we believe in being totally flexible around our guests' needs and are happy to adapt the day's itinerary to include specific requests.
    </Text>
    <Text variant="p">
      A popular day tour is to combine the city walking tour (which is a great orientation to Cape Town and South Africa) with Table Mountain and Kirstenbosch Botanical Gardens.
    </Text>
  </section>

  <!-- CTA Section -->
  <CallToActionSection
    title="Ready for Your Journey?"
    description="Explore the tours above and contact us if you have any questions or to book your adventure."
    buttonLabel="Contact Us"
    buttonHref="/contact"
  />
</div>

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

  .tour-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: clamp(1.5rem, 3vw, 2.5rem);
    justify-items: center;
  }

  .category-tag {
    display: inline-block;
    margin-top: auto;
    padding: 0.25rem 0.75rem;
    border-radius: 999px;
    background: var(--color-accent, #b8860b);
    color: #fff;
    font-size: 0.75rem;
    font-weight: 600;
    letter-spacing: 0.04em;
    text-transform: uppercase;
  }

  .flexibility-section {
    padding: clamp(2rem, 5vw, 4rem) clamp(1rem, 3vw, 2rem);
    background: var(--gradient-section);
    border-radius: 1rem;
    text-align: center;
    max-width: 900px;
    margin: 0 auto;
  }

  @media (max-width: 768px) {
    .tour-grid {
      grid-template-columns: 1fr;
      gap: 1.5rem;
    }
  }

  @media (max-width: 480px) {
    .tours-section {
      padding: clamp(1.5rem, 4vw, 2rem) 0;
    }

    .flexibility-section {
      padding: clamp(1.5rem, 4vw, 2rem) clamp(0.75rem, 2vw, 1rem);
    }
  }
</style>
