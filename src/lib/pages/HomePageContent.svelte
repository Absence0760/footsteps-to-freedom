<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import { Card } from "$lib/components/molecules";
import {
	CallToActionSection,
	HeroSection,
	TestimonialsSection,
} from "$lib/components/organisms";

const tourModules = import.meta.glob('/src/content/tours/*.md', { eager: true, query: '?raw', import: 'default' });

function parseFrontmatter(raw: string): Record<string, string> {
	const match = raw.match(/^---\r?\n([\s\S]*?)\r?\n---/);
	if (!match) return {};
	return Object.fromEntries(
		match[1].split('\n')
			.map(line => line.match(/^(\w+):\s*"?([^"]*)"?\s*$/))
			.filter(Boolean)
			.map(m => [m![1], m![2]])
	);
}

const allTours = Object.values(tourModules).map((raw: any) => parseFrontmatter(raw));
const featuredTours = allTours.slice(0, 3);
</script>

<div class="page-container">
  <!-- Enhanced Hero Section -->
  <HeroSection image="/south-africa-flag-traveler.png" alt="Scenic Cape Town landscape">
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
  </HeroSection>

  <!-- Mission/Story Section -->
  <section class="mission-section">
    <Text variant="h2">Our Mission: Bringing History to Life</Text>
    <Text variant="p">
      The natural beauty of the Cape is self-evident and widely recorded. Less evident are the human stories behind its scenic splendour.
      Against this backdrop, Footsteps to Freedom brings together curious travellers and well-informed guides to create enduring memories and firm friendships.
    </Text>
    <div class="mission-cta">
      <Link variant="inline-text" to="/about">Learn More About Us</Link>
    </div>
  </section>

  <!-- Tours Showcase -->
  <section class="tours-showcase">
    <Text variant="h2">Featured Tours</Text>
    <Text variant="p">Discover our most popular journeys through South African history</Text>

    <div class="tours-grid">
      {#each featuredTours as tour}
        <Card variant="primary" imageSrc={tour.image} slug={"tours/" + tour.slug}>
          <Text variant="h3">{tour.name}</Text>
          <Text variant="p">{tour.description}</Text>
          <div class="tour-meta">
            <Text variant="caption">{tour.duration}</Text>
          </div>
        </Card>
      {/each}
    </div>

    <div class="tours-cta">
      <Link variant="button" to="/tours">View All Tours</Link>
    </div>
  </section>

  <!-- Core Values Section -->
  <section class="values-section">
    <Text variant="h2">Our Core Values</Text>

    <div class="value-grid">
      <Card>
        <Text variant="h3">Authenticity</Text>
        <Text variant="p">We strive for historical accuracy and genuine cultural immersion in every tour.</Text>
      </Card>
      <Card>
        <Text variant="h3">Respect</Text>
        <Text variant="p">Respect for the past, the present, and the communities we visit is paramount.</Text>
      </Card>
      <Card>
        <Text variant="h3">Education</Text>
        <Text variant="p">Every journey is designed to be a profound learning experience.</Text>
      </Card>
      <Card>
        <Text variant="h3">Empathy</Text>
        <Text variant="p">We encourage understanding and connection with the human stories of history.</Text>
      </Card>
    </div>
  </section>

  <!-- Testimonials Section -->
  <section class="testimonials-wrapper">
    <TestimonialsSection />
  </section>

  <!-- Call to Action -->
  <section class="cta-wrapper">
    <CallToActionSection />
  </section>
</div>

<style>
  .hero-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    align-items: center;
    margin-top: 1.5rem;
    flex-wrap: wrap;
  }

  .mission-section {
    padding: clamp(2rem, 5vw, 4rem) 0;
    text-align: center;
    max-width: 900px;
    margin: 0 auto;
  }

  .mission-cta {
    margin-top: 1.5rem;
  }

  .tours-showcase {
    padding: clamp(2rem, 5vw, 4rem) clamp(1rem, 3vw, 2rem);
    background: var(--gradient-section);
    border-radius: 1rem;
    text-align: center;
  }

  .tours-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
    gap: clamp(1.5rem, 3vw, 2.5rem);
    margin: 2rem 0;
    justify-items: center;
  }

  .tour-meta {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    margin-top: auto;
    color: var(--text-muted);
  }

  .tours-cta {
    text-align: center;
    margin-top: 2rem;
  }

  .values-section {
    padding: clamp(2rem, 5vw, 4rem) 0;
    text-align: center;
  }

  .value-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: clamp(1.5rem, 3vw, 2rem);
    justify-items: center;
    margin-top: 2rem;
  }

  .testimonials-wrapper,
  .cta-wrapper {
    width: 100%;
  }

  @media (max-width: 768px) {
    .tours-grid {
      grid-template-columns: 1fr;
    }

    .value-grid {
      grid-template-columns: 1fr;
      gap: 1.5rem;
    }

    .hero-buttons {
      flex-direction: column;
      width: 100%;
      max-width: 300px;
      margin-left: auto;
      margin-right: auto;
    }
  }

  @media (max-width: 600px) {
    .tours-showcase {
      padding: clamp(1.5rem, 4vw, 2rem) clamp(0.75rem, 2vw, 1rem);
    }
  }
</style>
