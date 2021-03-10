package generativeart

import (
	"github.com/fogleman/gg"
	"github.com/jdxyw/generativeart/common"
	"math"
	"math/rand"
)

type oceanFish struct {
	lineNum int
	fishNum int
}

// NewOceanFish returns a oceanFish object.
func NewOceanFish(lineNum, fishNum int) *oceanFish {
	return &oceanFish{
		lineNum: lineNum,
		fishNum: fishNum,
	}
}

// Generative draws a ocean and fish images.
func (o *oceanFish) Generative(c *canva) {
	ctex := gg.NewContextForRGBA(c.img)

	o.drawlines(ctex, c)

	for i := 0; i < o.fishNum; i++ {
		ctex.Push()

		theta := float64(360*i) / float64(o.fishNum)
		r := float64(c.width) / 4.0

		ctex.Push()
		ctex.Translate(float64(c.width/2)+r*math.Cos(gg.Radians(theta)), float64(c.height/2)+r*math.Sin(gg.Radians(theta)))
		ctex.Rotate(gg.Radians(theta + 90))
		o.drawfish(ctex, c, 0, 0, float64(c.width)/10)
		ctex.Pop()

		ctex.Clip()
		o.drawlines(ctex, c)
		ctex.Pop()
		ctex.ClearPath()
		ctex.ResetClip()
	}
}

func (o *oceanFish) drawlines(ctx *gg.Context, c *canva) {
	for i := 0; i < o.lineNum; i++ {
		cl := c.opts.colorSchema[rand.Intn(len(c.opts.colorSchema))]
		ctx.SetColor(cl)
		ctx.SetLineWidth(common.RandomRangeFloat64(3, 20))
		y := rand.Float64() * float64(c.height)
		ctx.DrawLine(0, y+common.RandomRangeFloat64(-50, 50), float64(c.width), y+common.RandomRangeFloat64(-50, 50))
		ctx.Stroke()
	}
}

func (o *oceanFish) drawfish(ctex *gg.Context, c *canva, ox, oy, r float64) {
	ctex.Push()
	ctex.Translate(ox, oy)
	ctex.Rotate(gg.Radians(180))

	ctex.MoveTo(r*math.Cos(gg.Radians(0))-r*math.Pow(math.Sin(gg.Radians(0)), 2)/math.Sqrt(2),
		r*math.Cos(gg.Radians(0))*math.Sin(gg.Radians(0)))
	for theta := 1.0; theta < 361.0; theta += 1.0 {
		x := r*math.Cos(gg.Radians(theta)) - r*math.Pow(math.Sin(gg.Radians(theta)), 2)/math.Sqrt(2)
		y := r * math.Cos(gg.Radians(theta)) * math.Sin(gg.Radians(theta))
		ctex.LineTo(x, y)
	}
	ctex.ClosePath()
	ctex.Pop()
}
