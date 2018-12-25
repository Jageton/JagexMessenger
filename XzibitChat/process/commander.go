package process

import "XzibitChat/commands"

type Commander map[string]commands.Command


var Commands = Commander{
	"create" : &CreateDialogCommand{},
	"end" : &EndSession{},
	"enter" : &EnterDialogCommand{},
	"exit" : &ExitDialogCommand{},
	"dialogs" : &GetDialogListCommand{},
	"invite" : &InviteUserCommand{},
	"leave" : &LeaveDialogCommand{},
	"login" : &LoginCommand{},
	"regr" : &RegistrationCommand{},
	"send" : &SendMessageCommand{},
	"help" : &HelpCommand{},
	"status" : &StatusCommand{},
}


func GetCommand(cmdName string) (commands.Command, bool){
	cmd, ok := Commands[cmdName]
	if !ok {
		return nil, ok
	}
	newCmd := cmd
	return newCmd, true
}
