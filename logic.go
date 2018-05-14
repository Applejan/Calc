package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//混凝土轴心抗压强度标准值
var gfcks = []float64{16.7, 20.1, 23.4, 26.8}

//混凝土轴心抗压强度设计值
var gfc = []float64{11.9, 14.3, 16.7, 19.1}

//混凝土轴心抗拉强度设计值
var gft = []float64{1.27, 1.43, 1.57, 1.71}

//混凝土弹性模量
var gec = []float64{2.8e+4, 3.0e+4, 3.15e+4, 3.25e+4}

var fy = 360.0
var es = 2.0e+5

//CalQEquations return two value
func CalQEquations(xa, xb, xc float64) (xval []float64, err error) {
	b := math.Pow(xb, 2) - 4*xa*xc
	if b < 0 {
		err = errors.New("二次方程无合理解！终止。")
		return
	}
	xval = []float64{(-xb + math.Sqrt(b)) / 2 / xa, (-xb - math.Sqrt(b)) / 2 / xa}
	return
}

func userjoin(args ...string) string {
	return strings.Join(args, " ")
}

//Calc is the main feture to calc tha value
func (con *Cons) Calc() (outstring string) {

	fc := gfc[con.DegreeID]
	h0 := con.Height - con.As0
	xa := fc * con.Width / 2
	xb := -fc * con.Width * h0
	scale, _ := strconv.ParseFloat(con.Scale, 64)
	xc := con.Moment * 1e+6 * scale
	fcuk := gfcks[con.DegreeID] * 1.5
	ecu := 0.0033 - (fcuk-50)*1e-5
	if ecu > 0.0033 {
		ecu = 0.0033
	}
	b1 := 0.8
	eb := b1 / (1 + fy/ecu/es)
	var x float64
	xslice, err := CalQEquations(xa, xb, xc)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)

	}
	for _, v := range xslice {
		if v > 0 && v < eb*h0 {
			x = v
			outstring = fmt.Sprintf("现在受压区高度为%v\n", strconv.FormatFloat(x, 'f', 2, 64))
		}
	}
	asvalue := fc * con.Width * x / fy
	p0 := math.Max(45*gft[con.DegreeID]/fy, 0.2)
	if x > eb*h0 {
		outstring = userjoin(outstring, "X超过最大受压区高度，增大截面\n")
	} else {
		outstring = userjoin(outstring, fmt.Sprintf("构造配筋As为%v\n", strconv.FormatFloat(p0/100*con.Width*con.Height, 'f', 2, 64)))
		outstring = userjoin(outstring, fmt.Sprintf("配筋面积As为%v\n", strconv.FormatFloat(asvalue, 'f', 2, 64)))
	}
	return
}
