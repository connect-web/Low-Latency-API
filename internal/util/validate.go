package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var (
	SkillsMap = map[string]struct{}{
		"attack":       {},
		"defence":      {},
		"strength":     {},
		"hitpoints":    {},
		"ranged":       {},
		"prayer":       {},
		"magic":        {},
		"cooking":      {},
		"woodcutting":  {},
		"fletching":    {},
		"fishing":      {},
		"firemaking":   {},
		"crafting":     {},
		"smithing":     {},
		"mining":       {},
		"herblore":     {},
		"agility":      {},
		"thieving":     {},
		"slayer":       {},
		"farming":      {},
		"runecraft":    {},
		"hunter":       {},
		"construction": {},
	}

	MinigamesMap = map[string]struct{}{
		"Bounty Hunter - Hunter":            {},
		"Bounty Hunter - Rogue":             {},
		"Bounty Hunter (Legacy) - Hunter":   {},
		"Bounty Hunter (Legacy) - Rogue":    {},
		"Clue Scrolls (all)":                {},
		"Clue Scrolls (beginner)":           {},
		"Clue Scrolls (easy)":               {},
		"Clue Scrolls (medium)":             {},
		"Clue Scrolls (hard)":               {},
		"Clue Scrolls (elite)":              {},
		"Clue Scrolls (master)":             {},
		"LMS - Rank":                        {},
		"PvP Arena - Rank":                  {},
		"Soul Wars Zeal":                    {},
		"Rifts closed":                      {},
		"Abyssal Sire":                      {},
		"Alchemical Hydra":                  {},
		"Artio":                             {},
		"Barrows Chests":                    {},
		"Bryophyta":                         {},
		"Callisto":                          {},
		"Calvar'ion":                        {},
		"Cerberus":                          {},
		"Chambers of Xeric":                 {},
		"Chambers of Xeric: Challenge Mode": {},
		"Chaos Elemental":                   {},
		"Chaos Fanatic":                     {},
		"Commander Zilyana":                 {},
		"Corporeal Beast":                   {},
		"Crazy Archaeologist":               {},
		"Dagannoth Prime":                   {},
		"Dagannoth Rex":                     {},
		"Dagannoth Supreme":                 {},
		"Deranged Archaeologist":            {},
		"Duke Sucellus":                     {},
		"General Graardor":                  {},
		"Giant Mole":                        {},
		"Grotesque Guardians":               {},
		"Hespori":                           {},
		"Kalphite Queen":                    {},
		"King Black Dragon":                 {},
		"Kraken":                            {},
		"Kree'Arra":                         {},
		"K'ril Tsutsaroth":                  {},
		"Mimic":                             {},
		"Nex":                               {},
		"Nightmare":                         {},
		"Phosani's Nightmare":               {},
		"Obor":                              {},
		"Phantom Muspah":                    {},
		"Sarachnis":                         {},
		"Scorpia":                           {},
		"Skotizo":                           {},
		"Spindel":                           {},
		"Tempoross":                         {},
		"The Gauntlet":                      {},
		"The Corrupted Gauntlet":            {},
		"The Leviathan":                     {},
		"The Whisperer":                     {},
		"Theatre of Blood":                  {},
		"Theatre of Blood: Hard Mode":       {},
		"Thermonuclear Smoke Devil":         {},
		"Tombs of Amascut":                  {},
		"Tombs of Amascut: Expert Mode":     {},
		"TzKal-Zuk":                         {},
		"TzTok-Jad":                         {},
		"Vardorvis":                         {},
		"Venenatis":                         {},
		"Vet'ion":                           {},
		"Vorkath":                           {},
		"Wintertodt":                        {},
		"Zalcano":                           {},
		"Zulrah":                            {},
		"Deadman Points":                    {},
		"League Points":                     {},
		"Colosseum Glory":                   {},
	}
)

var Title = cases.Title(language.English)

func ValidMinigame(minigame string) bool {
	_, valid := MinigamesMap[minigame]
	return valid
}

func ValidateSkills(skills []string) (validSkills []string) {
	validSkills = make([]string, 0)
	repeatPrevention := map[string]struct{}{}

	for _, skill := range skills {
		skill_lower := strings.ToLower(skill)

		// prevent duplicate entries
		_, isRepeated := repeatPrevention[skill_lower]
		if isRepeated {
			continue
		}

		// validate skill from pre-defined map
		_, valid := SkillsMap[skill_lower]
		if valid {
			validSkills = append(validSkills, Title.String(skill_lower))
			repeatPrevention[skill_lower] = struct{}{}
		}
	}
	return validSkills
}
