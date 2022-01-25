package del2pakke

import "rsc.io/quote"

func TestQuote() string {
	return quote.Glass() + "\n" + quote.Go();
}
