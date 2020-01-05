// Package colornames provides named colors as defined in the SVG 1.1 spec.
// The package is inspired by golang.org/x/image/colornames
//
// See http://www.w3.org/TR/SVG/types.html#ColorKeywords
// See https://github.com/golang/image/tree/master/colornames
package colornames

import (
	"github.com/jacekolszak/pixiq"
)

var (
	// Aliceblue            = pixiq.RGB(240, 248, 255)
	Aliceblue               = pixiq.RGB(240, 248, 255)
	// Antiquewhite         = pixiq.RGB(250, 235, 215)
	Antiquewhite            = pixiq.RGB(250, 235, 215)
	// Aqua                 = pixiq.RGB(0, 255, 255)
	Aqua                    = pixiq.RGB(0, 255, 255)
	// Aquamarine           = pixiq.RGB(127, 255, 212)
	Aquamarine              = pixiq.RGB(127, 255, 212)
	// Azure                = pixiq.RGB(240, 255, 255)
	Azure                   = pixiq.RGB(240, 255, 255)
	// Beige                = pixiq.RGB(245, 245, 220)
	Beige                   = pixiq.RGB(245, 245, 220)
	// Bisque               = pixiq.RGB(255, 228, 196)
	Bisque                  = pixiq.RGB(255, 228, 196)
	// Black                = pixiq.RGB(0, 0, 0)
	Black                   = pixiq.RGB(0, 0, 0)
	// Blanchedalmond       = pixiq.RGB(255, 235, 205)
	Blanchedalmond          = pixiq.RGB(255, 235, 205)
	// Blue                 = pixiq.RGB(0, 0, 255)
	Blue                    = pixiq.RGB(0, 0, 255)
	// Blueviolet           = pixiq.RGB(138, 43, 226)
	Blueviolet              = pixiq.RGB(138, 43, 226)
	// Brown                = pixiq.RGB(165, 42, 42)
	Brown                   = pixiq.RGB(165, 42, 42)
	// Burlywood            = pixiq.RGB(222, 184, 135)
	Burlywood               = pixiq.RGB(222, 184, 135)
	// Cadetblue            = pixiq.RGB(95, 158, 160)
	Cadetblue               = pixiq.RGB(95, 158, 160)
	// Chartreuse           = pixiq.RGB(127, 255, 0)
	Chartreuse              = pixiq.RGB(127, 255, 0)
	// Chocolate            = pixiq.RGB(210, 105, 30)
	Chocolate               = pixiq.RGB(210, 105, 30)
	// Coral                = pixiq.RGB(255, 127, 80)
	Coral                   = pixiq.RGB(255, 127, 80)
	// Cornflowerblue       = pixiq.RGB(100, 149, 237)
	Cornflowerblue          = pixiq.RGB(100, 149, 237)
	// Cornsilk             = pixiq.RGB(255, 248, 220)
	Cornsilk                = pixiq.RGB(255, 248, 220)
	// Crimson              = pixiq.RGB(220, 20, 60)
	Crimson                 = pixiq.RGB(220, 20, 60)
	// Cyan                 = pixiq.RGB(0, 255, 255)
	Cyan                    = pixiq.RGB(0, 255, 255)
	// Darkblue             = pixiq.RGB(0, 0, 139)
	Darkblue                = pixiq.RGB(0, 0, 139)
	// Darkcyan             = pixiq.RGB(0, 139, 139)
	Darkcyan                = pixiq.RGB(0, 139, 139)
	// Darkgoldenrod        = pixiq.RGB(184, 134, 11)
	Darkgoldenrod           = pixiq.RGB(184, 134, 11)
	// Darkgray             = pixiq.RGB(169, 169, 169)
	Darkgray                = pixiq.RGB(169, 169, 169)
	// Darkgreen            = pixiq.RGB(0, 100, 0)
	Darkgreen               = pixiq.RGB(0, 100, 0)
	// Darkgrey             = pixiq.RGB(169, 169, 169)
	Darkgrey                = pixiq.RGB(169, 169, 169)
	// Darkkhaki            = pixiq.RGB(189, 183, 107)
	Darkkhaki               = pixiq.RGB(189, 183, 107)
	// Darkmagenta          = pixiq.RGB(139, 0, 139)
	Darkmagenta             = pixiq.RGB(139, 0, 139)
	// Darkolivegreen       = pixiq.RGB(85, 107, 47)
	Darkolivegreen          = pixiq.RGB(85, 107, 47)
	// Darkorange           = pixiq.RGB(255, 140, 0)
	Darkorange              = pixiq.RGB(255, 140, 0)
	// Darkorchid           = pixiq.RGB(153, 50, 204)
	Darkorchid              = pixiq.RGB(153, 50, 204)
	// Darkred              = pixiq.RGB(139, 0, 0)
	Darkred                 = pixiq.RGB(139, 0, 0)
	// Darksalmon           = pixiq.RGB(233, 150, 122)
	Darksalmon              = pixiq.RGB(233, 150, 122)
	// Darkseagreen         = pixiq.RGB(143, 188, 143)
	Darkseagreen            = pixiq.RGB(143, 188, 143)
	// Darkslateblue        = pixiq.RGB(72, 61, 139)
	Darkslateblue           = pixiq.RGB(72, 61, 139)
	// Darkslategray        = pixiq.RGB(47, 79, 79)
	Darkslategray           = pixiq.RGB(47, 79, 79)
	// Darkslategrey        = pixiq.RGB(47, 79, 79)
	Darkslategrey           = pixiq.RGB(47, 79, 79)
	// Darkturquoise        = pixiq.RGB(0, 206, 209)
	Darkturquoise           = pixiq.RGB(0, 206, 209)
	// Darkviolet           = pixiq.RGB(148, 0, 211)
	Darkviolet              = pixiq.RGB(148, 0, 211)
	// Deeppink             = pixiq.RGB(255, 20, 147)
	Deeppink                = pixiq.RGB(255, 20, 147)
	// Deepskyblue          = pixiq.RGB(0, 191, 255)
	Deepskyblue             = pixiq.RGB(0, 191, 255)
	// Dimgray              = pixiq.RGB(105, 105, 105)
	Dimgray                 = pixiq.RGB(105, 105, 105)
	// Dimgrey              = pixiq.RGB(105, 105, 105)
	Dimgrey                 = pixiq.RGB(105, 105, 105)
	// Dodgerblue           = pixiq.RGB(30, 144, 255)
	Dodgerblue              = pixiq.RGB(30, 144, 255)
	// Firebrick            = pixiq.RGB(178, 34, 34)
	Firebrick               = pixiq.RGB(178, 34, 34)
	// Floralwhite          = pixiq.RGB(255, 250, 240)
	Floralwhite             = pixiq.RGB(255, 250, 240)
	// Forestgreen          = pixiq.RGB(34, 139, 34)
	Forestgreen             = pixiq.RGB(34, 139, 34)
	// Fuchsia              = pixiq.RGB(255, 0, 255)
	Fuchsia                 = pixiq.RGB(255, 0, 255)
	// Gainsboro            = pixiq.RGB(220, 220, 220)
	Gainsboro               = pixiq.RGB(220, 220, 220)
	// Ghostwhite           = pixiq.RGB(248, 248, 255)
	Ghostwhite              = pixiq.RGB(248, 248, 255)
	// Gold                 = pixiq.RGB(255, 215, 0)
	Gold                    = pixiq.RGB(255, 215, 0)
	// Goldenrod            = pixiq.RGB(218, 165, 32)
	Goldenrod               = pixiq.RGB(218, 165, 32)
	// Gray                 = pixiq.RGB(128, 128, 128)
	Gray                    = pixiq.RGB(128, 128, 128)
	// Green                = pixiq.RGB(0, 128, 0)
	Green                   = pixiq.RGB(0, 128, 0)
	// Greenyellow          = pixiq.RGB(173, 255, 47)
	Greenyellow             = pixiq.RGB(173, 255, 47)
	// Grey                 = pixiq.RGB(128, 128, 128)
	Grey                    = pixiq.RGB(128, 128, 128)
	// Honeydew             = pixiq.RGB(240, 255, 240)
	Honeydew                = pixiq.RGB(240, 255, 240)
	// Hotpink              = pixiq.RGB(255, 105, 180)
	Hotpink                 = pixiq.RGB(255, 105, 180)
	// Indianred            = pixiq.RGB(205, 92, 92)
	Indianred               = pixiq.RGB(205, 92, 92)
	// Indigo               = pixiq.RGB(75, 0, 130)
	Indigo                  = pixiq.RGB(75, 0, 130)
	// Ivory                = pixiq.RGB(255, 255, 240)
	Ivory                   = pixiq.RGB(255, 255, 240)
	// Khaki                = pixiq.RGB(240, 230, 140)
	Khaki                   = pixiq.RGB(240, 230, 140)
	// Lavender             = pixiq.RGB(230, 230, 250)
	Lavender                = pixiq.RGB(230, 230, 250)
	// Lavenderblush        = pixiq.RGB(255, 240, 245)
	Lavenderblush           = pixiq.RGB(255, 240, 245)
	// Lawngreen            = pixiq.RGB(124, 252, 0)
	Lawngreen               = pixiq.RGB(124, 252, 0)
	// Lemonchiffon         = pixiq.RGB(255, 250, 205)
	Lemonchiffon            = pixiq.RGB(255, 250, 205)
	// Lightblue            = pixiq.RGB(173, 216, 230)
	Lightblue               = pixiq.RGB(173, 216, 230)
	// Lightcoral           = pixiq.RGB(240, 128, 128)
	Lightcoral              = pixiq.RGB(240, 128, 128)
	// Lightcyan            = pixiq.RGB(224, 255, 255)
	Lightcyan               = pixiq.RGB(224, 255, 255)
	// Lightgoldenrodyellow = pixiq.RGB(250, 250, 210)
	Lightgoldenrodyellow    = pixiq.RGB(250, 250, 210)
	// Lightgray            = pixiq.RGB(211, 211, 211)
	Lightgray               = pixiq.RGB(211, 211, 211)
	// Lightgreen           = pixiq.RGB(144, 238, 144)
	Lightgreen              = pixiq.RGB(144, 238, 144)
	// Lightgrey            = pixiq.RGB(211, 211, 211)
	Lightgrey               = pixiq.RGB(211, 211, 211)
	// Lightpink            = pixiq.RGB(255, 182, 193)
	Lightpink               = pixiq.RGB(255, 182, 193)
	// Lightsalmon          = pixiq.RGB(255, 160, 122)
	Lightsalmon             = pixiq.RGB(255, 160, 122)
	// Lightseagreen        = pixiq.RGB(32, 178, 170)
	Lightseagreen           = pixiq.RGB(32, 178, 170)
	// Lightskyblue         = pixiq.RGB(135, 206, 250)
	Lightskyblue            = pixiq.RGB(135, 206, 250)
	// Lightslategray       = pixiq.RGB(119, 136, 153)
	Lightslategray          = pixiq.RGB(119, 136, 153)
	// Lightslategrey       = pixiq.RGB(119, 136, 153)
	Lightslategrey          = pixiq.RGB(119, 136, 153)
	// Lightsteelblue       = pixiq.RGB(176, 196, 222)
	Lightsteelblue          = pixiq.RGB(176, 196, 222)
	// Lightyellow          = pixiq.RGB(255, 255, 224)
	Lightyellow             = pixiq.RGB(255, 255, 224)
	// Lime                 = pixiq.RGB(0, 255, 0)
	Lime                    = pixiq.RGB(0, 255, 0)
	// Limegreen            = pixiq.RGB(50, 205, 50)
	Limegreen               = pixiq.RGB(50, 205, 50)
	// Linen                = pixiq.RGB(250, 240, 230)
	Linen                   = pixiq.RGB(250, 240, 230)
	// Magenta              = pixiq.RGB(255, 0, 255)
	Magenta                 = pixiq.RGB(255, 0, 255)
	// Maroon               = pixiq.RGB(128, 0, 0)
	Maroon                  = pixiq.RGB(128, 0, 0)
	// Mediumaquamarine     = pixiq.RGB(102, 205, 170)
	Mediumaquamarine        = pixiq.RGB(102, 205, 170)
	// Mediumblue           = pixiq.RGB(0, 0, 205)
	Mediumblue              = pixiq.RGB(0, 0, 205)
	// Mediumorchid         = pixiq.RGB(186, 85, 211)
	Mediumorchid            = pixiq.RGB(186, 85, 211)
	// Mediumpurple         = pixiq.RGB(147, 112, 219)
	Mediumpurple            = pixiq.RGB(147, 112, 219)
	// Mediumseagreen       = pixiq.RGB(60, 179, 113)
	Mediumseagreen          = pixiq.RGB(60, 179, 113)
	// Mediumslateblue      = pixiq.RGB(123, 104, 238)
	Mediumslateblue         = pixiq.RGB(123, 104, 238)
	// Mediumspringgreen    = pixiq.RGB(0, 250, 154)
	Mediumspringgreen       = pixiq.RGB(0, 250, 154)
	// Mediumturquoise      = pixiq.RGB(72, 209, 204)
	Mediumturquoise         = pixiq.RGB(72, 209, 204)
	// Mediumvioletred      = pixiq.RGB(199, 21, 133)
	Mediumvioletred         = pixiq.RGB(199, 21, 133)
	// Midnightblue         = pixiq.RGB(25, 25, 112)
	Midnightblue            = pixiq.RGB(25, 25, 112)
	// Mintcream            = pixiq.RGB(245, 255, 250)
	Mintcream               = pixiq.RGB(245, 255, 250)
	// Mistyrose            = pixiq.RGB(255, 228, 225)
	Mistyrose               = pixiq.RGB(255, 228, 225)
	// Moccasin             = pixiq.RGB(255, 228, 181)
	Moccasin                = pixiq.RGB(255, 228, 181)
	// Navajowhite          = pixiq.RGB(255, 222, 173)
	Navajowhite             = pixiq.RGB(255, 222, 173)
	// Navy                 = pixiq.RGB(0, 0, 128)
	Navy                    = pixiq.RGB(0, 0, 128)
	// Oldlace              = pixiq.RGB(253, 245, 230)
	Oldlace                 = pixiq.RGB(253, 245, 230)
	// Olive                = pixiq.RGB(128, 128, 0)
	Olive                   = pixiq.RGB(128, 128, 0)
	// Olivedrab            = pixiq.RGB(107, 142, 35)
	Olivedrab               = pixiq.RGB(107, 142, 35)
	// Orange               = pixiq.RGB(255, 165, 0)
	Orange                  = pixiq.RGB(255, 165, 0)
	// Orangered            = pixiq.RGB(255, 69, 0)
	Orangered               = pixiq.RGB(255, 69, 0)
	// Orchid               = pixiq.RGB(218, 112, 214)
	Orchid                  = pixiq.RGB(218, 112, 214)
	// Palegoldenrod        = pixiq.RGB(238, 232, 170)
	Palegoldenrod           = pixiq.RGB(238, 232, 170)
	// Palegreen            = pixiq.RGB(152, 251, 152)
	Palegreen               = pixiq.RGB(152, 251, 152)
	// Paleturquoise        = pixiq.RGB(175, 238, 238)
	Paleturquoise           = pixiq.RGB(175, 238, 238)
	// Palevioletred        = pixiq.RGB(219, 112, 147)
	Palevioletred           = pixiq.RGB(219, 112, 147)
	// Papayawhip           = pixiq.RGB(255, 239, 213)
	Papayawhip              = pixiq.RGB(255, 239, 213)
	// Peachpuff            = pixiq.RGB(255, 218, 185)
	Peachpuff               = pixiq.RGB(255, 218, 185)
	// Peru                 = pixiq.RGB(205, 133, 63)
	Peru                    = pixiq.RGB(205, 133, 63)
	// Pink                 = pixiq.RGB(255, 192, 203)
	Pink                    = pixiq.RGB(255, 192, 203)
	// Plum                 = pixiq.RGB(221, 160, 221)
	Plum                    = pixiq.RGB(221, 160, 221)
	// Powderblue           = pixiq.RGB(176, 224, 230)
	Powderblue              = pixiq.RGB(176, 224, 230)
	// Purple               = pixiq.RGB(128, 0, 128)
	Purple                  = pixiq.RGB(128, 0, 128)
	// Red                  = pixiq.RGB(255, 0, 0)
	Red                     = pixiq.RGB(255, 0, 0)
	// Rosybrown            = pixiq.RGB(188, 143, 143)
	Rosybrown               = pixiq.RGB(188, 143, 143)
	// Royalblue            = pixiq.RGB(65, 105, 225)
	Royalblue               = pixiq.RGB(65, 105, 225)
	// Saddlebrown          = pixiq.RGB(139, 69, 19)
	Saddlebrown             = pixiq.RGB(139, 69, 19)
	// Salmon               = pixiq.RGB(250, 128, 114)
	Salmon                  = pixiq.RGB(250, 128, 114)
	// Sandybrown           = pixiq.RGB(244, 164, 96)
	Sandybrown              = pixiq.RGB(244, 164, 96)
	// Seagreen             = pixiq.RGB(46, 139, 87)
	Seagreen                = pixiq.RGB(46, 139, 87)
	// Seashell             = pixiq.RGB(255, 245, 238)
	Seashell                = pixiq.RGB(255, 245, 238)
	// Sienna               = pixiq.RGB(160, 82, 45)
	Sienna                  = pixiq.RGB(160, 82, 45)
	// Silver               = pixiq.RGB(192, 192, 192)
	Silver                  = pixiq.RGB(192, 192, 192)
	// Skyblue              = pixiq.RGB(135, 206, 235)
	Skyblue                 = pixiq.RGB(135, 206, 235)
	// Slateblue            = pixiq.RGB(106, 90, 205)
	Slateblue               = pixiq.RGB(106, 90, 205)
	// Slategray            = pixiq.RGB(112, 128, 144)
	Slategray               = pixiq.RGB(112, 128, 144)
	// Slategrey            = pixiq.RGB(112, 128, 144)
	Slategrey               = pixiq.RGB(112, 128, 144)
	// Snow                 = pixiq.RGB(255, 250, 250)
	Snow                    = pixiq.RGB(255, 250, 250)
	// Springgreen          = pixiq.RGB(0, 255, 127)
	Springgreen             = pixiq.RGB(0, 255, 127)
	// Steelblue            = pixiq.RGB(70, 130, 180)
	Steelblue               = pixiq.RGB(70, 130, 180)
	// Tan                  = pixiq.RGB(210, 180, 140)
	Tan                     = pixiq.RGB(210, 180, 140)
	// Teal                 = pixiq.RGB(0, 128, 128)
	Teal                    = pixiq.RGB(0, 128, 128)
	// Thistle              = pixiq.RGB(216, 191, 216)
	Thistle                 = pixiq.RGB(216, 191, 216)
	// Tomato               = pixiq.RGB(255, 99, 71)
	Tomato                  = pixiq.RGB(255, 99, 71)
	// Turquoise            = pixiq.RGB(64, 224, 208)
	Turquoise               = pixiq.RGB(64, 224, 208)
	// Violet               = pixiq.RGB(238, 130, 238)
	Violet                  = pixiq.RGB(238, 130, 238)
	// Wheat                = pixiq.RGB(245, 222, 179)
	Wheat                   = pixiq.RGB(245, 222, 179)
	// White                = pixiq.RGB(255, 255, 255)
	White                   = pixiq.RGB(255, 255, 255)
	// Whitesmoke           = pixiq.RGB(245, 245, 245)
	Whitesmoke              = pixiq.RGB(245, 245, 245)
	// Yellow               = pixiq.RGB(255, 255, 0)
	Yellow                  = pixiq.RGB(255, 255, 0)
	// Yellowgreen          = pixiq.RGB(154, 205, 50)
	Yellowgreen             = pixiq.RGB(154, 205, 50)
)
