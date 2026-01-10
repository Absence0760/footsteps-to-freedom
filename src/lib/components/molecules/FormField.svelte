<script lang="ts">
import { Input, Text } from "$lib/components/atoms";

import type { Variant } from "$lib/components/atoms";
import type { Snippet } from "svelte";

interface Props {
	variant?: Variant;
	label?: string;
	id?: string;
	type?: string;
	placeholder?: string;
	tag?: "input" | "textarea";
	disabled?: boolean;
	value: string | number;
	required?: boolean;
	rows?: number;
	children?: Snippet;
}

let {
	label = "",
	id = "",
	type = "text",
	placeholder = "",
	tag = "input",
	disabled = false,
	value = $bindable(""),
	required = false,
	rows,
	variant = "outlined",
	children,
	...restProps
}: Props = $props();
</script>

<div class="form-field-group">
  {#if label}
    <Text variant="label">{label}</Text>
  {/if}
  <Input
    {variant}
    {id}
    {type}
    {placeholder}
    {disabled}
    {required}
    {rows}
    bind:value
    {...restProps}
  />
  {#if children}
    {@render children()}
  {/if}
</div>

<style>
  .form-field-group {
    display: flex;
    flex-direction: column;
  }
</style>
