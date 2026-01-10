<script>
import { Link, Text } from "$lib/components/atoms";

export const year = new Date().getFullYear();
export const siteName = "Footsteps to Freedom";
export const footerLinks = [
	{ name: "Privacy Policy", href: "/privacy" },
	{ name: "Terms of Service", href: "/terms" },
	{ name: "FAQs", href: "/faqs" },
];
</script>

<footer>
  <div class="footer-content">
    <Text variant="p">
      &copy; {year} {siteName}. All rights reserved.
    </Text>

    <ul>
      {#each footerLinks as link}
        <li>
          <Link variant="footer" to={link.href}> {link.name} </Link> 
        </li>
      {/each}
    </ul>
  </div>
</footer>

<style>
  footer {
    background-color: var(--footer-bg);
    color: var(--text-footer);
    padding-top: 2rem; /* py-8 */
    padding-bottom: 2rem; /* py-8 */
    margin-top: 3rem; /* mt-12 */
    border-top-left-radius: 0.5rem; /* rounded-t-lg */
    border-top-right-radius: 0.5rem; /* rounded-t-lg */
    box-shadow: var(--shadow-inner);
    font-family: 'Inter', sans-serif;
  }

  .footer-content {
    max-width: 1200px;
    margin: 0 auto;
    padding-left: 1rem;
    padding-right: 1rem;
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  footer ul {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    list-style: none;
    padding: 0;
    margin: 0;
    column-gap: 1.5rem;
  }

  @media (min-width: 768px) {
    .footer-content {
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
    }
    footer ul {
      justify-content: flex-end;
    }
  }
</style>
