package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/42wim/matterbridge/bridge/config"
	"github.com/42wim/matterbridge/gateway"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
	r        *gateway.Router
	userid   string
	teamid   string
	channels []string
}

func (p *Plugin) OnActivate() error {
	fmt.Println("reading config")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	cfg := config.NewConfig(exPath + "/" + "matterbridge.toml")
	fmt.Println("matterbridge config read")
	r, err := gateway.NewRouter(cfg)
	if err != nil {
		log.Fatalf("Starting gateway failed: %s", err)
	}
	p.r = r
	go func() {
		err = r.Start()
		if err != nil {
			log.Fatalf("Starting gateway failed: %s", err)
		}
		fmt.Printf("Gateway(s) started succesfully. Now relaying messages")
		select {}
	}()
	go func() {
		// wait until activation is done, otherwise API doesn't seem to work?
		time.Sleep(time.Second)
		user, err := p.API.GetUserByUsername(cfg.Mattermost["plugin"].Login)
		if err != nil {
			fmt.Println("username", err.Error())
		}
		p.userid = user.Id
		fmt.Println("found userid", p.userid)
		team, err := p.API.GetTeamByName(cfg.Mattermost["plugin"].Team)
		if err != nil {
			fmt.Println("team", err.Error())
		}
		p.teamid = team.Id
		fmt.Println("found teamid", p.teamid)
	}()
	go func() {
		for msg := range p.r.MattermostPlugin {
			fmt.Printf("Got message %#v\n", msg)
			channel, _ := p.API.GetChannelByName(p.teamid, msg.Channel, false)
			props := make(map[string]interface{})
			props["matterbridge_"+p.userid] = true
			post := &model.Post{UserId: p.userid, ChannelId: channel.Id, Message: msg.Username + msg.Text, Props: props}
			fmt.Printf("Posting %#v\n", post)
			p.API.CreatePost(post)
		}
	}()
	return nil
}

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	if post.Props != nil {
		if _, ok := post.Props["matterbridge_"+p.userid].(bool); ok {
			fmt.Println("sent by matterbridge, ignoring")
			return
		}
	}
	ch, _ := p.API.GetChannel(post.ChannelId)
	u, _ := p.API.GetUser(post.UserId)
	msg := config.Message{Username: u.Nickname, UserID: post.UserId, Channel: ch.Name, Text: post.Message, ID: post.Id, Account: "mattermost.plugin", Protocol: "mattermost", Gateway: "plugin"}
	fmt.Printf("sending message %#v", msg)
	p.r.Message <- msg
}

func main() {
	plugin.ClientMain(&Plugin{})
}
