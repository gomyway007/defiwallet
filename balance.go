package service

import (
    "context"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"singo/abi"
	"singo/model"
	"singo/serializer"
    "log"
)

type UserBalance struct {
	TokenAddress string `form:"token" json:"token" binding:"required,min=5,max=80"`
	UserAddress string `form:"user" json:"user" binding:"required,min=8,max=80"`
}

//https://eth-sepolia.g.alchemy.com/v2/8LJnxzA9YOQVCh6G6v_OwccZZfEfGkXZ
//wss://eth-sepolia.g.alchemy.com/v2/8LJnxzA9YOQVCh6G6v_OwccZZfEfGkXZ
func (s *UserBalance)GetBalance() serializer.Response{
   
    url := "https://eth-sepolia.g.alchemy.com/v2/8LJnxzA9YOQVCh6G6v_OwccZZfEfGkXZ"
    client, err := ethclient.Dial(url)
    if err != nil {
       log.Fatalf("Could not connect to Infura with ethclient: %s", err)
    }
    log.Println("connect success")

    ctx := context.Background()
    chainId, err := client.ChainID(ctx)
    if err != nil {
       log.Fatalf("get chainId error: %s", err)
    }
    log.Printf("chainId: %s", chainId)
    bn, err := client.BlockNumber(ctx)
    if err != nil {
       log.Fatalf("get chainId error: %s", err)
    }
    log.Printf("blocknumber: %d", bn)

	address := common.HexToAddress(s.TokenAddress)//0x779877A7B0D9E8603169DdbD7836e478b4624789
	token, err := erc20.NewERC20(address, client)

	account := common.HexToAddress(s.UserAddress)//
	balance, err := token.BalanceOf(nil, account)
	if err != nil {
	  log.Fatalf("get balance error: %s", err)
	}
	log.Printf("account %s has balance %s", account, balance)

	b:=model.Balance{}
	b.SetBalance(balance)
	return serializer.BuildBalanceResponse(b)
}

type UserBalances struct {
	UserAddress string `form:"user" json:"user" binding:"required,min=8,max=80"`
}

func (s *UserBalances)GetBalances() serializer.Response{
   
    url := "https://eth-sepolia.g.alchemy.com/v2/8LJnxzA9YOQVCh6G6v_OwccZZfEfGkXZ"
    client, err := ethclient.Dial(url)
    if err != nil {
       log.Fatalf("Could not connect to Infura with ethclient: %s", err)
    }
    log.Println("connect success")

    ctx := context.Background()
    chainId, err := client.ChainID(ctx)
    if err != nil {
       log.Fatalf("get chainId error: %s", err)
    }
    log.Printf("chainId: %s", chainId)
    bn, err := client.BlockNumber(ctx)
    if err != nil {
       log.Fatalf("get chainId error: %s", err)
    }
    log.Printf("blocknumber: %d", bn)

	//get token list
	t:=model.Token{}
	ts:=t.GetToken()

	var bs []model.Balances
	account := common.HexToAddress(s.UserAddress)//
	for _,ti:=range ts{
		address := common.HexToAddress(ti.TokenAddress)//0x779877A7B0D9E8603169DdbD7836e478b4624789
		token, err := erc20.NewERC20(address, client)
	
		balance, err := token.BalanceOf(nil, account)
		if err != nil {
		  log.Fatalf("get balance error: %s", err)
		}
		log.Printf("account %s has token %s balance %s", account, ti.Symbol,balance)

		var b model.Balances
		b.Symbol = ti.Symbol
		b.TokenAddress = ti.TokenAddress
		b.Balance = balance
		bs = append(bs,b)
	}



	return serializer.BuildBalancesResponse(bs)
}