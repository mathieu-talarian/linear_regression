package main

import (
	"mathmoul/linear_regression/src/common"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func points(datas datas) plotter.XYs {
	pts := make(plotter.XYs, len(datas))
	for k, v := range datas {
		pts[k].X = v.km
		pts[k].Y = v.price
	}
	return pts
}

func normalized(datas datas, limits common.Limits, thetas common.Thetas) *plotter.Function {
	return plotter.NewFunction(func(x float64) float64 {
		normed := (x - limits.Min) / (limits.Max - limits.Min)
		return thetas.Zero + (thetas.One * normed)
	})
}

func graph(datas datas, limits common.Limits, thetas common.Thetas) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.X.Label.Text = "Km"
	p.Y.Label.Text = "Price"
	plotutil.AddLinePoints(p, points(datas))
	p.Add(normalized(datas, limits, thetas))
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "datas.png"); err != nil {
		panic(err)
	}
}
