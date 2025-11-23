<script>
  import { slide } from 'svelte/transition';

  export let weapon = null;
  export let showLegend = false; // Bound to parent

  function getBarWidth(value, max) {
    return Math.min(100, Math.max(0, (value / max) * 100)) + "%";
  }
</script>

{#if weapon}
<div class="weapon-card-replica">
  <div class="wr-header">
    <img src={weapon.icon_url} alt={weapon.name} />
    <div class="wr-curve"></div>
    <button class="help-btn" on:click={() => showLegend = !showLegend}>
      {showLegend ? '✕' : '?'}
    </button>
  </div>

  <div class="wr-body">
    {#if !showLegend}
      <!-- STATS VIEW -->
      <div class="content-wrapper" transition:slide|local>
        <div class="wr-tags">
          <span class="wr-tag type">Weapon</span>
          <span class="wr-tag rarity {weapon.rarity.toLowerCase()}">{weapon.rarity}</span>
        </div>
        <h3 class="wr-title">{weapon.name.toUpperCase()}</h3>
        <p class="wr-desc">{weapon.description}</p>
        {#if weapon.flavor_text}
        <div class="wr-infobox"><div class="wr-info-icon">ⓘ</div><p>{weapon.flavor_text}</p></div>
        {/if}
        <div class="wr-text-stats">
          <div class="wr-row"><span>Ammo Type</span> <strong>{weapon.ammo_type}</strong></div>
          <div class="wr-row"><span>Magazine Size</span> <strong>{weapon.mag_size}</strong></div>
          <div class="wr-row"><span>Firing Mode</span> <strong>Single</strong></div>
          <div class="wr-row"><span>Range</span> <strong>{weapon.range_far}m</strong></div>
        </div>
        <hr class="wr-divider">
        <div class="wr-bars-grid">
          <div class="wr-col">
            <div class="wr-bar-item"><div class="wr-label"><span>Damage</span> <span>{weapon.damage_close}</span></div><div class="wr-track"><div class="wr-fill" style="width: {getBarWidth(weapon.damage_close, 80)}"></div></div></div>
            <div class="wr-bar-item"><div class="wr-label"><span>Range</span> <span>{weapon.range_far}</span></div><div class="wr-track"><div class="wr-fill" style="width: {getBarWidth(weapon.range_far, 100)}"></div></div></div>
            <div class="wr-bar-item"><div class="wr-label"><span>Agility</span> <span>{weapon.agility}</span></div><div class="wr-track"><div class="wr-fill" style="width: {getBarWidth(weapon.agility, 100)}"></div></div></div>
          </div>
          <div class="wr-col">
            <div class="wr-bar-item"><div class="wr-label"><span>Fire Rate</span> <span>{weapon.display_fire_rate}</span></div><div class="wr-track"><div class="wr-fill" style="width: {getBarWidth(weapon.display_fire_rate, 40)}"></div></div></div>
            <div class="wr-bar-item"><div class="wr-label"><span>Stability</span> <span>{weapon.stability}</span></div><div class="wr-track"><div class="wr-fill" style="width: {getBarWidth(weapon.stability, 100)}"></div></div></div>
            <div class="wr-bar-item"><div class="wr-label"><span>Stealth</span> <span>{weapon.stealth}</span></div><div class="wr-track"><div class="wr-fill" style="width: {getBarWidth(weapon.stealth, 100)}"></div></div></div>
          </div>
        </div>
        <div class="wr-footer">
          <div class="wr-f-item"><svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C9 2 8 4 8 4H4V22H20V4H16C16 4 15 2 12 2ZM10 4H14C14 4 13.5 3 12 3C10.5 3 10 4 10 4ZM6 6H18V20H6V6Z"/></svg><strong>{weapon.weight} KG</strong></div>
          <div class="wr-f-item"><svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2ZM12 20C7.59 20 4 16.41 4 12C4 7.59 7.59 4 12 4C16.41 4 20 7.59 20 12C20 16.41 16.41 20 12 20ZM11 14H13V14.5C13 14.78 12.78 15 12.5 15H11.5C11.22 15 11 15.22 11 15.5V17H13V18H11V19H10V18H9.5C8.67 18 8 17.33 8 16.5V15.5C8 14.67 8.67 14 9.5 14H10V13.5C10 13.22 10.22 13 10.5 13H11.5C11.78 13 12 12.78 12 12.5V11H10V10H12V9H13V10H13.5C14.33 10 15 10.67 15 11.5V12.5C15 13.33 14.33 14 13.5 14H13V14Z"/></svg><strong>{weapon.value}</strong></div>
        </div>
      </div>
    {:else}
      <!-- LEGEND VIEW -->
      <div class="legend-view" transition:slide|local>
        <h3 class="legend-title">STATISTICS LEGEND</h3>
        <div class="legend-grid">
          <div class="l-item"><h4>Stability</h4><p>Reduces vertical recoil and improves shot consistency.</p></div>
          <div class="l-item"><h4>Agility</h4><p>Improves ADS speed and weapon swap time.</p></div>
          <div class="l-item"><h4>Stealth</h4><p>Reduces noise radius and visual detection signature.</p></div>
          <div class="l-item"><h4>Range</h4><p>Distance before damage falloff begins.</p></div>
          <div class="l-item"><h4>Weight</h4><p>Heavier items drain Stamina faster while sprinting.</p></div>
          <div class="l-item"><h4>Fire Rate</h4><p>A score (0-100) representing rounds per minute.</p></div>
        </div>
      </div>
    {/if}
  </div>
</div>
{/if}

<style>
  .weapon-card-replica { background-color: #ecece7; border-radius: 12px; overflow: hidden; color: #222; font-family: 'Inter', sans-serif; box-shadow: 0 10px 25px rgba(0,0,0,0.5); border: 1px solid #555; display: flex; flex-direction: column; transition: height 0.3s ease; }
  .wr-header { background-color: #2f4538; height: 150px; position: relative; display: flex; justify-content: center; align-items: center; flex-shrink: 0; }
  .wr-header img { max-height: 130px; filter: drop-shadow(0 8px 8px rgba(0,0,0,0.5)); transform: rotate(-4deg) translateY(10px); z-index: 2; }
  .wr-curve { position: absolute; bottom: 0; left: 0; width: 40px; height: 40px; background: #ecece7; clip-path: polygon(0 100%, 100% 100%, 0 0); z-index: 1; }
  .wr-body { padding: 1.2rem; padding-top: 0.5rem; flex-grow: 1; }
  .wr-tags { display: flex; gap: 5px; margin-bottom: 1rem; }
  .wr-tag { font-size: 0.7rem; padding: 3px 8px; text-transform: uppercase; font-weight: 700; color: #1a202c; }
  .wr-tag.type { background: #68d391; }
  .wr-tag.rarity.common { background: #a0aec0; }
  .wr-tag.rarity.uncommon { background: #68d391; }
  .wr-tag.rarity.rare { background: #63b3ed; }
  .wr-tag.rarity.epic { background: #9f7aea; }
  .wr-tag.rarity.legendary { background: #ecc94b; }
  .wr-title { font-size: 1.5rem; font-weight: 900; margin: 0 0 0.5rem 0; letter-spacing: -0.5px; }
  .wr-desc { font-size: 0.85rem; color: #555; line-height: 1.4; margin: 0 0 1rem 0; }
  .wr-infobox { background: #d8dad5; padding: 10px; border-radius: 6px; display: flex; gap: 10px; align-items: center; margin-bottom: 1.2rem; }
  .wr-info-icon { font-weight: bold; font-size: 1.2rem; opacity: 0.7; }
  .wr-infobox p { margin: 0; font-size: 0.8rem; font-weight: 700; }
  .wr-text-stats { display: flex; flex-direction: column; gap: 8px; }
  .wr-row { display: flex; justify-content: space-between; border-bottom: 1px solid #ccc; padding-bottom: 4px; font-size: 0.9rem; }
  .wr-row span { color: #444; }
  .wr-row strong { color: #000; }
  .wr-divider { border: 0; border-top: 2px solid #ccc; margin: 1.5rem 0; }
  .wr-bars-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 1.5rem; }
  .wr-col { display: flex; flex-direction: column; gap: 0.8rem; }
  .wr-bar-item { display: flex; flex-direction: column; gap: 2px; }
  .wr-label { display: flex; justify-content: space-between; font-size: 0.8rem; font-weight: 600; color: #222; }
  .wr-track { height: 8px; background: #ccc; border-radius: 4px; width: 100%; overflow: hidden; }
  .wr-fill { height: 100%; background: #000; border-radius: 4px; }
  .wr-footer { margin-top: 1.5rem; border-top: 2px solid #ccc; display: flex; padding-top: 1rem; }
  .wr-f-item { flex: 1; display: flex; flex-direction: column; align-items: center; gap: 5px; font-size: 0.9rem; font-weight: 900; color: #000; }
  .wr-f-item:first-child { border-right: 2px solid #ccc; }
  .help-btn { position: absolute; top: 10px; right: 10px; width: 24px; height: 24px; background: rgba(255,255,255,0.2); border: 1px solid rgba(255,255,255,0.4); border-radius: 50%; color: white; font-weight: bold; cursor: pointer; display: flex; justify-content: center; align-items: center; font-size: 0.8rem; z-index: 10; }
  .help-btn:hover { background: white; color: #2f4538; }
  .legend-view { padding: 0.5rem; }
  .legend-title { text-align: center; font-size: 1rem; font-weight: 900; margin-bottom: 1.5rem; color: #2f4538; text-transform: uppercase; border-bottom: 2px solid #ccc; padding-bottom: 0.5rem; }
  .legend-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 1.2rem; }
  .l-item h4 { margin: 0 0 4px 0; font-size: 0.8rem; text-transform: uppercase; color: #2f4538; font-weight: 800; }
  .l-item p { margin: 0; font-size: 0.75rem; color: #555; line-height: 1.3; }
</style>