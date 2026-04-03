<script lang="ts">
  import { Text } from "$lib/components/atoms";
  import { ContactForm } from "$lib/components/organisms";
  import { HeroSection } from "$lib/components/organisms";
</script>

<div class="page-container">
  <!-- Hero Section -->
  <HeroSection image="/elephants-walking-out-of-dried-pong.jpg" mobileImage="/vertical-africa.jpg" alt="Contact Footsteps to Freedom">
    <Text variant="h1">Get in Touch</Text>
    <Text variant="h2">Let's plan your unforgettable journey together</Text>
  </HeroSection>

  <!-- Contact Split -->
  <section class="contact-split">
    <!-- Left: info panel -->
    <div class="info-panel">
      <p class="info-eyebrow">Reach Out</p>
      <h2 class="info-heading">We'd Love to Hear From You</h2>
      <p class="info-body">
        Whether you have questions about our tours, need help planning your
        itinerary, or simply want to share your interests — we're here to help.
      </p>

      <div class="info-details">
        <div class="info-item">
          <span class="info-label">Call or WhatsApp</span>
          <div class="info-value-row">
            <span class="info-value">+27-83-309 5378</span>
            <a
              href="https://wa.me/27833095378"
              target="_blank"
              rel="noopener noreferrer"
              class="whatsapp-btn"
              aria-label="Chat on WhatsApp"
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16" height="16" fill="currentColor" aria-hidden="true">
                <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
              </svg>
              WhatsApp
            </a>
          </div>
        </div>

        <div class="info-item">
          <span class="info-label">Email</span>
          <a href="mailto:tours@footstepstofreedom.co.za" class="info-link">
            tours@footstepstofreedom.co.za
          </a>
        </div>

        <div class="info-item">
          <span class="info-label">Location</span>
          <span class="info-value">Cape Town, South Africa</span>
        </div>

        <div class="info-item">
          <span class="info-label">Response Time</span>
          <span class="info-value">Within 24 hours</span>
        </div>
      </div>

      <!-- Pull quote -->
      <blockquote class="info-quote">
        <p>"Footsteps to Freedom leaves footprints on your soul."</p>
        <cite>B. Bruyninckx, Belgium</cite>
      </blockquote>

      <!-- Social links -->
      <div class="social-row">
        <a
          href="https://www.facebook.com/footsteps2freedom.tours/"
          target="_blank"
          rel="noopener noreferrer"
          class="social-btn social-btn--facebook"
          aria-label="Footsteps to Freedom on Facebook"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20" fill="currentColor" aria-hidden="true">
            <path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>
          </svg>
        </a>

        <a
          href="https://www.instagram.com/footstepstofreedomtours/"
          target="_blank"
          rel="noopener noreferrer"
          class="social-btn social-btn--instagram"
          aria-label="Footsteps to Freedom on Instagram"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
            <rect x="2" y="2" width="20" height="20" rx="5" ry="5"/>
            <path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"/>
            <line x1="17.5" y1="6.5" x2="17.51" y2="6.5"/>
          </svg>
        </a>

        <a
          href="https://www.tripadvisor.com/Attraction_Review-g312659-d2098098-Reviews-Footsteps_to_Freedom_Tours-Cape_Town_Central_Western_Cape.html"
          target="_blank"
          rel="noopener noreferrer"
          class="social-btn social-btn--tripadvisor"
          aria-label="Footsteps to Freedom on TripAdvisor"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20" fill="currentColor" aria-hidden="true">
            <path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10 10-4.477 10-10S17.523 2 12 2zm0 18c-4.418 0-8-3.582-8-8s3.582-8 8-8 8 3.582 8 8-3.582 8-8 8zm-3.5-9a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm7 0a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm-3.5-2c-1.2 0-2.3.4-3.2 1H7l-1.5-1.5C6.8 7.2 9.3 6 12 6s5.2 1.2 6.5 3L17 10.5c-.9-.6-2-1-3.1-1H13c-.3-.3-.9-.5-1-.5zm0 7c-1.1 0-2-.4-2.7-1h5.4c-.7.6-1.6 1-2.7 1z"/>
          </svg>
        </a>
      </div>
    </div>

    <!-- Right: form panel -->
    <div class="form-panel">
      <ContactForm />
    </div>
  </section>
</div>

<style>
  /* ── Contact split layout ───────────────────────────────────────── */
  .contact-split {
    width: 100vw;
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    display: grid;
    grid-template-columns: 2fr 3fr;
    min-height: 600px;
  }

  /* ── Info panel (dark navy left col) ────────────────────────────── */
  .info-panel {
    background: var(--text-heading);
    padding: clamp(2.5rem, 5vw, 4rem) clamp(2rem, 4vw, 3.5rem);
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .info-eyebrow {
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--primary);
    margin: 0;
  }

  .info-heading {
    font-size: clamp(1.5rem, 3vw, 2.25rem);
    font-weight: 800;
    color: #ffffff;
    line-height: 1.2;
    margin: 0;
    letter-spacing: -0.02em;
  }

  .info-body {
    font-size: clamp(0.9rem, 1.5vw, 1rem);
    color: rgba(255, 255, 255, 0.7);
    line-height: 1.75;
    margin: 0;
  }

  /* ── Contact detail rows ────────────────────────────────────────── */
  .info-details {
    display: flex;
    flex-direction: column;
    gap: 0;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
  }

  .info-item {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;
    padding: 0.9rem 0;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  .info-label {
    font-size: 0.7rem;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--accent-gold);
  }

  .info-value {
    font-size: 0.95rem;
    color: rgba(255, 255, 255, 0.9);
  }

  .info-value-row {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    flex-wrap: wrap;
  }

  .info-link {
    font-size: 0.95rem;
    color: rgba(255, 255, 255, 0.9);
    text-decoration: none;
    transition: color 0.2s ease;
  }

  .info-link:hover {
    color: var(--primary);
    text-decoration: underline;
  }

  .whatsapp-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.3rem;
    color: #25D366;
    text-decoration: none;
    font-weight: 600;
    font-size: 0.85rem;
    padding: 0.2rem 0.55rem;
    border-radius: 0.375rem;
    border: 1.5px solid #25D366;
    transition: background-color 0.2s ease, color 0.2s ease;
    white-space: nowrap;
  }

  .whatsapp-btn:hover {
    background-color: #25D366;
    color: #ffffff;
  }

  /* ── Pull quote ─────────────────────────────────────────────────── */
  .info-quote {
    margin: 0.5rem 0 0;
    padding: 1rem 0 0;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    flex: 1;
  }

  .info-quote p {
    font-size: clamp(0.95rem, 1.5vw, 1.1rem);
    font-style: italic;
    color: rgba(255, 255, 255, 0.65);
    line-height: 1.7;
    margin: 0 0 0.5rem;
  }

  .info-quote cite {
    font-size: 0.8rem;
    font-style: normal;
    font-weight: 600;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--accent-gold);
  }

  /* ── Social buttons ─────────────────────────────────────────────── */
  .social-row {
    display: flex;
    gap: 0.75rem;
    align-items: center;
    padding-top: 0.5rem;
  }

  .social-btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 50%;
    text-decoration: none;
    color: #ffffff;
    transition: transform 0.2s ease, opacity 0.2s ease;
  }

  .social-btn:hover {
    transform: scale(1.1);
    opacity: 0.9;
  }

  .social-btn--facebook { background-color: #1877F2; }
  .social-btn--instagram {
    background: radial-gradient(circle at 30% 107%, #fdf497 0%, #fd5949 45%, #d6249f 60%, #285AEB 90%);
  }
  .social-btn--tripadvisor { background-color: #34E0A1; }

  /* ── Form panel (right col) ─────────────────────────────────────── */
  .form-panel {
    background: var(--body-bg);
    padding: clamp(2.5rem, 5vw, 4rem) clamp(2rem, 4vw, 4rem);
    display: flex;
    align-items: center;
  }

  /* Override ContactForm internal box-shadow/bg since the panel provides the container */
  .form-panel :global(.contact-form) {
    box-shadow: none;
    background: transparent;
    padding: 0;
    max-width: 100%;
    width: 100%;
  }

  /* ── Responsive ─────────────────────────────────────────────────── */
  @media (max-width: 860px) {
    .contact-split {
      grid-template-columns: 1fr;
    }

    .info-quote {
      display: none;
    }
  }

  @media (max-width: 540px) {
    .info-panel,
    .form-panel {
      padding: clamp(2rem, 5vw, 3rem) clamp(1.25rem, 4vw, 2rem);
    }
  }
</style>
