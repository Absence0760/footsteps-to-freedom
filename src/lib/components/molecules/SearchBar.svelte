<script lang="ts">
import { Button, Icon, Input } from "$lib/components/atoms";
import type { Snippet } from "svelte";
import SearchIcon from "~icons/material-symbols/search-rounded";

interface Props {
	query?: string;
	placeholder?: string;
	onSearch?: (query: string) => void;
	children?: Snippet;
}

const {
	query = $bindable(""),
	placeholder = "Search tours...",
	onSearch = () => {},
	children,
}: Props = $props();
</script>

<div class="search-bar">
  <Input
    value={query}
    placeholder={placeholder}
    id="search-input"
    name="search"
    label="Search tours"
  >
    <Icon
      icon="tour"
      size="1.25rem"
      color="var(--text-body, #4b5563)"
    />
  </Input>
  <Button
    variant="primary"
    onclick={() => onSearch(query)}
    label="Submit search"
  >
    Search
  </Button>
  {#if children}
    {@render children()}
  {/if}
</div>

<style>
  .search-bar {
    display: flex;
    gap: 1rem;
    align-items: center;
    max-width: 600px;
    margin: 0 auto;
    padding: 1rem;
  }
</style>