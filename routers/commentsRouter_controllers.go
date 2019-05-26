package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:DemoUserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:DemoUserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:DemoUserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:DemoUserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id([0-9]+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Getini",
            Router: `/ini/:initKey`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Getinis",
            Router: `/inis/:initKey`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xlgwr/demoApigo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
