<script>
  export let text;
</script>

<span class="tooltip-wrapper">
  <slot></slot>
  <div class="tooltip-box">
    {@html text}
  </div>
</span>

<style>
.tooltip-wrapper {
  position: relative;
  display: inline-block;
  cursor: help;
}

.tooltip-box {
  opacity: 0;
  visibility: hidden;
  position: absolute;
  left: 50%;
  top: 120%;
  transform: translateX(-50%);
  min-width: 260px;

  background: #1a1d23;
  color: #e2e8f0;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  font-size: 0.75rem;
  line-height: 1.25rem;
  text-align: left;
  box-shadow: 0 4px 14px rgba(0,0,0,0.4);

  border: 1px solid rgba(255,255,255,0.08);

  transition: opacity 0.15s ease;
  z-index: 20;
}

.tooltip-wrapper:hover .tooltip-box {
  opacity: 1;
  visibility: visible;
}
</style>
