package simulation

import (
	"arc-simulator/internal/models"
)

func SingleShot(w models.Weapon, t models.Target, dist float64) models.Target {
	dmg := CalculateDamageAtDistance(w, dist)
	return ApplyDamage(t, dmg)
}

func ShotsToKill(w models.Weapon, t models.Target, dist float64) (int, models.Target) {
	shots := 0
	for t.Health > 0 {
		t = SingleShot(w, t, dist)
		shots++
		if shots > 500 {
			break
		}
	}
	return shots, t
}

func TTK(w models.Weapon, stk int) float64 {
	if stk <= 0 {
		return 0
	}
	return float64(stk-1) / w.FireRate
}
