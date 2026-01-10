<script lang="ts">
import { Link, Text } from "$lib/components/atoms";
import DOMPurify from 'isomorphic-dompurify';

const { data } = $props();
const member = $derived(data.member);

// Sanitize HTML content to prevent XSS attacks
const sanitizedFullBio = $derived(member ? DOMPurify.sanitize(member.fullBio) : '');
</script>

<svelte:head>
  <title>{member.name} - Footsteps to Freedom Tours</title>
  <meta name="description" content="Learn more about {member.name}, {member.role} at Footsteps to Freedom Tours." />
</svelte:head>

{#if member}
<div class="bio-page">
  <Link to="/about">
    ← Back to About
  </Link>
  <div class="bio-container">
    <img
      src={member.image}
      alt={`Portrait of ${member.name}, ${member.role}`}
      class="bio-image"
      loading="lazy"
    />
    <Text
      tag="h1"
      class="bio-name"
      color="var(--text-heading, #1f2937)"
      fontSize="2rem"
      fontWeight="bold"
      marginBottom="1rem"
      textAlign="center"
    >
      {member.name}
    </Text>
    <Text
      tag="p"
      class="bio-role"
      color="var(--text-body, #4b5563)"
      fontSize="1.25rem"
      marginBottom="1.5rem"
      textAlign="center"
    >
      {member.role}
    </Text>
    <Text
      tag="div"
      class="bio-content"
      color="var(--text-body, #4b5563)"
      fontSize="1rem"
      lineHeight="1.7"
      textAlign="left"
    >
      {@html sanitizedFullBio}
    </Text>
  </div>
</div>
{:else}
  <p>Loading tour details...</p>
{/if}

<style>
  .bio-page {
    max-width: 1200px;
    margin: 0 auto;
    padding: 2rem;
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
    display: inline-block;
    margin-bottom: 2.5rem; /* Increased spacing */
    font-size: 1rem;
    font-weight: 500;
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