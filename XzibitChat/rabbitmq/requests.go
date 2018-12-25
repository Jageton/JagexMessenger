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
		id: 5,
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
	cmd, ok:= r.requests[answer.Id]
	if !ok {
		return
	}
	err := cmd.ParseAnswer([]byte(answer.Data))
	if err != nil {
		go cmd.Fail(err.Error())
	}else{
		go r.requests[answer.Id].Execute()
	}
	r.Remove(answer.Id)
}

