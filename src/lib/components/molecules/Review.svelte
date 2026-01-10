<script lang="ts">
import { Icon, Text } from "$lib/components/atoms";
import type { Snippet } from "svelte";
//   import StarIcon from '~icons/material-symbols/star-filled';

interface Props {
	rating: number;
	comment: string;
	author: string;
	children?: Snippet;
}

const { rating, comment, author, children }: Props = $props();

const maxRating = 5;
const stars = $derived(
	Array(maxRating)
		.fill(0)
		.map((_, i) => i < rating),
);
</script>

<div class="review">
  <div class="rating">
    {#each stars as isFilled, i}
      <Icon
        icon="star"
        size="1rem"
        color={isFilled ? 'var(--rating-filled, #f59e0b)' : 'var(--rating-empty, #d1d5db)'}
        ariaLabel={isFilled ? 'Filled star' : 'Empty star'}
      />
    {/each}
    <span class="sr-only">{rating} out of {maxRating} stars</span>
  </div>
  <Text variant="p">
    {comment}
  </Text>
  <Text variant="caption">
    — {author}
  </Text>
  {#if children}
    {@render children()}
  {/if}
</div>

<style>
  .review {
    background-color: var(--card-bg, #ffffff);
    padding: 1rem;
    border-radius: 0.5rem;
    box-shadow: var(--shadow-sm, 0 1px 2px 0 rgba(0, 0, 0, 0.05));
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .rating {
    display: flex;
    gap: 0.25rem;
  }

  .sr-only {
    position: absolute;
    width: 1px;
    height: 1px;
    padding: 0;
    margin: -1px;
    overflow: hidden;
    clip: rect(0, 0, 0, 0);
    border: 0;
  }
</style>