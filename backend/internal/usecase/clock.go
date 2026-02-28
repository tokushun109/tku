package usecase

import "time"

type Clock interface {
	Now() time.Time
}
