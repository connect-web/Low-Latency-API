package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

var (
	MAX_LIMIT = 100
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
