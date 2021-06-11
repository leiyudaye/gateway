/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-11 22:28:42
 * @LastEditors: lly
 * @LastEditTime: 2021-06-12 03:02:33
 */
package redis

import (
	"encoding/json"
	"testing"
)

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSet(t *testing.T) {
	type args struct {
		conf string
		key  string
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RedisSet",
			args: args{
				conf: "127.0.0.1:6379",
				key:  "redis_demo_user",
				v:    TestData{Name: "leiyudaydafdsfsdfsdfdsfsdfwewer", Age: 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			data, _ := json.Marshal(tt.args.v)
			t.Log(data)
			if err := Set(tt.args.conf, tt.args.key, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		conf string
		key  string
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RedisSet",
			args: args{
				conf: "127.0.0.1:6379",
				key:  "redis_demo_user",
				v:    &TestData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Get(tt.args.conf, tt.args.key, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			data := tt.args.v.(*TestData)
			t.Log(data.Name, data.Age)
		})
	}
}
