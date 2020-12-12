/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package event

import "github.com/EngoEngine/ecs"

const (
	ABILITY_SCORE_TOTAL_UPDATED = "ability_score_total_updated"
)

type AbilityScoreTotalUpdatedMsg struct {
	*ecs.BasicEntity
	Before int
	After  int
}

func (a AbilityScoreTotalUpdatedMsg) Type() string {
	return ABILITY_SCORE_TOTAL_UPDATED
}
