<script>
  import katex from "katex";
  import "katex/dist/katex.min.css"; // Import CSS so fonts work

  export let latex = "";
  
  // Render function
  $: html = katex.renderToString(latex, {
    throwOnError: false,
    displayMode: false, // Inline math
    colorIsTextColor: true
  });
</script>

<span class="math-container">
  {@html html}
</span>

<style>
  .math-container {
    font-size: 1rem;
    color: #718096; /* Muted grey text */
    opacity: 0.8;
    margin-top: 4px;
    display: block;
  }
</style>