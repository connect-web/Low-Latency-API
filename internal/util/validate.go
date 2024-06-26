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
)

var Title = cases.Title(language.English)

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
