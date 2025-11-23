package models

type Weapon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Rarity string `json:"rarity"`

	// Physics Stats (Used for Calculation)
	DamageClose float64 `json:"damage_close"`
	DamageFar   float64 `json:"damage_far"`
	RangeClose  float64 `json:"range_close"`
	RangeFar    float64 `json:"range_far"`
	FireRate    float64 `json:"fire_rate"` // Real RPM (e.g. 100)

	// Display Stats (Used for UI Bars)
	DisplayFireRate float64 `json:"display_fire_rate"` //this is only ui score not real rpm

	// Math Fields
	MagSize    int     `json:"mag_size"`
	ReloadTime float64 `json:"reload_time"`

	// Visuals
	IconURL     string  `json:"icon_url"`
	Description string  `json:"description"`
	FlavorText  string  `json:"flavor_text"`
	AmmoType    string  `json:"ammo_type"`
	Weight      float64 `json:"weight"`
	Value       int     `json:"value"`

	Agility   float64 `json:"agility"`
	Stability float64 `json:"stability"`
	Stealth   float64 `json:"stealth"`
}
