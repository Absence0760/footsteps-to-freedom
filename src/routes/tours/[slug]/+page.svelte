<script lang="ts">
import { Link, Text } from "$lib/components/atoms";

const { data } = $props();
const tour = $derived(data.tour);
</script>

<svelte:head>
  <title>{tour.name} - Footsteps to Freedom</title>
  <meta name="description" content="Learn more about {tour.name} at Footsteps to Freedom Tours." />
</svelte:head>

{#if tour}
<div class="page">
  <div class="bio-container">
    <img
      src={tour.image}
      alt={tour.name}
      class="bio-image"
      loading="lazy"
    />
    <div class="tour-content">
      <svelte:component this={tour.component} />
    </div>
    <div class="back-link">
      <Link variant="button" to="/tours">← Back to Tours</Link>
    </div>
  </div>
</div>
{:else}
  <p>Loading tour details...</p>
{/if}

<style>
  .page {
    padding-top: 5rem;
    padding-inline: 1rem;
  }

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

  .tour-content {
    line-height: 1.7;
  }

  .back-link {
    margin-top: 2rem;
  }

  @media (max-width: 768px) {
    .bio-container {
      padding: 1.5rem;
    }

    .bio-image {
      width: 150px;
      height: 150px;
    }
  }

  @media (max-width: 480px) {
    .bio-image {
      width: 120px;
      height: 120px;
    }
  }
</style>
