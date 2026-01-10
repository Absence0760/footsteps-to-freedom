<script lang="ts">
interface Props {
	orientation?: "horizontal" | "vertical";
	color?: string;
	thickness?: string;
}

const {
	orientation = "horizontal",
	color = "var(--input-border, #d1d5db)",
	thickness = "1px",
}: Props = $props();
</script>

<hr
  class="divider"
  style:border-color={color}
  style:width={orientation === 'vertical' ? thickness : '100%'}
  style:height={orientation === 'horizontal' ? thickness : '100%'}
/>

<style>
  .divider {
    border: none;
    background-color: currentColor;
    margin: 1rem 0;
  }
</style>