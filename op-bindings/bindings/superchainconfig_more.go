// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SuperchainConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SuperchainConfig.sol:SuperchainConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SuperchainConfig.sol:SuperchainConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SuperchainConfigStorageLayout = new(solc.StorageLayout)

var SuperchainConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106100885760003560e01c80636da663551161005b5780636da663551461012a5780637fbf7b6a1461013d578063c23a451a14610153578063c4d66de81461015b57600080fd5b80633f4ba83a1461008d578063452a93201461009757806354fd4d50146100c95780635c975abb14610112575b600080fd5b61009561016e565b005b61009f610294565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101056040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100c09190610768565b61011a6102cd565b60405190151581526020016100c0565b6100956101383660046107b1565b610305565b610145610435565b6040519081526020016100c0565b610145610463565b610095610169366004610880565b61048e565b610176610294565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610235576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f5375706572636861696e436f6e6669673a206f6e6c7920677561726469616e2060448201527f63616e20756e706175736500000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61026961026360017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76108b6565b60009055565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b60006102c86102c460017fd30e835d3f35624761057ff5b27d558f97bd5be034621e62240e5c0b784abe696108b6565b5490565b905090565b60006102fd6102c460017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76108b6565b600114905090565b61030d610294565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f5375706572636861696e436f6e6669673a206f6e6c7920677561726469616e2060448201527f63616e2070617573650000000000000000000000000000000000000000000000606482015260840161022c565b6103fb6103f560017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76108b6565b60019055565b7fc32e6d5d6d1de257f64eac19ddb1f700ba13527983849c9486b1ab007ea283818160405161042a9190610768565b60405180910390a150565b61046060017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76108b6565b81565b61046060017fd30e835d3f35624761057ff5b27d558f97bd5be034621e62240e5c0b784abe696108b6565b600054610100900460ff16158080156104ae5750600054600160ff909116105b806104c85750303b1580156104c8575060005460ff166001145b610554576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161022c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105b257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6105bb82610642565b801561061e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b9055565b61067561067060017fd30e835d3f35624761057ff5b27d558f97bd5be034621e62240e5c0b784abe696108b6565b829055565b60006040805173ffffffffffffffffffffffffffffffffffffffff841660208201527f7b743789cff01dafdeae47739925425aab5dfd02d0c8229e4a508bcd2b9f42bb9101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526106f291610768565b60405180910390a250565b6000815180845260005b8181101561072357602081850181015186830182015201610707565b81811115610735576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061077b60208301846106fd565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156107c357600080fd5b813567ffffffffffffffff808211156107db57600080fd5b818401915084601f8301126107ef57600080fd5b81358181111561080157610801610782565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561084757610847610782565b8160405282815287602084870101111561086057600080fd5b826020860160208301376000928101602001929092525095945050505050565b60006020828403121561089257600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461077b57600080fd5b6000828210156108ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50039056fea164736f6c634300080f000a"


var SuperchainConfigImmutableReferencesJSON = ""

func init() {
	if err := json.Unmarshal([]byte(SuperchainConfigStorageLayoutJSON), SuperchainConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SuperchainConfig"] = SuperchainConfigStorageLayout
	deployedBytecodes["SuperchainConfig"] = SuperchainConfigDeployedBin
	immutableReferences["SuperchainConfig"] = SuperchainConfigImmutableReferencesJSON
}
