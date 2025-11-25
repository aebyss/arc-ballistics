<script>
  import { slide } from 'svelte/transition';
  import MathTex from "../Math.svelte";
  import FormulaTooltip from "./FormulaTooltip.svelte";
  import { formulaInfo } from "./formulaText.js"; // optional
  export let stats = {};
  export let shield = {};
  export let compact = false; // Controlled by parent (showLegend)
</script>

<div class="card stats-display" class:compact={compact}>
  <h2>Simulation Results</h2>
  <div class="stat-grid">
    <script>
  
</script>
    <div class="stat-box big">
      <span class="label">Shots to Kill</span>
      <FormulaTooltip text={formulaInfo}>
        <span class="info-hover">ⓘ</span>
      </FormulaTooltip>
      <div class="value text-highlight">{stats.shotsToKill}</div>
      
      {#if !compact}
        <div class="sub-text" transition:slide>{(stats.damagePerShot * stats.shotsToKill).toFixed(0)} Total Dmg</div>
        <div transition:slide><MathTex latex={'n + \\lceil \\frac{H - n(1-f)D}{D} \\rceil'} /></div>
      {/if}
    </div>
    
    <div class="stat-box">
      <div class="wip-overlay">WIP</div>
       <span class="label">Burst DPS</span>
       
       <div class="value">{(stats.burst_dps || 0).toFixed(0)}</div>
       <span class="unit-label">Firing Only</span>
       {#if !compact}
          <div transition:slide><MathTex latex={'D \\times \\frac{RPM}{60}'} /></div>
       {/if}
    </div>

    <div class="stat-box">
      <div class="wip-overlay">WIP</div>
       <span class="label">Cyclic DPS</span>
       
       <div class="value">{(stats.cyclic_dps || 0).toFixed(0)}</div>
       <span class="unit-label">With Reload</span>
       {#if !compact}
          <div transition:slide><MathTex latex={'\\frac{D \\cdot M}{T_{dump} + T_{load}}'} /></div>
       {/if}
    </div>
    
    <div class="stat-box">
      <span class="label">Dmg / Shot</span>
      <div class="value">{stats.damagePerShot.toFixed(1)}</div>
      {#if !compact}
          <div transition:slide><MathTex latex={'D_{base} \\times (1 - DR)'} /></div>
      {/if}
    </div>
    
    <div class="stat-box">
      <span class="label">Target DR</span>
      <div class="value">{(shield.dr * 100).toFixed(0)}%</div>
      {#if !compact}
          <div transition:slide><MathTex latex={'\\text{Shield Mitigation}'} /></div>
      {/if}
    </div>
  </div>
</div>

<style>
  .card { background: #1a202c; border: 1px solid #2d3748; border-radius: 12px; padding: 1.5rem; }
  h2 { font-size: 0.9rem; text-transform: uppercase; color: #718096; border-bottom: 1px solid #2d3748; padding-bottom: 0.5rem; margin-bottom: 1.5rem; }
  
  .stats-display { display: flex; flex-direction: column; justify-content: flex-start; transition: all 0.3s ease; height: auto; }
  .stat-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 1.5rem; transition: gap 0.3s ease; }
  .stat-box { background: #232936; padding: 2rem 1.5rem; border-radius: 12px; text-align: center; display: flex; flex-direction: column; justify-content: center; box-shadow: 0 4px 6px rgba(0,0,0,0.2); transition: padding 0.3s ease; }
  .stat-box.big { grid-column: span 2; background: rgba(237, 137, 54, 0.08); border: 1px solid rgba(237, 137, 54, 0.3); padding: 2.5rem; transition: padding 0.3s ease; }
  .stat-box .label { font-size: 0.95rem; text-transform: uppercase; color: #a0aec0; letter-spacing: 1px; font-weight: 600; transition: font-size 0.3s; }
  .stat-box .value { font-size: 2.8rem; font-weight: 800; color: white; margin: 0.5rem 0; line-height: 1; transition: font-size 0.3s; }
  .text-highlight { color: #ed8936; font-size: 6rem !important; text-shadow: 0 0 20px rgba(237, 137, 54, 0.2); transition: font-size 0.3s; }
  .sub-text { font-size: 1rem; color: #cbd5e0; margin-bottom: 10px; font-weight: 500; }
  .unit-label { font-size: 0.75rem; color: #718096; text-transform: uppercase; margin-top: 4px; font-weight: 700; }
  .info-hover {
  position: relative; top: 1px; right: 10px; width: 24px; height: 24px; background: rgba(255,255,255,0.2); border: 1px solid rgba(255,255,255,0.4); border-radius: 50%; color: white; font-weight: bold; cursor: pointer; display: flex; justify-content: center; align-items: center; font-size: 0.8rem; z-index: 10; }
.info-hover:hover {
  opacity: 0.5;
}

.wip-overlay {
  position: relative;
  inset: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 2.2rem;
  font-weight: 700;
  color: rgba(224, 121, 11, 0.08);
  pointer-events: none; 
  text-transform: uppercase;
}
  /* COMPACT MODE */
  .stats-display.compact { padding: 1rem; }
  .stats-display.compact .stat-grid { gap: 0.8rem; }
  .stats-display.compact .stat-box { padding: 1rem 0.5rem; }
  .stats-display.compact .stat-box.big { padding: 1.5rem; }
  .stats-display.compact .text-highlight { font-size: 3.5rem !important; }
  .stats-display.compact .value { font-size: 1.8rem; }
  .stats-display.compact .label { font-size: 0.7rem; }
</style>