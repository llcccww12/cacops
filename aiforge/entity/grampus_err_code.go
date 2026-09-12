package entity

var GrampusJobCanNotRestart = &ErrCode{CodeVal: "5005", CodeMsg: "Job can not restart", CodeTrCode: "ai_task.can_not_restart"}
var GrampusJobNotExistInCenter = &ErrCode{CodeVal: "5010", CodeMsg: "Job doesn't exist in center, can not be restarted", CodeTrCode: "ai_task.not_in_center"}
