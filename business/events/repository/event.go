package repository

type IEventRepo interface {
}

type EventRepo struct {
}

func NewEventRepo() IEventRepo {
	return EventRepo{}
}
