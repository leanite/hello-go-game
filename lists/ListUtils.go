package lists

import (
	"hello.go.game/characters"
)

func RemoveNpcFromList(item *characters.Npc, list *[]*characters.Npc) {
	for i := 0; i < len(*list); i++ {
		if item.Equal((*list)[i]) {
			if len(*list) > 1 {
				*list = append((*list)[:i], (*list)[i+1:]...)
			} else {
				*list = []*characters.Npc{}
			}
			return
		}
	}
}
