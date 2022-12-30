package gnostr

import (
	"github.com/nbd-wtf/go-nostr"
)

type RankingTarget struct {
	EventID    string
	Event      nostr.Event
	Score      int64
	Mentions   int64
	MentionMap map[string]struct{}
}

//1. get ALL kind 1 events
//2. get ALL mentions of the kind 1 events
//3. sort by most mentions
