package bot

import (
	"fmt"
	"strings"
)

const (
	joinCommand = "join"
	joinUsage   = "Usage: !join #channel pass"
	joinMessage = "Hello humans o/ My master %s called me here"
)

func join(c *Cmd, channel string, user *User, conn connection) {
	channelToJoin := strings.TrimSpace(c.FullArg)
	if channelToJoin == "" {
		conn.Privmsg(channel, joinUsage)
	} else {
		conn.Join(channelToJoin)
		conn.Privmsg(c.Args[0], fmt.Sprintf(joinMessage, user.Nick))
	}
}
