<script lang="ts">
import { Link, Text } from "$lib/components/atoms";

const {
	title = "Ready to Explore History?",
	description = "Join us on an unforgettable journey. Have questions? We're here to help!",
	buttonLabel = "Contact Us",
	buttonHref = "/contact",
	...restProps
} = $props();
</script>

<section
  class="call-to-action-section"
  {...restProps}
  aria-labelledby="cta-title"
>
  <Text variant="h2">
    {title}
  </Text>

  <Text variant="p">
    {description}
  </Text>

  <div class="cta-button">
    <Link variant="button" to={buttonHref}>
      {buttonLabel}
    </Link>
  </div>
</section>

<style>
  .call-to-action-section {
    padding: clamp(2rem, 6vw, 4rem) 1rem;
    max-width: 900px;
    margin: 2rem auto;
    text-align: center;
    background-color: var(--secondary-bg, #f9fafb);
    border-radius: 1rem;
    box-shadow: var(--shadow-md, 0 4px 6px -1px rgba(0, 0, 0, 0.1));
    transition:
      transform 0.3s ease,
      box-shadow 0.3s ease;
  }

  .call-to-action-section:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-lg, 0 10px 15px -3px rgba(0, 0, 0, 0.1));
  }

  .cta-button {
    margin-top: 1.5rem;
  }

  @media (max-width: 600px) {
    .call-to-action-section {
      padding: 2rem 1rem;
      margin: 1rem;
      border-radius: 0.5rem;
    }
  }
</style>
