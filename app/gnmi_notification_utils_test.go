package app

import (
	"reflect"
	"testing"

	"github.com/openconfig/gnmi/proto/gnmi"
)

func TestNotificationBuilder(t *testing.T) {
	var targetName string = "127.0.0.1:57400"
	var prefix *gnmi.Path
	var path *gnmi.Path = &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "state"}, {Name: "aaa"},
		}}

	var nb *NotificationBuilder
	var err error
	nb, err = NewNotificationBuilder(targetName, prefix, path)
	if err != nil {
		t.Errorf("Error when creating NotificationBuilder %q", err)
	}

	var expectedNotification *gnmi.Notification = nb.BuildNotification()

	notification := gnmi.Notification{
		Prefix: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "state"}, {Name: "aaa"},
			}},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "radius"},
					},
				},
				Val: &gnmi.TypedValue{
					Value: &gnmi.TypedValue_JsonVal{
						JsonVal: []byte(`14`),
					},
				},
			},
		},
	}

	nb.AppendNotification(&notification)

	updatedNotification := nb.BuildNotification()

	expectedNotification.Update[0].Val = &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: []byte(`{"radius":"14"}`)}}
	expectedNotification.Timestamp = updatedNotification.GetTimestamp()
	if !reflect.DeepEqual(updatedNotification, expectedNotification) {
		t.Errorf("got %q, expected %q", updatedNotification, expectedNotification)
	}
}
