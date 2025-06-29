### Template:
- package template : implements data-driven templates for generating HTML output.(html/tempate more secure then text/template)

	//	t, err := template.New("webpage").Parse(tpl)
	t := template.Must(template.New("webpage").Funcs(funcMap).Parse(tpl)) // it will panic immediately when there is an error , so no need to declare an err as I use it in template.New alone.
	//	checkError(err)

* Funcs adds the elements of  the argument map to the template's function map. it must be called before the template is parsed: like toUpper, convert data to specific format extra

```go
funcMap := map[string]interface{}{
  "toUpper": strings.ToUpper,
}
/// in the template:
{{ toUpper .Title }}

```
	
