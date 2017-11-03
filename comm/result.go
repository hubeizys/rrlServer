package comm

type Result struct {
	Ret map[string]interface{}
}

func (cst *Result) Construct() *Result {
	cst.Ret = make(map[string]interface{})
	cst.SetValue("", 0, "")
	return cst
}

func (cst *Result) SetValue(p_err string, p_num int64, p_result interface{}) {
	cst.Ret["err"] = p_err
	cst.Ret["num"] = p_num
	cst.Ret["result"] = p_result
}

func (cst *Result) SetResult(p_result interface{}) {
	cst.Ret["result"] = p_result
}

func (cst *Result) Get() map[string]interface{} {
	return cst.Ret
}

func init() {
}
