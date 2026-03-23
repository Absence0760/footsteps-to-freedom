<script lang="ts">
import { Image, Link } from "$lib/components/atoms";
import { onMount } from "svelte"

let scrolled = false;

onMount(() => {
    const handleScroll = () => {
      scrolled = window.scrollY > 50;
    };
    
    window.addEventListener('scroll', handleScroll);
    
    // Cleanup
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  });
</script>

<header class="header" class:scrolled>
  <div class="logo-container">
    <Image
      src="/hero-image4.png"
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
</header>

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
    z-index: 10;
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
</style>
