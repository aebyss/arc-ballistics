# ARC Ballistics

ARC Ballistics is a lightweight damage simulator for the game *ARC Raiders*. It calculates damage output considering damage reduction (DR), shields, and reload times to provide accurate burst and cyclic DPS metrics.

The project features a single-binary architecture with a **Go (Gin)** backend and a static **Svelte** frontend.

## Features

- **Detailed Calculations:** Accurate Shots-to-Kill (STK)
- **DPS Metrics:** Calculates both Burst DPS (firing only) and Cyclic DPS (including reload time).
- **BURST DPS** Is not currently correct since RPM on weapons is only for UI purposes and is not representing correct stat.
- **Defense Simulation:** Accounts for target Shields and Damage Reduction (DR).
- **Local Cache:** Implements a local weapon cache for fast retrieval.
- **Single Binary:** The frontend is embedded or served statically, allowing the app to run as a single executable.