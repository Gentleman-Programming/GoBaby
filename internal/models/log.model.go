package models

import (
	"GoBaby/internal/utils"
	"time"
)

type Log struct {
	Date     time.Time // Date of the log
	Duration int       // Duration of the intake
}

func (c Log) FormatDuration(duration int) string {
	return utils.FormatDuration(duration)
}
