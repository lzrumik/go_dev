// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ch84_beego_api_project/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		
		beego.NSNamespace("/FORTUNE_REPAYMENT_PLAN",
			beego.NSInclude(
				&controllers.FORTUNEREPAYMENTPLANController{},
			),
		),

		beego.NSNamespace("/FORTUNE_REPAYMENT_PLAN_EXEC_LOG",
			beego.NSInclude(
				&controllers.FORTUNEREPAYMENTPLANEXECLOGController{},
			),
		),

		beego.NSNamespace("/activity_subdivision_user_base_info",
			beego.NSInclude(
				&controllers.ActivitySubdivisionUserBaseInfoController{},
			),
		),

		beego.NSNamespace("/activity_subdivision_user_base_info_copy",
			beego.NSInclude(
				&controllers.ActivitySubdivisionUserBaseInfoCopyController{},
			),
		),

		beego.NSNamespace("/dev_property_configure",
			beego.NSInclude(
				&controllers.DevPropertyConfigureController{},
			),
		),

		beego.NSNamespace("/fortune_activity",
			beego.NSInclude(
				&controllers.FortuneActivityController{},
			),
		),

		beego.NSNamespace("/fortune_activity_custom",
			beego.NSInclude(
				&controllers.FortuneActivityCustomController{},
			),
		),

		beego.NSNamespace("/fortune_activity_extend",
			beego.NSInclude(
				&controllers.FortuneActivityExtendController{},
			),
		),

		beego.NSNamespace("/fortune_activity_result",
			beego.NSInclude(
				&controllers.FortuneActivityResultController{},
			),
		),

		beego.NSNamespace("/fortune_activity_signin",
			beego.NSInclude(
				&controllers.FortuneActivitySigninController{},
			),
		),

		beego.NSNamespace("/fortune_activity_signup",
			beego.NSInclude(
				&controllers.FortuneActivitySignupController{},
			),
		),

		beego.NSNamespace("/fortune_appropriate_approve_history",
			beego.NSInclude(
				&controllers.FortuneAppropriateApproveHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_appropriate_info",
			beego.NSInclude(
				&controllers.FortuneAppropriateInfoController{},
			),
		),

		beego.NSNamespace("/fortune_approval_biz_ident",
			beego.NSInclude(
				&controllers.FortuneApprovalBizIdentController{},
			),
		),

		beego.NSNamespace("/fortune_approval_flow_relevance",
			beego.NSInclude(
				&controllers.FortuneApprovalFlowRelevanceController{},
			),
		),

		beego.NSNamespace("/fortune_approval_flow_step",
			beego.NSInclude(
				&controllers.FortuneApprovalFlowStepController{},
			),
		),

		beego.NSNamespace("/fortune_area",
			beego.NSInclude(
				&controllers.FortuneAreaController{},
			),
		),

		beego.NSNamespace("/fortune_asset_proof",
			beego.NSInclude(
				&controllers.FortuneAssetProofController{},
			),
		),

		beego.NSNamespace("/fortune_asset_proof_attachment",
			beego.NSInclude(
				&controllers.FortuneAssetProofAttachmentController{},
			),
		),

		beego.NSNamespace("/fortune_asset_proof_history",
			beego.NSInclude(
				&controllers.FortuneAssetProofHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_bank_auth_approve_history",
			beego.NSInclude(
				&controllers.FortuneBankAuthApproveHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_bank_auth_history",
			beego.NSInclude(
				&controllers.FortuneBankAuthHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_bank_auth_info",
			beego.NSInclude(
				&controllers.FortuneBankAuthInfoController{},
			),
		),

		beego.NSNamespace("/fortune_bank_auth_info_很危险的",
			beego.NSInclude(
				&controllers.FortuneBankAuthInfo很危险的Controller{},
			),
		),

		beego.NSNamespace("/fortune_bank_auth_info_cstm",
			beego.NSInclude(
				&controllers.FortuneBankAuthInfoCstmController{},
			),
		),

		beego.NSNamespace("/fortune_bank_list",
			beego.NSInclude(
				&controllers.FortuneBankListController{},
			),
		),

		beego.NSNamespace("/fortune_batch_change_task",
			beego.NSInclude(
				&controllers.FortuneBatchChangeTaskController{},
			),
		),

		beego.NSNamespace("/fortune_batch_change_task_detail",
			beego.NSInclude(
				&controllers.FortuneBatchChangeTaskDetailController{},
			),
		),

		beego.NSNamespace("/fortune_call_history_attachment_info",
			beego.NSInclude(
				&controllers.FortuneCallHistoryAttachmentInfoController{},
			),
		),

		beego.NSNamespace("/fortune_canal_subscribe",
			beego.NSInclude(
				&controllers.FortuneCanalSubscribeController{},
			),
		),

		beego.NSNamespace("/fortune_canal_subscribe_allot_record",
			beego.NSInclude(
				&controllers.FortuneCanalSubscribeAllotRecordController{},
			),
		),

		beego.NSNamespace("/fortune_change_official_history",
			beego.NSInclude(
				&controllers.FortuneChangeOfficialHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_classify_cycle_associate",
			beego.NSInclude(
				&controllers.FortuneClassifyCycleAssociateController{},
			),
		),

		beego.NSNamespace("/fortune_classify_cycle_task",
			beego.NSInclude(
				&controllers.FortuneClassifyCycleTaskController{},
			),
		),

		beego.NSNamespace("/fortune_classify_task",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskController{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfoController{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201808",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201808Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201809",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201809Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201810",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201810Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201811",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201811Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201812",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201812Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201901",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201901Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201902",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201902Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_task_record_info_201906",
			beego.NSInclude(
				&controllers.FortuneClassifyTaskRecordInfo201906Controller{},
			),
		),

		beego.NSNamespace("/fortune_classify_user_int_id",
			beego.NSInclude(
				&controllers.FortuneClassifyUserIntIdController{},
			),
		),

		beego.NSNamespace("/fortune_classify_user_int_id_min",
			beego.NSInclude(
				&controllers.FortuneClassifyUserIntIdMinController{},
			),
		),

		beego.NSNamespace("/fortune_company",
			beego.NSInclude(
				&controllers.FortuneCompanyController{},
			),
		),

		beego.NSNamespace("/fortune_company_info",
			beego.NSInclude(
				&controllers.FortuneCompanyInfoController{},
			),
		),

		beego.NSNamespace("/fortune_company_info_attach",
			beego.NSInclude(
				&controllers.FortuneCompanyInfoAttachController{},
			),
		),

		beego.NSNamespace("/fortune_customer_book",
			beego.NSInclude(
				&controllers.FortuneCustomerBookController{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_2010614",
			beego.NSInclude(
				&controllers.FortuneCustomerBook2010614Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_20180201",
			beego.NSInclude(
				&controllers.FortuneCustomerBook20180201Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_20180621",
			beego.NSInclude(
				&controllers.FortuneCustomerBook20180621Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_20180704",
			beego.NSInclude(
				&controllers.FortuneCustomerBook20180704Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_20181011",
			beego.NSInclude(
				&controllers.FortuneCustomerBook20181011Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_20181015",
			beego.NSInclude(
				&controllers.FortuneCustomerBook20181015Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_20190306",
			beego.NSInclude(
				&controllers.FortuneCustomerBook20190306Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_allocation_hisotory",
			beego.NSInclude(
				&controllers.FortuneCustomerBookAllocationHisotoryController{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_allot_record",
			beego.NSInclude(
				&controllers.FortuneCustomerBookAllotRecordController{},
			),
		),

		beego.NSNamespace("/fortune_customer_book_history",
			beego.NSInclude(
				&controllers.FortuneCustomerBookHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_customer_call_history",
			beego.NSInclude(
				&controllers.FortuneCustomerCallHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_customer_call_history_20180201",
			beego.NSInclude(
				&controllers.FortuneCustomerCallHistory20180201Controller{},
			),
		),

		beego.NSNamespace("/fortune_customer_call_plan",
			beego.NSInclude(
				&controllers.FortuneCustomerCallPlanController{},
			),
		),

		beego.NSNamespace("/fortune_customer_call_plan_attach",
			beego.NSInclude(
				&controllers.FortuneCustomerCallPlanAttachController{},
			),
		),

		beego.NSNamespace("/fortune_customer_change_relation_history",
			beego.NSInclude(
				&controllers.FortuneCustomerChangeRelationHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_customer_change_relation_history_copy",
			beego.NSInclude(
				&controllers.FortuneCustomerChangeRelationHistoryCopyController{},
			),
		),

		beego.NSNamespace("/fortune_customer_change_relation_history_tools",
			beego.NSInclude(
				&controllers.FortuneCustomerChangeRelationHistoryToolsController{},
			),
		),

		beego.NSNamespace("/fortune_customer_subscribe",
			beego.NSInclude(
				&controllers.FortuneCustomerSubscribeController{},
			),
		),

		beego.NSNamespace("/fortune_customer_subscribe_allot_history",
			beego.NSInclude(
				&controllers.FortuneCustomerSubscribeAllotHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_customer_subscribe_copy",
			beego.NSInclude(
				&controllers.FortuneCustomerSubscribeCopyController{},
			),
		),

		beego.NSNamespace("/fortune_customer_task",
			beego.NSInclude(
				&controllers.FortuneCustomerTaskController{},
			),
		),

		beego.NSNamespace("/fortune_customer_task_anomaly",
			beego.NSInclude(
				&controllers.FortuneCustomerTaskAnomalyController{},
			),
		),

		beego.NSNamespace("/fortune_customer_task_anomaly_copy",
			beego.NSInclude(
				&controllers.FortuneCustomerTaskAnomalyCopyController{},
			),
		),

		beego.NSNamespace("/fortune_customer_task_copy",
			beego.NSInclude(
				&controllers.FortuneCustomerTaskCopyController{},
			),
		),

		beego.NSNamespace("/fortune_customer_wealth_planne",
			beego.NSInclude(
				&controllers.FortuneCustomerWealthPlanneController{},
			),
		),

		beego.NSNamespace("/fortune_customer_wealth_planne_allot_record",
			beego.NSInclude(
				&controllers.FortuneCustomerWealthPlanneAllotRecordController{},
			),
		),

		beego.NSNamespace("/fortune_customers",
			beego.NSInclude(
				&controllers.FortuneCustomersController{},
			),
		),

		beego.NSNamespace("/fortune_customers_address",
			beego.NSInclude(
				&controllers.FortuneCustomersAddressController{},
			),
		),

		beego.NSNamespace("/fortune_customers_copy",
			beego.NSInclude(
				&controllers.FortuneCustomersCopyController{},
			),
		),

		beego.NSNamespace("/fortune_customers_copy20180718",
			beego.NSInclude(
				&controllers.FortuneCustomersCopy20180718Controller{},
			),
		),

		beego.NSNamespace("/fortune_customers_cstm",
			beego.NSInclude(
				&controllers.FortuneCustomersCstmController{},
			),
		),

		beego.NSNamespace("/fortune_customers_tools",
			beego.NSInclude(
				&controllers.FortuneCustomersToolsController{},
			),
		),

		beego.NSNamespace("/fortune_customers_xxx",
			beego.NSInclude(
				&controllers.FortuneCustomersXxxController{},
			),
		),

		beego.NSNamespace("/fortune_customize_subscribe",
			beego.NSInclude(
				&controllers.FortuneCustomizeSubscribeController{},
			),
		),

		beego.NSNamespace("/fortune_customize_subscribe_allot_record",
			beego.NSInclude(
				&controllers.FortuneCustomizeSubscribeAllotRecordController{},
			),
		),

		beego.NSNamespace("/fortune_evaluation_history",
			beego.NSInclude(
				&controllers.FortuneEvaluationHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_evaluation_new",
			beego.NSInclude(
				&controllers.FortuneEvaluationNewController{},
			),
		),

		beego.NSNamespace("/fortune_group",
			beego.NSInclude(
				&controllers.FortuneGroupController{},
			),
		),

		beego.NSNamespace("/fortune_ident_sign",
			beego.NSInclude(
				&controllers.FortuneIdentSignController{},
			),
		),

		beego.NSNamespace("/fortune_intl_account_info",
			beego.NSInclude(
				&controllers.FortuneIntlAccountInfoController{},
			),
		),

		beego.NSNamespace("/fortune_loan",
			beego.NSInclude(
				&controllers.FortuneLoanController{},
			),
		),

		beego.NSNamespace("/fortune_loan_achievement_datum",
			beego.NSInclude(
				&controllers.FortuneLoanAchievementDatumController{},
			),
		),

		beego.NSNamespace("/fortune_loan_attachment_info",
			beego.NSInclude(
				&controllers.FortuneLoanAttachmentInfoController{},
			),
		),

		beego.NSNamespace("/fortune_loan_bak",
			beego.NSInclude(
				&controllers.FortuneLoanBakController{},
			),
		),

		beego.NSNamespace("/fortune_loan_bak_20180428",
			beego.NSInclude(
				&controllers.FortuneLoanBak20180428Controller{},
			),
		),

		beego.NSNamespace("/fortune_loan_bak_copy",
			beego.NSInclude(
				&controllers.FortuneLoanBakCopyController{},
			),
		),

		beego.NSNamespace("/fortune_loan_contract_queue",
			beego.NSInclude(
				&controllers.FortuneLoanContractQueueController{},
			),
		),

		beego.NSNamespace("/fortune_loan_copy11111",
			beego.NSInclude(
				&controllers.FortuneLoanCopy11111Controller{},
			),
		),

		beego.NSNamespace("/fortune_loan_cstm",
			beego.NSInclude(
				&controllers.FortuneLoanCstmController{},
			),
		),

		beego.NSNamespace("/fortune_loan_history",
			beego.NSInclude(
				&controllers.FortuneLoanHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_loan_period",
			beego.NSInclude(
				&controllers.FortuneLoanPeriodController{},
			),
		),

		beego.NSNamespace("/fortune_loan_period_bak",
			beego.NSInclude(
				&controllers.FortuneLoanPeriodBakController{},
			),
		),

		beego.NSNamespace("/fortune_loan_period_calculate_info",
			beego.NSInclude(
				&controllers.FortuneLoanPeriodCalculateInfoController{},
			),
		),

		beego.NSNamespace("/fortune_loan_period_calculate_info_bak",
			beego.NSInclude(
				&controllers.FortuneLoanPeriodCalculateInfoBakController{},
			),
		),

		beego.NSNamespace("/fortune_loan_reorder_manage",
			beego.NSInclude(
				&controllers.FortuneLoanReorderManageController{},
			),
		),

		beego.NSNamespace("/fortune_log_oversea_join_golden",
			beego.NSInclude(
				&controllers.FortuneLogOverseaJoinGoldenController{},
			),
		),

		beego.NSNamespace("/fortune_log_oversea_join_loan",
			beego.NSInclude(
				&controllers.FortuneLogOverseaJoinLoanController{},
			),
		),

		beego.NSNamespace("/fortune_log_oversea_reserve",
			beego.NSInclude(
				&controllers.FortuneLogOverseaReserveController{},
			),
		),

		beego.NSNamespace("/fortune_mail",
			beego.NSInclude(
				&controllers.FortuneMailController{},
			),
		),

		beego.NSNamespace("/fortune_market_bind_workorder",
			beego.NSInclude(
				&controllers.FortuneMarketBindWorkorderController{},
			),
		),

		beego.NSNamespace("/fortune_market_bind_workorder_report",
			beego.NSInclude(
				&controllers.FortuneMarketBindWorkorderReportController{},
			),
		),

		beego.NSNamespace("/fortune_market_unbind_workorder",
			beego.NSInclude(
				&controllers.FortuneMarketUnbindWorkorderController{},
			),
		),

		beego.NSNamespace("/fortune_market_unbind_workorder_allot_record",
			beego.NSInclude(
				&controllers.FortuneMarketUnbindWorkorderAllotRecordController{},
			),
		),

		beego.NSNamespace("/fortune_market_unbind_workorder_report",
			beego.NSInclude(
				&controllers.FortuneMarketUnbindWorkorderReportController{},
			),
		),

		beego.NSNamespace("/fortune_msg",
			beego.NSInclude(
				&controllers.FortuneMsgController{},
			),
		),

		beego.NSNamespace("/fortune_msg_sendrecord",
			beego.NSInclude(
				&controllers.FortuneMsgSendrecordController{},
			),
		),

		beego.NSNamespace("/fortune_node",
			beego.NSInclude(
				&controllers.FortuneNodeController{},
			),
		),

		beego.NSNamespace("/fortune_node_button",
			beego.NSInclude(
				&controllers.FortuneNodeButtonController{},
			),
		),

		beego.NSNamespace("/fortune_order_detail",
			beego.NSInclude(
				&controllers.FortuneOrderDetailController{},
			),
		),

		beego.NSNamespace("/fortune_overseas_register",
			beego.NSInclude(
				&controllers.FortuneOverseasRegisterController{},
			),
		),

		beego.NSNamespace("/fortune_pay_report",
			beego.NSInclude(
				&controllers.FortunePayReportController{},
			),
		),

		beego.NSNamespace("/fortune_pay_report_attachment_download_task",
			beego.NSInclude(
				&controllers.FortunePayReportAttachmentDownloadTaskController{},
			),
		),

		beego.NSNamespace("/fortune_pay_report_attachment_info",
			beego.NSInclude(
				&controllers.FortunePayReportAttachmentInfoController{},
			),
		),

		beego.NSNamespace("/fortune_pay_report_zip_archive",
			beego.NSInclude(
				&controllers.FortunePayReportZipArchiveController{},
			),
		),

		beego.NSNamespace("/fortune_pf_question_evaluation",
			beego.NSInclude(
				&controllers.FortunePfQuestionEvaluationController{},
			),
		),

		beego.NSNamespace("/fortune_pf_question_evaluation_copy",
			beego.NSInclude(
				&controllers.FortunePfQuestionEvaluationCopyController{},
			),
		),

		beego.NSNamespace("/fortune_pf_question_evaluation_history",
			beego.NSInclude(
				&controllers.FortunePfQuestionEvaluationHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_pf_question_evaluation_history_copy",
			beego.NSInclude(
				&controllers.FortunePfQuestionEvaluationHistoryCopyController{},
			),
		),

		beego.NSNamespace("/fortune_plan_call_msg_history",
			beego.NSInclude(
				&controllers.FortunePlanCallMsgHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_position",
			beego.NSInclude(
				&controllers.FortunePositionController{},
			),
		),

		beego.NSNamespace("/fortune_qr_code_info",
			beego.NSInclude(
				&controllers.FortuneQrCodeInfoController{},
			),
		),

		beego.NSNamespace("/fortune_qualified_flag",
			beego.NSInclude(
				&controllers.FortuneQualifiedFlagController{},
			),
		),

		beego.NSNamespace("/fortune_question",
			beego.NSInclude(
				&controllers.FortuneQuestionController{},
			),
		),

		beego.NSNamespace("/fortune_question_bank",
			beego.NSInclude(
				&controllers.FortuneQuestionBankController{},
			),
		),


		beego.NSNamespace("/fortune_question_evaluation",
			beego.NSInclude(
				&controllers.FortuneQuestionEvaluationController{},
			),
		),

		beego.NSNamespace("/fortune_question_evaluation_history",
			beego.NSInclude(
				&controllers.FortuneQuestionEvaluationHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_recommend_fc_setting",
			beego.NSInclude(
				&controllers.FortuneRecommendFcSettingController{},
			),
		),

		beego.NSNamespace("/fortune_recommend_fc_setting_detail",
			beego.NSInclude(
				&controllers.FortuneRecommendFcSettingDetailController{},
			),
		),

		beego.NSNamespace("/fortune_recommend_fc_setting_detail_deal",
			beego.NSInclude(
				&controllers.FortuneRecommendFcSettingDetailDealController{},
			),
		),

		beego.NSNamespace("/fortune_record_approval_biz",
			beego.NSInclude(
				&controllers.FortuneRecordApprovalBizController{},
			),
		),

		beego.NSNamespace("/fortune_report_import_service",
			beego.NSInclude(
				&controllers.FortuneReportImportServiceController{},
			),
		),

		beego.NSNamespace("/fortune_role",
			beego.NSInclude(
				&controllers.FortuneRoleController{},
			),
		),

		beego.NSNamespace("/fortune_subscribe_approval_biz",
			beego.NSInclude(
				&controllers.FortuneSubscribeApprovalBizController{},
			),
		),

		beego.NSNamespace("/fortune_sundry_config",
			beego.NSInclude(
				&controllers.FortuneSundryConfigController{},
			),
		),

		beego.NSNamespace("/fortune_temp",
			beego.NSInclude(
				&controllers.FortuneTempController{},
			),
		),

		beego.NSNamespace("/fortune_temp_cstm",
			beego.NSInclude(
				&controllers.FortuneTempCstmController{},
			),
		),

		beego.NSNamespace("/fortune_usa_account_info",
			beego.NSInclude(
				&controllers.FortuneUsaAccountInfoController{},
			),
		),

		beego.NSNamespace("/fortune_user",
			beego.NSInclude(
				&controllers.FortuneUserController{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_approve_info",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionApproveInfoController{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_attachment_info",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionAttachmentInfoController{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_attachment_info_bak",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionAttachmentInfoBakController{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_info",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionInfoController{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_info_bak",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionInfoBakController{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_info_bak_20180509",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionInfoBak20180509Controller{},
			),
		),

		beego.NSNamespace("/fortune_user_subscription_info_copy",
			beego.NSInclude(
				&controllers.FortuneUserSubscriptionInfoCopyController{},
			),
		),

		beego.NSNamespace("/fortune_user_tools",
			beego.NSInclude(
				&controllers.FortuneUserToolsController{},
			),
		),

		beego.NSNamespace("/fortune_vars",
			beego.NSInclude(
				&controllers.FortuneVarsController{},
			),
		),

		beego.NSNamespace("/fortune_video_records",
			beego.NSInclude(
				&controllers.FortuneVideoRecordsController{},
			),
		),

		beego.NSNamespace("/fortune_video_records_history",
			beego.NSInclude(
				&controllers.FortuneVideoRecordsHistoryController{},
			),
		),

		beego.NSNamespace("/fortune_video_records_work_orders",
			beego.NSInclude(
				&controllers.FortuneVideoRecordsWorkOrdersController{},
			),
		),

		beego.NSNamespace("/fortune_work_hours",
			beego.NSInclude(
				&controllers.FortuneWorkHoursController{},
			),
		),

		beego.NSNamespace("/global_region",
			beego.NSInclude(
				&controllers.GlobalRegionController{},
			),
		),

		beego.NSNamespace("/make_unique_id",
			beego.NSInclude(
				&controllers.MakeUniqueIdController{},
			),
		),

		beego.NSNamespace("/test_tb1",
			beego.NSInclude(
				&controllers.TestTb1Controller{},
			),
		),

		beego.NSNamespace("/test_tb2",
			beego.NSInclude(
				&controllers.TestTb2Controller{},
			),
		),

		beego.NSNamespace("/trigger_test1",
			beego.NSInclude(
				&controllers.TriggerTest1Controller{},
			),
		),

		beego.NSNamespace("/trigger_test1_depart",
			beego.NSInclude(
				&controllers.TriggerTest1DepartController{},
			),
		),

		beego.NSNamespace("/trigger_test2",
			beego.NSInclude(
				&controllers.TriggerTest2Controller{},
			),
		),

		beego.NSNamespace("/trigger_test_loan",
			beego.NSInclude(
				&controllers.TriggerTestLoanController{},
			),
		),

		beego.NSNamespace("/trigger_test_order",
			beego.NSInclude(
				&controllers.TriggerTestOrderController{},
			),
		),

		beego.NSNamespace("/trigger_test_subscription",
			beego.NSInclude(
				&controllers.TriggerTestSubscriptionController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
