package service

import(
	"singo/model"
	"singo/serializer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"singo/abi"
	"log"
)

type Token struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=20"`
	Symbol string `form:"symbol" json:"symbol" binding:"required,min=2,max=20"`
	TokenAddress string `form:"address" json:"token" binding:"required,min=5,max=80"`
	
}

func (t *Token) SetToken() serializer.Response{
	token:=model.Token{Name:t.Name,Symbol:t.Symbol,TokenAddress:t.TokenAddress}
	token.SetToken(t.Name,t.Symbol,t.TokenAddress)
	return serializer.BuildTokenResponse(token)
}

type Address struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=20"`

	Address string `form:"address" json:"token" binding:"required,min=5,max=80"`
	
}

func (t *Address) SetAddress() serializer.Response{
	a:=model.Address{Name:t.Name,Address:t.Address}
	a.SetAddress(t.Name,t.Address)
	return serializer.BuildAddressResponse(a)
}

func (t *Token) ListenTransfer(){

	url := "wss://eth-sepolia.g.alchemy.com/v2/8LJnxzA9YOQVCh6G6v_OwccZZfEfGkXZ"
	ec, err := ethclient.Dial(url)
	if err != nil {
	  log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}
  
	address := common.HexToAddress(t.TokenAddress)
	token, err := erc20.NewERC20(address, ec)
	if err != nil {
	  log.Fatalf("new erc20 error: %s", err)
	}
	sink := make(chan *erc20.ERC20Transfer)
	sub, err := token.WatchTransfer(nil, sink, nil, nil)
	if err != nil {
	  log.Fatalf("subscribe transfer event error: %s", err)
	}

	log.Printf("ListenTransfer start,token: %s",t.Symbol)
	for {
	  select {
	  case <-sub.Err():
		return
	  case event := <-sink:
		from := event.From.Hex()
		to := event.To.Hex()
		value := event.Value.String()

		transfer := model.Transfer{Address:address.Hex(),From:from,To:to,Value:value}
		transfer.SetTransfer(address.Hex(),from,to,value)


	  }
	}
}

func (t *Token) ListenApprove(){

	url := "wss://eth-sepolia.g.alchemy.com/v2/8LJnxzA9YOQVCh6G6v_OwccZZfEfGkXZ"
	ec, err := ethclient.Dial(url)
	if err != nil {
	  log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}
  
	address := common.HexToAddress(t.TokenAddress)
	token, err := erc20.NewERC20(address, ec)
	if err != nil {
	  log.Fatalf("new erc20 error: %s", err)
	}
	sink := make(chan *erc20.ERC20Approval)
	sub, err := token.WatchApproval(nil, sink, nil, nil)
	if err != nil {
	  log.Fatalf("subscribe Approval event error: %s", err)
	}

	log.Printf("ListenApproval start,token: %s",t.Symbol)
	for {
	  select {
	  case <-sub.Err():
		return
	  case event := <-sink:

		owner := event.Owner.Hex()
		spender := event.Spender.Hex()
		value := event.Value.String()

		a := model.Approve{Address:address.Hex(),Owner:owner,Spender:spender,Value:value}
		a.SetApprove(address.Hex(),owner,spender,value)


	  }
	}
}