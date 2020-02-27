package template_cache

import (
	"amp-templates/server/services/configuration"
	"amp-templates/server/services/hot_reloading"
	"amp-templates/server/services/log"
	"bytes"
	"html/template"
	"os"
	"time"
)

var (
	cache         *template.Template
	TemplateFuncs = &FuncMap{template.FuncMap{}}
)

type FuncMap struct {
	template.FuncMap
}

func (t *FuncMap) AddFunc(k string, f func(...interface{}) (string, error)) {
	t.FuncMap[k] = f
}

func Lookup(template string) *template.Template {
	return cache.Lookup(template + ".html")
}

func LookupTextTmpl(template string) *template.Template {
	return cache.Lookup(template + ".html")
}

var lastModTime = time.Unix(0, 0)

func Configure() {
	tc, _ := buildTemplateCache()
	cache = tc

	if !configuration.Config.Production {
		go func() {
			for range time.Tick(1 * time.Second) {
				templateCache, isUpdated := buildTemplateCache()
				if isUpdated {
					cache = templateCache
				}
			}
		}()
	}
}

func buildTemplateCache() (*template.Template, bool) {
	needUpdate := false
	f, _ := os.Open(configuration.Config.Templates)
	fileInfos, _ := f.Readdir(-1)
	fileNames := make([]string, len(fileInfos))
	for idx, fi := range fileInfos {
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needUpdate = true
			log.Info("Adding Template: " + fi.Name())
		}
		fileNames[idx] = "templates/" + fi.Name()
	}

	var tc *template.Template
	if needUpdate {
		log.Info("Template change detected, updating...")
		tc = template.Must(template.New("").Funcs(TemplateFuncs.FuncMap).ParseFiles(fileNames...))
		log.Info("template update complete")
		hot_reloading.ForceReload()
	}

	return tc, needUpdate
}

func ExecuteTemplate(t string, data interface{}) string {
	tc := Lookup(t)

	buff := &bytes.Buffer{}
	if tc != nil {
		if err := tc.Execute(buff, data); err != nil {
			log.Error("Error executing template: " + t + " err: " + err.Error())
			return ""
		}
	} else {
		log.Error("Couldn't find template in cache: " + t)
		return ""
	}
	return buff.String()
}
