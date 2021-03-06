package chessUtils

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGetRoomPass(t *testing.T) {
	a := GetRoomPass(5)
	t.Logf("a :%v-%v", a)
	b := GetRoomPass(3)
	t.Logf("b :%v-%v", b)
	c := GetRoomPass(3)
	t.Logf("c :%v-%v", c)
	d := GetRoomPass(9)
	t.Logf("d :%v-%v", d)
}

func TestGetRoomPassV2(t *testing.T) {
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(12))
	t.Log(GetRoomPassV2(12))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(39))
	t.Log(GetRoomPassV2(99))
	t.Log(GetRoomPassV2(1))
	t.Log(GetRoomPassV2(2))
	t.Log(GetRoomPassV2(3))
}

//测试： 洗牌
func TestShuffle(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		time.AfterFunc(time.Second/100, func() {
			defer wg.Done()
			arr := Xipai(0, 51)
			fmt.Println(arr)
		})
		wg.Wait()
	}

}

func TestXipai(t *testing.T) {
	list := Xipai(0, 3)
	t.Log(list)
	t.Log(list[len(list)-1])
	list = list[:len(list)-1]
	t.Log(list)
	return
	for i := 0; i < 10000000; i++ {
		list := Xipai(0, 51)
		m := map[int32]int{}

		for _, v := range list {
			if _, ok := m[v]; ok {
				m[v]++
				t.Log(list, m)
				return
			} else {
				m[v] = 1
			}
		}
	}

}
