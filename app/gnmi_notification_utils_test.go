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

	nb, err := NewNotificationBuilder(targetName, "JSON_IETF", prefix, path)
	if err != nil {
		t.Errorf("Error when creating NotificationBuilder %q", err)
	}

	expectedNotification, err := nb.BuildNotification()
	if err != nil {
		t.Errorf("Error when Building Notification %q", err)
	}

	notification := gnmi.Notification{
		Prefix: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "state"}, {Name: "aaa"},
			}},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{
					Elem: []*gnmi.PathElem{},
				},
				Val: &gnmi.TypedValue{
					Value: &gnmi.TypedValue_JsonIetfVal{
						JsonIetfVal: []byte(`{"radius":14}`),
					},
				},
			},
		},
	}

	if err = nb.AppendNotification(&notification); err != nil {
		t.Errorf("Error when appending Notification %q", err)
	}

	updatedNotification, err := nb.BuildNotification()
	if err != nil {
		t.Errorf("Error when Building Notification %q", err)
	}

	expectedNotification.Update[0].Val = &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonIetfVal{JsonIetfVal: []byte(`{"state":{"aaa":{"radius":14}}}`)}}
	expectedNotification.Timestamp = updatedNotification.GetTimestamp()
	if !reflect.DeepEqual(updatedNotification, expectedNotification) {
		t.Errorf("got %q, expected %q", updatedNotification, expectedNotification)
	}
}
