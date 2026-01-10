<script lang="ts">
import { Icon } from "$lib/components/atoms";
import type { Snippet } from "svelte";

export type Variant = "outlined" | "filled" | "minimal";
type IconName = "plus" | "home" | "about" | "contact" | "tour";

interface Props {
	variant?: Variant;
	type?: string;
	value?: string | number;
	placeholder?: string;
	required?: boolean;
	id?: string;
	name?: string;
	rows?: number | null;
	disabled?: boolean;
	label?: string;
	error?: string;
	icon?: IconName;
	children?: Snippet;
}

let {
	variant = "outlined",
	type = "text",
	value = $bindable(""),
	placeholder = "",
	required = false,
	id = "",
	name = "",
	rows = null,
	disabled = false,
	label = "",
	error = "",
	icon,
	children,
}: Props = $props();
</script>

<div class="input-wrapper {variant}" class:error>
  {#if label}
    <label for={id} class="input-label">{label}</label>
  {/if}
  <div class="input-container">
    {#if icon}
      <Icon
        {icon}
        size="1.25rem"
        color="var(--text-input, #6b7280)"
        ariaLabel={`${label || placeholder || 'input'} icon`}
      />
    {/if}
    {#if type === 'textarea'}
      <textarea
        {id}
        {name}
        {placeholder}
        {required}
        {disabled}
        {rows}
        bind:value
        class="input textarea {variant}"
        class:with-icon={icon}
        class:error
        aria-label={label || placeholder || undefined}
        aria-invalid={error ? 'true' : undefined}
        aria-describedby={error && id ? `${id}-error` : undefined}
      ></textarea>
    {:else}
      <input
        {type}
        {id}
        {name}
        {placeholder}
        {required}
        {disabled}
        bind:value
        class="input {variant}"
        class:with-icon={icon}
        class:error
        aria-label={label || placeholder || undefined}
        aria-invalid={error ? 'true' : undefined}
        aria-describedby={error && id ? `${id}-error` : undefined}
      />
    {/if}
  </div>
  {#if error}
    <span class="error-message" id={id ? `${id}-error` : undefined}>{error}</span>
  {/if}
  {#if children}
    {@render children()}
  {/if}
</div>

<style>
  .input-wrapper {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .input-label {
    font-size: 0.875rem;
    font-weight: 500;
    color: var(--text-label, #374151);
  }

  .input-container {
    position: relative;
    display: flex;
    align-items: center;
  }

  .input {
    width: 100%;
    padding: 0.75rem 1rem;
    font-size: 1rem;
    color: var(--text-input, #1f2937);
    transition: all 0.2s ease;
  }

  /* Outlined Variant */
  .input.outlined {
    border: 1px solid var(--input-border, #d1d5db);
    border-radius: 0.5rem;
    background-color: var(--input-bg, #ffffff);
  }

  .input.outlined:focus {
    outline: none;
    border-color: var(--focus-ring-color, #60a5fa);
    box-shadow: 0 0 0 3px rgba(96, 165, 250, 0.3);
  }

  /* Filled Variant */
  .input.filled {
    border: none;
    border-radius: 0.375rem;
    background-color: var(--input-bg, #e5e7eb);
  }

  .input.filled:focus {
    outline: none;
    background-color: var(--input-bg-focus, #d1d5db);
    box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.4);
  }

  /* Minimal Variant */
  .input.minimal {
    border: none;
    border-bottom: 2px solid var(--input-border, #d1d5db);
    background-color: transparent;
    padding: 0.5rem 0.75rem;
  }

  .input.minimal:focus {
    outline: none;
    border-color: var(--focus-ring-color, #60a5fa);
  }

  .input.with-icon {
    padding-left: 2.5rem;
  }

  .input.textarea {
    resize: vertical;
    min-height: 80px;
  }

  .input.textarea.minimal {
    padding: 0.75rem;
  }

  .input.error {
    border-color: var(--error-color, #ef4444);
  }

  .input.outlined.error:focus {
    box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.3);
  }

  .input.filled.error:focus {
    box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.3);
  }

  .input[disabled] {
    cursor: not-allowed;
    opacity: 0.7;
  }

  .input.outlined[disabled] {
    background-color: var(--input-disabled-bg, #e5e7eb);
  }

  .input.filled[disabled] {
    background-color: var(--input-disabled-bg, #d1d5db);
  }

  .input.minimal[disabled] {
    border-color: var(--input-disabled-bg, #e0e0e0);
  }

  .error-message {
    font-size: 0.75rem;
    color: var(--error-color, #ef4444);
  }

  @media (max-width: 600px) {
    .input {
      padding: 0.5rem 0.75rem;
    }

    .input.textarea {
      min-height: 60px;
    }

    .input.with-icon {
      padding-left: 2rem;
    }

    :global(.icon-wrapper) {
      left: 0.5rem;
    }
  }
</style>