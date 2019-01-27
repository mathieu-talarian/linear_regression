package main

import (
	"image/color"
	"log"
	"mathmoul/linear_regression/src/common"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
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

func normalized(datas datas, thetas common.Thetas) *plotter.Function {
	f := plotter.NewFunction(func(x float64) float64 {
		return thetas.Zero + (thetas.One * x)
	})
	f.Color = color.RGBA{0, 255, 0, 255}
	return f
}

func graph(datas datas, thetas common.Thetas) {
	p, err := plot.New()
	p.Add(plotter.NewGrid())
	if err != nil {
		panic(err)
	}
	p.X.Label.Text = "Km"
	p.Y.Label.Text = "Price"
	s, err := plotter.NewScatter(points(datas))
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	p.Add(s)
	p.Legend.Add("Datas", s)
	f := normalized(datas, thetas)
	p.Add(f)
	p.Legend.Add("Normalised", f)
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "datas.png"); err != nil {
		panic(err)
	}
}
