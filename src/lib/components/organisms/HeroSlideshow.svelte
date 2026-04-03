<script lang="ts">
import { base } from '$app/paths';
import type { Snippet } from 'svelte';

interface Slide {
	image: string;
	mobileImage?: string;
	alt: string;
	caption?: string;
}

interface Props {
	slides: Slide[];
	interval?: number;
	children?: Snippet;
}

const { slides, interval = 5000, children }: Props = $props();

const resolve = (src: string) =>
	src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;

const resolvedSlides = $derived(
	slides.map((s) => ({
		...s,
		image: resolve(s.image),
		mobileImage: s.mobileImage ? resolve(s.mobileImage) : resolve(s.image),
	}))
);

let current = $state(0);
let paused = $state(false);
let reducedMotion = $state(false);

// Touch tracking
let touchStartX = $state(0);
let touchStartY = $state(0);

function prev() {
	current = (current - 1 + resolvedSlides.length) % resolvedSlides.length;
}

function next() {
	current = (current + 1) % resolvedSlides.length;
}

function goTo(index: number) {
	current = index;
}

function onTouchStart(e: TouchEvent) {
	touchStartX = e.touches[0].clientX;
	touchStartY = e.touches[0].clientY;
}

function onTouchEnd(e: TouchEvent) {
	const dx = e.changedTouches[0].clientX - touchStartX;
	const dy = e.changedTouches[0].clientY - touchStartY;
	// Only swipe if horizontal movement is dominant
	if (Math.abs(dx) > Math.abs(dy) && Math.abs(dx) > 40) {
		if (dx < 0) next();
		else prev();
	}
}

$effect(() => {
	// Check prefers-reduced-motion
	const mq = window.matchMedia('(prefers-reduced-motion: reduce)');
	reducedMotion = mq.matches;
	const handler = (e: MediaQueryListEvent) => { reducedMotion = e.matches; };
	mq.addEventListener('change', handler);

	return () => mq.removeEventListener('change', handler);
});

$effect(() => {
	if (paused || reducedMotion || resolvedSlides.length <= 1) return;
	const id = setInterval(next, interval);
	return () => clearInterval(id);
});
</script>

<section
	class="slideshow"
	aria-label="Hero image slideshow"
	aria-roledescription="carousel"
	onmouseenter={() => (paused = true)}
	onmouseleave={() => (paused = false)}
	ontouchstart={onTouchStart}
	ontouchend={onTouchEnd}
>
	<!-- Slides -->
	{#each resolvedSlides as slide, i}
		<div
			class="slide"
			class:active={i === current}
			class:reduced={reducedMotion}
			aria-hidden={i !== current}
			role="group"
			aria-roledescription="slide"
			aria-label="{i + 1} of {resolvedSlides.length}: {slide.alt}"
			style:--slide-bg="url({slide.image})"
			style:--slide-mobile-bg="url({slide.mobileImage})"
		></div>
	{/each}

	<!-- Dark overlay -->
	<div class="overlay" aria-hidden="true"></div>

	<!-- Content -->
	<div class="content">
		{#if children}
			{@render children()}
		{/if}
		{#if resolvedSlides[current].caption}
			<p class="caption">{resolvedSlides[current].caption}</p>
		{/if}
	</div>

	<!-- Previous / Next arrows -->
	{#if resolvedSlides.length > 1}
		<button
			class="arrow arrow-prev"
			onclick={prev}
			aria-label="Previous slide"
		>
			<svg viewBox="0 0 24 24" aria-hidden="true" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
				<polyline points="15 18 9 12 15 6" />
			</svg>
		</button>
		<button
			class="arrow arrow-next"
			onclick={next}
			aria-label="Next slide"
		>
			<svg viewBox="0 0 24 24" aria-hidden="true" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
				<polyline points="9 18 15 12 9 6" />
			</svg>
		</button>

		<!-- Dot indicators -->
		<div class="dots" role="tablist" aria-label="Slide navigation">
			{#each resolvedSlides as _, i}
				<button
					class="dot"
					class:dot-active={i === current}
					role="tab"
					aria-selected={i === current}
					aria-label="Go to slide {i + 1}"
					onclick={() => goTo(i)}
				></button>
			{/each}
		</div>
	{/if}
</section>

<style>
	.slideshow {
		position: relative;
		width: 100%;
		min-height: 100vh;
		height: 60vh;
		overflow: hidden;
		display: flex;
		align-items: center;
		justify-content: center;
		padding-top: 5rem;
	}

	/* ---- Slides ---- */
	.slide {
		position: absolute;
		inset: 0;
		background-image: var(--slide-bg);
		background-size: cover;
		background-position: center;
		background-attachment: fixed;
		background-repeat: no-repeat;
		opacity: 0;
		transition: opacity 1s ease-in-out;
		will-change: opacity;
	}

	.slide.active {
		opacity: 1;
	}

	/* Skip transition when reduced motion is preferred */
	.slide.reduced {
		transition: none;
	}

	/* ---- Overlay ---- */
	.overlay {
		position: absolute;
		inset: 0;
		background-color: rgba(0, 0, 0, 0.45);
		z-index: 1;
		pointer-events: none;
	}

	/* ---- Content ---- */
	.content {
		position: relative;
		z-index: 2;
		text-align: center;
		padding: clamp(1.5rem, 4vw, 2rem);
		max-width: 800px;
		margin: 0 auto;
		color: #fff;
	}

	.content :global(.text) {
		color: #ffffff;
		font-size: 1.1rem;
		font-weight: 600;
		line-height: 1.5;
		margin-bottom: 0.5rem;
	}

	.caption {
		margin-top: 1rem;
		font-size: 0.85rem;
		letter-spacing: 0.06em;
		text-transform: uppercase;
		color: rgba(255, 255, 255, 0.75);
		font-weight: 500;
	}

	/* ---- Arrows ---- */
	.arrow {
		position: absolute;
		top: 50%;
		transform: translateY(-50%);
		z-index: 3;
		background: rgba(0, 0, 0, 0.35);
		border: none;
		border-radius: 50%;
		width: 2.75rem;
		height: 2.75rem;
		display: flex;
		align-items: center;
		justify-content: center;
		cursor: pointer;
		color: #fff;
		transition: background 0.2s ease;
		padding: 0;
	}

	.arrow:hover,
	.arrow:focus-visible {
		background: rgba(0, 0, 0, 0.65);
		outline: 2px solid rgba(255, 255, 255, 0.8);
		outline-offset: 2px;
	}

	.arrow svg {
		width: 1.25rem;
		height: 1.25rem;
	}

	.arrow-prev {
		left: 1rem;
	}

	.arrow-next {
		right: 1rem;
	}

	/* ---- Dots ---- */
	.dots {
		position: absolute;
		bottom: 1.25rem;
		left: 50%;
		transform: translateX(-50%);
		z-index: 3;
		display: flex;
		gap: 0.5rem;
		align-items: center;
	}

	.dot {
		width: 0.6rem;
		height: 0.6rem;
		border-radius: 50%;
		border: 2px solid rgba(255, 255, 255, 0.8);
		background: transparent;
		cursor: pointer;
		padding: 0;
		transition: background 0.2s ease, transform 0.2s ease;
	}

	.dot:hover,
	.dot:focus-visible {
		outline: 2px solid rgba(255, 255, 255, 0.9);
		outline-offset: 2px;
	}

	.dot-active {
		background: #fff;
		transform: scale(1.25);
	}

	/* ---- Responsive ---- */
	@media (max-width: 600px) {
		.slideshow {
			height: 50vh;
		}

		.slide {
			background-image: var(--slide-mobile-bg);
			background-attachment: scroll;
		}

		.arrow {
			width: 2.25rem;
			height: 2.25rem;
		}

		.arrow-prev {
			left: 0.5rem;
		}

		.arrow-next {
			right: 0.5rem;
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.slide {
			background-attachment: scroll;
		}
	}
</style>
