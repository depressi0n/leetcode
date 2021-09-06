package question

import (
	"container/heap"
	"fmt"
	"math"
)

func Answer0743(times [][]int, n int, k int) int {
	return networkDelayTime(times, n, k)
}

const INF = math.MaxInt64 / 2

func networkDelayTime(times [][]int, n int, k int) int {
	// 单源最短路径 dijkstra算法
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = INF
		}
	}
	for _, edge := range times {
		g[edge[0]-1][edge[1]-1] = edge[2]
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = INF
	}
	dis[k-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		// 找出一个最小的
		for y, u := range used {
			if !u && (x == -1 || dis[y] < dis[x]) {
				x = y
			}
		}
		used[x] = true
		// 更新长度
		for y, cost := range g[x] {
			dis[y] = min(dis[y], dis[x]+cost)
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dis[i])
	}
	if res > INF/2 {
		return -1
	}
	return res
}

func networkDelayTime2(times [][]int, n int, k int) int {
	const INF = 0x3f3f3f3f
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	w := make([][]int, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				w[i][j] = 0
			} else {
				w[i][j] = INF
			}
		}
	}
	fmt.Println(w)
	for _, time := range times {
		w[time[0]-1][time[1]-1] = time[2]
	}
	// Floyd算法 - 邻接矩阵
	floyd := func(w [][]int, n int) {
		for p := 0; p < n; p++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					// 处理最大值的情况
					if w[i][p] == INF || w[p][j] == INF {
						continue
					}
					w[i][j] = min(w[i][j], w[i][p]+w[p][j])
				}
			}
		}
	}
	floyd(w, n)
	fmt.Println(w)
	res := 0
	for i := 0; i < n; i++ {
		if w[k-1][i] == INF {
			return -1
		}
		res = max(res, w[k-1][i])
	}
	return res
}

// Dijkstra - 堆优化
type pair0743 struct{ d, x int }
type hp0743 []pair0743

func (h hp0743) Len() int              { return len(h) }
func (h hp0743) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp0743) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp0743) Push(v interface{})   { *h = append(*h, v.(pair0743)) }
func (h *hp0743) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func networkDelayTime3(times [][]int, n int, k int) int {
	type edge struct{ to, time int }
	g := make([][]edge, n)
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x] = append(g[x], edge{y, t[2]})
	}
	const INF = math.MaxInt64/2
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dis:=make([]int,n)
	for i := 0; i < n; i++ {
		dis[i]=INF
	}
	dis[k-1]=0
	h := &hp0743{{0, k - 1}}
	for h.Len() > 0 {
		// 弹出一个最小的
		p := heap.Pop(h).(pair0743)
		x := p.x
		if dis[x] < p.d {
			continue
		}
		// 如果距离更小，则将周围的点更新进来
		for _, e := range g[x] {
			y := e.to
			if d := dis[x] + e.time; d < dis[y] {
				dis[y] = d
				heap.Push(h, pair0743{d, y})
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dis[i] == INF {
			return -1
		}
		res = max(res, dis[i])
	}
	return res
}

//TODO：Bellman Ford
//class Solution {
//    class Edge {
//        int a, b, c;
//        Edge(int _a, int _b, int _c) {
//            a = _a; b = _b; c = _c;
//        }
//    }
//    int N = 110, M = 6010;
//    int[] dist = new int[N];
//    boolean[] vis = new boolean[N];
//    int INF = 0x3f3f3f3f;
//    int n, m, k;
//    List<Edge> es = new ArrayList<>();
//    public int networkDelayTime(int[][] ts, int _n, int _k) {
//        n = _n; k = _k;
//        m = ts.length;
//        for (int[] t : ts) {
//            int u = t[0], v = t[1], c = t[2];
//            es.add(new Edge(u, v, c));
//        }
//        bf();
//        int ans = 0;
//        for (int i = 1; i <= n; i++) {
//            ans = Math.max(ans, dist[i]);
//        }
//        return ans > INF / 2 ? -1 : ans;
//    }
//    void bf() {
//        Arrays.fill(dist, INF);
//        dist[k] = 0;
//        for (int p = 1; p <= n; p++) {
//            int[] prev = dist.clone();
//            for (Edge e : es) {
//                int a = e.a, b = e.b, c = e.c;
//                dist[b] = Math.min(dist[b], prev[a] + c);
//            }
//        }
//    }
//}
//

//TODO：SPFA（邻接表） -- Bellman Ford 的优化实现，可以使用队列进行优化，也可以使用栈进行优化。
//class Solution {
//    int N = 110, M = 6010;
//    int[] he = new int[N], e = new int[M], ne = new int[M], w = new int[M];
//    int[] dist = new int[N];
//    boolean[] vis = new boolean[N];
//    int INF = 0x3f3f3f3f;
//    int n, k, idx;
//    void add(int a, int b, int c) {
//        e[idx] = b;
//        ne[idx] = he[a];
//        he[a] = idx;
//        w[idx] = c;
//        idx++;
//    }
//    public int networkDelayTime(int[][] ts, int _n, int _k) {
//        n = _n; k = _k;
//        Arrays.fill(he, -1);
//        for (int[] t : ts) {
//            int u = t[0], v = t[1], c = t[2];
//            add(u, v, c);
//        }
//        spfa();
//        int ans = 0;
//        for (int i = 1; i <= n; i++) {
//            ans = Math.max(ans, dist[i]);
//        }
//        return ans > INF / 2 ? -1 : ans;
//    }
//    void spfa() {
//        Arrays.fill(vis, false);
//        Arrays.fill(dist, INF);
//        dist[k] = 0;
//        Deque<Integer> d = new ArrayDeque<>();
//        d.addLast(k);
//        vis[k] = true;
//        while (!d.isEmpty()) {
//            int poll = d.pollFirst();
//            vis[poll] = false;
//            for (int i = he[poll]; i != -1; i = ne[i]) {
//                int j = e[i];
//                if (dist[j] > dist[poll] + w[i]) {
//                    dist[j] = dist[poll] + w[i];
//                    if (vis[j]) continue;
//                    d.addLast(j);
//                    vis[j] = true;
//                }
//            }
//        }
//    }
//}
//

//class Solution {
//public:
//    int networkDelayTime(vector<vector<int>>& times, int N, int K) {
//        vector<int> dis(N+1,-1);
//        dis[K]=0;
//        using Pair=pair<int,int>;//first是距离，second是目标点
//        priority_queue<Pair,vector<Pair>,greater<Pair>> pq;
//        pq.emplace(0,K);//起点先入队
//
//        while(!pq.empty()){
//            auto e=pq.top();pq.pop();//e为连接visited和unvisited的最小边
//            if(e.first>dis[e.second]) continue;//如果e的权比K到e.second还大，就不可能缩短路径了
//            for(int i=0;i<times.size();i++){
//                if(times[i][0]==e.second){//遍历一遍所有以e.second为起点的边，做relax，并将relax之后的点入队
//                    int v=times[i][1];
//                    int w=e.first+times[i][2];
//                    if(dis[v]==-1||dis[v]>w){
//                        dis[v]=w;
//                        pq.emplace(w,v);
//                    }
//                }
//            }
//
//        }
//
//        int ans=0;
//        for(int i=1;i<=N;i++){
//            if(dis[i]==-1) return -1;
//            ans=max(ans,dis[i]);
//        }
//        return ans;
//    }
//};
//class Solution {
//public:
//    int networkDelayTime(vector<vector<int>>& times, int N, int K) {
//        vector<vector<long long>> graph(N+1,vector<long long>(N+1,INT_MAX));
//        for(int i=1;i<=N;i++)    graph[i][i]=0;
//        for(auto e:times)    graph[e[0]][e[1]]=e[2];
//        vector<bool> visited(N+1,false);    visited[K]=true;
//
//        for(int i=1;i<N;i++){//进行一次表示，从K到j，经过一个点能不能缩短路径。一共进行N-1就可以。
//            int min_id=0,min_dis=INT_MAX;//每次在unused的点中找到K最近的那个
//            for(int j=1;j<=N;j++){
//                if(visited[j]==false && graph[K][j]<min_dis){
//                    min_dis=graph[K][j];
//                    min_id=j;
//                }
//            }
//            visited[min_id]=true;//把这个点标记为“使用过了”
//            for(int j=1;j<=N;j++){//relax
//                if(graph[K][min_id]+graph[min_id][j]<graph[K][j]){
//                    graph[K][j]=graph[K][min_id]+graph[min_id][j];
//                }
//            }
//        }
//        int ans=0;
//        for(int i=1;i<=N;i++){
//            if(graph[K][i]==INT_MAX) return -1;
//            ans=max(ans,(int)graph[K][i]);
//        }
//        return ans;
//    }
//};
//class Solution {
//public:
//    int networkDelayTime(vector<vector<int>>& times, int N, int K) {
//        vector<vector<long long>> graph(N+1,vector<long long>(N+1,INT_MAX));
//        for(int i=1;i<=N;i++)    graph[i][i]=0;
//        for(auto e:times)    graph[e[0]][e[1]]=e[2];
//
//        for(int k=1;k<=N;k++)//k放在最外层
//            for(int i=1;i<=N;i++)
//                for(int j=1;j<=N;j++)
//                    graph[i][j]=min(graph[i][j],graph[i][k]+graph[k][j]);
//
//        int ans=0;
//        for(int i=1;i<=N;i++){
//            if(graph[K][i]==INT_MAX) return -1;
//            ans=max(ans,(int)graph[K][i]);
//        }
//        return ans;
//    }
//};
