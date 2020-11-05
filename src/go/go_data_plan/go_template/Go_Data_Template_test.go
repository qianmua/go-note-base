package go_template

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"testing"
)

/**
 * @author
 * @date 2020/11/5 15:33
 * create by Goland
 * @version 1.0
 */ 
 
// template
// MVC模式中V的处理

// Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象
// {{.FieldName}} 段必须是导出的(字段首字母必须是大写的)


type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func TestTemplateM1(t *testing.T) {
	t2 := template.New("template demo")
	t2, _ = t2.Parse("hello {{.UserName}}!")
	p := Person{UserName:"jackson"}
	t2.Execute(os.Stdout , p)

}

// 输出嵌套字段内容

func TestTemplateM2(t *testing.T) {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t3 := template.New("fieldname example")
	t3, _ = t3.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t3.Execute(os.Stdout, p)
}


// 条件处理
// if里面无法使用条件判断，例如.Mail=="astaxie@gmail.com"，这样的判断是不正确的，if里面只能是bool值
func TestTemplateM3(t *testing.T) {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

// pipeline
// 管道
// {{. | html}}

// 模板变量
// $variable := pipeline
/**
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
 */


// 模板函数

// 每一个模板函数都有一个唯一值的名字，然后与一个Go函数关联，通过如下的方式来关联
// type FuncMap map[string]interface{}

// 例如，如果我们想要的email函数的模板函数名是emailDeal，
//它关联的Go函数名称是EmailDealWith,那么我们可以通过下面的方式来注册这个函数
// t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})

// EmailDealWith这个函数的参数和返回值定义如下：
// func EmailDealWith(args …interface{}) string

func TestTemplateM4(t *testing.T) {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t2 := template.New("fieldname example")
	t2 = t2.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t2, _ = t2.Parse(`hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
				{{end}}
				`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t2.Execute(os.Stdout, p)
}
func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return substrs[0] + " at " + substrs[1]
}


// 嵌套模板

// 比如thymeleaf 模板
// header body footer

// {{define "子模板名称"}}内容{{end}}
// {{template "子模板名称"}}

func TestTemplateM5(t *testing.T) {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}