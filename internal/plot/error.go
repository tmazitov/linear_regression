package plot

import "errors"

var (
	ErrPlotConfiguration = errors.New("plot configuration error:")
	ErrPlotSave          = errors.New("plot save error:")
)
