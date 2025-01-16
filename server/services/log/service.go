package log

import "github.com/crowci/crow/v3/server/model"

type Service interface {
	LogFind(step *model.Step) ([]*model.LogEntry, error)
	LogAppend(step *model.Step, logEntries []*model.LogEntry) error
	LogDelete(step *model.Step) error
}
