package goqueue

import (
	"testing"
	"time"
)

var queue = NewInstance(RedisConfig{Host: "localhost", Port: 6379, DB: 0})

var key = "charactor"

func TestPush(test *testing.T) {
	_, b := queue.Push(key, "卡布达")
	_, b = queue.Push(key, "孙悟饭")
	_, b = queue.Push(key, "鲨鱼辣椒")
	_, b = queue.Push(key, "蝎子莱莱")
	if b != nil {
		test.Fatal(b)
	}
}

func TestGetLength(test *testing.T) {
	a, b := queue.GetLength(key)
	if b != nil {
		test.Fatal(b)
	}
	if a != 4 {
		test.Fatal("GetLength获取到的数目错误,期望 4,得到", a)
	}
}

func TestGetRange(test *testing.T) {
	_, b := queue.GetRange(key, 0, 10)
	if b != nil {
		test.Fatal(b)
	}
}

func TestPop(test *testing.T) {
	for {
		start := time.Now()
		_, b := queue.Pop(1*time.Second, key)
		end := time.Now()
		subS := end.Sub(start)
		if subS.Seconds() >= 1.0 { // 判断是否超时
			break
		}
		if b != nil {
			test.Error("damn!", b)
			break
		}
	}
}

func TestGetFinshedLength(test *testing.T) {
	a, b := queue.GetLength(key)
	if b != nil {
		test.Fatal(b)
	}
	if a != 0 {
		test.Fatal("GetLength获取到的数目错误,期望 0,得到", a)
	}
}
