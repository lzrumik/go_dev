package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:ActivitySubdivisionUserBaseInfoCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:DevPropertyConfigureController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FORTUNEREPAYMENTPLANEXECLOGController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityCustomController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityExtendController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivityResultController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySigninController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneActivitySignupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateApproveHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAppropriateInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalBizIdentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowRelevanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneApprovalFlowStepController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAreaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofAttachmentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneAssetProofHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthApproveHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfoCstmController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankAuthInfo很危险的Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBankListController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneBatchChangeTaskDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCallHistoryAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCanalSubscribeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneChangeOfficialHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleAssociateController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyCycleTaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201808Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201809Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201810Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201811Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201812Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201901Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201902Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfo201906Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyTaskRecordInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneClassifyUserIntIdMinController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoAttachController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCompanyInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook2010614Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180201Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180621Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20180704Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181011Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20181015Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBook20190306Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllocationHisotoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookAllotRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerBookHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistory20180201Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanAttachController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerCallPlanController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerChangeRelationHistoryToolsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeAllotHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerSubscribeCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskAnomalyCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerTaskCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneAllotRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomerWealthPlanneController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersAddressController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopy20180718Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersCstmController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersToolsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomersXxxController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeAllotRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneCustomizeSubscribeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneEvaluationNewController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneGroupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIdentSignController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneIntlAccountInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAchievementDatumController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBak20180428Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanBakCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanContractQueueController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCopy11111Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanCstmController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodBakController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoBakController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodCalculateInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanPeriodController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLoanReorderManageController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinGoldenController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaJoinLoanController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneLogOverseaReserveController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketBindWorkorderReportController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderAllotRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMarketUnbindWorkorderReportController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneMsgSendrecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeButtonController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneNodeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOrderDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneOverseasRegisterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentDownloadTaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePayReportZipArchiveController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePfQuestionEvaluationHistoryCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePlanCallMsgHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortunePositionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQrCodeInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQualifiedFlagController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionBankController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneQuestionEvaluationHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecommendFcSettingDetailDealController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRecordApprovalBizController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneReportImportServiceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneRoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSubscribeApprovalBizController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneSundryConfigController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneTempCstmController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUsaAccountInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionApproveInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoBakController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionAttachmentInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBak20180509Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoBakController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserSubscriptionInfoCopyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneUserToolsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVarsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsHistoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneVideoRecordsWorkOrdersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:FortuneWorkHoursController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:GlobalRegionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:MakeUniqueIdController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb1Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TestTb2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest1DepartController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTest2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestLoanController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestOrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"] = append(beego.GlobalControllerRouter["ch84_beego_api_project/controllers:TriggerTestSubscriptionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
