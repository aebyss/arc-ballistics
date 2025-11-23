<script>
  import { onMount } from "svelte";
  
  // Import Child Components
  import ConfigPanel from "./lib/ConfigPanel.svelte";
  import WeaponCard from "./lib/WeaponCard.svelte";
  import StatsPanel from "./lib/StatsPanel.svelte";

  // --- State ---
  let weapons = {}; 
  let selectedWeaponKey = "";
  let isLoading = true;
  let selectedShieldKey = "light"; 
  
  // UI State
  let showLegend = false; 

  let stats = {
    shotsToKill: 0,
    damagePerShot: 0,
    burst_dps: 0,
    cyclic_dps: 0
  };

  const shields = {
    none:   { name: "No Shield", shield: 0,   dr: 0.00 },
    light:  { name: "Light",     shield: 40,  dr: 0.40 },
    medium: { name: "Medium",    shield: 70,  dr: 0.425 },
    heavy:  { name: "Heavy",     shield: 80,  dr: 0.525 },
    green:  { name: "Green",     shield: 50,  dr: 0.23 },
    blue:   { name: "Blue",      shield: 80,  dr: 0.45 },
    purple: { name: "Purple",    shield: 120, dr: 0.65 },
    gold:   { name: "Gold",      shield: 170, dr: 0.80 }
  };

  // --- Reactivity ---
  // 1. Selected Weapon Logic
  $: selectedWeapon = weapons[selectedWeaponKey] || { 
      name: "Loading...", rarity: "Common", icon_url: "", description: "Fetching...",
      fire_rate: 0, display_fire_rate: 0, range_close: 0, range_far: 0,
      damage_close: 0, mag_size: 0, weight: 0, value: 0
  };
  

  $: selectedShield = shields[selectedShieldKey];
  
  $: if(selectedWeapon && selectedWeapon.name !== "Loading..." && selectedShield) { 
      updateStats(); 
  }

  onMount(async () => {
    try {
      const res = await fetch("/api/config/weapons");
      weapons = await res.json();
      const keys = Object.keys(weapons);
      if (keys.length > 0) {
        keys.sort(); 
        selectedWeaponKey = keys[0];
      }
      isLoading = false;
    } catch (e) { console.error(e); }
  });

  async function updateStats() {
    if (!selectedWeaponKey) return;
    try {
      const res = await fetch("/api/simulate/stk", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          weapon: selectedWeapon,
          target: { health: 100, shield: selectedShield.shield, dr: selectedShield.dr },
          distance: 0 
        })
      });
      const data = await res.json();
      stats = {
        shotsToKill: data.shots_to_kill,
        damagePerShot: data.damage,
        burst_dps: data.burst_dps,
        cyclic_dps: data.cyclic_dps
      };
    } catch (e) { console.error(e); }
  }
</script>

<main>
  <div class="layout">
    <header>
      <h1>ARC <span class="highlight">BALLISTICS</span></h1>
      <p class="status-line">STATUS: {isLoading ? 'Loading...' : `Ready (${Object.keys(weapons).length} Items)`}</p>
      <a class="repo-link" href="https://github.com/aebyss/arc-ballistics" target="_blank">
    GitHub Repository
  </a>
    </header>

    <div class="grid-container">
      
      <div class="left-panel-wrapper">
        <ConfigPanel 
            bind:weaponKey={selectedWeaponKey} 
            bind:shieldKey={selectedShieldKey}
            {weapons}
            {shields}
            {isLoading}
        />

        <WeaponCard 
            weapon={selectedWeapon} 
            bind:showLegend={showLegend} 
        />
      </div>

      <div class="center-stack">
    <StatsPanel 
        {stats} 
        shield={selectedShield} 
        compact={showLegend} 
    />

    <div class="formula-legend">
        <h3>Formula Legend</h3>

        <p><b>D</b> — Damage per shot</p>
        <p><b>n</b> — Number of shots</p>
        <p><b>H</b> — Health</p>
        <p><b>DR</b> — Damage Reduction (0–1)</p>
        <p><b>RPM</b> — Rounds per minute</p>
        <p><b>T_load</b> — Reload Time</p>
    </div>
</div>
    </div>
  </div>
</main>

<style>
  :global(body) { background: #0f1115; color: #e2e8f0; font-family: 'Inter', sans-serif; margin: 0; }
  main { padding: 2rem; display: flex; justify-content: center; }
  .layout { width: 100%; max-width: 1200px; display: flex; flex-direction: column; gap: 1.5rem; }
  
  header h1 { text-align: center; font-size: 2.5rem; margin: 0; text-transform: uppercase; letter-spacing: -1px; }
  .highlight { color: #ed8936; }
  .status-line { text-align: center; color: #718096; font-size: 0.8rem; margin-top: 0.5rem; }

  .grid-container {
  display: grid;
  grid-template-columns: 400px 450px 300px;
  justify-content: center;
  gap: 2rem;
  align-items: center;
}
  .repo-link {
  display: block;
  text-align: center;
  margin-top: 0.5rem;
  font-size: 0.85rem;
  color: #ed8936;
  text-decoration: none;
  transition: opacity 0.2s ease;
}

.repo-link:hover {
  opacity: 0.7;
}
  .left-panel-wrapper { display: flex; flex-direction: column; gap: 1.5rem; }

  @media(max-width: 900px) { .grid-container { grid-template-columns: 1fr; } }

 .center-stack {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.formula-legend {
  background: #1a1d24;
  padding: 1.5rem;
  border-radius: 12px;
  border: 1px solid #2a2e36;
  color: #e2e8f0;
  font-size: 0.9rem;
}

.formula-legend h3 {
  margin: 0 0 1rem 0;
  font-size: 1rem;
  text-transform: uppercase;
  color: #ed8936;
}
</style>
