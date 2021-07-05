package service

func (ac *ActionContext) Done() {
	ac.done = true
}

func (ac *ActionContext) HasDone() bool {
	return ac.done
}

func (ac *ActionContext) Copy() *ActionContext {
	return &ActionContext{Logger: ac.Logger}
}

