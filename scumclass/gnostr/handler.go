package gnostr

import (
	"mindmachine/mindmachine"
)

func HandleEvent(event mindmachine.Event) (h mindmachine.HashSeq, b bool) {
	//if mind, _ := mindmachine.WhichMindForKind(event.Kind); mind == "gnostr" || event.Kind == 1 {
	//switch event.Kind {
	//case 1:
	currentState.mutex.Lock()
	defer currentState.mutex.Unlock()
	currentState.upsert(
		RankingTarget{
			EventID:    event.ID,
			Event:      event.Nostr(),
			MentionMap: make(map[string]struct{}),
		})
	//}
	//}
	return
}
