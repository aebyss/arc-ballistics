# ARC Ballistics  
Damage and TTK Simulator for ARC Raiders

ARC Ballistics is a lightweight web application that simulates weapon performance in ARC Raiders.  
It provides calculations for damage, shots-to-kill, burst DPS, cyclic DPS, and armor mitigation.  
The frontend is implemented in Svelte, and the backend is written in Go using Gin.

---

## Features

- Live weapon configuration  
- Shield and damage-reduction simulation  
- Shots-to-Kill calculation  
- Burst DPS and Cyclic DPS models  
- Mathematical formulas displayed in the UI  
- Full offline cache for weapon data  
- Automatic deployment to Raspberry Pi via GitHub Actions  
- Static Svelte frontend served directly by the Go server  

---

## Repository Structure

arc-ballistics/
│
├── frontend/ # Svelte application
│ ├── src/
│ ├── dist/ # Production build
│ └── vite.config.mjs
│
├── internal/
│ ├── api/ # HTTP handlers
│ ├── models/ # Weapon, target and request models
│ └── simulation/ # Damage and DPS calculations
│
├── weapons_cache.json # Local weapon data synchronized at startup
├── main.go # Server entrypoint
├── go.mod / go.sum
└── workflows/
└── deploy.yml # GitHub Action to deploy to Raspberry Pi

yaml
Copy code

---

## Build and Run

### 1. Frontend

cd frontend
npm install
npm run build

bash
Copy code

This produces a `frontend/dist/` directory that the Go backend serves.

### 2. Backend

go mod tidy
go build -o arcballistics main.go
./arcballistics

nginx
Copy code

The server listens on:

http://localhost:1337

yaml
Copy code

---

## API Endpoints

### `GET /api/config/weapons`
Returns all weapon definitions.

### `POST /api/simulate/stk`
Computes all statistics:

Request body:
```json
{
  "weapon": { ... },
  "target": {
    "health": 100,
    "shield": 40,
    "dr": 0.4
  },
  "distance": 0
}
Response:

json
Copy code
{
  "shots_to_kill": 3,
  "damage": 40.0,
  "burst_dps": 67,
  "cyclic_dps": 40
}
Mathematical Formulas
This section explains all calculations used in the simulator.

1. Damage Per Shot
Damage is reduced by the target’s Damage Reduction (DR):

𝐷
shot
=
𝐷
base
×
(
1
−
𝐷
𝑅
)
D 
shot
​
 =D 
base
​
 ×(1−DR)
Where:

D_base is the weapon’s base damage

DR is between 0 and 1

2. Shots to Kill (STK)
𝑛
=
⌈
𝐻
+
𝑆
𝐷
shot
⌉
n=⌈ 
D 
shot
​
 
H+S
​
 ⌉
Where:

H is health

S is shield value

D_shot is damage after mitigation

ceil ensures whole shots

3. Burst DPS (no reload)
Burst DPS assumes firing rate with no reloads:

Burst DPS
=
𝐷
shot
×
𝑅
𝑃
𝑀
60
Burst DPS=D 
shot
​
 × 
60
RPM
​
 
Where:

RPM is rounds per minute

4. Cyclic DPS (including reload)
Cyclic DPS includes magazine size and reload time:

Cyclic DPS
=
𝑀
⋅
𝐷
shot
𝑀
𝑅
𝑃
𝑀
/
60
+
𝑇
load
Cyclic DPS= 
RPM/60
M
​
 +T 
load
​
 
M⋅D 
shot
​
 
​
 
Where:

M is magazine size

T_load is reload time

5. Armor and Shield Mitigation
𝐷
𝑅
=
target.damage_reduction
DR=target.damage_reduction
Displayed in UI as:

𝐷
𝑅
%
=
𝐷
𝑅
×
100
DR 
%
​
 =DR×100
Deployment (Raspberry Pi)
A self-hosted GitHub Actions runner is used for automated deployment.

When pushing to master, the following occurs:

Svelte frontend is rebuilt

Go binary is built on the Pi

Systemd service is stopped

New binary is copied to the deployment directory

Service is restarted

Workflow file:
workflows/deploy.yml

Systemd Service
Place in /etc/systemd/system/arcballistics.service:

ini
Copy code
[Unit]
Description=Arc Ballistics
After=network.target

[Service]
User=edo
WorkingDirectory=/home/edo/arc-ballistics
ExecStart=/home/edo/arc-ballistics/arcballistics
Restart=always

[Install]
WantedBy=multi-user.target
Enable and start:

pgsql
Copy code
sudo systemctl daemon-reload
sudo systemctl enable arcballistics
sudo systemctl start arcballistics
License
MIT License.

Credits
Created by EF
Weapon data sourced from MetaForge and cached locally at runtime.