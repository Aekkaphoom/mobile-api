package dateutils

import "time"

type setLayouts struct {
	layout string
}

type SetFormatDate interface {
	ConvertDateStringToSQLDate(dateStr time.Time) time.Time
}

func New(format ...string) SetFormatDate {
	defaultFormat := "2006-01-02 15:04:05"
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	return &setLayouts{
		layout: defaultFormat,
	}
}

func (f *setLayouts) ConvertDateStringToSQLDate(dateStr time.Time) time.Time {
	t, err := time.Parse(f.layout, dateStr.Format(f.layout))
	if err != nil {
		return time.Time{}
	}
	return t
}
