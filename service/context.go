package service

func (ac *ActionNodeContext) Done() {
	ac.done = true
}

func (ac *ActionNodeContext) HasDone() bool {
	return ac.done
}

func (ac *ActionNodeContext) Copy() *ActionNodeContext {
	return &ActionNodeContext{Logger: ac.Logger}
}

