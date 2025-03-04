// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package photo

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Form() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"min-w-80 flex flex-col flex-1 items-center justify-center h-full\"><form hx-post=\"/photos/add-new\" hx-encoding=\"multipart/form-data\" class=\"flex flex-col gap-8 p-4 bg-zinc-300 dark:bg-zinc-600 rounded-md shadow-md\"><div><label for=\"title\">Title</label> <input type=\"text\" name=\"title\" value=\"\" class=\"w-full\" required></div><div><label for=\"description\">Description</label> <textarea type=\"textarea\" name=\"description\" rows=\"6\" value=\"\" class=\"w-full\" required></textarea></div><input type=\"file\" name=\"file\" accept=\"image/jpeg\"> <button type=\"submit\" class=\"bg-zinc-700 text-zinc-200 dark:bg-zinc-200 dark:text-zinc-700 py-1 px-2 mt-2 rounded-md\">Send</button></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
