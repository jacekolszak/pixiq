package blend

import (
	"github.com/jacekolszak/pixiq/image"
)

type ColorBlender interface {
	BlendSourceToTargetColor(source, target image.Color) image.Color
}

func New(colorBlender ColorBlender) *Tool {
	if colorBlender == nil {
		panic("nil colorBlender")
	}
	return &Tool{
		colorBlender: colorBlender,
	}
}

// TODO For optimization purposes there can be a dedicated tool which will override
// colors without using a ColorBlender
func NewSource() *Source {
	s := &Source{}
	s.Tool = New(s)
	return s
}

type Source struct {
	*Tool
}

func (s Source) BlendSourceToTargetColor(source, target image.Color) image.Color {
	return source
}

func NewSourceOver() *SourceOver {
	return NewSourceOverWithOpacity(100)
}

func NewSourceOverWithOpacity(opacity int) *SourceOver {
	tool := &SourceOver{opacity: opacity}
	tool.Tool = New(tool)
	return tool
}

// Aka Normal
type SourceOver struct {
	*Tool
	opacity int
}

func (s *SourceOver) BlendSourceToTargetColor(source, target image.Color) image.Color {
	return source
}

func (s *SourceOver) SetOpacity(opacity int) {
	s.opacity = opacity
}

type Tool struct {
	colorBlender ColorBlender
}

// BlendSourceToTarget blends source into target.
//
// If target has 0x0 size then whole source is blended, otherwise source is clamped.
func (t *Tool) BlendSourceToTarget(source, target image.Selection) {
	//sourceLines := source.Lines()
	//if sourceLines.Length() == 0 {
	//	return
	//}
	//if len(sourceLines.LineForRead(0)) == 0 {
	//	return
	//}
	//targetLines := target.Lines()
	//lines := sourceLines.Length()
	//if targetLines.Length() > 0 && lines > targetLines.Length() {
	//	lines = targetLines.Length()
	//}
	//for y := 0; y < lines; y++ {
	//	line := sourceLines.LineForRead(y)
	//	length := len(line)
	//	if target.Width() > 0 && length > target.Width() {
	//		length = target.Width()
	//	}
	//	for x := 0; x < length; x++ {
	//		sourceColor := line[x]
	//		targetColor := target.Color(x, y)
	//		target.SetColor(x, y, t.colorBlender.BlendSourceToTargetColor(sourceColor, targetColor))
	//	}
	//}
	width := source.Width()
	if target.Width() > 0 && target.Width() < width {
		width = target.Width()
	}
	height := source.Height()
	if target.Height() > 0 && target.Height() < height {
		height = target.Height()
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sourceColor := source.Color(x, y)
			targetColor := target.Color(x, y)
			color := t.colorBlender.BlendSourceToTargetColor(sourceColor, targetColor)
			target.SetColor(x, y, color)
		}
	}
}