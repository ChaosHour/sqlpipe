package formatters

import (
	"database/sql"

	"github.com/sqlpipe/sqlpipe/internal/engine/transfers/formatters/shared"
)

var SnowflakeCreateFormatters = map[string]func(column *sql.ColumnType, terminator string) (string, error){
	"SQL_UNKNOWN_TYPE":    shared.TextCreateFormatter,
	"SQL_CHAR":            shared.TextCreateFormatter,
	"SQL_NUMERIC":         shared.NumberCreateFormatter,
	"SQL_DECIMAL":         shared.NumberCreateFormatter,
	"SQL_INTEGER":         shared.IntCreateFormatter,
	"SQL_SMALLINT":        shared.SmallIntCreateFormatter,
	"SQL_FLOAT":           shared.FloatCreateFormatter,
	"SQL_REAL":            shared.FloatCreateFormatter,
	"SQL_DOUBLE":          shared.FloatCreateFormatter,
	"SQL_DATETIME":        shared.TimestampCreateFormatter,
	"SQL_TIME":            shared.TimeCreateFormatter,
	"SQL_VARCHAR":         shared.TextCreateFormatter,
	"SQL_TYPE_DATE":       shared.DateCreateFormatter,
	"SQL_TYPE_TIME":       shared.TimeCreateFormatter,
	"SQL_TYPE_TIMESTAMP":  shared.TimestampCreateFormatter,
	"SQL_TIMESTAMP":       shared.TimestampCreateFormatter,
	"SQL_LONGVARCHAR":     shared.TextCreateFormatter,
	"SQL_BINARY":          shared.BinaryCreateFormatter,
	"SQL_VARBINARY":       shared.BinaryCreateFormatter,
	"SQL_LONGVARBINARY":   shared.BinaryCreateFormatter,
	"SQL_BIGINT":          shared.BigIntCreateFormatter,
	"SQL_TINYINT":         shared.SmallIntCreateFormatter,
	"SQL_BIT":             shared.BooleanCreateFormatter,
	"SQL_WCHAR":           shared.TextCreateFormatter,
	"SQL_WVARCHAR":        shared.TextCreateFormatter,
	"SQL_WLONGVARCHAR":    shared.TextCreateFormatter,
	"SQL_GUID":            shared.TextCreateFormatter,
	"SQL_SIGNED_OFFSET":   shared.TextCreateFormatter,
	"SQL_UNSIGNED_OFFSET": shared.TextCreateFormatter,
	"SQL_SS_XML":          shared.TextCreateFormatter,
	"SQL_SS_TIME2":        shared.TimeCreateFormatter,
}

var SnowflakeValFormatters = map[string]func(value interface{}, terminator string) (formattedValue string, err error){
	"SQL_UNKNOWN_TYPE":    shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_CHAR":            shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_NUMERIC":         shared.RawXnull,
	"SQL_DECIMAL":         shared.RawXnull,
	"SQL_INTEGER":         shared.RawXnull,
	"SQL_SMALLINT":        shared.RawXnull,
	"SQL_FLOAT":           shared.RawXnull,
	"SQL_REAL":            shared.RawXnull,
	"SQL_DOUBLE":          shared.RawXnull,
	"SQL_DATETIME":        shared.CastToTimeFormatToTimetampStringXnull,
	"SQL_TIME":            shared.CastToTimeFormatToTimeStringXnull,
	"SQL_VARCHAR":         shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_TYPE_DATE":       shared.CastToTimeFormatToSnowflakeDateStringXnull,
	"SQL_TYPE_TIME":       shared.CastToTimeFormatToTimeStringXnull,
	"SQL_TYPE_TIMESTAMP":  shared.CastToTimeFormatToTimetampStringXnull,
	"SQL_TIMESTAMP":       shared.CastToTimeFormatToTimetampStringXnull,
	"SQL_LONGVARCHAR":     shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_BINARY":          shared.CastToBytesCastToStringPrintQuotedHexXnull,
	"SQL_VARBINARY":       shared.CastToBytesCastToStringPrintQuotedHexXnull,
	"SQL_LONGVARBINARY":   shared.CastToBytesCastToStringPrintQuotedHexXnull,
	"SQL_BIGINT":          shared.RawXnull,
	"SQL_TINYINT":         shared.RawXnull,
	"SQL_BIT":             shared.CastToBoolWriteTextEquivalentXnull,
	"SQL_WCHAR":           shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_WVARCHAR":        shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_WLONGVARCHAR":    shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_GUID":            shared.QuotedXnull,
	"SQL_SIGNED_OFFSET":   shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_UNSIGNED_OFFSET": shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_SS_XML":          shared.CastToBytesCastToStringPrintQuotedXnull,
	"SQL_SS_TIME2":        shared.CastToBytesCastToStringPrintQuotedXnull,
}
