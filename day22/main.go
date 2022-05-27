package main

import (
    "errors"
    "fmt"
    "hash/crc32"
    "sort"
    "strconv"
    "sync"
)

type Hash func(data []byte) uint32

//排序切片
type Uint32Slice []uint32

func (uis Uint32Slice) Len() int {
    return len(uis)
}

func (uis Uint32Slice) Less(i, j int) bool {
    return uis[i] < uis[j]
}

func (uis Uint32Slice) Swap(i, j int)  {
    uis[i], uis[j] = uis[j], uis[i]
}

type ConsistentHashBanlance struct {
    rwx sync.RWMutex
    hash Hash
    replicas int // 虚拟节点数量，虚拟节点越多 节点分布更趋向均匀
    keys Uint32Slice //节点排序
    hashMap map[uint32]string  //节点哈希和Key的map,键是hash值，值是节点key
}


// 验证是否为空
func (c *ConsistentHashBanlance) IsEmpty() bool {
    return len(c.keys) == 0
}


// Add 方法用来添加缓存节点，参数为节点key，比如使用IP
func (c *ConsistentHashBanlance) Add(params ...string) error {
    if len(params) == 0{
        return errors.New("param len 1 at least")
    }
    c.rwx.Lock()
    defer c.rwx.Unlock()
    for _, addr := range params{
        // 结合 虚拟节点数量 计算所有虚拟节点的hash值，并存入m.keys中，同时在m.hashMap中保存哈希值和key的映射
        for i := 0; i < c.replicas; i++ {
            fmt.Println(strconv.Itoa(i) + addr)
            hash := c.hash([]byte(strconv.Itoa(i) + addr))
            c.keys = append(c.keys, hash)
            c.hashMap[hash] = addr
        }
    }

    // 对所有虚拟节点的哈希值进行排序，方便之后进行二分查找
    sort.Sort(c.keys)
    return nil
}

// Get 方法根据给定的对象获取最靠近它的那个节点
func (c *ConsistentHashBanlance) Get(key string) (string, error) {
    if c.IsEmpty() {
        return "", errors.New("node is empty")
    }
    fmt.Println(c.hashMap)
    fmt.Println(c.keys)
    //crc32
    hash := c.hash([]byte(key))
    //通过二分查找获取最优节点。第一个服务器hash大于数据hash值的。就是最优服务器节点
    idx := sort.Search(len(c.keys), func(i int) bool {
        return c.keys[i] >= hash
    })
    //如果查找结果 大于 服务器节点哈希数组的最大索引，表示此时该对象哈希值位于最后一个节点之后，那么放入第一个节点中
    if idx == len(c.keys) {
        idx = 0
    }
    c.rwx.RLock()
    defer c.rwx.RUnlock()
    return c.hashMap[c.keys[idx]], nil
}

func getConsistentHashBanlance(replicas int, fn Hash) *ConsistentHashBanlance {
    m := &ConsistentHashBanlance{
        replicas: replicas,
        hash: fn,
        hashMap: make(map[uint32]string),
    }

    if m.hash == nil {
        //最多32位,保证是一个2^32-1环
        m.hash = crc32.ChecksumIEEE
    }
    return m
}
func main2()  {
    var replicas int
    replicas = 3
    ban := getConsistentHashBanlance(replicas, nil)
    ban.Add("127.0.0.1", "127.0.0.2", "127.0.0.3")
    my, _ := ban.Get("fuck")
    //my2, _ := ban.Get("fuck2")
    //my3, _ := ban.Get("fuck3")
    //fmt.Println(my, my2, my3)
    fmt.Println(my)
}

type SortSlice []int64

func (ss SortSlice) Len() int {
    return len(ss)
}

func (ss SortSlice) Less(i, j int) bool {
    return ss[i] < ss[j]
}

func (ss SortSlice) Swap(i, j int)  {
    ss[i], ss[j] = ss[j], ss[i]
}

func main()  {
    sortSlice := SortSlice{1,2,44,55,2,5,6,8,0,3}
    sort.Sort(sortSlice)
    fmt.Println(sortSlice)
    //num := binarySearch(55, sortSlice)
    num1 := search(44, sortSlice)
    //fmt.Println(num)
    fmt.Println(num1)
}

//非递归
func binarySearch(target int64, nums []int64) int {
    left := 0
    right := len(nums)
    for left < right {
        mid := (left +right) / 2
        fmt.Println(right, left, mid)
        if target == nums[mid] {
            return mid
        }

        if target > nums[mid] {
            left = mid + 1
            continue
        }

        if target < nums[mid] {
            right = mid - 1
            continue
        }
    }
    return -1
}

//递归
func search(target int64, nums []int64) int {
    return binarySearchRecursive(target, nums, 0, len(nums))
}

func binarySearchRecursive(target int64, nums []int64, left, right int) int {
    if left > right {
        return -1
    }

    mid := int(uint(left + right) >> 1)
    fmt.Println(right, left, mid)
    if target == nums[mid] {
        return mid
    }

    if target > nums[mid] {
        return binarySearchRecursive(target, nums, mid + 1, right)
    }
    if target < nums[mid] {
        return binarySearchRecursive(target, nums, mid, right - 1)
    }

    return -1
}
