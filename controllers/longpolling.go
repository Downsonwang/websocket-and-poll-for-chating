package controllers

import "charroom/models"

type LongPollingController struct {
	baseController
}

func (this *LongPollingController) Join() {
	//Safe check
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	//Join chat room
	Join(uname, nil)

	this.TplName = "longpolling.html"
	this.Data["IsLongPolling"] = true
	this.Data["userName"] = uname
}

// Post method handles receive messages requests for LongPollingController.
func (this *LongPollingController) Post() {
	this.TplName = "longpolling.html"

	uname := this.GetString("uname")
	content := this.GetString("content")

	if len(uname) == 0 || len(content) == 0 {
		return
	}
	publish <- newEvent(models.EVENT_MESSAGE, uname, content)
}

// Fetch method handles fetch archives requests for LongPollingController.

func (this *LongPollingController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	if err != nil {
		return
	}
	events := models.GetEvents(int(lastReceived))
	if len(events) > 0 {
		this.Data["json"] = events
		this.ServeJSON()
		return
	}

	//Wait for new message
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch
	this.Data["json"] = models.GetEvents(int(lastReceived))
	this.ServeJSON()
}
