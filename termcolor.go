// termcolor.go
// Derek Rhodes, physci@gmail.com
// based on python library: termcolors by
// Konstantin Lepa <konstantin lepa at gmail com>
// license GPLv3, see COPYING.txt

package termcolor;
import . "fmt";

// if anyone wants the other attributes, let me know.

func Colorize(s, color string) string{
	code := 7;

	switch color{
	case "black":   code = 0; break;
	case "red":     code = 1; break;
	case "green":   code = 2; break;
	case "yellow":  code = 3; break;
	case "blue":    code = 4; break;
	case "magenta": code = 5; break;
	case "cyan":    code = 6; break;
	case "white":   code = 7; break;
	default: return s // couldn't find your color
	}

	reset := "\033[0m";
	return Sprintf("\033[1;3%dm%s%s", code, s, reset);
}
