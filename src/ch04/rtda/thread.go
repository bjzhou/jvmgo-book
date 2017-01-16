package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.pop()
}
