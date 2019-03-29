package service

import (
	"reflect"
	"testing"
)

func TestExractMetaData(t *testing.T) {
	t1 := "++ Battalion Detachment +5CP (Imperium - Adeptus Custodes)[79 PL, 1552pts, -1CP] ++No Force Org Slot +Open the Vaults (1 Relic) [-1CP]Use Beta RulesHQ +++ Total: [101 PL, 5CP, 2000pts] ++"
	t2 := "++ Battalion Detachment +5CP (Imperium - Adeptus Mechanicus) (Imperium - Adeptus Mechanicus) (Imperium - Adeptus Custodes)[79 PL, 1552pts, -1CP] ++No Force Org Slot +Open the Vaults (1 Relic) [-1CP]Use Beta RulesHQ +++ Total: [101 PL, 5CP, 2000pts] ++"
	meta := ListMetaData{CP: 5, PL: 101, PTS: 2000, Faction: "Adeptus Custodes"}
	result := ListMetaData{CP: 5, PL: 101, PTS: 2000, Faction: "Adeptus Mechanicus"}
	type args struct {
		list string
	}
	extract := []struct {
		name string
		args args
		want ListMetaData
	}{
		{"should return the list meta data", args{list: t1}, meta},
		{"should return the faction with the most occurences data", args{list: t2}, result},
	}

	for _, tt := range extract {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExractMetaData(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExractMetaData() = %v, want %v", got, tt.want)
			}
		})
	}
}
