package main

type App struct {
	Name     string
	Title    string
	Apex     string
	Intro    string
	Features []string
	ImageUrl string
}

var apps map[string]App = make(map[string]App)

func init() {

	apps["imagelicious.org"] = App{
		Name:  "imagelicious",
		Title: "Imagelicious",
		Apex:  "imagelicious.org",
		Intro: "Simple to use personal photo gallery built on Firebase.",
		Features: []string{
			"Uses Firebase Authentication",
			"A Single-Page App built with ReactJS",
			"Bundled using Browserify",
			"Creates thumbnails inside the browser",
		},
	}

	apps["publish.li"] = App{
		Name:  "publish",
		Title: "publish.li",
		Apex:  "publish.li",
		Intro: "Article publishing site - no login - just write and publish!",
		Features: []string{
			"Written in Go",
			"Uses BoltDB embedded key/value store",
			"Frontend written with Vue.js",
		},
	}

	apps["bcrypt.org"] = App{
		Name:  "bcrypt",
		Title: "bcrypt.org",
		Apex:  "bcrypt.org",
		Intro: "Simple server with API to check your bcrypt password strategy.",
		Features: []string{
			"Written in Go",
			"Simple API",
			"Use in your unit tests",
			"Independent validation of your passwords",
		},
	}

}
