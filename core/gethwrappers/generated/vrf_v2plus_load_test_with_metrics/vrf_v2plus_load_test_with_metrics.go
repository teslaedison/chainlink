// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_v2plus_load_test_with_metrics

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var VRFV2PlusLoadTestWithMetricsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"OnlyOwnerOrCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"}],\"name\":\"CoordinatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"getRequestBlockTimes\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"_nativePayment\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_numWords\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestCount\",\"type\":\"uint16\"}],\"name\":\"requestRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_averageResponseTimeInBlocksMillions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_averageResponseTimeInSecondsMillions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fastestResponseTimeInBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fastestResponseTimeInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestBlockTimes\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_responseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_slowestResponseTimeInBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_slowestResponseTimeInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_vrfCoordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorV2Plus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060055560006006556103e760075560006008556103e76009556000600a553480156200003157600080fd5b50604051620017cf380380620017cf8339810160408190526200005491620001dd565b803380600081620000ac5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000df57620000df8162000131565b5050506001600160a01b0381166200010a5760405163d92e233d60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055506200020f565b6001600160a01b0381163314156200018c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a3565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001f057600080fd5b81516001600160a01b03811681146200020857600080fd5b9392505050565b6115b0806200021f6000396000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80638ea98117116100d8578063ad00fe611161008c578063d8a4676f11610066578063d8a4676f14610334578063dc1670db14610359578063f2fde38b1461036257600080fd5b8063ad00fe611461031a578063b1e2174914610323578063d826f88f1461032c57600080fd5b80639eccacf6116100bd5780639eccacf614610286578063a168fa89146102a6578063a4c52cf51461031157600080fd5b80638ea981171461024b578063958cccb71461025e57600080fd5b8063557d2e921161012f57806379ba50971161011457806379ba5097146101fb57806381a4342c146102035780638da5cb5b1461020c57600080fd5b8063557d2e92146101df5780636846de20146101e857600080fd5b80631742748e116101605780631742748e146101b85780631fe543e3146101c157806339aea80a146101d657600080fd5b806301e5f8281461017c5780630b26348614610198575b600080fd5b61018560065481565b6040519081526020015b60405180910390f35b6101ab6101a636600461127c565b610375565b60405161018f919061129e565b610185600a5481565b6101d46101cf36600461110e565b610471565b005b61018560075481565b61018560045481565b6101d46101f63660046111fd565b6104f7565b6101d4610717565b61018560055481565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018f565b6101d461025936600461109f565b610814565b61027161026c3660046110dc565b61099e565b60405163ffffffff909116815260200161018f565b6002546102269073ffffffffffffffffffffffffffffffffffffffff1681565b6102e76102b43660046110dc565b600d602052600090815260409020805460028201546003830154600484015460059094015460ff90931693919290919085565b6040805195151586526020860194909452928401919091526060830152608082015260a00161018f565b61018560095481565b61018560085481565b610185600b5481565b6101d46109d8565b6103476103423660046110dc565b610a11565b60405161018f969594939291906112e8565b61018560035481565b6101d461037036600461109f565b610af6565b606060006103838385611414565b600c549091508111156103955750600c545b60006103a185836114a4565b67ffffffffffffffff8111156103b9576103b9611574565b6040519080825280602002602001820160405280156103e2578160200160208202803683370190505b509050845b8281101561046857600c818154811061040257610402611545565b6000918252602090912060088204015460079091166004026101000a900463ffffffff168261043188846114a4565b8151811061044157610441611545565b63ffffffff9092166020928302919091019091015280610460816114dd565b9150506103e7565b50949350505050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146104e9576002546040517f1cf993f400000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b6104f38282610b0a565b5050565b6104ff610c81565b60005b8161ffff168161ffff16101561070d5760006040518060c001604052808881526020018a81526020018961ffff1681526020018763ffffffff1681526020018563ffffffff1681526020016105666040518060200160405280891515815250610d02565b90526002546040517f9b1c385e00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff90911690639b1c385e906105c4908590600401611354565b602060405180830381600087803b1580156105de57600080fd5b505af11580156105f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061691906110f5565b600b81905590506000610627610dbe565b6040805160c08101825260008082528251818152602080820185528084019182524284860152606084018390526080840186905260a08401839052878352600d815293909120825181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690151517815590518051949550919390926106b5926001850192910190610ff3565b506040820151600282015560608201516003820155608082015160048083019190915560a09092015160059091015580549060006106f2836114dd565b91905055505050508080610705906114bb565b915050610502565b5050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610798576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016104e0565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590610854575060025473ffffffffffffffffffffffffffffffffffffffff163314155b156108d8573361087960005473ffffffffffffffffffffffffffffffffffffffff1690565b6002546040517f061db9c100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff938416600482015291831660248301529190911660448201526064016104e0565b73ffffffffffffffffffffffffffffffffffffffff8116610925576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be69060200160405180910390a150565b600c81815481106109ae57600080fd5b9060005260206000209060089182820401919006600402915054906101000a900463ffffffff1681565b6000600581905560068190556103e76007819055600a829055600882905560095560048190556003819055610a0f90600c9061103e565b565b6000818152600d60209081526040808320815160c081018352815460ff1615158152600182018054845181870281018701909552808552606095879586958695869586959194929385840193909290830182828015610a8f57602002820191906000526020600020905b815481526020019060010190808311610a7b575b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815250509050806000015181602001518260400151836060015184608001518560a001519650965096509650965096505091939550919395565b610afe610c81565b610b0781610e5b565b50565b6000828152600d6020908152604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811782558351610b59939290910191840190610ff3565b506000828152600d6020526040902042600390910155610b77610dbe565b6000838152600d6020526040812060058101839055600401549091610b9c91906114a4565b6000848152600d6020526040812060028101546003909101549293509091610bc491906114a4565b9050610bdb82600754600654600554600354610f51565b600555600755600655600954600854600a54600354610bff93859390929091610f51565b600a5560095560085560038054906000610c18836114dd565b9091555050600c80546001810182556000919091527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c76008820401805460079092166004026101000a63ffffffff81810219909316949092169190910292909217909155505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016104e0565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa82604051602401610d3b91511515815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915292915050565b600046610dca81610fcc565b15610e5457606473ffffffffffffffffffffffffffffffffffffffff1663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610e1657600080fd5b505afa158015610e2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4e91906110f5565b91505090565b4391505090565b73ffffffffffffffffffffffffffffffffffffffff8116331415610edb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016104e0565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000808080610f6389620f4240611467565b905086891115610f71578896505b878910610f7e5787610f80565b885b97506000808611610f915781610fbb565b610f9c866001611414565b82610fa7888a611467565b610fb19190611414565b610fbb919061142c565b979a98995096979650505050505050565b600061a4b1821480610fe0575062066eed82145b80610fed575062066eee82145b92915050565b82805482825590600052602060002090810192821561102e579160200282015b8281111561102e578251825591602001919060010190611013565b5061103a92915061105f565b5090565b508054600082556007016008900490600052602060002090810190610b0791905b5b8082111561103a5760008155600101611060565b803561ffff8116811461108657600080fd5b919050565b803563ffffffff8116811461108657600080fd5b6000602082840312156110b157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146110d557600080fd5b9392505050565b6000602082840312156110ee57600080fd5b5035919050565b60006020828403121561110757600080fd5b5051919050565b6000806040838503121561112157600080fd5b8235915060208084013567ffffffffffffffff8082111561114157600080fd5b818601915086601f83011261115557600080fd5b81358181111561116757611167611574565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811085821117156111aa576111aa611574565b604052828152858101935084860182860187018b10156111c957600080fd5b600095505b838610156111ec5780358552600195909501949386019386016111ce565b508096505050505050509250929050565b600080600080600080600060e0888a03121561121857600080fd5b8735965061122860208901611074565b95506040880135945061123d6060890161108b565b93506080880135801515811461125257600080fd5b925061126060a0890161108b565b915061126e60c08901611074565b905092959891949750929550565b6000806040838503121561128f57600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b818110156112dc57835163ffffffff16835292840192918401916001016112ba565b50909695505050505050565b600060c082018815158352602060c08185015281895180845260e086019150828b01935060005b8181101561132b5784518352938301939183019160010161130f565b505060408501989098525050506060810193909352608083019190915260a09091015292915050565b6000602080835283518184015280840151604084015261ffff6040850151166060840152606084015163ffffffff80821660808601528060808701511660a0860152505060a084015160c08085015280518060e086015260005b818110156113cb57828101840151868201610100015283016113ae565b818111156113de57600061010083880101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169390930161010001949350505050565b6000821982111561142757611427611516565b500190565b600082611462577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561149f5761149f611516565b500290565b6000828210156114b6576114b6611516565b500390565b600061ffff808316818114156114d3576114d3611516565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561150f5761150f611516565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV2PlusLoadTestWithMetricsABI = VRFV2PlusLoadTestWithMetricsMetaData.ABI

var VRFV2PlusLoadTestWithMetricsBin = VRFV2PlusLoadTestWithMetricsMetaData.Bin

func DeployVRFV2PlusLoadTestWithMetrics(auth *bind.TransactOpts, backend bind.ContractBackend, _vrfCoordinator common.Address) (common.Address, *types.Transaction, *VRFV2PlusLoadTestWithMetrics, error) {
	parsed, err := VRFV2PlusLoadTestWithMetricsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV2PlusLoadTestWithMetricsBin), backend, _vrfCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV2PlusLoadTestWithMetrics{address: address, abi: *parsed, VRFV2PlusLoadTestWithMetricsCaller: VRFV2PlusLoadTestWithMetricsCaller{contract: contract}, VRFV2PlusLoadTestWithMetricsTransactor: VRFV2PlusLoadTestWithMetricsTransactor{contract: contract}, VRFV2PlusLoadTestWithMetricsFilterer: VRFV2PlusLoadTestWithMetricsFilterer{contract: contract}}, nil
}

type VRFV2PlusLoadTestWithMetrics struct {
	address common.Address
	abi     abi.ABI
	VRFV2PlusLoadTestWithMetricsCaller
	VRFV2PlusLoadTestWithMetricsTransactor
	VRFV2PlusLoadTestWithMetricsFilterer
}

type VRFV2PlusLoadTestWithMetricsCaller struct {
	contract *bind.BoundContract
}

type VRFV2PlusLoadTestWithMetricsTransactor struct {
	contract *bind.BoundContract
}

type VRFV2PlusLoadTestWithMetricsFilterer struct {
	contract *bind.BoundContract
}

type VRFV2PlusLoadTestWithMetricsSession struct {
	Contract     *VRFV2PlusLoadTestWithMetrics
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV2PlusLoadTestWithMetricsCallerSession struct {
	Contract *VRFV2PlusLoadTestWithMetricsCaller
	CallOpts bind.CallOpts
}

type VRFV2PlusLoadTestWithMetricsTransactorSession struct {
	Contract     *VRFV2PlusLoadTestWithMetricsTransactor
	TransactOpts bind.TransactOpts
}

type VRFV2PlusLoadTestWithMetricsRaw struct {
	Contract *VRFV2PlusLoadTestWithMetrics
}

type VRFV2PlusLoadTestWithMetricsCallerRaw struct {
	Contract *VRFV2PlusLoadTestWithMetricsCaller
}

type VRFV2PlusLoadTestWithMetricsTransactorRaw struct {
	Contract *VRFV2PlusLoadTestWithMetricsTransactor
}

func NewVRFV2PlusLoadTestWithMetrics(address common.Address, backend bind.ContractBackend) (*VRFV2PlusLoadTestWithMetrics, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV2PlusLoadTestWithMetricsABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV2PlusLoadTestWithMetrics(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetrics{address: address, abi: abi, VRFV2PlusLoadTestWithMetricsCaller: VRFV2PlusLoadTestWithMetricsCaller{contract: contract}, VRFV2PlusLoadTestWithMetricsTransactor: VRFV2PlusLoadTestWithMetricsTransactor{contract: contract}, VRFV2PlusLoadTestWithMetricsFilterer: VRFV2PlusLoadTestWithMetricsFilterer{contract: contract}}, nil
}

func NewVRFV2PlusLoadTestWithMetricsCaller(address common.Address, caller bind.ContractCaller) (*VRFV2PlusLoadTestWithMetricsCaller, error) {
	contract, err := bindVRFV2PlusLoadTestWithMetrics(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetricsCaller{contract: contract}, nil
}

func NewVRFV2PlusLoadTestWithMetricsTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV2PlusLoadTestWithMetricsTransactor, error) {
	contract, err := bindVRFV2PlusLoadTestWithMetrics(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetricsTransactor{contract: contract}, nil
}

func NewVRFV2PlusLoadTestWithMetricsFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV2PlusLoadTestWithMetricsFilterer, error) {
	contract, err := bindVRFV2PlusLoadTestWithMetrics(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetricsFilterer{contract: contract}, nil
}

func bindVRFV2PlusLoadTestWithMetrics(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV2PlusLoadTestWithMetricsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2PlusLoadTestWithMetrics.Contract.VRFV2PlusLoadTestWithMetricsCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.VRFV2PlusLoadTestWithMetricsTransactor.contract.Transfer(opts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.VRFV2PlusLoadTestWithMetricsTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2PlusLoadTestWithMetrics.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.contract.Transfer(opts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) GetRequestBlockTimes(opts *bind.CallOpts, offset *big.Int, quantity *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "getRequestBlockTimes", offset, quantity)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) GetRequestBlockTimes(offset *big.Int, quantity *big.Int) ([]uint32, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.GetRequestBlockTimes(&_VRFV2PlusLoadTestWithMetrics.CallOpts, offset, quantity)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) GetRequestBlockTimes(offset *big.Int, quantity *big.Int) ([]uint32, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.GetRequestBlockTimes(&_VRFV2PlusLoadTestWithMetrics.CallOpts, offset, quantity)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

	error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(GetRequestStatus)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.RequestTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RequestBlockNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentBlockNumber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.GetRequestStatus(&_VRFV2PlusLoadTestWithMetrics.CallOpts, _requestId)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.GetRequestStatus(&_VRFV2PlusLoadTestWithMetrics.CallOpts, _requestId)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) Owner() (common.Address, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.Owner(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) Owner() (common.Address, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.Owner(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SAverageResponseTimeInBlocksMillions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_averageResponseTimeInBlocksMillions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SAverageResponseTimeInBlocksMillions() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SAverageResponseTimeInBlocksMillions(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SAverageResponseTimeInBlocksMillions() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SAverageResponseTimeInBlocksMillions(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SAverageResponseTimeInSecondsMillions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_averageResponseTimeInSecondsMillions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SAverageResponseTimeInSecondsMillions() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SAverageResponseTimeInSecondsMillions(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SAverageResponseTimeInSecondsMillions() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SAverageResponseTimeInSecondsMillions(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SFastestResponseTimeInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_fastestResponseTimeInBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SFastestResponseTimeInBlocks() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SFastestResponseTimeInBlocks(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SFastestResponseTimeInBlocks() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SFastestResponseTimeInBlocks(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SFastestResponseTimeInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_fastestResponseTimeInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SFastestResponseTimeInSeconds() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SFastestResponseTimeInSeconds(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SFastestResponseTimeInSeconds() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SFastestResponseTimeInSeconds(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SLastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SLastRequestId() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SLastRequestId(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SLastRequestId() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SLastRequestId(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SRequestBlockTimes(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_requestBlockTimes", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SRequestBlockTimes(arg0 *big.Int) (uint32, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SRequestBlockTimes(&_VRFV2PlusLoadTestWithMetrics.CallOpts, arg0)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SRequestBlockTimes(arg0 *big.Int) (uint32, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SRequestBlockTimes(&_VRFV2PlusLoadTestWithMetrics.CallOpts, arg0)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SRequestCount() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SRequestCount(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SRequestCount() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SRequestCount(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

	error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(SRequests)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RequestTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequestBlockNumber = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentBlockNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SRequests(&_VRFV2PlusLoadTestWithMetrics.CallOpts, arg0)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SRequests(&_VRFV2PlusLoadTestWithMetrics.CallOpts, arg0)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SResponseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_responseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SResponseCount() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SResponseCount(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SResponseCount() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SResponseCount(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SSlowestResponseTimeInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_slowestResponseTimeInBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SSlowestResponseTimeInBlocks() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SSlowestResponseTimeInBlocks(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SSlowestResponseTimeInBlocks() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SSlowestResponseTimeInBlocks(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SSlowestResponseTimeInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_slowestResponseTimeInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SSlowestResponseTimeInSeconds() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SSlowestResponseTimeInSeconds(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SSlowestResponseTimeInSeconds() (*big.Int, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SSlowestResponseTimeInSeconds(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCaller) SVrfCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2PlusLoadTestWithMetrics.contract.Call(opts, &out, "s_vrfCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SVrfCoordinator() (common.Address, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SVrfCoordinator(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsCallerSession) SVrfCoordinator() (common.Address, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SVrfCoordinator(&_VRFV2PlusLoadTestWithMetrics.CallOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.AcceptOwnership(&_VRFV2PlusLoadTestWithMetrics.TransactOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.AcceptOwnership(&_VRFV2PlusLoadTestWithMetrics.TransactOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.RawFulfillRandomWords(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, requestId, randomWords)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.RawFulfillRandomWords(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, requestId, randomWords)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactor) RequestRandomWords(opts *bind.TransactOpts, _subId *big.Int, _requestConfirmations uint16, _keyHash [32]byte, _callbackGasLimit uint32, _nativePayment bool, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.contract.Transact(opts, "requestRandomWords", _subId, _requestConfirmations, _keyHash, _callbackGasLimit, _nativePayment, _numWords, _requestCount)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) RequestRandomWords(_subId *big.Int, _requestConfirmations uint16, _keyHash [32]byte, _callbackGasLimit uint32, _nativePayment bool, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.RequestRandomWords(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, _subId, _requestConfirmations, _keyHash, _callbackGasLimit, _nativePayment, _numWords, _requestCount)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorSession) RequestRandomWords(_subId *big.Int, _requestConfirmations uint16, _keyHash [32]byte, _callbackGasLimit uint32, _nativePayment bool, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.RequestRandomWords(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, _subId, _requestConfirmations, _keyHash, _callbackGasLimit, _nativePayment, _numWords, _requestCount)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.contract.Transact(opts, "reset")
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) Reset() (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.Reset(&_VRFV2PlusLoadTestWithMetrics.TransactOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorSession) Reset() (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.Reset(&_VRFV2PlusLoadTestWithMetrics.TransactOpts)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactor) SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.contract.Transact(opts, "setCoordinator", _vrfCoordinator)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SetCoordinator(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, _vrfCoordinator)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.SetCoordinator(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, _vrfCoordinator)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.TransferOwnership(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, to)
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2PlusLoadTestWithMetrics.Contract.TransferOwnership(&_VRFV2PlusLoadTestWithMetrics.TransactOpts, to)
}

type VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator struct {
	Event *VRFV2PlusLoadTestWithMetricsCoordinatorSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusLoadTestWithMetricsCoordinatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusLoadTestWithMetricsCoordinatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusLoadTestWithMetricsCoordinatorSet struct {
	VrfCoordinator common.Address
	Raw            types.Log
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) FilterCoordinatorSet(opts *bind.FilterOpts) (*VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator, error) {

	logs, sub, err := _VRFV2PlusLoadTestWithMetrics.contract.FilterLogs(opts, "CoordinatorSet")
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator{contract: _VRFV2PlusLoadTestWithMetrics.contract, event: "CoordinatorSet", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) WatchCoordinatorSet(opts *bind.WatchOpts, sink chan<- *VRFV2PlusLoadTestWithMetricsCoordinatorSet) (event.Subscription, error) {

	logs, sub, err := _VRFV2PlusLoadTestWithMetrics.contract.WatchLogs(opts, "CoordinatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusLoadTestWithMetricsCoordinatorSet)
				if err := _VRFV2PlusLoadTestWithMetrics.contract.UnpackLog(event, "CoordinatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) ParseCoordinatorSet(log types.Log) (*VRFV2PlusLoadTestWithMetricsCoordinatorSet, error) {
	event := new(VRFV2PlusLoadTestWithMetricsCoordinatorSet)
	if err := _VRFV2PlusLoadTestWithMetrics.contract.UnpackLog(event, "CoordinatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator struct {
	Event *VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusLoadTestWithMetrics.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator{contract: _VRFV2PlusLoadTestWithMetrics.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusLoadTestWithMetrics.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested)
				if err := _VRFV2PlusLoadTestWithMetrics.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested, error) {
	event := new(VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested)
	if err := _VRFV2PlusLoadTestWithMetrics.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator struct {
	Event *VRFV2PlusLoadTestWithMetricsOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusLoadTestWithMetricsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusLoadTestWithMetricsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusLoadTestWithMetricsOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusLoadTestWithMetrics.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator{contract: _VRFV2PlusLoadTestWithMetrics.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2PlusLoadTestWithMetricsOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusLoadTestWithMetrics.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusLoadTestWithMetricsOwnershipTransferred)
				if err := _VRFV2PlusLoadTestWithMetrics.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetricsFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferred, error) {
	event := new(VRFV2PlusLoadTestWithMetricsOwnershipTransferred)
	if err := _VRFV2PlusLoadTestWithMetrics.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetRequestStatus struct {
	Fulfilled             bool
	RandomWords           []*big.Int
	RequestTimestamp      *big.Int
	FulfilmentTimestamp   *big.Int
	RequestBlockNumber    *big.Int
	FulfilmentBlockNumber *big.Int
}
type SRequests struct {
	Fulfilled             bool
	RequestTimestamp      *big.Int
	FulfilmentTimestamp   *big.Int
	RequestBlockNumber    *big.Int
	FulfilmentBlockNumber *big.Int
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetrics) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV2PlusLoadTestWithMetrics.abi.Events["CoordinatorSet"].ID:
		return _VRFV2PlusLoadTestWithMetrics.ParseCoordinatorSet(log)
	case _VRFV2PlusLoadTestWithMetrics.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV2PlusLoadTestWithMetrics.ParseOwnershipTransferRequested(log)
	case _VRFV2PlusLoadTestWithMetrics.abi.Events["OwnershipTransferred"].ID:
		return _VRFV2PlusLoadTestWithMetrics.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV2PlusLoadTestWithMetricsCoordinatorSet) Topic() common.Hash {
	return common.HexToHash("0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6")
}

func (VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV2PlusLoadTestWithMetricsOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VRFV2PlusLoadTestWithMetrics *VRFV2PlusLoadTestWithMetrics) Address() common.Address {
	return _VRFV2PlusLoadTestWithMetrics.address
}

type VRFV2PlusLoadTestWithMetricsInterface interface {
	GetRequestBlockTimes(opts *bind.CallOpts, offset *big.Int, quantity *big.Int) ([]uint32, error)

	GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAverageResponseTimeInBlocksMillions(opts *bind.CallOpts) (*big.Int, error)

	SAverageResponseTimeInSecondsMillions(opts *bind.CallOpts) (*big.Int, error)

	SFastestResponseTimeInBlocks(opts *bind.CallOpts) (*big.Int, error)

	SFastestResponseTimeInSeconds(opts *bind.CallOpts) (*big.Int, error)

	SLastRequestId(opts *bind.CallOpts) (*big.Int, error)

	SRequestBlockTimes(opts *bind.CallOpts, arg0 *big.Int) (uint32, error)

	SRequestCount(opts *bind.CallOpts) (*big.Int, error)

	SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

		error)

	SResponseCount(opts *bind.CallOpts) (*big.Int, error)

	SSlowestResponseTimeInBlocks(opts *bind.CallOpts) (*big.Int, error)

	SSlowestResponseTimeInSeconds(opts *bind.CallOpts) (*big.Int, error)

	SVrfCoordinator(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, _subId *big.Int, _requestConfirmations uint16, _keyHash [32]byte, _callbackGasLimit uint32, _nativePayment bool, _numWords uint32, _requestCount uint16) (*types.Transaction, error)

	Reset(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCoordinatorSet(opts *bind.FilterOpts) (*VRFV2PlusLoadTestWithMetricsCoordinatorSetIterator, error)

	WatchCoordinatorSet(opts *bind.WatchOpts, sink chan<- *VRFV2PlusLoadTestWithMetricsCoordinatorSet) (event.Subscription, error)

	ParseCoordinatorSet(log types.Log) (*VRFV2PlusLoadTestWithMetricsCoordinatorSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2PlusLoadTestWithMetricsOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV2PlusLoadTestWithMetricsOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
