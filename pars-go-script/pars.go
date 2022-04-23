package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	// Loop through every single Pen in the known universe.
	// For now, just get one to operate on.
	Pen1 := getPen1()
	fmt.Println("-----------------------")
	fmt.Println("Starting Processing of", Pen1.title)

	// Get a slice of all the PARS URLs on the Pen (that we actually care about).
	CSSPARSonPen := getCSSPARs(Pen1)

	if len(CSSPARSonPen) > 0 {

		for _, v := range CSSPARSonPen {
			stringParts := strings.Split(v, "/")
			slug := stringParts[5]
			fmt.Println("The slug of the PARS is: ", slug)

			// Pretend to query for the PARS using the slug
			Pen2 := getPen2()

			// Logic to test for matching preprocessors (and it's not vanilla CSS)
			if Pen1.css_processor == Pen2.css_processor && Pen1.css_processor != "none" {
				fmt.Println("MATCH: use the processor extension", v)
				// We'd need a map of correct file extensions to use...
				fmt.Println("We should update the PARS URL to: ", v+"."+Pen1.css_processor)

			} else {
				fmt.Println("NO MATCH: use the .css URL extension")
				fmt.Println("We should update the PARS URL to: ", v+".css")
			}

		}

	} else {
		fmt.Println("No PARS found, do nothing")
	}
}

func cleanUrl(inURL string) string {
	// fmt.Println("Before cleaning: ", inURL)
	u, err := url.Parse(inURL)
	if err != nil {
		return inURL
	}

	// Rip off the query string and the hash
	u.RawQuery = ""
	u.Fragment = ""
	// fmt.Println("After cleaning: ", u)
	return u.String()
}

func getCSSPARs(pen pen) []string {
	allCSSPARS := []string{}

	for _, v := range pen.external_css {
		// rudimentary test for a PARS URL
		// 1) Is CodePen URL
		// 2) Is a /pen/ URL
		// 3) Doesn't already have a valid URL extension
		if strings.HasPrefix(v, "https://codepen.io") && strings.Contains(v, "/pen/") && !strings.HasSuffix(v, ".css") && !strings.HasSuffix(v, ".less") && !strings.HasSuffix(v, ".scss") && !strings.HasSuffix(v, ".sass") && !strings.HasSuffix(v, ".stylus") && !strings.HasSuffix(v, ".postcss") {
			fmt.Println("RESOURCE: Needs updating", v)
			allCSSPARS = append(allCSSPARS, cleanUrl(v))

		} else {
			fmt.Println("RESOURCE: Nothing to update about it.", v)
		}
	}

	fmt.Println("# of PARS that need logic processing:", len(allCSSPARS))
	return allCSSPARS
}
