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
        <Text variant="h4">Call or WhatsApp</Text>
        <div class="info-value">
          <span class="phone-number">+27-83-309 5378</span>
          <a
            href="https://wa.me/27833095378"
            target="_blank"
            rel="noopener noreferrer"
            class="whatsapp-link"
            aria-label="Chat on WhatsApp"
          >
            <!-- WhatsApp icon -->
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20" fill="currentColor" aria-hidden="true">
              <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
            </svg>
            WhatsApp
          </a>
        </div>
      </div>

      <div class="info-row">
        <Text variant="h4">Email</Text>
        <a href="mailto:tours@footstepstofreedom.co.za" class="contact-link">
          tours@footstepstofreedom.co.za
        </a>
      </div>

      <div class="info-row">
        <Text variant="h4">Location</Text>
        <Text variant="p">Cape Town, South Africa</Text>
      </div>

      <div class="info-row">
        <Text variant="h4">Response Time</Text>
        <Text variant="p">We typically respond within 24 hours</Text>
      </div>

      <!-- Social Media -->
      <div class="social-section">
        <Text variant="h4">Follow Us</Text>
        <div class="social-icons">
          <a
            href="https://www.facebook.com/footsteps2freedom.tours/"
            target="_blank"
            rel="noopener noreferrer"
            class="social-link social-link--facebook"
            aria-label="Footsteps to Freedom on Facebook"
          >
            <!-- Facebook icon -->
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" fill="currentColor" aria-hidden="true">
              <path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>
            </svg>
          </a>

          <a
            href="https://www.instagram.com/footstepstofreedomtours/"
            target="_blank"
            rel="noopener noreferrer"
            class="social-link social-link--instagram"
            aria-label="Footsteps to Freedom on Instagram"
          >
            <!-- Instagram icon -->
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
              <rect x="2" y="2" width="20" height="20" rx="5" ry="5"/>
              <path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"/>
              <line x1="17.5" y1="6.5" x2="17.51" y2="6.5"/>
            </svg>
          </a>

          <a
            href="https://www.tripadvisor.com/Attraction_Review-g312659-d2098098-Reviews-Footsteps_to_Freedom_Tours-Cape_Town_Central_Western_Cape.html"
            target="_blank"
            rel="noopener noreferrer"
            class="social-link social-link--tripadvisor"
            aria-label="Footsteps to Freedom on TripAdvisor"
          >
            <!-- TripAdvisor owl icon -->
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" fill="currentColor" aria-hidden="true">
              <path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10 10-4.477 10-10S17.523 2 12 2zm0 18c-4.418 0-8-3.582-8-8s3.582-8 8-8 8 3.582 8 8-3.582 8-8 8zm-3.5-9a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm7 0a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm-3.5-2c-1.2 0-2.3.4-3.2 1H7l-1.5-1.5C6.8 7.2 9.3 6 12 6s5.2 1.2 6.5 3L17 10.5c-.9-.6-2-1-3.1-1H13c-.3-.3-.9-.5-1-.5zm0 7c-1.1 0-2-.4-2.7-1h5.4c-.7.6-1.6 1-2.7 1z"/>
            </svg>
          </a>
        </div>
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

  .info-value {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    flex-wrap: wrap;
    justify-content: flex-end;
  }

  .phone-number {
    font-size: 1rem;
    color: var(--text-body);
  }

  .whatsapp-link {
    display: inline-flex;
    align-items: center;
    gap: 0.35rem;
    color: #25D366;
    text-decoration: none;
    font-weight: 600;
    font-size: 0.95rem;
    padding: 0.25rem 0.6rem;
    border-radius: 0.375rem;
    border: 1.5px solid #25D366;
    transition: background-color 0.2s ease, color 0.2s ease;
  }

  .whatsapp-link:hover {
    background-color: #25D366;
    color: #ffffff;
  }

  .contact-link {
    color: var(--primary);
    text-decoration: none;
    font-size: 1rem;
    transition: color 0.2s ease, text-decoration 0.2s ease;
  }

  .contact-link:hover {
    color: var(--primary-hover);
    text-decoration: underline;
  }

  .social-section {
    margin-top: 1.5rem;
    padding-top: 1.5rem;
    border-top: 1px solid var(--border-color, #e5e7eb);
  }

  .social-section :global(.text--h4) {
    margin-bottom: 1rem;
  }

  .social-icons {
    display: flex;
    justify-content: center;
    gap: 1rem;
    margin-top: 0.75rem;
  }

  .social-link {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 2.75rem;
    height: 2.75rem;
    border-radius: 50%;
    text-decoration: none;
    transition: transform 0.2s ease, background-color 0.2s ease, color 0.2s ease;
    color: #ffffff;
  }

  .social-link:hover {
    transform: scale(1.1);
  }

  .social-link--facebook {
    background-color: #1877F2;
  }

  .social-link--facebook:hover {
    background-color: #0c5fd4;
  }

  .social-link--instagram {
    background: radial-gradient(circle at 30% 107%, #fdf497 0%, #fdf497 5%, #fd5949 45%, #d6249f 60%, #285AEB 90%);
    color: #ffffff;
  }

  .social-link--instagram:hover {
    opacity: 0.9;
    transform: scale(1.1);
  }

  .social-link--tripadvisor {
    background-color: #34E0A1;
    color: #ffffff;
  }

  .social-link--tripadvisor:hover {
    background-color: #25c98e;
  }

  @media (max-width: 768px) {
    .info-grid {
      grid-template-columns: 1fr;
      gap: 1rem;
    }

    .contact-intro {
      margin-bottom: 2rem;
    }

    .info-row {
      flex-direction: column;
      align-items: flex-start;
      gap: 0.5rem;
    }

    .info-value {
      justify-content: flex-start;
    }

    .info-row :global(.text--p) {
      text-align: left;
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
