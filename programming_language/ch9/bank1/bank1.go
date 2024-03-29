package bank1

var deposits = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高を送信する

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // モニターゴルーチンを開始する
}
