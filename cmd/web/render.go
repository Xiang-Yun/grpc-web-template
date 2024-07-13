package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int64
	FloatMap        map[string]float32
	Data            map[string]interface{}
	Flash           string
	Warning         string
	Error           string // for custom show on screen -> show on page (function: notify)
	IsAuthenticated int
	UserID          int
	Backend         string
	Host            string
	ProxyAddr       string
}

var functions = template.FuncMap{
	"formatCurrency": formatCurrency,
	"dateCheck":      dateCheck,
}

func formatCurrency(n int) string {
	f := float32(n) / float32(100)
	return fmt.Sprintf("$%.2f", f)
}

func dateCheck(curr string) bool {
	return true
}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	// td.Flash = app.Session.PopString(r.Context(), "flash")
	// td.Warning = app.Session.PopString(r.Context(), "warning")
	// td.Error = app.Session.PopString(r.Context(), "error")
	td.Host = r.Host

	routeIP := strings.Split(td.Host, ":")[0]
	td.Backend = "http://" + routeIP + app.config.backend
	td.ProxyAddr = "http://" + routeIP + app.config.apiProxy

	fmt.Println(td.Backend)

	if app.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = 1
		td.UserID = app.Session.GetInt(r.Context(), "userID")
	} else {
		td.IsAuthenticated = 0
		td.UserID = 0
	}
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.templateCache[templateToRender]

	if templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}

	// parse desired base layout
	baseLayouts := []string{}
	if strings.Contains(page, "admin") {
		baseLayouts = append(baseLayouts, "templates/admin.layout.gohtml")
	} else {
		baseLayouts = append(baseLayouts, "templates/base.layout.gohtml")
	}

	if len(partials) > 0 {
		baseLayouts = append(baseLayouts, templateToRender, strings.Join(partials, ","), templateToRender)
	} else {
		baseLayouts = append(baseLayouts, templateToRender)
	}

	// render all the html page
	t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, baseLayouts...)
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
