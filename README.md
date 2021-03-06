# golang_leetcode

leetcode in golang  

go 文件 : <https://tour.go-zh.org/list>  

## 常用go 指令

初始化mod檔  
`go mod init xxxx.com/projectName`  
查看go環境變數  
`go env`  

## import 注意事項

路徑位置的root為mod名稱/最上層路徑  
ex.  
mod = bartektao.com  
最上層目錄名稱 = golang_leetcode  
=> import 資料夾的根目錄為 bartektao.com/golang_leetcode

## defer

於資源申請後，可立即執行defer以預先關閉資源，這樣就不會忘記釋放資源，並且拋出錯誤時，也會被執行  

## method

可視為 C# 中，class 的擴充方法，限制該方法必須存在同一個package中  
`func (m Model) Abs() float64{ }` => 非指標method => 其方法內容部會改變 Model 本身的值，操作時操作的為該 Model 副本  
若為指標擴充(method)，則操作的如下  
`func (m *Model) Scale(f float64) { }` => 剛方法內容若改動到 Model，則會變動 Model 本身  

## 常見介面使用

```
type error interface{
    Error() string
}
```

可自定義 error 輸出內容，如下  

```
type MyError struct {
 When time.Time
 What string
}

func (e *MyError) Error() string {
 return fmt.Sprintf("at %v, %s",
  e.When, e.What)
}
```

MyError 等同於 error type，如下  

```
var e error
e = &MyError{
 time.Now(),
 "it didn't work",
}
```

## make 使用方法

建立 int array 且長度為 5 的切片  
`array := make([]int, 5)`  

建立 int array，長度為 0 ,容量為 5  
`array := make([]int, 0, 5) // len(array) = 0, cpa(array) = 5`  

初始化 map 使用(方法之一)

```
var dic map[string]int // 此時 m 為空的map，make 後才能 insert value
dic := make(map[string]int) 
dic["a"] = 1
```

## map 使用方法

直接初始化 map 並設值的方式  
`dic := map[string]int{"a":1, "b",2}`  
`v := map`  

insert  
`dic["key"]=123`  

get  
`var num int = dic["key"]`  

delete  
`delete(dic, "key")`  

test key is present in map  
`elem, ok = dic["key"]`  
if "key" is in dic, elem will set to the value, and ok = true  
if "key" isn't in dic, elem is the zero value for the map's element type, and ok = false  

## go

## sync.Mutex

Lock  
Unlock  
配合 `defer mux.Unlock()` 可確保解鎖  

## []rune vs []byte

`type rune = int32 // 4個byte`  
`type byte = uint8 // 1個byte`  
string 轉為 rune or byte 時，會轉為 unicode  

英文字: 1個byte  
中文字: 3個byte  
在顯示中文時，slice 的長度要正確才能顯示  

## ...的用法

...為 golang的語法糖，放於 array 後面，代表將該 array slice 後進行操作
ex. `func append(slice []Type, elems ...Type) []Type` => 以string為例，可傳入 `append([]string, string)` or `append([]string, string, string)` or `append([]string, []string{"1", "2", "3"}...)`

## go get vs git clone

* go get  
該命令可以理解為兩個操作 git clone + go install , 執行會將遠端程式碼及其依賴項 clone 到【$GOPATH/src】 目錄中  

```
-d 只下載不安裝
-f 只有在你包含了-u引數的時候才有效，
   不讓-u去驗證import中的每一個都已經獲取了，
   這對於本地fork的包特別有用
-fix 在獲取原始碼之後先執行fix，然後再去做其他的事情
-t 同時也下載需要為執行測試所需要的包
-u 強制使用網路去更新包和它的依賴包
-v 顯示執行的命令
```  

* git clone  
該命令可以理解為將程式碼克隆到目錄中, 執行會將遠端程式碼 clone 到當前目錄中

## 全域變數 !?!?

## channel 不等待所有 goroutine 完成

```go
func GetReachableWebUrl(urlstr string) (string, error) {
 hasHttp := make(chan bool)
 hasHttps := make(chan bool)
 finishChan := make(chan struct{})
 var tcpError error
 wg := sync.WaitGroup{}
 wg.Add(2)
 go func() {
  defer wg.Done()
  v, err := IsHostReachable(urlstr, "80")
  hasHttp <- v
  tcpError = err
 }()
 go func() {
  defer wg.Done()
  v, err := IsHostReachable(urlstr, "443")
  hasHttps <- v
  tcpError = err
 }()
 go func() {
  wg.Wait()
  close(finishChan)
 }()

Loop:
 for {
  select {
  case <-hasHttp:
   if strings.HasPrefix(urlstr, "https://") {
    return strings.Replace(urlstr, "https://", "http://", 1), nil
   }
   break Loop
  case <-hasHttps:
   if strings.HasPrefix(urlstr, "http://") {
    return strings.Replace(urlstr, "http://", "https://", 1), nil
   }
   break Loop
  case <-finishChan:
   if tcpError != nil {
    return urlstr, tcpError
   }
   break Loop
  case <-time.After(10 * time.Second):
   break Loop
  }
 }
 return urlstr, fmt.Errorf("timeout, url: %v", urlstr)
}

func IsHostReachable(urlstr string, port string) (bool, error) {
 urlInfo, err := url.Parse(urlstr)
 if err != nil {
  logrus.Errorln("Site unreachable, error: ", err)
  return false, err
 }
 timeout := 1 * time.Second
 conn, err := net.DialTimeout("tcp", urlInfo.Host+":"+port, timeout)
 if err != nil {
  logrus.Errorln("Site unreachable, host: ", urlInfo.Host, ", error: ", err)
  return false, err
 }
 defer conn.Close()
 logrus.Println(conn.LocalAddr().String())
 logrus.Println(conn.RemoteAddr().String())
 return true, nil
}
```
