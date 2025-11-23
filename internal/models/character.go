package models

type Character struct {
	Stamina		float64 `json:"stamina"`
	MaxStamina	float64 `json:"max_stamina"`
	RegenRate	float64 `json:"regen_rate"`
	Talents		[]Talent `json:"talents"`
}
