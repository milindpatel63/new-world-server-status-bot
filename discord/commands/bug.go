package commands

type BugCommand struct {
}

func (b BugCommand) Check(parameters []string) bool {
	return parameters[0] == "bug" || parameters[0] == "Bug" || parameters[0] == "BUG"
}

func (b BugCommand) Handle(parameters []string) (string, error) {
	return "If you have encountered an issue with this bot," +
		"please open an issue on github at <https://github.com/Siiiimon/new-world-server-status-bot>." +
		" Otherwise contact Dot#8089 on Discord.", nil
}
