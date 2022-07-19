package codewriter

// func Test_MakeCommand(t *testing.T){
// 	type args struct {
// 		filePath string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
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