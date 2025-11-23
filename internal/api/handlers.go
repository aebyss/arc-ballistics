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

func calculateShotsToKill(D float64, H float64, S float64, DR float64) int {
	if D <= 0 {
		return 999 // safety
	}

	// Damage dealt while shield is active
	shieldedDamage := D * (1 - DR)
	if shieldedDamage <= 0 {
		shieldedDamage = 0.1
	}

	// Shots needed to break shield (shield takes full damage)
	n := math.Ceil(S / D)

	// target dies before shield breaks
	shotsToKillShielded := math.Ceil(H / shieldedDamage)
	if shotsToKillShielded <= n {
		return int(shotsToKillShielded)
	}

	// shield breaks, continue with full damage
	healthRemaining := H - n*shieldedDamage
	if healthRemaining < 0 {
		healthRemaining = 0
	}

	shotsAfterBreak := math.Ceil(healthRemaining / D)

	return int(n + shotsAfterBreak)
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

	D := calculateBaseDamage(req.Weapon, req.Distance)
	H := req.Target.Health
	S := req.Target.Shield
	DR := req.Target.DR

	// Correct calculation
	shotsToKill := calculateShotsToKill(D, H, S, DR)

	// --- DPS Calculations ---
	rpm := req.Weapon.FireRate
	if rpm <= 0 {
		rpm = 1
	}
	rps := rpm / 60.0
	burstDPS := D * rps

	mag := req.Weapon.MagSize
	if mag <= 0 {
		mag = 30
	}

	magDumpTime := float64(mag-1) / rps
	damagePerMag := D * float64(mag)
	cyclicDPS := damagePerMag / (magDumpTime + req.Weapon.ReloadTime)

	c.JSON(http.StatusOK, gin.H{
		"shots_to_kill": shotsToKill,
		"damage":        D,
		"burst_dps":     burstDPS,
		"cyclic_dps":    cyclicDPS,
	})
}
