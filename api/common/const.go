package common

const (
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/id.md#%E6%B8%B8%E6%88%8Fid

	GameIdBH3 = "1" // 崩坏3
	GameIdYS  = "2" // 原神
	GameIdBH2 = "3" // 崩坏学园2
	GameIdWD  = "4" // 未定事件簿
	GameIdDBY = "5" // 大别野
	GameIdSR  = "6" // 崩坏：星穹铁道
	GameIdZZZ = "8" // 绝区零

	GameNameBH3 = "bh3"
	GameNameYS  = "hk4e"
	GameNameBH2 = "bh2"
	GameNameWD  = "nxx"
	GameNameDBY = "plat"
	GameNameSR  = "hkrpg"
	GameNameZZZ = "nap"

	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/id.md#%E6%B8%B8%E6%88%8F%E6%A0%87%E8%AF%86%E7%AC%A6

	GameBizBH3CN = GameNameBH3 + "_" + cn
	GameBizYSCN  = GameNameYS + "_" + cn
	GameBizBH2CN = GameNameBH2 + "_" + cn
	GameBizWDCN  = GameNameWD + "_" + cn
	GameBizSRCN  = GameNameSR + "_" + cn
	GameBizZZZCN = GameNameZZZ + "_" + cn

	cn = "cn"

	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/id.md#%E8%AE%BA%E5%9D%9Bid

	ForumIdSR = "53"

	XRpcSignGame = "x-rpc-signgame"
)
