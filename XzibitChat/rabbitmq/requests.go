package rabbitmq

import (
	"XzibitChat/commands"
	"sync"
)

type Requests struct {
	requests map[int64]commands.Command
	id       int64
	mutex    *sync.RWMutex
}

func NewReqests() *Requests {
	return &Requests{
		requests: make(map[int64]commands.Command),
		id: 0,
		mutex: &sync.RWMutex{},
	}
}

func (r *Requests) Add(cmd commands.Command) int64 {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	defer func () { r.id++ }()
	r.requests[r.id] = cmd
	return r.id
}

func (r *Requests) Remove(id int64) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.requests, id)
}

func (r *Requests) Execute(answer *Answer) {
	cmd := r.requests[answer.Request]
	err := cmd.ParseAnswer(answer.Value)
	if err != nil {
		go cmd.Fail(err.Error())
	}else{
		go r.requests[answer.Request].Execute()
	}
	r.Remove(answer.Request)
}

