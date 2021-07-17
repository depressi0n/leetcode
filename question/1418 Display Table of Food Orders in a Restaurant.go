package question

import (
	"sort"
	"strconv"
)

func Answer1418(orders [][]string) [][]string {
	return displayTable(orders)
}
func displayTable(orders [][]string) [][]string {
	// 行表示一个table
	// 第一列表示table number，后续按照字母排序字母表排序
	// 第一行是Table，食物1，食物2，...
	// 行按照数字增序向下
	res := [][]string{{"Table"}}
	// 用一个map缓存每个table的foodItem
	m := map[string]map[string]int{}
	// foodItem
	foodItemM := map[string]struct{}{}
	foodItems := make([]string, 0, len(orders))
	// tableNum
	tableNums := make([]string, 0, len(orders))
	for i := 0; i < len(orders); i++ {
		//customerName
		//orders[i][0]
		//tableNumber ,foodItem
		if _, ok1 := m[orders[i][1]]; ok1 {
			if _, ok2 := m[orders[i][1]][orders[i][2]]; ok2 {
				m[orders[i][1]][orders[i][2]] += 1
			} else {
				m[orders[i][1]][orders[i][2]] = 1
			}
		} else {
			tableNums = append(tableNums, orders[i][1])
			m[orders[i][1]] = map[string]int{}
			m[orders[i][1]][orders[i][2]] = 1
		}
		if _, ok3 := foodItemM[orders[i][2]]; !ok3 {
			foodItemM[orders[i][2]] = struct{}{}
			foodItems = append(foodItems, orders[i][2])
		}
	}
	// 对foodItem排序
	sort.Slice(foodItems, func(i, j int) bool {
		return foodItems[i] < foodItems[j]
	})
	// 对tableNum排序
	sort.Slice(tableNums, func(i, j int) bool {
		x, _ := strconv.Atoi(tableNums[i])
		y, _ := strconv.Atoi(tableNums[j])
		return x < y
	})
	for i := 0; i < len(foodItems); i++ {
		res[0] = append(res[0], foodItems[i])
	}
	for i := 0; i < len(tableNums); i++ {
		res = append(res, []string{tableNums[i]})
		for j := 0; j < len(foodItems); j++ {
			res[len(res)-1] = append(res[len(res)-1], strconv.Itoa(m[tableNums[i]][foodItems[j]]))
		}
	}
	return res
}
