# golang_leetcode
leetcode in golang

go 文件 : https://tour.go-zh.org/list

## 常用go 指令
初始化mod檔
`go init mod xxxx.com/projectName`
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

