package app

import (
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/openconfig/gnmi/path"
	"github.com/openconfig/gnmi/proto/gnmi"
)

type NotificationBuilder struct {
	TargetName   string
	Prefix       *gnmi.Path
	CompletePath []string
	JSONObj      *gabs.Container
}

func NewNotificationBuilder(notifTargetName string, notifPrefix *gnmi.Path, notifPath *gnmi.Path) (*NotificationBuilder, error) {
	notifBuilder := new(NotificationBuilder)

	notifBuilder.TargetName = notifTargetName

	if notifPrefix == nil {
		notifPrefix = new(gnmi.Path)
	} else {
		if notifPrefix.GetTarget() == "" {
			notifPrefix.Target = notifTargetName
		}
	}
	notifBuilder.Prefix = notifPrefix

	var fp []string
	var err error
	fp, err = path.CompletePath(notifPrefix, notifPath)
	if err != nil {
		return nil, err
	}
	notifBuilder.CompletePath = fp

	notifBuilder.JSONObj = gabs.New()

	return notifBuilder, err
}

func (nb NotificationBuilder) AppendNotification(notification *gnmi.Notification) {
	if notification.Prefix != nil {
		var path []string
		for _, prefixPathElmt := range notification.Prefix.Elem {
			path = append(path, prefixPathElmt.Name)
		}

		var value string
		for _, updt := range notification.Update {
			path = append(path, updt.Path.Elem[0].Name)
			value = string(updt.Val.GetJsonVal())
		}

		jsonObj := gabs.New()
		jsonObj.Set(value, path[len(nb.CompletePath):]...)

		nb.JSONObj.Merge(jsonObj)
	}
}

func (nb NotificationBuilder) GetCompletePath() []string {
	return nb.CompletePath
}

func (nb NotificationBuilder) BuildNotification() *gnmi.Notification {
	notification := gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    nb.Prefix,
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: strings.Join(nb.CompletePath, "/")},
					},
				},
				Val: &gnmi.TypedValue{
					Value: &gnmi.TypedValue_JsonVal{
						JsonVal: []byte(nb.JSONObj.String()),
					},
				},
			},
		},
	}

	return &notification
}
