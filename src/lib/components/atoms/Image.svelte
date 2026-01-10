<script lang="ts">
import { base } from '$app/paths';

interface Props {
	src: string;
	alt?: string;
	width?: string;
	height?: string;
	borderRadius?: string;
	shadow?: string;
	class?: string;
}

const {
	src: rawSrc,
	alt = "/hero-image4.png",
	width = "100%",
	height = "auto",
	borderRadius = "0.5rem",
	shadow = "var(--shadow-sm, 0 1px 2px 0 rgba(0, 0, 0, 0.05))",
	class: className = "",
}: Props = $props();

// Prepend base path to relative URLs (starting with /)
const src = rawSrc.startsWith('/') && !rawSrc.startsWith('//') ? `${base}${rawSrc}` : rawSrc;
</script>

<img
  {src}
  {alt}
  class={className}
  style:width={width}
  style:height={height}
  style:border-radius={borderRadius}
  style:box-shadow={shadow}
/>

<style>
  .image {
    display: block;
    object-fit: cover;
  }
</style>