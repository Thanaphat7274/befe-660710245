package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Lottery struct {
	Date   string          `json:"date"`
	Prizes []PrizeCategory `json:"prizes"`
}
type PrizeCategory struct {
	Tier        string   `json:"tier"`
	PrizeAmount int      `json:"prizeamount"`
	RewardCount int      `json:"rewardcount"`
	Numbers     []string `json:"numbers"`
}

var lottery = []Lottery{
	{Date: "2025-08-16",
		Prizes: []PrizeCategory{
			{Tier: "รางวัลที่ 1",
				PrizeAmount: 6000000,
				RewardCount: 1,
				Numbers:     []string{"994865"},
			},
			{Tier: "เลขหน้า 3 ตัว",
				PrizeAmount: 4000,
				RewardCount: 2,
				Numbers:     []string{"247", "602"},
			},
			{Tier: "เลขท้าย 3 ตัว",
				PrizeAmount: 4000,
				RewardCount: 2,
				Numbers:     []string{"834", "989"},
			},
			{Tier: "เลขท้าย 2 ตัว",
				PrizeAmount: 2000,
				RewardCount: 1,
				Numbers:     []string{"63"},
			},
			{Tier: "รางวัลที่ 2",
				PrizeAmount: 200000,
				RewardCount: 5,
				Numbers:     []string{"125498", "186972", "214818", "470306", "614735"},
			},
			{Tier: "รางวัลที่ 3",
				PrizeAmount: 80000,
				RewardCount: 10,
				Numbers:     []string{"143483", "266113", "287931", "513832", "529861", "637633", "745255", "863033", "882121", "999696"},
			},
			{Tier: "รางวัลที่ 4",
				PrizeAmount: 40000,
				RewardCount: 50,
				Numbers:     []string{"013278", "085364", "238882", "321760", "391483", "468121", "563319", "625936", "744031", "892195"},
			},
			{Tier: "รางวัลที่ 5",
				PrizeAmount: 20000,
				RewardCount: 100,
				Numbers:     []string{"009012", "137282", "224856", "282595", "344137", "450505", "603032", "685690", "785675", "901273"},
			},
		},
	},
}

func check_lottery(c *gin.Context) {
	number := c.Query("numbers")
	if number != "" {
		for _, lot := range lottery {
			for _, prize := range lot.Prizes {
				for _, n := range prize.Numbers {
					if n == number {
						c.JSON(http.StatusOK, gin.H{
							"status":     "คุณคือผู้โชคดี",
							"tier":       prize.Tier,
							"prizeAmoun": prize.PrizeAmount,
							"date":       lot.Date,
							"numbers":    number,
						})
						return
					}
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "คุณไม่ถูกรางวัล",
			"numbers": number,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "เกิดข้อผิดพลาด",
	})
}

func main() {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/lottery", check_lottery)
	}
	r.Run(":8080")
}
