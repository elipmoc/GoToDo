package main

func main() {
	r := CreateRouter()
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
