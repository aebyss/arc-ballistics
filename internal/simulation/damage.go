package simulation

import (
    "arc-simulator/internal/models"
)

func CalculateDamageAtDistance(w models.Weapon, dist float64) float64 {
    if dist <= w.RangeClose {
        return w.DamageClose
    }
    if dist >= w.RangeFar {
        return w.DamageFar
    }
    t := (dist - w.RangeClose) / (w.RangeFar - w.RangeClose)
    return w.DamageClose + t*(w.DamageFar - w.DamageClose)
}


func ApplyDamage(t models.Target, dmg float64) models.Target {
    if t.Shield > 0 {
        t.Shield -= dmg

        t.Health -= dmg * t.Dr

        if t.Shield < 0 {
            bleed := -t.Shield
            t.Shield = 0
            t.Health -= bleed
        }

    } else {
        t.Health -= dmg
    }

    if t.Health < 0 {
        t.Health = 0
    }

    return t
}



func SimulateHit(w models.Weapon, t models.Target, dist float64) models.Target {
    dmg := CalculateDamageAtDistance(w, dist)
    return ApplyDamage(t, dmg)
}

