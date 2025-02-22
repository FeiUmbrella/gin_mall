package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

// ProductViewKey 返回某一商品在Redis中存储的key
func ProductViewKey(id uint) string {
	return fmt.Sprintf("View:Product:%s", strconv.Itoa(int(id)))
}
