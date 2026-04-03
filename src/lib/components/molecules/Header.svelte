<script lang="ts">
import { Image, Link } from "$lib/components/atoms";
import { asset } from "$lib/utils/assets";

let menuOpen = false;

function toggleMenu() {
  menuOpen = !menuOpen;
}

function closeMenu() {
  menuOpen = false;
}
</script>

<header class="header">
  <nav class="nav nav-left">
    <Link variant="header" to="/">Home</Link>
    <Link variant="header" to="/about">Why Footsteps to Freedom</Link>
  </nav>

  <div class="logo-container">
    <Image
      src={asset('/logo.png')}
      alt="FTF Logo"
      width="65px"
      height="65px"
    />
    <div class="brand-text">
      <Link variant="header" to="/">FOOTSTEPS TO FREEDOM</Link>
      <p class="payoff-line">Discover the true spirit of South Africa</p>
    </div>
  </div>

  <nav class="nav nav-right">
    <Link variant="header" to="/tours">Tours</Link>
    <Link variant="header" to="/contact">Contact</Link>
    <Link variant="header" to="/reviews">Reviews</Link>
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
      <Link variant="header" to="/" onclick={closeMenu}>Home</Link>
      <Link variant="header" to="/about" onclick={closeMenu}>Why Footsteps to Freedom</Link>
      <Link variant="header" to="/tours" onclick={closeMenu}>Tours</Link>
      <Link variant="header" to="/contact" onclick={closeMenu}>Contact</Link>
      <Link variant="header" to="/reviews" onclick={closeMenu}>Reviews</Link>
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
    display: grid;
    grid-template-columns: 1fr auto 1fr;
    align-items: center;
    padding: 1rem 2rem;
    background-color: white;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    z-index: 100;
    --text-header: #1f2937;
    min-height: 80px;
  }

  .logo-container {
    display: flex;
    align-items: center;
    gap: 1rem;
    justify-self: center;
    text-align: center;
  }

  .brand-text {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

.payoff-line {
    margin: 0;
    font-size: 0.75rem;
    font-style: italic;
    color: var(--text-header);
    opacity: 0.9;
  }

  .nav {
    display: flex;
    gap: 0.1rem;
    align-items: center;
  }

  .nav-left {
    justify-self: start;
  }

  .nav-right {
    justify-self: end;
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
    .header {
      grid-template-columns: 1fr auto;
      justify-content: space-between;
      display: flex;
    }

    .nav {
      display: none;
    }

    .hamburger {
      display: flex;
    }

    .logo-container {
      justify-self: start;
    }

    .logo-container :global(.link.header) {
      font-size: 0.8rem;
    }

    .payoff-line {
      font-size: 0.65rem;
    }
  }
</style>
