<script lang="ts">
import { base } from '$app/paths';
import type { Snippet } from 'svelte';

export interface Slide {
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

const { slides, interval = 5500, children }: Props = $props();

const resolve = (src: string) =>
	src.startsWith('/') && !src.startsWith('//') ? `${base}${src}` : src;

let current = $state(0);
let paused = $state(false);

let touchStartX = 0;
let touchStartY = 0;

function next() {
	current = (current + 1) % slides.length;
}

function prev() {
	current = (current - 1 + slides.length) % slides.length;
}

function goTo(index: number) {
	current = index;
}

function handleTouchStart(e: TouchEvent) {
	touchStartX = e.touches[0].clientX;
	touchStartY = e.touches[0].clientY;
}

function handleTouchEnd(e: TouchEvent) {
	const dx = e.changedTouches[0].clientX - touchStartX;
	const dy = e.changedTouches[0].clientY - touchStartY;
	if (Math.abs(dx) > Math.abs(dy) && Math.abs(dx) > 40) {
		if (dx < 0) next();
		else prev();
	}
}

function handleKeydown(e: KeyboardEvent) {
	if (e.key === 'ArrowLeft') prev();
	else if (e.key === 'ArrowRight') next();
}

$effect(() => {
	const id = setInterval(() => {
		if (!paused) next();
	}, interval);
	return () => clearInterval(id);
});
</script>

<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<section
	class="hero-slideshow"
	aria-label="Hero image slideshow"
	onmouseenter={() => (paused = true)}
	onmouseleave={() => (paused = false)}
	ontouchstart={handleTouchStart}
	ontouchend={handleTouchEnd}
	onkeydown={handleKeydown}
>
	{#each slides as slide, i}
		<div class="slide" class:active={current === i} aria-hidden={current !== i}>
			<div
				class="slide-bg"
				style:--slide-bg={`url(${resolve(slide.image)})`}
				style:--slide-mobile-bg={slide.mobileImage
					? `url(${resolve(slide.mobileImage)})`
					: `url(${resolve(slide.image)})`}
				role="img"
				aria-label={slide.alt}
			></div>
		</div>
	{/each}

	<div class="hero-overlay"></div>

	<div class="hero-content">
		{#if children}
			{@render children()}
		{/if}
		<p class="slide-caption" aria-live="polite">
			{slides[current].caption ?? ''}
		</p>
	</div>

	<button class="nav-btn prev" onclick={prev} aria-label="Previous slide">&#8249;</button>
	<button class="nav-btn next" onclick={next} aria-label="Next slide">&#8250;</button>

	<div class="dots" role="group" aria-label="Slide navigation">
		{#each slides as _, i}
			<button
				class="dot"
				class:active={current === i}
				onclick={() => goTo(i)}
				aria-label={`Go to slide ${i + 1}`}
				aria-pressed={current === i}
			></button>
		{/each}
	</div>
</section>

<style>
	.hero-slideshow {
		position: relative;
		min-height: 100vh;
		height: 60vh;
		padding-top: 5rem;
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 1;
		overflow: hidden;
	}

	/* === Slides === */
	.slide {
		position: absolute;
		inset: 0;
		opacity: 0;
		transition: opacity 1s ease-in-out;
	}

	.slide.active {
		opacity: 1;
	}

	.slide-bg {
		position: absolute;
		inset: 0;
		background-image: var(--slide-bg);
		background-size: cover;
		background-position: center;
		background-repeat: no-repeat;
	}

	/* Ken Burns animations — 4 variants, one per slide */
	@keyframes ken-burns-1 {
		from { transform: scale(1) translate(0, 0); }
		to   { transform: scale(1.12) translate(-2%, -1%); }
	}
	@keyframes ken-burns-2 {
		from { transform: scale(1.1) translate(2%, 1%); }
		to   { transform: scale(1) translate(0, 0); }
	}
	@keyframes ken-burns-3 {
		from { transform: scale(1) translate(1%, 0); }
		to   { transform: scale(1.1) translate(-1%, -2%); }
	}
	@keyframes ken-burns-4 {
		from { transform: scale(1.08) translate(-1%, 1%); }
		to   { transform: scale(1) translate(1%, 0); }
	}

	/* Animation restarts from the beginning each time .active is applied */
	.slide:nth-child(1).active .slide-bg { animation: ken-burns-1 8s ease-in-out forwards; }
	.slide:nth-child(2).active .slide-bg { animation: ken-burns-2 8s ease-in-out forwards; }
	.slide:nth-child(3).active .slide-bg { animation: ken-burns-3 8s ease-in-out forwards; }
	.slide:nth-child(4).active .slide-bg { animation: ken-burns-4 8s ease-in-out forwards; }

	/* === Overlay === */
	.hero-overlay {
		position: absolute;
		inset: 0;
		background-color: rgba(0, 0, 0, 0.45);
		z-index: 1;
	}

	/* === Content === */
	.hero-content {
		position: relative;
		z-index: 2;
		text-align: center;
		padding: clamp(1.5rem, 4vw, 2rem);
		max-width: 800px;
		margin: 0 auto;
	}

	.hero-content :global(.text) {
		color: #ffffff;
		font-size: 1.1rem;
		font-weight: 600;
		line-height: 1.5;
		margin-bottom: 0.5rem;
	}

	.slide-caption {
		color: rgba(255, 255, 255, 0.75);
		font-size: 0.85rem;
		font-style: italic;
		letter-spacing: 0.05em;
		margin-top: 1.25rem;
		min-height: 1.2em;
	}

	/* === Navigation arrows === */
	.nav-btn {
		position: absolute;
		top: 50%;
		transform: translateY(-50%);
		z-index: 3;
		background: rgba(0, 0, 0, 0.35);
		border: 2px solid transparent;
		color: #fff;
		font-size: 2.5rem;
		line-height: 1;
		padding: 0.25rem 0.8rem 0.35rem;
		cursor: pointer;
		border-radius: 4px;
		transition: background 0.2s, border-color 0.2s;
		user-select: none;
	}

	.nav-btn:hover,
	.nav-btn:focus-visible {
		background: rgba(0, 0, 0, 0.6);
		border-color: #fff;
		outline: none;
	}

	.nav-btn.prev { left: 1rem; }
	.nav-btn.next { right: 1rem; }

	/* === Dot indicators === */
	.dots {
		position: absolute;
		bottom: 1.25rem;
		left: 50%;
		transform: translateX(-50%);
		z-index: 3;
		display: flex;
		gap: 0.65rem;
	}

	.dot {
		width: 10px;
		height: 10px;
		border-radius: 50%;
		border: 2px solid rgba(255, 255, 255, 0.7);
		background: transparent;
		cursor: pointer;
		padding: 0;
		transition: background 0.25s, border-color 0.25s;
	}

	.dot.active {
		background: #fff;
		border-color: #fff;
	}

	.dot:hover,
	.dot:focus-visible {
		border-color: #fff;
		background: rgba(255, 255, 255, 0.5);
		outline: none;
	}

	/* === Mobile === */
	@media (max-width: 600px) {
		.hero-slideshow {
			height: 50vh;
		}

		.slide-bg {
			background-image: var(--slide-mobile-bg);
		}

		.nav-btn {
			font-size: 2rem;
			padding: 0.2rem 0.6rem 0.3rem;
		}

		.nav-btn.prev { left: 0.5rem; }
		.nav-btn.next { right: 0.5rem; }
	}

	/* === Reduced motion — disable Ken Burns and crossfade === */
	@media (prefers-reduced-motion: reduce) {
		.slide { transition: none; }
		.slide-bg { animation: none !important; }
	}
</style>
