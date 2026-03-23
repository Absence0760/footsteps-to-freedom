<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import { Card } from "$lib/components/molecules";
import { HeroSection } from "$lib/components/organisms";
const teamModules = import.meta.glob('/src/content/team/*.md', { eager: true });
const teamMembers = Object.values(teamModules).map((mod: any) => mod.metadata);

const { title = "Meet Our Dedicated Team", ...restProps } = $props();
</script>

<div class="page-container">
  <!-- Hero Section -->
  <HeroSection image="/runner-with-table-mountain.jpg" mobileImage="/vertical-africa.jpg" alt="Cape Town heritage sites">
    <Text variant="h1">Our Story</Text>
    <Text variant="h2">Passionate guides sharing the human stories of history</Text>
  </HeroSection>

  <!-- Mission Section -->
  <section class="mission-section">
    <Text variant="h2">Our Mission: History Unveiled</Text>

    <div class="mission-content">
      <Text variant="p">
        At <strong>Footsteps To Freedom Tours</strong>, we believe history isn't just about dates and names; it's about the vibrant human stories that shaped our world. Our mission is to bring these narratives to life through immersive, respectful, and deeply engaging travel experiences. We specialize in journeys that explore pivotal moments and untold tales of freedom, resistance, and human resilience.
      </Text>

      <Text variant="p">
        In 2002, the Footsteps to Freedom city walking tour of Cape Town was developed and launched by Garth Angus. It was then that he met tourist guide and ex-teacher (History and English) Mary-Ann Wenham, who instilled in him a love of history, our city, its stories, and its people. Since then, Garth has met many other talented tourist guides, some of whom today form the core of Footsteps to Freedom. We share a love of our city and its people, our country, and, importantly, a love of learning new things.
      </Text>
    </div>
  </section>

  <!-- Philosophy Section -->
  <section class="philosophy-section">
    <Text variant="h2">Our Touring Philosophy</Text>

    <Text variant="p" style="text-align: center; max-width: 700px; margin: 0 auto 2rem;">
      Our touring philosophy is simple, and it revolves around people:
    </Text>

    <div class="philosophy-grid">
      <Card variant="primary">
        <Text variant="h3">Personalized Matching</Text>
        <Text variant="p">
          We attempt to match the unique interests of our curious travellers with one of our passionate, enquiring tourist guides.
        </Text>
      </Card>

      <Card variant="primary">
        <Text variant="h3">Flexible Tours</Text>
        <Text variant="p">
          Straitjacket tours are for mass market tour operators. Our tours are individual and flexible.
        </Text>
      </Card>

      <Card variant="primary">
        <Text variant="h3">Your Interests Matter</Text>
        <Text variant="p">
          We ask you to tell us about yourself; your work & leisure interests, hobbies, age bracket, nationality, as this will help us design an itinerary you will enjoy.
        </Text>
      </Card>
    </div>
  </section>

  <!-- Team Section -->
  <section class="team-section">
    <Text variant="h2">{title}</Text>

    <div class="team-grid">
      {#each teamMembers as member}
        <Card
          variant="primary"
          imageSrc={member.image}
          slug={"about/" + member.slug}
        >
          <Text variant="h3">{member.name}</Text>
          <Text variant="p">{member.bio}</Text>
        </Card>
      {/each}
    </div>
  </section>

  <!-- CTA Section -->
  <section class="cta-section">
    <Text variant="h2">Ready to Explore with Us?</Text>
    <Text variant="p">
      Let's create an unforgettable journey through South Africa's rich history together.
    </Text>
    <div class="cta-buttons">
      <Link variant="button" to="/tours">View Our Tours</Link>
      <Link
        variant="button"
        to="/contact"
        style="--button-primary-bg: var(--secondary); --button-primary-hover-bg: var(--secondary-hover);"
      >
        Contact Us
      </Link>
    </div>
  </section>
</div>

<style>
  .mission-section {
    padding: clamp(2rem, 5vw, 4rem) 0;
    text-align: center;
    max-width: 900px;
    margin: 0 auto;
  }

  .mission-content {
    margin-top: 2rem;
    text-align: left;
  }

  .philosophy-section {
    padding: clamp(2rem, 5vw, 4rem) clamp(1rem, 3vw, 2rem);
    background: var(--gradient-section);
    border-radius: 1rem;
    text-align: center;
  }

  .philosophy-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: clamp(1.5rem, 3vw, 2rem);
    justify-items: center;
    margin-top: 2rem;
  }

  .team-section {
    padding: clamp(2rem, 5vw, 4rem) clamp(1rem, 4vw, 3rem);
    text-align: center;
  }

  .team-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: clamp(1.5rem, 3vw, 2.5rem);
    margin-top: 2rem;
  }

  .team-grid :global(.card.primary) {
    max-width: 100%;
    min-width: 0;
  }

  .cta-section {
    padding: clamp(2rem, 5vw, 3rem) clamp(1rem, 3vw, 2rem);
    background: var(--secondary-bg);
    border-radius: 1rem;
    text-align: center;
    margin-top: 2rem;
  }

  .cta-section :global(.text--p) {
    text-align: center;
    margin-left: auto;
    margin-right: auto;
  }

  .cta-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    align-items: center;
    margin-top: 1.5rem;
    flex-wrap: wrap;
  }

  @media (max-width: 768px) {
    .team-grid {
      grid-template-columns: 1fr;
      gap: 1.5rem;
    }

    .philosophy-grid {
      grid-template-columns: 1fr;
      gap: 1.5rem;
    }

    .cta-buttons {
      flex-direction: column;
      width: 100%;
      max-width: 300px;
      margin-left: auto;
      margin-right: auto;
    }
  }

  @media (max-width: 480px) {
    .philosophy-section {
      padding: clamp(1.5rem, 4vw, 2rem) clamp(0.75rem, 2vw, 1rem);
    }
  }
</style>
