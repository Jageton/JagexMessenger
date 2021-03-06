package process

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type Dialog struct {
	Id   int64
	Name string
}

func (d *Dialog) String() string {
    return d.Name + ": " + strconv.FormatInt(d.Id, 10)
}

type GetDialogListCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	DialogsList []Dialog
}

func (g *GetDialogListCommand) ParseAnswer(args []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(args, &m)
	if err != nil {
		return err
	}
	return mapstructure.Decode(m, &g)
}

func (g GetDialogListCommand) CommandName() string {
    return "get_dialogs"
}

func ( GetDialogListCommand) IsGlobal() bool {
    return true
}

func (g *GetDialogListCommand) PreExecuteCheck() bool {
	msg := ""
	isOk := true
	switch  {
	case !g.user.IsLogon():
		msg = "You are not logged in"
		isOk = false
	case g.user.InHub():
		msg = "Please leave dialog!"
		isOk = false
	}
	if !isOk {
		g.Fail(msg)
	}
	return isOk
}

func (g *GetDialogListCommand) Execute() {
	str := ""
	for i, d := range g.DialogsList {
		str += strconv.Itoa(i) + ". " + d.String() + "\n"
	}
	SendMessage(g.user.ID(), str)
}

func (g *GetDialogListCommand) GetParsedJSON() []byte {
    m := map[string]int64{
    	"UserId" : g.user.ID(),
    }
    bytes, _ := json.Marshal(m)
    return bytes
}

func ( GetDialogListCommand) ParseData(*tgbotapi.Message) bool {
    return true
}

func ( GetDialogListCommand) Help() string {
    return "/dialogs"
}

