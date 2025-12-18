package limiter

import (
	"errors"
	"time"
)

type DailyLimiter struct {
	limit int
	count int
	day   time.Time
}

func NewDailyLimiter(limit int) *DailyLimiter {
	return &DailyLimiter{
		limit: limit,
		day:   startOfDay(time.Now()),
	}
}

func (l *DailyLimiter) Allow() error {
	now := time.Now()
	if !sameDay(l.day, now) {
		l.day = startOfDay(now)
		l.count = 0
	}
	if l.count >= l.limit {
		return errors.New("daily connection limit reached")
	}
	l.count++
	return nil
}

func startOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
