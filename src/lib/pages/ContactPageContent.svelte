<script lang="ts">
  import { Text } from "$lib/components/atoms";
  import { ContactForm } from "$lib/components/organisms";
  import { HeroSection } from "$lib/components/organisms";
</script>

<div class="page-container">
  <!-- Hero Section -->
  <HeroSection image="/downtown.jpg" alt="Contact Footsteps to Freedom">
    <Text variant="h1">Get in Touch</Text>
    <Text variant="h2">Let's plan your unforgettable journey together</Text>
  </HeroSection>

  <!-- Contact Section -->
  <section class="contact-section">
    <div class="contact-intro">
      <Text variant="h2">We'd Love to Hear From You</Text>
      <Text variant="p">
        Whether you have questions about our tours, need help planning your
        itinerary, or want to share your interests with us, we're here to help.
        Fill out the form below and we'll get back to you as soon as possible.
      </Text>
    </div>

    <ContactForm />
  </section>

  <!-- Contact Info Section -->
  <section class="info-section">
    <Text variant="h3">Other Ways to Reach Us</Text>
    <div class="info-grid">
      <div class="info-item">
        <Text variant="h4">Email</Text>
        <Text variant="p">info@footstepstofreedom.co.za</Text>
      </div>
      <div class="info-item">
        <Text variant="h4">Location</Text>
        <Text variant="p">Cape Town, South Africa</Text>
      </div>
      <div class="info-item">
        <Text variant="h4">Response Time</Text>
        <Text variant="p">We typically respond within 24 hours</Text>
      </div>
    </div>
  </section>
</div>

<style>
  .contact-section {
    padding: clamp(2rem, 5vw, 4rem) 0;
    max-width: 800px;
    margin: 0 auto;
  }

  .contact-intro {
    text-align: center;
    margin-bottom: 3rem;
  }

  .form-wrapper {
    background: var(--card-bg);
    padding: clamp(1.5rem, 4vw, 2.5rem);
    border-radius: 1rem;
    box-shadow: var(--shadow-lg);
  }

  .info-section {
    padding: clamp(2rem, 5vw, 4rem) clamp(1rem, 3vw, 2rem);
    background: var(--gradient-section);
    border-radius: 1rem;
    text-align: center;
    max-width: 900px;
    margin: 2rem auto 0;
  }

  .info-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: clamp(1.5rem, 3vw, 2rem);
    margin-top: 2rem;
  }

  .info-item {
    padding: 1.5rem;
    background: var(--card-bg);
    border-radius: 0.75rem;
    box-shadow: var(--shadow-sm);
  }

  @media (max-width: 768px) {
    .info-grid {
      grid-template-columns: 1fr;
      gap: 1rem;
    }

    .contact-intro {
      margin-bottom: 2rem;
    }
  }

  @media (max-width: 480px) {
    .form-wrapper {
      padding: clamp(1rem, 3vw, 1.5rem);
    }

    .info-section {
      padding: clamp(1.5rem, 4vw, 2rem) clamp(0.75rem, 2vw, 1rem);
    }
  }
</style>
