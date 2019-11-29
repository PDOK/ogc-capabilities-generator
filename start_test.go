package main

import (
	"fmt"
	"pdok-capabilities-gen/builder"
	"testing"
)

func Test_build(t *testing.T) {
	type args struct {
		builder builder.Builder
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_envString(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := envString(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("envString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		typeStr    string
		versionStr string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"wms_is_ok", args{
			typeStr:    "wms",
			versionStr: "1.3.0",
		}, nil},
		{"wms_is_not_ok", args{
			typeStr:    "wmds",
			versionStr: "1.3.0",
		}, fmt.Errorf("service type unknown : 'wmds', possible values are : wfs, wms, wmts")},
		{"wMts_is_ok", args{
			typeStr:    "wMts",
			versionStr: "1.0.0",
		}, nil},
		{"wFS_is_ok", args{
			typeStr:    "wFS",
			versionStr: "2.0.0",
		}, nil},
		{"wms_version_is_not_ok", args{
			typeStr:    "wms",
			versionStr: "123.0.0",
		}, fmt.Errorf("version unknown for wms['123.0.0'], possible versions are 1.3.0")},
		{"wfs_version_is_not_ok", args{
			typeStr:    "wfs",
			versionStr: "123.0.0",
		}, fmt.Errorf("version unknown for wfs['123.0.0'], possible versions are 1.1.0, 2.0.0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.args.typeStr, tt.args.versionStr)

			if err == nil && tt.want == nil {
				return
			}

			if tt.want != nil && err == nil || err != nil && tt.want == nil {
				t.Errorf("Got: \"%v\", but want: \"%v\"", err, tt.want)
				return
			}

			if err.Error() != tt.want.Error() {
				t.Errorf("Got: \"%v\", but want: \"%v\"", err, tt.want)
			}
		})
	}
}
