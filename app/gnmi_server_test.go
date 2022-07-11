package app

import (
	"testing"

	"github.com/openconfig/gnmi/proto/gnmi"
)

func TestPathSortReorder(t *testing.T) {
	// First case: first /a/b, then /a in the Update part of a SetRequest
	req := gnmi.SetRequest{}
	target := "tgt1"
	refDel := map[string][]string{target: []string{}}
	refRepl := map[string][]string{target: []string{"UUID1", "UUID2"}}
	refUpd := map[string][]string{target: []string{}}
	req.Replace = append(req.Replace, &gnmi.Update{})
	req.Replace[0].Path = &gnmi.Path{Elem: []*gnmi.PathElem{&gnmi.PathElem{Name: "a"}, &gnmi.PathElem{Name: "b"}}}
	req.Replace = append(req.Replace, &gnmi.Update{})
	req.Replace[1].Path = &gnmi.Path{Elem: []*gnmi.PathElem{&gnmi.PathElem{Name: "a"}}}

	pathSort(&req, refDel, refRepl, refUpd, target)

	// Ensure that the request itself was changed
	if len(req.Replace[0].Path.Elem) != 1 {
		t.Errorf("got %d, expected %d", len(req.Replace[0].Path.Elem), 1)
	}
	if req.Replace[0].Path.Elem[0].Name != "a" {
		t.Errorf("got %s, expected %s", req.Replace[0].Path.Elem[0].Name, "a")
	}
	if len(req.Replace[1].Path.Elem) != 2 {
		t.Errorf("got %d, expected %d", len(req.Replace[1].Path.Elem), 2)
	}
	if req.Replace[1].Path.Elem[0].Name != "a" {
		t.Errorf("got %s, expected %s", req.Replace[1].Path.Elem[0].Name, "a")
	}
	if req.Replace[1].Path.Elem[1].Name != "b" {
		t.Errorf("got %s, expected %s", req.Replace[1].Path.Elem[1].Name, "b")
	}

	// Ensure that all references are updated correctly
	if refRepl[target][0] != "UUID2" {
		t.Errorf("got %s, expected %s", refRepl[target][0], "UUID2")
	}
	if refRepl[target][1] != "UUID1" {
		t.Errorf("got %s, expected %s", refRepl[target][0], "UUID1")
	}

	// Check that no new target was added
	if len(refUpd) != 1 {
		t.Errorf("got %d, expected %d", len(refUpd), 1)
	}
	if _, found := refUpd[target]; !found {
		t.Errorf("did not have expected key %s", target)
	}
	if len(refRepl) != 1 {
		t.Errorf("got %d, expected %d", len(refRepl), 1)
	}
	if _, found := refRepl[target]; !found {
		t.Errorf("did not have expected key %s", target)
	}
	if len(refDel) != 1 {
		t.Errorf("got %d, expected %d", len(refDel), 1)
	}
	if _, found := refDel[target]; !found {
		t.Errorf("did not have expected key %s", target)
	}

	// Check that no entries for that target have been added
	if len(refUpd[target]) != 0 {
		t.Errorf("got %d, expected %d", len(refUpd[target]), 0)
	}
	if len(refDel[target]) != 0 {
		t.Errorf("got %d, expected %d", len(refDel[target]), 0)
	}
	if len(refRepl[target]) != 2 {
		t.Errorf("got %d, expected %d", len(refRepl[target]), 2)
	}
}

func TestPathSortSame(t *testing.T) {
	// Second case: first /a, then /a/b in the Update part of a SetRequest
	req := gnmi.SetRequest{}
	target := "tgt1"
	refDel := map[string][]string{target: []string{}}
	refRepl := map[string][]string{target: []string{"UUID1", "UUID2"}}
	refUpd := map[string][]string{target: []string{}}
	req.Replace = append(req.Replace, &gnmi.Update{})
	req.Replace[0].Path = &gnmi.Path{Elem: []*gnmi.PathElem{&gnmi.PathElem{Name: "a"}}}
	req.Replace = append(req.Replace, &gnmi.Update{})
	req.Replace[1].Path = &gnmi.Path{Elem: []*gnmi.PathElem{&gnmi.PathElem{Name: "a"}, &gnmi.PathElem{Name: "b"}}}

	pathSort(&req, refDel, refRepl, refUpd, target)

	// Ensure that the request itself was changed
	if len(req.Replace[0].Path.Elem) != 1 {
		t.Errorf("got %d, expected %d", len(req.Replace[0].Path.Elem), 1)
	}
	if req.Replace[0].Path.Elem[0].Name != "a" {
		t.Errorf("got %s, expected %s", req.Replace[0].Path.Elem[0].Name, "a")
	}
	if len(req.Replace[1].Path.Elem) != 2 {
		t.Errorf("got %d, expected %d", len(req.Replace[1].Path.Elem), 2)
	}
	if req.Replace[1].Path.Elem[0].Name != "a" {
		t.Errorf("got %s, expected %s", req.Replace[1].Path.Elem[0].Name, "a")
	}
	if req.Replace[1].Path.Elem[1].Name != "b" {
		t.Errorf("got %s, expected %s", req.Replace[1].Path.Elem[1].Name, "b")
	}

	// Ensure that all references are updated correctly
	if refRepl[target][0] != "UUID1" {
		t.Errorf("got %s, expected %s", refRepl[target][0], "UUID1")
	}
	if refRepl[target][1] != "UUID2" {
		t.Errorf("got %s, expected %s", refRepl[target][0], "UUID2")
	}

	// Check that no new target was added
	if len(refUpd) != 1 {
		t.Errorf("got %d, expected %d", len(refUpd), 1)
	}
	if _, found := refUpd[target]; !found {
		t.Errorf("did not have expected key %s", target)
	}
	if len(refRepl) != 1 {
		t.Errorf("got %d, expected %d", len(refRepl), 1)
	}
	if _, found := refRepl[target]; !found {
		t.Errorf("did not have expected key %s", target)
	}
	if len(refDel) != 1 {
		t.Errorf("got %d, expected %d", len(refDel), 1)
	}
	if _, found := refDel[target]; !found {
		t.Errorf("did not have expected key %s", target)
	}

	// Check that no entries for that target have been added
	if len(refUpd[target]) != 0 {
		t.Errorf("got %d, expected %d", len(refUpd[target]), 0)
	}
	if len(refDel[target]) != 0 {
		t.Errorf("got %d, expected %d", len(refDel[target]), 0)
	}
	if len(refRepl[target]) != 2 {
		t.Errorf("got %d, expected %d", len(refRepl[target]), 2)
	}
}
