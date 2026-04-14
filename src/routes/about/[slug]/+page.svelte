<script lang="ts">
import { base } from '$app/paths';
import { Link, Text } from "$lib/components/atoms";

const { data } = $props();
const member = $derived(data.member);
</script>

<svelte:head>
  <title>{member.name} - Footsteps to Freedom Tours</title>
  <meta name="description" content="Learn more about {member.name}, {member.role} at Footsteps to Freedom Tours." />
</svelte:head>

{#if member}
<div class="bio-page">
  <div class="bio-container">
    <img
      src="{base}{member.image}"
      alt={`Portrait of ${member.name}, ${member.role}`}
      class="bio-image"
      loading="lazy"
    />
    <Text
      tag="h1"
      class="bio-name"
      style="color: var(--text-heading, #1f2937); font-size: 2rem; font-weight: bold; margin-bottom: 1rem; text-align: center;"
    >
      {member.name}
    </Text>
    <Text
      tag="p"
      class="bio-role"
      style="color: var(--text-body, #4b5563); font-size: 1.25rem; margin-bottom: 1.5rem; text-align: center;"
    >
      {member.role}
    </Text>
    <div class="bio-content">
      <svelte:component this={member.component} />
    </div>
    <div class="back-link">
      <button class="back-button" onclick={() => history.back()}>← Back</button>
    </div>
  </div>
</div>
{:else}
  <p>Loading tour details...</p>
{/if}

<style>
  .bio-page {
    max-width: 1200px;
    margin: 0 auto;
    padding: 5rem 2rem 2rem;
    box-sizing: border-box;
    background-color: transparent; /* Remove explicit background */
  }

  .bio-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 2rem;
    background-color: var(--card-bg, #f7f7f7); /* Matches TeamMemberCard */
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
    border: 3px solid var(--button-primary-bg, #4f46e5); /* Matches TeamMemberCard */
  }

  .back-link {
    margin-top: 2rem;
  }

  .back-button {
    display: inline-block;
    padding: 0.625rem 1.25rem;
    border-radius: 0.375rem;
    background: var(--button-primary-bg, #4f46e5);
    color: var(--button-primary-text, #fff);
    font-size: 1rem;
    font-weight: 600;
    border: none;
    cursor: pointer;
  }

  .back-button:hover {
    opacity: 0.9;
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

  }
</style>