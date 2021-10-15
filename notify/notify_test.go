package notify

// TODO: rework this
//import (
//	"github.com/google/go-cmp/cmp"
//	"testing"
//)
//
//func TestGetNewlyOnlineServers(t *testing.T) {
//	type test struct {
//		name        string
//		inPrev      map[string]string
//		inCurrent   map[string]string
//		wantServers []string
//		wantChanged bool
//	}
//
//	tests := []test{
//		{
//			name:        "Offline -> Online",
//			inPrev:      map[string]string{"Delphnius": "Offline"},
//			inCurrent:   map[string]string{"Delphnius": "Online"},
//			wantServers: []string{"Delphnius"},
//			wantChanged: true,
//		},
//		{
//			name:        "Online -> Online",
//			inPrev:      map[string]string{"Delphnius": "Online"},
//			inCurrent:   map[string]string{"Delphnius": "Online"},
//			wantServers: nil,
//			wantChanged: false,
//		},
//		{
//			name:        "Online -> Offline",
//			inPrev:      map[string]string{"Delphnius": "Online"},
//			inCurrent:   map[string]string{"Delphnius": "Offline"},
//			wantServers: nil,
//			wantChanged: false,
//		},
//	}
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			gotServers, gotChanged := getNewlyOnlineServers(tc.inPrev, tc.inCurrent)
//			if !cmp.Equal(tc.wantServers, gotServers) {
//				t.Errorf("\nexpected: %v\ngot: %v\nDiff:\n%s", tc.wantServers, gotServers, cmp.Diff(tc.wantServers, gotServers))
//			}
//			if gotChanged != tc.wantChanged {
//				t.Fatalf("\nexpected: %v\ngot: %v\n", tc.wantChanged, gotChanged)
//			}
//		})
//	}
//}
