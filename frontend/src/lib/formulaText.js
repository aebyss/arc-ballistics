export const formulaInfo = `
  <b>Model Assumptions</b><br><br>

  <b>D</b> — Damage per shot<br>
  <b>f</b> — Shield damage reduction (0–1)<br>
  <b>H</b> — Health points<br>
  <b>S</b> — Shield HP<br>
  <b>n = ⌈S / D⌉</b> — Shots until shield breaks<br><br>

  <b>Effective shielded damage:</b><br>
  D<sub>shield</sub> = (1 − f)D<br><br>

  While the shield is active, HP is reduced by (1−f)D per shot.
`;
