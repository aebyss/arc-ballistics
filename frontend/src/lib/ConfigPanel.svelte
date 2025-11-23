<script>
  export let weapons = {};
  export let shields = {};
  export let weaponKey = ""; // Bound to parent
  export let shieldKey = ""; // Bound to parent
  export let isLoading = true;
</script>

<div class="card inputs-card">
  <h2>Configuration</h2>
  
  <div class="control-group">
    <label>Weapon System</label>
    <select bind:value={weaponKey} disabled={isLoading}>
      {#each Object.keys(weapons).sort() as key}
        <option value={key}>{weapons[key].name}</option>
      {/each}
    </select>
  </div>

  <div class="control-group">
    <label>Target Shield</label>
    <select bind:value={shieldKey}>
      {#each Object.keys(shields) as s}
        <option value={s}>{shields[s].name}</option>
      {/each}
    </select>
  </div>
</div>

<style>
  .card { background: #1a202c; border: 1px solid #2d3748; border-radius: 12px; padding: 1.5rem; }
  h2 { font-size: 0.9rem; text-transform: uppercase; color: #718096; border-bottom: 1px solid #2d3748; padding-bottom: 0.5rem; margin-bottom: 1.5rem; }
  .control-group { margin-bottom: 1rem; }
  label { display: block; font-size: 0.85rem; font-weight: 600; margin-bottom: 0.5rem; color: #cbd5e0; }
  select { width: 100%; padding: 0.75rem; background: #2d3748; border: 1px solid #4a5568; color: white; border-radius: 6px; cursor: pointer; }
</style>