<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import { Card } from "$lib/components/molecules";
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
    padding: clamp(1.5rem, 3vw, 2rem) 0 0;
    text-align: center;
    max-width: 900px;
    margin: 0 auto;
  }

  .mission-cta {
    margin-top: 1.5rem;
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


</style>
