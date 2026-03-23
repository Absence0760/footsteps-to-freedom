<script lang="ts">
  import { Text } from "$lib/components/atoms";
  import { ContactForm } from "$lib/components/organisms";
  import { HeroSection } from "$lib/components/organisms";
</script>

<div class="page-container">
  <!-- Hero Section -->
  <HeroSection image="/elephants-walking-out-of-dried-pong.jpg" mobileImage="/vertical-penguin.jpg" alt="Contact Footsteps to Freedom">
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
    <div class="info-card">
      <Text variant="h3">Other Ways to Reach Us</Text>
      <div class="info-row">
        <Text variant="h4">Email</Text>
        <Text variant="p">info@footstepstofreedom.co.za</Text>
      </div>
      <div class="info-row">
        <Text variant="h4">Location</Text>
        <Text variant="p">Cape Town, South Africa</Text>
      </div>
      <div class="info-row">
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
    max-width: 800px;
    margin: 0 auto;
  }

  .info-card {
    padding: clamp(1.5rem, 4vw, 2.5rem);
    background: var(--card-bg);
    border-radius: 1rem;
    box-shadow: var(--shadow-md);
    text-align: center;
  }

  .info-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 1rem 0;
    border-top: 1px solid var(--border-color, #e5e7eb);
    gap: 1rem;
  }

  .info-row:first-of-type {
    margin-top: 1rem;
  }

  .info-row :global(.text--h4) {
    text-align: left;
    margin-bottom: 0;
    white-space: nowrap;
  }

  .info-row :global(.text--p) {
    text-align: right;
    margin-bottom: 0;
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
