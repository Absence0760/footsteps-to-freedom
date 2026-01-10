<script lang="ts">
import { Icon } from "$lib/components/atoms";
import type { IconName } from "$lib/components/atoms";
import type { Snippet } from "svelte";
import type { MouseEventHandler } from "svelte/elements";

interface ButtonProps {
	variant?: "default" | "primary" | "icon-light" | "icon-dark";
	label?: string;
	disabled?: boolean;
	icon?: IconName;
	type?: "button" | "submit" | "reset";
	onclick?: MouseEventHandler<HTMLButtonElement>;
	children?: Snippet;
}

const {
	variant = "default",
	label = "",
	disabled = false,
	icon,
	type = "button",
	onclick,
	children,
}: ButtonProps = $props();
</script>

<button
  {type}
  class="button {variant}"
  {disabled}
  onclick={onclick}
  aria-label={label || 'Button'}
>
  {#if icon}
    <Icon icon={icon}/>
  {/if}
  {#if children}
    {@render children()}
  {:else if label}
    <span class="text">{label}</span>
  {/if}
</button>

<style>
  .button {
    position: relative;
    border-radius: 5px;
    width: 136px;
    height: 40px;
    border: none;
    cursor: pointer;
    font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    transition: all 0.3s ease;
  }

  .text {
    font-size: 12px;
    font-weight: 500;
    line-height: 1.5;
    text-align: center;
  }


  .default {
    background: var(--button-primary-bg, #FF6B35);
    color: var(--button-primary-text, #FFFFFF);
  }

  .default:hover:not(.disabled) {
    background: var(--button-primary-hover-bg, #E85A28);
    transform: translateY(-2px);
    box-shadow: var(--shadow-md);
  }

  .primary {
    background: var(--button-primary-bg, #FF6B35);
    color: var(--button-primary-text, #FFFFFF);
  }

  .primary:hover:not(.disabled) {
    background: var(--button-primary-hover-bg, #E85A28);
    transform: translateY(-2px);
    box-shadow: var(--shadow-md);
  }

  .icon-light {
    width: 40px; /* Match button height for square shape */
    height: 40px;
    background: var(--input-bg, #e3e3e3);
  }

  .icon-dark {
    width: 40px;
    height: 40px;
    background: var(--input-bg-dark, #1f2937);
  }

  .disabled {
    background: var(--button-disabled-bg, #9ca3af);
    color: var(--button-disabled-text, #ffffff);
    cursor: not-allowed;
    opacity: 0.6;
  }

  .button:focus {
    outline: none;
    box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.3);
  }

  @media (max-width: 600px) {
    .button {
      width: 120px;
      height: 36px;
      border-radius: 2px;
      padding: 0 clamp(0.5rem, 2vw, 0.75rem);
    }

    .text {
      font-size: clamp(0.625rem, 2vw, 0.75rem);
    }
  }
</style>