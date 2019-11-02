module github.com/summa-tx/burning-bitcoin/golang

go 1.13

require (
	bou.ke/monkey v1.0.1 // indirect
	github.com/cosmos/cosmos-sdk v0.35.0
	github.com/gorilla/mux v1.7.3
	github.com/otiai10/copy v0.0.0-20180813032824-7e9a647135a1 // indirect
	github.com/otiai10/curr v0.0.0-20150429015615-9b4961190c95 // indirect
	github.com/otiai10/mint v1.2.3 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/summa-tx/bitcoin-spv/golang v1.1.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.32.6
	github.com/tendermint/tm-db v0.2.0
	google.golang.org/appengine v1.4.0 // indirect
)

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.34.4-0.20191102053406-d1f6c30cc5ee
