<script lang="ts">
import { Alert, Button, Text } from "$lib/components/atoms";
import { FormField } from "$lib/components/molecules";

let name = $state("");
let email = $state("");
let noOfGuests = $state(1);
let message = $state("");
let toast = $state<{ message: string; type: "success" | "error" } | null>(null);
let toastTimer: ReturnType<typeof setTimeout>;
let isSubmitting = $state(false);

function showToast(message: string, type: "success" | "error") {
	clearTimeout(toastTimer);
	toast = { message, type };
	toastTimer = setTimeout(() => (toast = null), 5000);
}

let buttonLabel = $derived(isSubmitting ? "Sending..." : "Send Message");

async function handleSubmit(e: Event) {
	e.preventDefault();
	isSubmitting = true;

	if (!name || !email || !message) {
		showToast("Please fill in all required fields.", "error");
		isSubmitting = false;
		return;
	}

	const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
	if (!emailRegex.test(email)) {
		showToast("Please enter a valid email address.", "error");
		isSubmitting = false;
		return;
	}

	try {
		const res = await fetch("https://api.web3forms.com/submit", {
			method: "POST",
			headers: { "Content-Type": "application/json", Accept: "application/json" },
			body: JSON.stringify({
				access_key: "01753924-01ed-44cd-8810-41b6c47e539e",
				name,
				email,
				guests: noOfGuests,
				message,
			}),
		});
		const data = await res.json();
		if (data.success) {
			showToast("Thank you for your message! We'll get back to you soon.", "success");
			name = "";
			email = "";
			noOfGuests = 1;
			message = "";
		} else {
			showToast("Something went wrong. Please email us at info@footstepstofreedom.co.za", "error");
		}
	} catch (error) {
		showToast("An error occurred. Please email us at info@footstepstofreedom.co.za", "error");
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


  {#if toast}
    <div class="toast toast--{toast.type}" role="alert">
      {toast.message}
    </div>
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

  .toast {
    padding: 0.875rem 1.25rem;
    border-radius: 0.5rem;
    font-size: 0.95rem;
    font-weight: 500;
    animation: slide-in 0.2s ease;
  }

  .toast--success {
    background: #dcfce7;
    color: #166534;
    border: 1px solid #bbf7d0;
  }

  .toast--error {
    background: #fee2e2;
    color: #991b1b;
    border: 1px solid #fecaca;
  }

  @keyframes slide-in {
    from { opacity: 0; transform: translateY(-6px); }
    to   { opacity: 1; transform: translateY(0); }
  }
</style>