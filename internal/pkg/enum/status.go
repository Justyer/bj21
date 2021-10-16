package enum

const (
	DBStatusCreate = "create"
	DBStatusQueue  = "queue"
	DBStatusDoing  = "doing"
	DBStatusFail   = "fail"
	DBStatusSucc   = "succ"
	DBStatusCancel = "cancel"
)

const (
	APIStatusCreate = "create"
	APIStatusDoing  = "doing"
	APIStatusFail   = "fail"
	APIStatusSucc   = "succ"
)

const (
	CmdLogin     = "login"
	CmdLogout    = "logout"
	CmdTableList = "tablelist"
	CmdSitDown   = "sitdown"
)
