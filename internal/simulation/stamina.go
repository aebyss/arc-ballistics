package simulation

import "arc-simulator/internal/models"

func ApplyStaminaCost(c *models.Character, cost float64) {
    c.Stamina -= cost
    if c.Stamina < 0 {
        c.Stamina = 0
    }
}

func RegenerateStamina(c *models.Character) {
    c.Stamina += c.RegenRate
    if c.Stamina > c.MaxStamina {
        c.MaxStamina = c.Stamina
    }
}

