package question

import (
	"container/heap"
	"math"
	"sort"
)

func Answer0313(n int,primes []int)int{
	return nthSuperUglyNumber(n,primes)
}
type hp0313 struct {
	sort.IntSlice
}

func (h *hp0313) Push(v interface{})  {
	h.IntSlice=append(h.IntSlice,v.(int))
}
func (h *hp0313) Pop() interface{}  {
	a:=h.IntSlice
	v:=a[len(a)-1]
	h.IntSlice=a[:len(a)-1]
	return v
}
func nthSuperUglyNumber(n int, primes []int) int {
	// 素因子全在primes的第n个数字
	// 等价于要从小到达排列由primes中素因子能形成的所有数字，取第n个，从1开始
	seen:=map[int]struct{}{1: {}}
	h:=&hp0313{[]int{1}}
	var res int
	for i:=0;i<n;i++{
		res=heap.Pop(h).(int)
		for _,prime:=range primes{
			next:=res*prime
			if _,ok:=seen[prime];!ok{
				heap.Push(h,next)
			}
		}
	}
	return res
}


func nthSuperUglyNumber2(n int, primes []int) int {
	// 素因子全在primes的第n个数字
	// 等价于要从小到达排列由primes中素因子能形成的所有数字，取第n个，从1开始
	// 创建一个与数组primes相同的数组pointers，表示下一个是当前指针指向的数*对应的质因数，初始化，均为1
	// dp[i]=min{dp[pointers[j]*primes[ij])
	// 比较 dp[i] == dp[pointers[j]]*primes[j] ，若相等，则pointers[j]+=1
	dp:=make([]int,n+1)  // dp表示已经出现过的丑数
	dp[1]=1
	pointer:=make([]int,len(primes)) // pointer[k]表示前pointer[k]-1个丑数已经乘过primes[k]了
	for i:=0;i<len(pointer);i++{
		pointer[i]=1
	}
	min:=func(a,b int)int{
		if a<b{
			return  a
		}
		return b
	}
	for i:=2;i<len(dp);i++{
		tmp:=make([]int,len(primes))
		minNum:=math.MaxInt64
		for j,p:=range pointer{ // 这里这么做的目的是，找到下一个可能的丑数
			tmp[j]=dp[p]*primes[j]
			minNum=min(minNum,tmp[j])
		}
		dp[i]=minNum
		for j,num:=range tmp{
			if minNum==num{
				pointer[j]++
			}
		}
	}
	return dp[n]
}