package command

type CommandDispatcher interface {
	Dispatch()
}

type SimpleDispatcher struct {
}

func (sd *SimpleDispatcher) Dispatch() {

}
