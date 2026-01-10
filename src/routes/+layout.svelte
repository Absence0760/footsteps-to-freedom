<script lang="ts">
import "normalize.css";
import "../theme.css";
import { Footer, Header } from "$lib/components/organisms";
</script>

<svelte:head>
  <title>FTF</title>
</svelte:head>

<Header />

<slot />

<Footer />
