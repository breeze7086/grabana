package decoder

import (
	"github.com/breeze7086/grabana/dashboard"
	"github.com/breeze7086/grabana/row"
)

type DashboardRow struct {
	Name   string
	Panels []DashboardPanel
}

func (r DashboardRow) toOption() (dashboard.Option, error) {
	opts := []row.Option{}

	for _, panel := range r.Panels {
		opt, err := panel.toOption()
		if err != nil {
			return nil, err
		}

		opts = append(opts, opt)
	}

	return dashboard.Row(r.Name, opts...), nil
}
