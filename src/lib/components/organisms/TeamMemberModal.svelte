<script lang="ts">
import { base } from '$app/paths';

interface MemberData {
  name: string;
  role: string;
  bio: string;
  image: string;
  slug: string;
  component: any;
}

interface Props {
  member: MemberData | null;
  onClose: () => void;
}

let { member, onClose }: Props = $props();

function handleBackdropClick(e: MouseEvent) {
  if (e.target === e.currentTarget) onClose();
}

function handleKeydown(e: KeyboardEvent) {
  if (e.key === 'Escape') onClose();
}

const imageSrc = $derived(
  member?.image && member.image.startsWith('/') && !member.image.startsWith('//')
    ? `${base}${member.image}`
    : member?.image ?? ''
);
</script>

<svelte:window onkeydown={handleKeydown} />

{#if member}
<div
  class="modal-backdrop"
  onclick={handleBackdropClick}
  role="dialog"
  aria-modal="true"
  aria-label={member.name}
>
  <div class="modal">
    <button class="close-btn" onclick={onClose} aria-label="Close modal">&#x2715;</button>

    <div class="modal-hero">
      <img src={imageSrc} alt={member.name} class="hero-image" />
      <div class="hero-overlay">
        <span class="role-badge">{member.role}</span>
        <h2 class="member-name">{member.name}</h2>
      </div>
    </div>

    <div class="modal-body">
      <svelte:component this={member.component} />
    </div>
  </div>
</div>
{/if}

<style>
  .modal-backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.65);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    padding: 1rem;
    backdrop-filter: blur(2px);
  }

  .modal {
    background: var(--card-bg, #ffffff);
    border-radius: 1rem;
    width: 100%;
    max-width: 720px;
    max-height: 90vh;
    overflow-y: auto;
    position: relative;
    box-shadow: var(--shadow-2xl, 0 25px 50px -12px rgba(0, 0, 0, 0.35));
    display: flex;
    flex-direction: column;
  }

  .close-btn {
    position: absolute;
    top: 0.75rem;
    right: 0.75rem;
    z-index: 10;
    background: rgba(0, 0, 0, 0.5);
    color: #fff;
    border: none;
    border-radius: 50%;
    width: 2rem;
    height: 2rem;
    font-size: 0.875rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background 0.2s ease;
  }

  .close-btn:hover {
    background: rgba(0, 0, 0, 0.75);
  }

  .modal-hero {
    position: relative;
    width: 100%;
    flex-shrink: 0;
  }

  .hero-image {
    width: 100%;
    height: 320px;
    object-fit: cover;
    object-position: center top;
    display: block;
    border-radius: 1rem 1rem 0 0;
  }

  .hero-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(to top, rgba(0,0,0,0.75) 0%, rgba(0,0,0,0.1) 60%);
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    padding: 1.5rem;
    border-radius: 1rem 1rem 0 0;
  }

  .role-badge {
    display: inline-block;
    background: var(--primary, #FF6B35);
    color: #fff;
    font-size: 0.7rem;
    font-weight: 700;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    padding: 0.25rem 0.6rem;
    border-radius: 999px;
    margin-bottom: 0.5rem;
    align-self: flex-start;
  }

  .member-name {
    color: #fff;
    margin: 0;
    font-size: clamp(1.25rem, 4vw, 1.75rem);
    font-weight: 700;
    line-height: 1.2;
    text-shadow: 0 1px 4px rgba(0,0,0,0.4);
  }

  .modal-body {
    padding: 1.75rem 2rem;
    line-height: 1.75;
    color: var(--text-body, #2D3748);
    font-size: 0.95rem;
  }

  .modal-body :global(h1),
  .modal-body :global(h2),
  .modal-body :global(h3) {
    color: var(--text-heading, #1A1A2E);
    margin-top: 1.5rem;
    margin-bottom: 0.5rem;
  }

  .modal-body :global(p) {
    margin-bottom: 1rem;
  }

  .modal-body :global(ul) {
    padding-left: 1.5rem;
    margin-bottom: 1rem;
  }

  .modal-body :global(li) {
    margin-bottom: 0.4rem;
  }

  .modal-body :global(strong) {
    color: var(--text-heading, #1A1A2E);
    font-weight: 600;
  }

  @media (max-width: 600px) {
    .modal-backdrop {
      padding: 0;
      align-items: flex-end;
    }

    .modal {
      max-height: 92vh;
      border-radius: 1rem 1rem 0 0;
      max-width: 100%;
    }

    .hero-image {
      height: 220px;
      border-radius: 1rem 1rem 0 0;
    }

    .hero-overlay {
      border-radius: 1rem 1rem 0 0;
    }

    .modal-body {
      padding: 1.25rem 1rem;
    }
  }
</style>
