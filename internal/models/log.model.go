package models

import (
	"GoBaby/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	Date     primitive.DateTime `bson:"date"`
	Duration int                `bson:"duration"` // Duration of the intake
}

func (c Log) FormatDuration(duration int) string {
	return utils.FormatDuration(duration)
}

func (c Log) FormatPrimitiveDateTime(date primitive.DateTime) string {
	return utils.FormatPrimitiveDate(date)
}
