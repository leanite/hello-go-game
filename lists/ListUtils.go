package lists

import (
	"hello.go.game/characters"
)

func RemoveNpcFromList(item *characters.Npc, list []*characters.Npc) []*characters.Npc {
	// for i := 0; i < len(list); i++ {
	// 	if list[i].Equal(*item) {
	// 		if len(list) > 1 {
	// 			copy(list[i:], list[i+1:])
	// 			return list[:len(list)-1]
	// 		} else {
	// 			return []*characters.Npc{}
	// 		}
	// 	}
	// }
	var newList []*characters.Npc
	for _, npc := range list {
		if npc.Equal(item) == false {
			newList = append(newList, npc)
		}
	}
	return newList
}

func RemoveNpcFromListPtr(item *characters.Npc, list *[]*characters.Npc) {
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

/*
for i := 0; i < len(list); i++ {
		if *list[i] == *item {
			if len(list) > 1 {
				copy(list[i:], list[i+1:])
				return list[:len(list)-1]
			} else {
				return []*characters.Npc{}
			}
		}
	}
	return list
*/
