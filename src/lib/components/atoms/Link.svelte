<script lang="ts">
import { goto } from "$app/navigation";
import { base } from "$app/paths";
import { Icon } from "$lib/components/atoms";
import type { Snippet } from "svelte";

type Variant = "header" | "footer" | "inline-text" | "button";
type IconName = "plus" | "home" | "about" | "contact" | "tour";

interface Props {
	variant?: Variant;
	href?: string;
	target?: "_blank" | "_self" | "_parent" | "_top" | string;
	rel?: string;
	icon?: IconName;
	to?: string;
	style?: string;
	children?: Snippet;
}

const {
	variant = "inline-text",
	to,
	href: externalHref,
	target = "_self",
	rel = target === "_blank" ? "noopener noreferrer" : "",
	icon,
	children,
	...restProps
}: Props = $props();

// Prepend base path to internal links
const href = externalHref || (to ? `${base}${to}` : "#");

const handleClick = (e: MouseEvent) => {
	if (to && target === "_self") {
		e.preventDefault();
		goto(`${base}${to}`);
	}
};
</script>

<a
  {href}
  {target}
  {rel}
  class="link {variant}"
  onclick={handleClick}
  {...restProps}
>
  {#if icon}
    <Icon
      {icon}
      size={variant === 'header' || variant === 'button' ? '1.5rem' : '1rem'}
      color="currentColor"
      ariaLabel={`${children ? 'icon for link' : 'link icon'}`}
    />
  {/if}
  {#if children}
    {@render children()}
  {/if}
</a>

<style>
  .link {
    text-decoration: none;
    cursor: pointer;
    transition: color 0.2s ease, border 0.2s ease, background-color 0.2s ease;
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
  }

  /* Header Variant */
  .link.header {
    font-size: 1.125rem;
    font-weight: 600;
    color: var(--text-header, #1f2937);
    padding: 0.5rem 1rem;
    border-radius: 0.375rem;
  }

  .link.header:hover {
    color: var(--button-primary-bg, #4f46e5);
    background-color: var(--button-primary-bg-light, #e0e7ff);
  }

  .link.header:focus {
    outline: none;
    box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.3);
  }

  /* Footer Variant */
  .link.footer {
    font-size: 0.875rem;
    color: var(--text-footer, #6b7280);
  }

  .link.footer:hover {
    color: var(--button-primary-bg, #4f46e5);
    border-bottom: 1px solid var(--button-primary-bg, #4f46e5);
  }

  .link.footer:focus {
    outline: none;
    box-shadow: 0 0 0 2px rgba(79, 70, 229, 0.3);
  }

  /* Inline Text Variant */
  .link.inline-text {
    font-size: 1rem;
    color: var(--button-primary-bg, #4f46e5);
    text-decoration: underline;
  }

  .link.inline-text:hover {
    color: var(--button-primary-bg-dark, #6366f1);
    text-decoration: none;
  }

  .link.inline-text:focus {
    outline: none;
    box-shadow: 0 0 0 2px rgba(79, 70, 229, 0.3);
  }

  /* Button Variant */
  .link.button {
    font-size: 1rem;
    font-weight: 500;
    color: var(--button-text, #ffffff);
    background-color: var(--button-primary-bg, #4f46e5);
    padding: 0.75rem 1.5rem;
    border-radius: 0.5rem;
    text-align: center;
  }

  .link.button:hover {
    background-color: var(--button-primary-bg-dark, #6366f1);
  }

  .link.button:focus {
    outline: none;
    box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.3);
  }

  @media (max-width: 600px) {
    .link.header {
      font-size: 1rem;
      padding: 0.25rem 0.75rem;
    }

    .link.footer {
      font-size: 0.75rem;
    }

    .link.button {
      padding: 0.5rem 1rem;
    }
  }
</style>