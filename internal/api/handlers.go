package api

import (
	"math"
	"net/http"

	"arc-simulator/internal/models"

	"github.com/gin-gonic/gin"
)

// --- Request Model ---
type SimulationRequest struct {
	Weapon models.Weapon `json:"weapon"`
	Target struct {
		Health float64 `json:"health"`
		Shield float64 `json:"shield"`
		DR     float64 `json:"dr"`
	} `json:"target"`
	Distance float64 `json:"distance"`
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/config/weapons", HandleGetWeapons)
		api.POST("/simulate/damage", HandleSimulateDamage)
		api.POST("/simulate/stk", HandleSimulateTTK)
	}
}

func HandleGetWeapons(c *gin.Context) {
	weapons := Store.GetAll()
	c.JSON(http.StatusOK, weapons)
}

// 1. DAMAGE SIMULATOR (Single Shot)
func HandleSimulateDamage(c *gin.Context) {
	var req SimulationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dmg := calculateBaseDamage(req.Weapon, req.Distance)

	// Apply Shield Logic for single shot display
	actualDmg := dmg
	if req.Target.Shield > 0 {
		// If hitting shield, reduce damage
		actualDmg = dmg * (1 - req.Target.DR)
	}

	c.JSON(http.StatusOK, gin.H{
		"damage": actualDmg,
		"result": gin.H{
			"health": req.Target.Health, // Simplified for single shot view
			"shield": math.Max(0, req.Target.Shield-actualDmg),
		},
	})
}

// 2. STK & DPS SIMULATOR
// Helper for Range Falloff
func calculateBaseDamage(w models.Weapon, dist float64) float64 {
	damage := w.DamageClose
	if dist > w.RangeFar {
		damage = w.DamageFar
	} else if dist > w.RangeClose {
		ratio := (dist - w.RangeClose) / (w.RangeFar - w.RangeClose)
		damage = w.DamageClose - (ratio * (w.DamageClose - w.DamageFar))
	}
	return damage
}

// 3. POST /api/simulate/stk
func HandleSimulateTTK(c *gin.Context) {
	var req SimulationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	H := req.Target.Health
	f := req.Target.DR // Shield Block Rate
	D := calculateBaseDamage(req.Weapon, req.Distance)
	ShieldHP := req.Target.Shield

	// Safety check for D=0 to prevent divide by zero
	if D <= 0 {
		D = 1
	}

	// 2. Calculate 'n' (Landed shots while shield is active)
	// Shield takes full raw damage.
	// n = Ceil(ShieldHP / Damage)
	var n float64 = 0
	if ShieldHP > 0 {
		n = math.Ceil(ShieldHP / D)
	}

	var shotsToKill float64

	// Damage dealt per shot while shield is UP
	shieldedDamage := D * (1 - f)
	if shieldedDamage <= 0 {
		shieldedDamage = 0.1
	} // Safety

	// Check Case 1: Does the target die BEFORE the shield breaks?
	shotsToKillShielded := math.Ceil(H / shieldedDamage)

	if shotsToKillShielded <= n {
		// CASE 1: Target dies while shield is still active
		// Formula: Ceil[ H / ((1-f)D) ]
		shotsToKill = shotsToKillShielded
	} else {
		// CASE 2: Shield breaks, then target dies
		// Formula: n + Ceil[ (H - n(1-f)D) / D ]

		// Health lost during the 'n' shots
		healthLost := n * shieldedDamage

		// Remaining Health
		remainingHP := H - healthLost
		if remainingHP < 0 {
			remainingHP = 0
		}

		// Shots needed to finish off remaining HP (Unshielded)
		shotsAfterBreak := math.Ceil(remainingHP / D)
		shotsToKill = n + shotsAfterBreak
	}

	// 4. Calculate Stats for Frontend
	finalShots := int(shotsToKill)

	// Burst DPS
	rpm := req.Weapon.FireRate
	if rpm <= 0 {
		rpm = 1
	}
	rps := rpm / 60.0
	burstDPS := D * rps // Burst DPS // WIP not correct values for rpm

	// Cyclic DPS
	mag := req.Weapon.MagSize
	if mag <= 0 {
		mag = 30
	}

	magDumpTime := float64(mag-1) / rps
	damagePerMag := D * float64(mag)
	cyclicDPS := damagePerMag / (magDumpTime + req.Weapon.ReloadTime)

	c.JSON(http.StatusOK, gin.H{
		"shots_to_kill": finalShots,
		"damage":        D,
		"burst_dps":     burstDPS,
		"cyclic_dps":    cyclicDPS,
	})
}
