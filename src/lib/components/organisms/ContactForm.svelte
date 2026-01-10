<script lang="ts">
import { Alert, Button, Text } from "$lib/components/atoms";
import { FormField } from "$lib/components/molecules";

let name = $state("");
let email = $state("");
let noOfGuests = $state(1);
let message = $state("");
let statusMessage = $state("");
let isSubmitting = $state(false);

let buttonLabel = $derived(isSubmitting ? "Sending..." : "Send Message");

async function handleSubmit(e: Event) {
	e.preventDefault();
	isSubmitting = true;
	statusMessage = "";

	if (!name || !email || !message) {
		statusMessage = "Please fill in all required fields.";
		isSubmitting = false;
		return;
	}

	// Improved email validation regex - requires valid domain with 2+ character TLD
	const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
	if (!emailRegex.test(email)) {
		statusMessage = "Please enter a valid email address.";
		isSubmitting = false;
		return;
	}

	try {
		await new Promise((resolve) => setTimeout(resolve, 1500));
		statusMessage = "Thank you for your message! We will get back to you soon.";
		name = "";
		email = "";
		noOfGuests = 1;
		message = "";
	} catch (error) {
		// NOTE: In production, send errors to a secure logging service instead of console
		// Only show generic error messages to users to avoid leaking sensitive information
		statusMessage = "An error occurred. Please try again.";
	} finally {
		isSubmitting = false;
	}
}
</script>

<form onsubmit={handleSubmit} class="contact-form">
  <FormField
    variant="outlined"
    label="Name"
    id="name"
    type="text"
    placeholder="Your Name"
    bind:value={name}
    required
  />

  <FormField
    label="Email"
    id="email"
    type="email"
    placeholder="your.email@example.com"
    bind:value={email}
    required
  />

  <FormField
    label="Number of Guests"
    id="noOfGuests"
    type="text"
    placeholder=1
    bind:value={noOfGuests}
    required
  />

  <FormField
    label="Message"
    id="message"
    type="textarea"
    rows={10}
    placeholder="Your message..."
    bind:value={message}
    required
  />

  <Button
    type="submit"
    label={buttonLabel}
    disabled={isSubmitting}
  />


  {#if statusMessage}
    <Text>{statusMessage}</Text>
  {/if}
</form>

<style>
  .contact-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
    max-width: 600px;
    margin: 0 auto;
    background-color: var(--form-bg, #ffffff);
    padding: 2rem;
    border-radius: 0.75rem;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.2s ease;
  }

  .contact-form:hover {
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
  }
</style>