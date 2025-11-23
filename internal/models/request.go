package models

type SimRequest struct {
    Weapon   Weapon  `json:"weapon"`
    Target   Target  `json:"target"`
    Distance float64 `json:"distance"`
}

