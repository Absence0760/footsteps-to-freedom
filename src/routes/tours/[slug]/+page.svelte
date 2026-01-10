<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import DOMPurify from 'isomorphic-dompurify';

const { data } = $props();
const tour = $derived(data.tour);

// Sanitize HTML content to prevent XSS attacks
const sanitizedFullTour = $derived(tour ? DOMPurify.sanitize(tour.fullTour) : '');
</script>

<svelte:head>
  <title>{tour.name} - Footsteps to Freedom</title>
  <meta name="description" content="Learn more about {tour.name} at Footsteps to Freedom Tours." />
</svelte:head>

{#if tour}
<div>
  <Link to="/tours">
    ← Back to Tours
  </Link>
  
  <div class="bio-container">
    <img
      src={tour.image}
      alt={` ${tour.name}`}
      class="bio-image"
      loading="lazy"
    />
    <p>
      {@html sanitizedFullTour}
    </p>
  </div>
</div>
{:else}
  <p>Loading tour details...</p>
{/if}

<style>

  .bio-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 2rem;
    background-color: var(--card-bg, #f7f7f7); 
    border-radius: 0.75rem;
    box-shadow: var(--shadow-md, 0 4px 6px -1px rgba(0, 0, 0, 0.1));
  }

  .bio-image {
    width: 200px;
    height: 200px;
    border-radius: 50%;
    object-fit: cover;
    margin: 0 auto 1.5rem;
    display: block;
    border: 3px solid var(--button-primary-bg, #4f46e5);
  }

  @media (max-width: 768px) {
    .bio-container {
      padding: 1.5rem;
    }

    .bio-image {
      width: 150px;
      height: 150px;
    }

    .bio-name {
      font-size: 1.75rem;
    }

    .bio-role {
      font-size: 1.1rem;
    }

    .back-link {
      margin-bottom: 2rem; /* Slightly reduced for smaller screens */
    }
  }

  @media (max-width: 480px) {
    .bio-page {
      padding: 1rem;
    }

    .bio-image {
      width: 120px;
      height: 120px;
    }

    .bio-name {
      font-size: 1.5rem;
    }

    .bio-role {
      font-size: 1rem;
    }

    .bio-content {
      font-size: 0.95rem;
    }

    .back-link {
      margin-bottom: 1.5rem; /* Further reduced for mobile */
    }
  }
</style>