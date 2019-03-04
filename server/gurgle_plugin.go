package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// GurglePlugin plugin to gurgling things
type GurglePlugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

func (p *GurglePlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {

	p.API.LogDebug("GO GO gurgle-plugin")

	dir, err := os.Getwd()
	if err == nil {

		filePath := fmt.Sprintf("%s/client/plugins/com.gohabitate.gurgle-plugin/index.html", dir)
		p.API.LogDebug(filePath)
		data, err := ioutil.ReadFile(filePath)
		if err == nil {
			w.Write(data)
		}
	}
}

// MessageHasBeenPosted Message post event handler
func (p *GurglePlugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	configuration := p.getConfiguration()

	p.API.LogDebug(configuration.channelId)
	p.API.LogDebug(post.ChannelId)
	// Ignore posts not in the configured channel
	if post.ChannelId != configuration.channelId {
		return
	}

	// Ignore posts this plugin made.
	if sentByPlugin, _ := post.Props["sent_by_plugin"].(bool); sentByPlugin {
		return
	}

	// Ignore posts without a plea for help.
	if !strings.Contains(post.Message, "gurgle") {
		return
	}

	p.API.CreatePost(&model.Post{
		ChannelId: configuration.channelId,
		Message:   "Do you even Gurgle Bro?",
		Props: map[string]interface{}{
			"sent_by_plugin": true,
		},
	})

	p.API.SendEphemeralPost(post.UserId, &model.Post{
		ChannelId: configuration.channelId,
		Message:   "Good gurgley murgley?!!!",
		Props: map[string]interface{}{
			"sent_by_plugin": true,
		},
	})
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
