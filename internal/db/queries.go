package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

var (
	MAX_LIMIT = 1
)

// BuildPlayersByRatioQuery constructs the SQL query for fetching players based on a ratio.
func BuildPlayersByRatioQuery(skills []string, ratio float64) (string, []interface{}) {
	// Your existing query construction logic
	query := `
	SELECT
	    p.name,
	    player_live.skills_ratio,
	    player_live.minigames
	FROM player_live 
	LEFT JOIN players p ON p.id = player_live.playerid 
	WHERE `
	conditions := []string{}

	var paramNumber int
	params := []interface{}{}
	for i, skill := range skills {
		paramNumber = i + 1
		paramName := fmt.Sprintf("$%d", paramNumber)
		conditions = append(conditions, fmt.Sprintf("COALESCE(NULLIF((player_live.skills_ratio->>%s)::numeric, 'NaN'), 0)", paramName))
		params = append(params, skill)
	}

	// Join all conditions with + and compare to the ratio
	if len(conditions) > 0 {
		query += strings.Join(conditions, " + ") + fmt.Sprintf(" > CAST($%d AS numeric)", paramNumber+1)
		params = append(params, ratio)
	}

	// Execute the query
	query += fmt.Sprintf(" LIMIT %d;", MAX_LIMIT)
	return query, params
}

// BuildPlayersByExperienceQuery constructs the SQL query for fetching players based on a experience.
func BuildPlayersByExperienceQuery(skillExperienceThresholds map[string]int64) (string, []interface{}) {
	// Your existing query construction logic
	query := `
	SELECT
	    p.name,
	    player_live.skills_experience,
	    player_live.minigames
	FROM player_live 
	LEFT JOIN players p ON p.id = player_live.playerid 
	WHERE `
	conditions := []string{}

	params := []interface{}{}
	var paramNumber int
	for skill, experienceThreshold := range skillExperienceThresholds {
		conditions = append(conditions, fmt.Sprintf("CAST($%d AS numeric) < COALESCE(NULLIF((player_live.skills_experience->>CAST($%d AS TEXT))::numeric, 'NaN'), 0)", (paramNumber*2)+1, (paramNumber*2)+2))
		params = append(params, experienceThreshold)
		params = append(params, skill)

		paramNumber++
	}

	// Join all conditions with + and compare to the ratio
	if len(conditions) > 0 {
		query += strings.Join(conditions, " AND ")
	}

	// Execute the query
	query += fmt.Sprintf(" LIMIT %d;", MAX_LIMIT)
	return query, params
}

// BuildPlayersByExperienceQuery constructs the SQL query for fetching players based on a Level.
func BuildPlayersByLevelQuery(skillLevelThresholds map[string]int) (string, []interface{}) {
	// Your existing query construction logic
	query := `
	SELECT
	    p.name,
	    player_live.skills_Levels,
	    player_live.minigames
	FROM player_live 
	LEFT JOIN players p ON p.id = player_live.playerid 
	WHERE `
	conditions := []string{}

	params := []interface{}{}
	var paramNumber int
	for skill, LevelThreshold := range skillLevelThresholds {
		conditions = append(conditions, fmt.Sprintf("CAST($%d AS numeric) < COALESCE(NULLIF((player_live.skills_Levels->>CAST($%d AS TEXT))::numeric, 'NaN'), 0)", (paramNumber*2)+1, (paramNumber*2)+2))
		params = append(params, LevelThreshold)
		params = append(params, skill)

		paramNumber++
	}

	// Join all conditions with + and compare to the ratio
	if len(conditions) > 0 {
		query += strings.Join(conditions, " AND ")
	}

	// Execute the query
	query += fmt.Sprintf(" LIMIT %d;", MAX_LIMIT)
	return query, params
}

func RecentlyFoundAccountsNoMinigames() string {
	query := `
	select
		p.name,
		pls.total_level,
		pls.combat_level,
		player_live.skills_levels
	from player_live
	LEFT JOIN public.player_live_stats pls on player_live.playerid = pls.playerid
	LEFT JOIN PLAYERS P ON P.id = player_live.playerid
	WHERE DATE(player_live.last_updated) = DATE(NOW())
	AND player_live.minigames = '{}'
	ORDER BY pls.total_level DESC;
	`

	return query
}

func BuildUsersQuery(usernames map[string]struct{}) (string, []string) {
	query := `
	select
		p.name,
		pl.skills_levels,
		pl.minigames,
		gains.skills_experience
	FROM players p
	LEFT JOIN player_live pl on pl.playerid = p.id
	LEFT JOIN player_gains gains on gains.playerid = p.id
	where P.NAME = any($1);
	`
	uniqueUsernames := []string{}
	for username := range usernames {
		uniqueUsernames = append(uniqueUsernames, username)
	}

	return query, uniqueUsernames
}

func getMinCondition(skill string, level int, paramIndexInput int) (paramIndex int, conditionQuery string, params []interface{}) {
	params = []interface{}{
		level, skill,
	}
	conditionQuery = fmt.Sprintf(" and $%d <= COALESCE(NULLIF((PL.skills_levels->>$%d)::numeric, 'NaN'), 0)", paramIndexInput, paramIndexInput+1)
	return paramIndexInput + 2, conditionQuery, params
}
func getMaxCondition(skill string, level int, paramIndexInput int) (paramIndex int, conditionQuery string, params []interface{}) {
	params = []interface{}{
		level, skill,
	}
	conditionQuery = fmt.Sprintf(" and $%d >= COALESCE(NULLIF((PL.skills_levels->>$%d)::numeric, 'NaN'), 0)", paramIndexInput, paramIndexInput+1)
	return paramIndexInput + 2, conditionQuery, params
}

func getXpThresholdCondition(skill string, experience int, paramIndexInput int) (paramIndex int, conditionQuery string, params []interface{}) {
	params = []interface{}{
		experience, skill,
	}
	conditionQuery = fmt.Sprintf(" and $%d <= COALESCE(NULLIF((gains.skills_experience->>$%d)::numeric, 'NaN'), 0)", paramIndexInput, paramIndexInput+1)
	return paramIndexInput + 2, conditionQuery, params
}

func BuildSkillBotFinderQuery(selectedSkills map[string]struct{}, dailyXpThresholds, minLevels, maxLevels map[string]int) (string, []interface{}) {
	query := `
	select
		p.name,
		pl.skills_levels,
		pl.minigames,
		gains.skills_experience as daily_experience_gains
	
	from stats.pearson pn
	LEFT JOIN PLAYERS P ON P.ID = pn.PLAYERID
	LEFT JOIN player_live PL on PL.playerid = P.id
	LEFT JOIN player_gains gains on gains.playerid = P.id
	
	where
		50 < array_length(pn.linked_players, 1)
		%s
	
	ORDER BY array_length(pn.linked_players, 1) desc
	LIMIT 100;
	`

	params := []interface{}{}
	var conditions string
	paramIndex := 1

	for skill := range selectedSkills {
		xpThreshold, validXpThreshold := dailyXpThresholds[skill]
		minLevel, validminLevels := minLevels[skill]
		maxLevel, validmaxLevels := maxLevels[skill]

		if validXpThreshold {
			var conditionQuery string
			var newParams []interface{}

			paramIndex, conditionQuery, newParams = getXpThresholdCondition(skill, xpThreshold, paramIndex)
			conditions += conditionQuery
			params = append(params, newParams...)
		}

		if validminLevels {
			var conditionQuery string
			var newParams []interface{}

			paramIndex, conditionQuery, newParams = getMinCondition(skill, minLevel, paramIndex)
			conditions += conditionQuery
			params = append(params, newParams...)
		}

		if validmaxLevels {
			var conditionQuery string
			var newParams []interface{}

			paramIndex, conditionQuery, newParams = getMaxCondition(skill, maxLevel, paramIndex)
			conditions += conditionQuery
			params = append(params, newParams...)
		}

	}

	query = fmt.Sprintf(query, conditions)

	return query, params
}

func getMinigameCondition(minigame string, score int, paramIndexInput int) (paramIndex int, conditionQuery string, params []interface{}) {
	params = []interface{}{
		score, minigame,
	}
	conditionQuery = fmt.Sprintf(" and $%d <= COALESCE(NULLIF((gains.minigames->>$%d)::numeric, 'NaN'), 0)", paramIndexInput, paramIndexInput+1)
	return paramIndexInput + 2, conditionQuery, params
}

func BuildMinigameBotFinderQuery(selectedSkills map[string]struct{}, minigameThresholds, dailyXpThresholds, minLevels, maxLevels map[string]int) (string, []interface{}) {
	query := `
	select
		p.name,
		pl.skills_levels,
		pl.minigames,
		gains.skills_experience as daily_experience_gains,
		gains.minigames as daily_minigame_gains
	
	from stats.minigame_links pn
	LEFT JOIN PLAYERS P ON P.ID = pn.PLAYERID
	LEFT JOIN player_live PL on PL.playerid = P.id
	LEFT JOIN player_gains gains on gains.playerid = P.id
	
	where
		1=1
		%s
	
	ORDER BY pn.links desc
	LIMIT 100;
	`

	params := []interface{}{}
	var conditions string
	paramIndex := 1

	for skill := range selectedSkills {
		xpThreshold, validXpThreshold := dailyXpThresholds[skill]
		minLevel, validminLevels := minLevels[skill]
		maxLevel, validmaxLevels := maxLevels[skill]

		if validXpThreshold {
			var conditionQuery string
			var newParams []interface{}

			paramIndex, conditionQuery, newParams = getXpThresholdCondition(skill, xpThreshold, paramIndex)
			conditions += conditionQuery
			params = append(params, newParams...)
		}

		if validminLevels {
			var conditionQuery string
			var newParams []interface{}

			paramIndex, conditionQuery, newParams = getMinCondition(skill, minLevel, paramIndex)
			conditions += conditionQuery
			params = append(params, newParams...)
		}

		if validmaxLevels {
			var conditionQuery string
			var newParams []interface{}

			paramIndex, conditionQuery, newParams = getMaxCondition(skill, maxLevel, paramIndex)
			conditions += conditionQuery
			params = append(params, newParams...)
		}
	}

	for minigame, score := range minigameThresholds {
		var conditionQuery string
		var newParams []interface{}

		paramIndex, conditionQuery, newParams = getMinigameCondition(minigame, score, paramIndex)
		conditions += conditionQuery
		params = append(params, newParams...)
	}

	query = fmt.Sprintf(query, conditions)

	return query, params
}
