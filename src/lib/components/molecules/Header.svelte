<script lang="ts">
import { Image, Link } from "$lib/components/atoms";
import { onMount } from "svelte"

let scrolled = false;
let menuOpen = false;

onMount(() => {
    const handleScroll = () => {
      scrolled = window.scrollY > 50;
    };

    window.addEventListener('scroll', handleScroll);

    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  });

function toggleMenu() {
  menuOpen = !menuOpen;
}

function closeMenu() {
  menuOpen = false;
}
</script>

<header class="header" class:scrolled>
  <div class="logo-container">
    <Image
      src="/hero-image.png"
      alt="FTF Logo"
      width="50px"
      height="50px"
    />
    <Link variant="header" to="/">FOOTSTEPS TO FREEDOM</Link>
  </div>

  <nav class="nav">
    <Link variant="header" to="/about">About</Link>
    <Link variant="header" to="/tours">Tours</Link>
    <Link variant="header" to="/contact">Contact</Link>
  </nav>

  <button class="hamburger" onclick={toggleMenu} aria-label="Toggle menu" aria-expanded={menuOpen}>
    <span class:open={menuOpen}></span>
    <span class:open={menuOpen}></span>
    <span class:open={menuOpen}></span>
  </button>
</header>

{#if menuOpen}
  <!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
  <div class="modal-overlay" onclick={closeMenu}>
    <nav class="modal-nav" onclick={(e) => e.stopPropagation()}>
      <Link variant="header" to="/about" onclick={closeMenu}>About</Link>
      <Link variant="header" to="/tours" onclick={closeMenu}>Tours</Link>
      <Link variant="header" to="/contact" onclick={closeMenu}>Contact</Link>
    </nav>
  </div>
{/if}

<style>
  .header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    box-sizing: border-box;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    box-shadow: var(--shadow-sm, 0 1px 2px 0 rgba(0, 0, 0, 0.05));
    transition: background-color 0.3s ease, box-shadow 0.3s ease;
    z-index: 100;
    --text-header: #ffffff;
  }

  .header.scrolled {
    background-color: white;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    --text-header: #1f2937;
  }

  .logo-container {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .nav {
    display: flex;
    gap: 0.1rem;
    align-items: center;
  }

  /* Hamburger button — hidden on desktop */
  .hamburger {
    display: none;
    flex-direction: column;
    justify-content: center;
    gap: 5px;
    background: none;
    border: none;
    cursor: pointer;
    padding: 0.25rem;
  }

  .hamburger span {
    display: block;
    width: 24px;
    height: 2px;
    background-color: var(--text-header, #ffffff);
    transition: transform 0.2s ease, opacity 0.2s ease;
  }

  .hamburger span.open:nth-child(1) {
    transform: translateY(7px) rotate(45deg);
  }

  .hamburger span.open:nth-child(2) {
    opacity: 0;
  }

  .hamburger span.open:nth-child(3) {
    transform: translateY(-7px) rotate(-45deg);
  }

  /* Full-screen modal */
  .modal-overlay {
    position: fixed;
    inset: 0;
    background: transparent;
    z-index: 99;
    display: flex;
    align-items: flex-start;
    justify-content: flex-end;
  }

  .modal-nav {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 0.25rem;
    padding-top: 80px;
    padding-right: 1.5rem;
    --text-header: #ffffff;
  }

  .modal-nav :global(.link.header) {
    font-size: 1rem;
    font-weight: 600;
    padding: 0.5rem 1rem;
    text-shadow: 0 2px 8px rgba(0, 0, 0, 0.6);
  }

  .modal-nav :global(.link.header:hover) {
    background-color: rgba(255, 255, 255, 0.15);
    color: #ffffff;
  }

  @media (max-width: 640px) {
    .nav {
      display: none;
    }

    .hamburger {
      display: flex;
    }

    .logo-container :global(.link.header) {
      font-size: 0.8rem;
    }
  }
</style>
