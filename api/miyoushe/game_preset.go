package miyoushe

import (
	"strconv"

	"github.com/starudream/go-lib/core/v2/codec/json"
)

var AllGames = json.MustUnmarshalTo[[]*Game](`
[
  {
    "id": 1,
    "name": "崩坏3",
    "en_name": "bh3",
    "op_name": "bh3"
  },
  {
    "id": 2,
    "name": "原神",
    "en_name": "ys",
    "op_name": "hk4e"
  },
  {
    "id": 3,
    "name": "崩坏学园2",
    "en_name": "bh2",
    "op_name": "bh2"
  },
  {
    "id": 4,
    "name": "未定事件簿",
    "en_name": "wd",
    "op_name": "nxx"
  },
  {
    "id": 5,
    "name": "大别野",
    "en_name": "dby",
    "op_name": "plat"
  },
  {
    "id": 6,
    "name": "崩坏：星穹铁道",
    "en_name": "sr",
    "op_name": "hkrpg"
  },
  {
    "id": 8,
    "name": "绝区零",
    "en_name": "zzz",
    "op_name": "nap"
  }
]
`)

var AllGamesById = func() map[string]*Game {
	m := map[string]*Game{}
	for i := range AllGames {
		m[strconv.Itoa(AllGames[i].Id)] = AllGames[i]
	}
	return m
}()
