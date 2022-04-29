package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/openconfig/gnmi/path"
	"github.com/openconfig/gnmi/proto/gnmi"
)

type NotificationBuilder struct {
	TargetName   string
	Encoding     string
	Prefix       *gnmi.Path
	CompletePath []string
	PathElem     []*gnmi.PathElem
	JSONObj      *gabs.Container
}

func NewNotificationBuilder(notifTargetName string, encoding string, notifPrefix *gnmi.Path, notifPath *gnmi.Path) (*NotificationBuilder, error) {
	notifBuilder := new(NotificationBuilder)

	notifBuilder.TargetName = notifTargetName
	notifBuilder.Encoding = encoding

	if notifPrefix == nil {
		notifPrefix = new(gnmi.Path)
	} else {
		if notifPrefix.GetTarget() == "" {
			notifPrefix.Target = notifTargetName
		}
	}
	notifBuilder.Prefix = notifPrefix

	fp, err := path.CompletePath(notifPrefix, notifPath)
	if err != nil {
		return nil, err
	}
	notifBuilder.CompletePath = fp
	notifBuilder.PathElem = notifPath.Elem

	notifBuilder.JSONObj = gabs.New()

	return notifBuilder, err
}

func (nb NotificationBuilder) AppendNotification(notification *gnmi.Notification) error {
	if notification.Prefix != nil {
		var path []string = formatPathElem(notification.Prefix.Elem)

		for _, updt := range notification.Update {
			path = append(path, formatPathElem(updt.Path.Elem)...)

			var value []byte
			switch strings.ToUpper(nb.Encoding) {
				case "JSON_IETF":
					value = updt.Val.GetJsonIetfVal()
				case "JSON":
					value = updt.Val.GetJsonVal()
				default:
					return fmt.Errorf("Invalid encoding type %q, possible values are [ JSON, JSON_IETF ]", nb.Encoding)
			}

			jsonObj := gabs.New()
			jsonValue, err := gabs.ParseJSON(value)
			if err != nil {
				return err
			}
			jsonObj.Set(jsonValue.Data(), path...)
			nb.JSONObj.Merge(jsonObj)
		}
	}

	return nil
}

func (nb NotificationBuilder) GetCompletePath() []string {
	return nb.CompletePath
}

func (nb NotificationBuilder) BuildNotification() (*gnmi.Notification, error) {
	switch strings.ToUpper(nb.Encoding) {
		case "JSON_IETF":
			return (&gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    nb.Prefix,
				Update: []*gnmi.Update{
					{
						Path: &gnmi.Path{
							Elem: nb.PathElem,
						},
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_JsonIetfVal{
								JsonIetfVal: nb.JSONObj.EncodeJSON(),
							},
						},
					},
				},
			}), nil
		case "JSON":
			return (&gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    nb.Prefix,
				Update: []*gnmi.Update{
					{
						Path: &gnmi.Path{
							Elem: nb.PathElem,
						},
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_JsonVal{
								JsonVal: nb.JSONObj.EncodeJSON(),
							},
						},
					},
				},
			}), nil
		default:
			return nil, fmt.Errorf("Invalid encoding type %q, possible values are [ JSON, JSON_IETF ]", nb.Encoding)
	}
}

func formatPathElem(pathElems []*gnmi.PathElem) []string {
	var path []string

	for _, elmt := range pathElems {
		var keys string
		for k, v := range elmt.Key {
			keys += "["+k+"="+v+"]"
		}
		if len(keys) > 0 {
			path = append(path, elmt.Name+keys)
		} else {
			path = append(path, elmt.Name)
		}
	}

	return path
}
