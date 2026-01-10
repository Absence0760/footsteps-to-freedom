<script lang="ts">
import { goto } from "$app/navigation";

export type IconName = "plus" | "home" | "about" | "contact" | "tour" | "star";

interface IconDefinition {
	viewBox: string;
	paths: string[];
}

const icons: Record<IconName, IconDefinition> = {
	plus: {
		viewBox: "0 0 24 24",
		paths: [
			'<line x1="12" y1="5" x2="12" y2="19" />',
			'<line x1="5" y1="12" x2="19" y2="12" />',
		],
	},
	home: {
		viewBox: "0 0 24 24",
		paths: [
			'<path d="M3 12L5 10M5 10L12 3L19 10M5 10V20A1 1 0 0 0 6 21H9M19 10L21 12M19 10V20A1 1 0 0 1 18 21H15" />',
		],
	},
	about: {
		viewBox: "0 0 24 24",
		paths: [
			'<circle cx="12" cy="12" r="10" />',
			'<path d="M12 16V12M12 8H12.01" />',
		],
	},
	contact: {
		viewBox: "0 0 24 24",
		paths: [
			'<path d="M4 4H20C21.1 4 22 4.9 22 6V18C22 19.1 21.1 20 20 20H4C2.9 20 2 19.1 2 18V6C2 4.9 2.9 4 4 4Z" />',
			'<polyline points="22,6 12,13 2,6" />',
		],
	},
	tour: {
		viewBox: "0 0 24 24",
		paths: [
			'<path d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2Z" />',
			'<path d="M12 6V12L16 14" />',
		],
	},
	star: {
		viewBox: "0 0 24 24",
		paths: [
			'<polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2" />',
		],
	},
};

interface Props {
	icon: IconName;
	size?: string;
	color?: string;
	ariaLabel?: string;
	decorative?: boolean;
	id?: string;
	class?: string;
	onClick?: () => void;
}

const {
	icon,
	size = "1rem",
	color = "var(--text-body, #4b5563)",
	ariaLabel = `${icon} icon`,
	decorative = false,
	id = "",
	class: className = "",
	onClick = () => {},
}: Props = $props();

const handleSvgClick = () => {
	if (id) {
		goto(id);
	}
	onClick();
};

const selectedIcon = icons[icon];
const isFillIcon = icon === "star";
</script>

<div class="icon-wrapper {className}">
  <svg
    width={size}
    height={size}
    viewBox={selectedIcon.viewBox}
    fill={isFillIcon ? color : "none"}
    stroke={isFillIcon ? "none" : color}
    stroke-width={isFillIcon ? "0" : "2"}
    stroke-linecap="round"
    stroke-linejoin="round"
    aria-hidden={decorative || ariaLabel ? undefined : true}
    aria-label={ariaLabel || undefined}
    onclick={handleSvgClick}
  >
    {#each selectedIcon.paths as path}
      {@html path}
    {/each}
  </svg>
</div>

<style>
  .icon-wrapper {
    position: absolute;
    bottom: 1rem;
    left: 1rem;
    transition: transform 0.3s ease;
    cursor: pointer;
  }

  .icon-wrapper:hover {
    transform: scale(1.1);
  }

  @media (max-width: 600px) {
    .icon-wrapper {
      bottom: 0.5rem;
      left: 0.5rem;
    }

    .icon-wrapper svg {
      width: 28px;
      height: 28px;
    }
  }
</style>