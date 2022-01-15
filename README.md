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

