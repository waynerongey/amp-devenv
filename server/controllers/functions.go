package controllers

import (
	"amp-templates/server/services/configuration"
	"amp-templates/server/services/log"
	"amp-templates/server/services/mux"
	"amp-templates/server/services/template_cache"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)


var (
	AmpStylesCache template.CSS
)

//ConfigureTemplateFuncs functions available in templates
func ConfigureTemplateFuncs() {

	template_cache.TemplateFuncs.FuncMap["Route"] = RebuildRoute
	template_cache.TemplateFuncs.FuncMap["CanonicalRoute"] = RebuildCanonicalRoute
	template_cache.TemplateFuncs.FuncMap["GetAmpStyles"] = GetAmpStyles
	template_cache.TemplateFuncs.FuncMap["HotReloading"] = HotReloading
}

//Err checks the error and redirects if there is one
//path must be rebuilt
func Err(err error, w http.ResponseWriter, r *http.Request, path string, message string) bool {
	if err != nil {
		msg := message + ": " + err.Error()
		Redirect(w, r, path, msg)
		return true
	}
	return false
}

//Redirect redirects the request
func Redirect(w http.ResponseWriter, r *http.Request, path string, message string) {
	log.Info(fmt.Sprintf("Redirecting %s from %s to %s because '%s'", r.RemoteAddr, r.URL.Path, path, message))
	http.Redirect(w, r, path, http.StatusSeeOther)
}

//RebuildCanonicalRoute rebuilds the route with the key value pairs
//as a conanical route (with the hostname)
func RebuildCanonicalRoute(name string, pairs ...interface{}) string {
	route := RebuildRoute(name, pairs...)
	return "https://" + configuration.Config.Domain + route
}

//RebuildRoute rebuilds the relative route with key value pairs
func RebuildRoute(name string, pairs ...interface{}) string {
	stringPairs := make([]string, len(pairs))
	for i := 0; i < len(pairs); i++ {
		var val string
		t := pairs[i]
		switch t.(type) {
		case string:
			val = pairs[i].(string)
		case int:
			val = strconv.Itoa(pairs[i].(int))
		case uint:
			val = fmt.Sprint(pairs[i])
		}
		stringPairs = append(stringPairs, val)
	}
	route, err := mux.Router.Get(name).URL(stringPairs...)
	if err != nil {
		return "/"
	}
	return route.Path
}

func executeTemplate(t string, w http.ResponseWriter, r *http.Request, data interface{}) {
	tc := template_cache.Lookup(t)
	if tc != nil {
		if err := tc.Execute(w, data); err != nil {
			log.Error("Error executing template: " + t + " err: " + err.Error())
			return
		}
	} else {
		log.Error("Couldn't find template in cache: " + t)
		return
	}
}

func GetAmpStyles() template.CSS {
	if configuration.Config.Production && AmpStylesCache != "" {
		return AmpStylesCache
	}
	styles, err := ioutil.ReadFile(configuration.Config.AmpStylesFile)
	if err != nil {
		log.Error(fmt.Sprintf("Couldn't find amp styles file at %s. Check config.json.", configuration.Config.AmpStylesFile))
		return ""
	}
	AmpStylesCache = template.CSS(strings.ReplaceAll(string(styles), "!important", ""))
	return AmpStylesCache
}


func HotReloading() template.HTML {
	if configuration.Config.Production {
		return ""
	}

	return template.HTML(fmt.Sprintf(`
<script>
	window.addEventListener("load", function() {
		var qs = getQueryString();
		if (qs['x'] || qs['y']) {
			window.scrollTo(qs['x'], qs['y']); 
		}
		var ws = new WebSocket("ws://%s/api/dev/hr");
		ws.onmessage = function(evt) {
			if (evt.data === "reload") {
				var top  = window.pageYOffset || document.documentElement.scrollTop,
					left = window.pageXOffset || document.documentElement.scrollLeft;
				var l = location;
				location = l.protocol + '//' + l.hostname + ':' + l.port +  l.pathname + '?x=' + left + '&y=' + top;
			}
		}
	});

	function getQueryString() {
    var queryStringKeyValue = window.location.search.replace('?', '').split('&');
    var qsJsonObject = {};
    if (queryStringKeyValue != '') {
        for (i = 0; i < queryStringKeyValue.length; i++) {
            qsJsonObject[queryStringKeyValue[i].split('=')[0]] = queryStringKeyValue[i].split('=')[1];
        }
    }
    return qsJsonObject;
}
</script>
`, configuration.Config.Address))
}