// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/dev-sandip/go-auth/templates/layouts"
import "github.com/dev-sandip/go-auth/templates/components"

func Home() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"container\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.Markdown("# Why Choose Go and Gin for Backend Development?\n\n"+
				"Go (or Golang) and the Gin web framework have become go-to choices for building high-performance, scalable, and maintainable backend systems. Here's why:\n\n"+
				"## 🚀 Benefits of Using Go\n\n"+
				"- **Simplicity and Readability**\n"+
				"- **Concurrency Made Easy**\n"+
				"- **Performance**\n"+
				"- **Standard Library**\n"+
				"- **Static Typing**\n\n"+
				"## ⚙️ Benefits of Using Gin\n\n"+
				"- **High Performance**\n"+
				"- **Middleware Support**\n"+
				"- **JSON Validation and Binding**\n"+
				"- **Crash-Free**\n"+
				"- **Good Documentation and Community**\n\n"+
				"## 💻 Sample Gin Code\n\n"+
				"```go\n"+
				"package main\n\n"+
				"import (\n"+
				"    \"github.com/gin-gonic/gin\"\n"+
				")\n\n"+
				"func main() {\n"+
				"    r := gin.Default()\n\n"+
				"    r.GET(\"/\", func(c *gin.Context) {\n"+
				"        c.JSON(200, gin.H{\n"+
				"            \"message\": \"Hello from Gin!\",\n"+
				"        })\n"+
				"    })\n\n"+
				"    r.Run(\":8080\")\n"+
				"}\n"+
				"```\n\n"+
				"## 🔚 Conclusion\n\n"+
				"If you're looking for a backend stack that's fast, easy to scale, and simple to use — Go with Gin is an excellent choice!\n").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layouts.Base("Benefits of Go and Gin").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
