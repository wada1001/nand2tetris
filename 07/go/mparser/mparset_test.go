package mparser

import (
	"reflect"
	"testing"
)

// func Test_MakeCommand(t *testing.T){
// 	type args struct {
// 		filePath string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		wantErr bool
// 		want *Parser
// 	}{
// 		// TODO not exists in testing??
// 		{"success", args{"../MemoryAccess/BasicTest/BasicTest.vm"}, true, &Parser{}},
// 		{"not_exsits", args{""}, false, nil},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			_, err := MakeParser(test.args.filePath)
// 			if err != nil {
// 				t.Errorf("err occured = %v, want %v", err, test.wantErr)
// 				// t.Errorf("not expected err = %v, want %v", err, test.wantErr)
// 			}
// 		})
// 	}
// }

func Test_Command_MakeCommand(t *testing.T){
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		wantErr bool
		want *Command
	}{
		{"pattern_1", args{"push constant 10"}, false, &Command{
			command: "push",
			arg1: "constant",
			arg2: "10",
		}},
		{"pattern_2", args{"push constant"}, false, &Command{
			command: "push",
			arg1: "constant",
			arg2: "",
		}},
		{"pattern_3", args{"push"}, false, &Command{
			command: "push",
			arg1: "",
			arg2: "",
		}},
		{"empty", args{""}, true, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := MakeCommand(test.args.command)
			if test.wantErr && err == nil {
				t.Errorf("err occured. err = %v", err)	
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("not expected val = %v, want %v", got, test.want)
			}
		})
	}
}

