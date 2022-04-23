package main

type pen struct {
	id    string
	title string

	html string
	css  string
	js   string

	html_processor string
	html_classes   string
	stuff_for_head string

	css_processor    string
	css_base         string
	vendor_prefixing string
	external_css     []string

	js_processor string
	external_js  []string
}

func getPen1() pen {
	return pen{
		id:    "OJzGgPb",
		title: "Pen 1",

		html: "<div>Hello World</div>",
		css: `div { 
	color: red; 
}`,
		js: "alert('Hello World');",

		html_processor: "none",
		html_classes:   "",
		stuff_for_head: "",

		css_processor:    "none",
		css_base:         "none",
		vendor_prefixing: "none",
		external_css:     []string{"https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.0.2/css/bootstrap.min.css", "https://codepen.io/chriscoyier/pen/NWXmjZj", "https://codepen.io/chriscoyier/pen/dafDjk?stupid=stuff#hello"},

		js_processor: "none",
		external_js:  []string{},
	}
}

func getPen2() pen {
	return pen{
		id:    "NWXmjZj",
		title: "Pen 2",

		html: "",
		css: `body {
	font-size: 120%;
}`,
		js: "",

		html_processor: "none",
		html_classes:   "",
		stuff_for_head: "",

		css_processor:    "none",
		css_base:         "none",
		vendor_prefixing: "none",
		external_css:     []string{},

		js_processor: "none",
		external_js:  []string{},
	}
}
