package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "4で割り切れる, 100で割り切れる, 400で割り切れるtest", args: args{year: 2000}, want: true},
		{name: "4で割り切れる, 100で割り切れる, 400で割り切れるの境界値(-1)test", args: args{year: 1999}, want: false},
		{name: "4で割り切れる, 100で割り切れる, 400で割り切れるの境界値(+1)test", args: args{year: 2001}, want: false},
		{name: "4で割り切れる, 100で割り切れる, 400で割り切れないtest", args: args{year: 100}, want: false},
		{name: "4で割り切れる, 100で割り切れる, 400で割り切れないの境界値(-1)test", args: args{year: 99}, want: false},
		{name: "4で割り切れる, 100で割り切れる, 400で割り切れないの境界値(+1)test", args: args{year: 101}, want: false},
		{name: "4で割り切れる, 100で割り切れない, 400で割り切れないtest", args: args{year: 404}, want: true},
		{name: "4で割り切れる, 100で割り切れない, 400で割り切れないの境界値(-1)test", args: args{year: 403}, want: false},
		{name: "4で割り切れる, 100で割り切れない, 400で割り切れないの境界値(+1)test", args: args{year: 405}, want: false},
		{name: "4で割り切れない, 100で割り切れない, 400で割り切れないtest", args: args{year: 17}, want: false},
		{name: "4で割り切れない, 100で割り切れない, 400で割り切れないの境界値(-1)test", args: args{year: 16}, want: true},
		{name: "4で割り切れない, 100で割り切れない, 400で割り切れないの境界値(+1)test", args: args{year: 18}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeapYear(tt.args.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeapYear_useAssert(t *testing.T) {
	// 4で割り切れる, 100で割り切れる, 400で割り切れる
	assert.True(t, IsLeapYear(2000), "4で割り切れる, 100で割り切れる, 400で割り切れるtest")
	assert.False(t, IsLeapYear(1999), "4で割り切れる, 100で割り切れる, 400で割り切れる境界値(-1)test")
	assert.False(t, IsLeapYear(2001), "4で割り切れる, 100で割り切れる, 400で割り切れる境界値(+1)test")

	// 4で割り切れる, 100で割り切れる, 400で割り切れない
	assert.False(t, IsLeapYear(100), "4で割り切れる, 100で割り切れる, 400で割り切れないtest")
	assert.False(t, IsLeapYear(99), "4で割り切れる, 100で割り切れる, 400で割り切れない境界値(-1)test")
	assert.False(t, IsLeapYear(101), "4で割り切れる, 100で割り切れる, 400で割り切れない境界値(+1)test")

	// 4で割り切れる, 100で割り切れない, 400で割り切れない
	assert.True(t, IsLeapYear(404), "4で割り切れる, 100で割り切れない, 400で割り切れないtest")
	assert.False(t, IsLeapYear(403), "4で割り切れる, 100で割り切れない, 400で割り切れない境界値(-1)test")
	assert.False(t, IsLeapYear(405), "4で割り切れる, 100で割り切れない, 400で割り切れない境界値(+1)test")

	// 4で割り切れない, 100で割り切れない, 400で割り切れない
	assert.False(t, IsLeapYear(17), "4で割り切れない, 100で割り切れない, 400で割り切れないtest")
	assert.True(t, IsLeapYear(16), "4で割り切れない, 100で割り切れない, 400で割り切れない境界値(-1)test")
	assert.False(t, IsLeapYear(18), "4で割り切れない, 100で割り切れない, 400で割り切れない境界値(+1)test")
}

func TestMain(m *testing.M) {
	m.Run()
}
