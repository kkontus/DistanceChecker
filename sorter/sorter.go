package sorter

import (
	"DistanceChecker/entity"
)

type ByUserId []entity.Customer

func (a ByUserId) Len() int           { return len(a) }
func (a ByUserId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUserId) Less(i, j int) bool { return a[i].UserId < a[j].UserId }
