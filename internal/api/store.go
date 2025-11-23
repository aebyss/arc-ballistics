package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"arc-simulator/internal/models"
)

const (
	MetaForgeURL = "https://metaforge.app/api/arc-raiders/items"
	CacheFile    = "weapons_cache.json"
)

type WeaponStore struct {
	Weapons map[string]models.Weapon
	mutex   sync.RWMutex
}

var Store = &WeaponStore{
	Weapons: make(map[string]models.Weapon),
}

func InitStore() {
	if err := Store.loadFromDisk(); err != nil {
		fmt.Println("No local cache. Fetching from MetaForge...")
		Store.Sync()
	} else {
		fmt.Println("Loaded weapons from disk.")
	}

	go func() {
		for range time.Tick(24 * time.Hour) {
			Store.Sync()
		}
	}()
}

func (s *WeaponStore) GetAll() map[string]models.Weapon {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.Weapons
}

type metaForgeItem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ItemType string `json:"item_type"`
	Rarity   string `json:"rarity"`

	// Root Level Fields
	Icon        string      `json:"icon"`
	Description string      `json:"description"`
	FlavorText  string      `json:"flavor_text"`
	Value       interface{} `json:"value"` // Handle int/float
	Subcategory string      `json:"subcategory"`
	AmmoType    string      `json:"ammo_type"`

	StatBlock map[string]interface{} `json:"stat_block"`
}

func (s *WeaponStore) Sync() error {
	client := http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(MetaForgeURL + "?item_type=Weapon&limit=300")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResp struct {
		Data []metaForgeItem `json:"data"`
	}

	if err := json.Unmarshal(body, &apiResp); err != nil {
		return err
	}

	cleanWeapons := make(map[string]models.Weapon)

	for _, item := range apiResp.Data {
		if item.ItemType != "Weapon" {
			continue
		}

		getVal := func(key string) float64 {
			if v, ok := item.StatBlock[key]; ok {
				return toFloat(v)
			}
			return 0
		}

		// --- RPM & CLASS LOGIC ---
		subCat := item.Subcategory
		lowerCat := strings.ToLower(subCat)

		uiFireRate := getVal("fireRate") // e.g. 16.3

		realRPM := 600.0 // Default

		// Logic to map Weapon Class to Real Physics RPM WIP
		if strings.Contains(lowerCat, "hand cannon") {
			realRPM = 100 // Slow logic for Anvil
		} else if strings.Contains(lowerCat, "sniper") || strings.Contains(lowerCat, "bolt") {
			realRPM = 50
		} else if strings.Contains(lowerCat, "rocket") || strings.Contains(lowerCat, "launcher") {
			realRPM = 20
		} else if strings.Contains(lowerCat, "smg") {
			realRPM = 800
		} else if strings.Contains(lowerCat, "assault") {
			realRPM = 600
		} else if strings.Contains(lowerCat, "rifle") {
			realRPM = 300
		} else if uiFireRate > 0 {
			realRPM = uiFireRate * 10 // Fallback
		}

		// --- RELOAD LOGIC --- WIP
		reload := 2.5
		if strings.Contains(lowerCat, "hand cannon") {
			reload = 3.0
		}
		if strings.Contains(lowerCat, "sniper") {
			reload = 4.0
		}
		if strings.Contains(lowerCat, "smg") {
			reload = 2.0
		}

		damage := getVal("damage")
		rng := getVal("range")

		w := models.Weapon{
			ID:     item.ID,
			Name:   item.Name,
			Type:   item.ItemType,
			Rarity: item.Rarity,

			FireRate:        realRPM,
			DisplayFireRate: uiFireRate,

			DamageClose: damage,
			RangeClose:  rng * 0.3,
			RangeFar:    rng,
			DamageFar:   damage * 0.6,

			MagSize:    int(getVal("magazineSize")),
			ReloadTime: reload,

			IconURL:     item.Icon,
			Description: item.Description,
			FlavorText:  item.FlavorText,
			AmmoType:    item.AmmoType,
			Weight:      getVal("weight"),
			Value:       int(toFloat(item.Value)),

			Agility:   getVal("agility"),
			Stability: getVal("stability"),
			Stealth:   getVal("stealth"),
		}
		// DEFAULTS
		if w.DamageClose == 0 {
			w.DamageClose = 10
		}
		if w.MagSize == 0 {
			w.MagSize = 30
		}
		if w.RangeFar == 0 {
			w.RangeFar = 50
			w.RangeClose = 15
		}
		if w.IconURL == "" {
			w.IconURL = "https://placehold.co/200x100?text=No+Image"
		}
		if w.AmmoType == "" {
			w.AmmoType = "Standard"
		}

		key := strings.ToLower(strings.ReplaceAll(item.Name, " ", "_"))
		cleanWeapons[key] = w
	}

	s.mutex.Lock()
	s.Weapons = cleanWeapons
	s.mutex.Unlock()

	s.saveToDisk(cleanWeapons)
	fmt.Printf("Synced %d weapons (Root Fields Fixed).\n", len(cleanWeapons))
	return nil
}

func (s *WeaponStore) saveToDisk(data map[string]models.Weapon) {
	file, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(CacheFile, file, 0644)
}

func (s *WeaponStore) loadFromDisk() error {
	data, err := os.ReadFile(CacheFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.Weapons)
}

func toFloat(unk interface{}) float64 {
	if unk == nil {
		return 0
	}
	switch i := unk.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	default:
		return 0
	}
}
