package main 

import (
	"goprojects/hello/ex20/koreaPost"
	"goprojects/hello/ex20/fedex"

)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}