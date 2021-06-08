package googlechat

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

// GoogleChatRepository struct
type GoogleChatRepository struct {
	Client *http.Client
}

func New() appcontext.Component {
	logLocal := config.GetLogger()
	ctx := context.Background()
	data, err := ioutil.ReadFile(config.Values.GoogleChatSAPath)
	if err != nil {
		logLocal.Error(err)
	}
	creds, err := google.CredentialsFromJSON(
		ctx,
		data,
		"https://www.googleapis.com/auth/chat.bot",
	)
	if err != nil {
		logLocal.Error(err)
	}
	return GoogleChatRepository{Client: oauth2.NewClient(ctx, creds.TokenSource)}
}

func (repo GoogleChatRepository) SendMessageGoogleChat(event *chat.DeprecatedEvent, content string) error {
	logLocal := config.GetLogger()
	ctx := context.Background()
	// repo.Client
	service, err := chat.NewService(ctx, option.WithHTTPClient(repo.Client))
	if err != nil {
		logLocal.Errorf("failed to create chat service: %s", err)
		return err
	}
	msgService := chat.NewSpacesMessagesService(service)
	header := chat.CardHeader{
		Subtitle: fmt.Sprintf("Requested from %s", event.Message.Sender.DisplayName),
		Title:    config.Values.AppName,
	}
	keyValue1 := &chat.KeyValue{
		TopLabel:         event.Message.ArgumentText,
		Content:          content,
		ContentMultiline: true,
	}
	widget1 := chat.WidgetMarkup{
		KeyValue: keyValue1,
	}
	section1 := new(chat.Section)
	section1.Widgets = append(section1.Widgets, &widget1)

	// section2 := new(chat.Section)
	// button1 := new(chat.Button)
	// textButton := new(chat.TextButton)
	// textButton.Text = "Sensu Bot URL"
	// onclick := new(chat.OnClick)
	// openlink := new(chat.OpenLink)
	// onclick.OpenLink = openlink
	// openlink.Url = sensuURL
	// widget2 := chat.WidgetMarkup{}
	// widget2.Buttons = append(widget2.Buttons, button1)
	// section2.Widgets = append(section2.Widgets, &widget2)

	card := new(chat.Card)
	card.Header = &header
	card.Sections = []*chat.Section{section1}

	msg := &chat.Message{
		// Text: fmt.Sprintf("response %s", content),
		Cards: []*chat.Card{card},
	}
	fmt.Printf("%v", &msg)
	req := msgService.Create(event.Space.Name, msg)
	if event.ThreadKey != "" {
		logLocal.Debugf("configuring thread key %s", event.ThreadKey)
		req.ThreadKey(event.ThreadKey)
	} else {
		logLocal.Debugf("configuring using thread from message %s", event.Message.Thread.Name)
		msg.Thread = &chat.Thread{
			Name: event.Message.Thread.Name,
		}
	}
	_, err = req.Do()
	if err != nil {
		logLocal.Errorf("failed to create message: %s", err)
		return err
	}
	return nil
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.GoogleChatRepository, New)
	if config.Values.GoogleChatSAPath != "Absent" {
		appcontext.Current.Add(appcontext.GoogleChatRepository, New)
		logLocal := config.GetLogger()
		logLocal.Info("Connected to Gooogle Chat API")
		logLocal.Debug("Gooogle Chat Admin list ", config.Values.GoogleChatAdminList)
	}

}
