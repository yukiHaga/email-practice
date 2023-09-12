package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	// SMTPサーバーのアドレス
	mailServerAddr = "localhost:1025"
	// 送信者のメールアドレス
	mailSenderAddr = "samplehoge@sample.com"
	// 受信者のメールアドレス
	mailRecepterAddr = "hogehoge@sample.com"
	// CC:受信者のメールアドレス
	mailCCRecepterAddr = "hogehogecc@sample.com"
)

// SMTPクライアントのコード
// メール送信の処理をするとメール処理の時間が長いのでリクエストした時のレスポンスが遅くなってしまいます。これを防ぐためにメール処理は非同期に行うことがあります。
// そんな遅くもなかったから今回はゴルーチンを使わずに実装した
func main() {
	clientConn, err := smtp.Dial(mailServerAddr)
	if err != nil {
		log.Printf("fail to conn: %v\r\n", err)
		return
	}

	// メール送信者(送信元)の設定
	clientConn.Mail(mailSenderAddr)
	// メール受信者(送信先)の設定
	clientConn.Rcpt(mailRecepterAddr)
	// CCでのメール受信者の設定
	clientConn.Rcpt(mailCCRecepterAddr)

	// メールのボディを作成
	wc, err := clientConn.Data()
	if err != nil {
		log.Printf("fail to Data: %v\r\n", err)
		return
	}
	defer wc.Close()

	var mailMessage string

	mailMessage += fmt.Sprintf("To: %s\r\n", mailRecepterAddr)
	mailMessage += fmt.Sprintf("Cc: %s\r\n", mailCCRecepterAddr)
	mailMessage += fmt.Sprintf("From: %s\r\n", mailSenderAddr)
	mailMessage += fmt.Sprintf("Subject: %s\r\n", "ハローワールドの贈呈")

	// メールヘッダの終わり
	mailMessage += "\r\n"

	// 本文指定
	mailMessage += `
			おはようございます。
			HELLO WORLDを贈呈します
		`

	_, err = wc.Write([]byte(mailMessage))
	if err != nil {
		log.Printf("fail to Write: %v\r\n", err)
		return
	}

	// メールセッションを終了
	clientConn.Quit()
}
