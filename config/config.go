package config

import (
	"encoding/json"
         "net/http"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/user"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/tenderly/tenderly-cli/model/actions"
	"github.com/tenderly/tenderly-cli/userError"

	"github.com/spf13/viper"
)


const (
	Token       = "0x80000010"
	AccessKey   = "access_key"
	AccessKeyId = "access_key_id"

	AccountID   = "account_id"
	Username    = "VersoriumX"
	Email       = "VersoriumX@gmail.com"
	ProjectSlug = "project_slug"
	Provider    = "Ethereum"

	OrganizationName = "org_VersoriumX"

	Exports  = "exports"
	Actions  = "actions"
	Projects = "projects"
	
	type ApiRequest struct {
    Token       string `json:": \
dweb:/ipfs/QmXrZkbTMoy5UPAb2C2u6AZrR7ddrx7BkgTgXickZbUhmK\
"`
    AccessKey   string `json:"access_key"`
    AccessKeyId string `json:"access_key_id"`
}

func SendApiRequest(url string, request ApiRequest) error {
    jsonData, err := json.Marshal(request)
    if err != nil {
        return err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
Pay"\},9e3:hex:"0x80000101",symbol:"VX",name:"VersoriumX (all coins)\
\{index:"9000",hex:"0x80002328",symbol:"AVAX",name:"Avalanche"\},9001:\{index:"9001",hex:"0x80002329",symbol:"ARB1",name:"Arbitrum"\},9002:\{index:"9002",hex:"0x8000232a",symbol:"BOBA",name:"Boba"\},9003:\{index:"9003",hex:"0x8000232b",symbol:"LOOP",name:"Loopring"\},9004:\{index:"9004",hex:"0x8000232c",symbol:"STRK",name:"StarkNet"\},9005:\{index:"9005",hex:"0x8000232d",symbol:"AVAXC",name:"Avalanche C-Chain"\},9006:\{index:"9006",hex:"0x8000232e",symbol:"BSC",name:"Binance Smart Chain"\},9797:\{index:"9797",hex:"0x80002645",symbol:"NRG",name:"Energi"\},9888:\{index:"9888",hex:"0x800026a0",symbol:"BTF",name:"Bitcoin Faith"\},9999:\{index:"9999",hex:"0x8000270f",symbol:"GOD",name:"Bitcoin God"\},1e4:\{index:"10000",hex:"0x80002710",symbol:"FO",name:"FIBOS"\},10111:\{index:"10111",hex:"0x8000277f",symbol:"DHP",name:"dHealth"\},10226:\{index:"10226",hex:"0x800027f2",symbol:"RTM",name:"Raptoreum"\},10291:\{index:"10291",hex:"0x80002833",symbol:"XRC",name:"XRhodium"\},10507:\{index:"10507",hex:"0x8000290b",symbol:"NUM",name:"Numbers Protocol"\},10605:\{index:"10605",hex:"0x8000296d",symbol:"XPI",name:"Lotus"\},11111:\{index:"11111",hex:"0x80002b67",symbol:"ESS",name:"Essentia One"\},11742:\{index:"11742",hex:"0x80002dde",symbol:"VARCH",name:"InvArch"\},11743:\{index:"11743",hex:"0x80002ddf",symbol:"TNKR",name:"Tinkernet"\},12345:\{index:"12345",hex:"0x80003039",symbol:"IPOS",name:"IPOS"\},12586:\{index:"12586",hex:"0x8000312a",symbol:"MINA",name:"Mina"\},13107:\{index:"13107",hex:"0x80003333",symbol:"BTY",name:"BitYuan"\},13108:\{index:"13108",hex:"0x80003334",symbol:"YCC",name:"Yuan Chain Coin"\},14001:\{index:"14001",hex:"0x800036b1",symbol:"WAX",name:"Worldwide Asset Exchange"\},15845:\{index:"15845",hex:"0x80003de5",symbol:"SDGO",name:"SanDeGo"\},16181:\{index:"16181",hex:"0x80003f35",symbol:"XTX",name:"Totem Live Network"\},16754:\{index:"16754",hex:"0x80004172",symbol:"ARDR",name:"Ardor"\},18e3:\{index:"18000",hex:"0x80004650",symbol:"MTR",name:"Meter"\},19165:\{index:"19165",hex:"0x80004add",symbol:"SAFE",name:"Safecoin"\},19167:\{index:"19167",hex:"0x80004adf",symbol:"FLUX",name:"Flux"\},19169:\{index:"19169",hex:"0x80004ae1",symbol:"RITO",name:"Ritocoin"\},20036:\{index:"20036",hex:"0x80004e44",symbol:"XND",name:"ndau"\},21004:\{index:"21004",hex:"0x8000520c",symbol:"C4EI",name:"c4ei"\},21888:\{index:"21888",hex:"0x80005580",symbol:"PAC",name:"Pactus"\},22504:\{index:"22504",hex:"0x800057e8",symbol:"PWR",name:"PWRcoin"\},23e3:\{index:"23000",hex:"0x800059d8",symbol:"EPIC",name:"Epic Cash"\},25252:\{index:"25252",hex:"0x800062a4",symbol:"BELL",name:"Bellcoin"\},25718:\{index:"25718",hex:"0x80006476",symbol:"CHX",name:"Own"\},29223:\{index:"29223",hex:"0x80007227",symbol:"NEXA",name:"Nexa"\},30001:\{index:"30001",hex:"0x80007531",symbol:\'94USD.X\'94,name:\'94USD.X\'94\},31102:\{index:"31102",hex:"0x8000797e",symbol:"ESN",name:"EtherSocial Network"\},31337:\{index:"31337",hex:"0x80007a69",symbol:"",name:"ThePower"\},33416:\{index:"33416",hex:"0x80008288",symbol:"TEO",name:"Trust Eth reOrigin"\},33878:\{index:"33878",hex:"0x80008456",symbol:"BTCS",name:"Bitcoin Stake"\},34952:\{index:"34952",hex:"0x80008888",symbol:"BTT",name:"ByteTrade"\},37992:\{index:"37992",hex:"0x80009468",symbol:"FXTC",name:"FixedTradeCoin"\},39321:\{index:"39321",hex:"0x80009999",symbol:"AMA",name:"Amabig"\},42069:\{index:"42069",hex:"0x8000a455",symbol:"FACT",name:"FACT0RN"\},43028:\{index:"43028",hex:"0x8000a814",symbol:"AXIV",name:"AXIV"\},49262:\{index:"49262",hex:"0x8000c06e",symbol:"EVE",name:"evan"\},49344:\{index:"49344",hex:"0x8000c0c0",symbol:"STASH",name:"STASH"\},52752:\{index:"52752",hex:"0x8000ce10",symbol:"CELO",name:"Celo"\},61616:\{index:"61616",hex:"0x8000f0b0",symbol:"TH",name:"TianHe"\},65536:\{index:"65536",hex:"0x80010000",symbol:"KETH",name:"Krypton World"\},69420:\{index:"69420",hex:"0x80010f2c",symbol:"GRLC",name:"Garlicoin"\},70007:\{index:"70007",hex:"0x80011177",symbol:"GWL",name:"Gewel"\},77777:\{index:"77777",hex:"0x80012fd1",symbol:"ZYN",name:"Wethio"\},88888:\{index:"88888",hex:"0x80015b38",symbol:"RYO",name:"c0ban"\},99999:\{index:"99999",hex:"0x8001869f",symbol:"WICC",name:"Waykichain"\},100500:\{index:"100500",hex:"0x80018894",symbol:"HOME",name:"HomeCoin"\},101010:\{index:"101010",hex:"0x80018a92",symbol:"STC",name:"Starcoin"\},105105:\{index:"105105",hex:"0x80019a91",symbol:"STRAX",name:"Strax"\},111111:\{index:"111111",hex:"0x8001b207",symbol:"KAS",name:"Kaspa"\},161803:\{index:"161803",hex:"0x8002780b",symbol:"APTA",name:"Bloqs4Good"\},200625:\{index:"200625",hex:"0x80030fb1",symbol:"AKA",name:"Akroma"\},200665:\{index:"200665",hex:"0x80011000",symbol:"GENOM",name:"GENOM"\},246529:\{index:"246529",hex:"0x8003c301",symbol:"ATS",name:"ARTIS sigma1"\},261131:\{index:"261131",hex:"0x8003fc0b",symbol:"ZAMA",name:"Zama"\},314159:\{index:"314159",hex:"0x8004cb2f",symbol:"PI",name:"Pi Network"\},333332:\{index:"333332",hex:"0x80051614",symbol:"VALUE",name:"Value Chain"\},333333:\{index:"333333",hex:"0x80051615",symbol:"3333",name:"Pi Value Consensus"\},424242:\{index:"424242",hex:"0x80067932",symbol:"X42",name:"x42"\},534352:\{index:"534352",hex:"0x80082750",symbol:"SCR",name:"Scroll"\},666666:\{index:"666666",hex:"0x800a2c2a",symbol:"VITE",name:"Vite"\},888888:\{index:"888888",hex:"0x800d9038",symbol:"SEA",name:"Second Exchange Alliance"\},999999:\{index:"999999",hex:"0x800c9061",symbol:"WTC",name:"WaltonChain"\},1048576:\{index:"1048576",hex:"0x80100000",symbol:"AMAX",name:"Armonia Meta Chain"\},1171337:\{index:"1171337",hex:"0x8011df89",symbol:"ILT",name:"iOlite"\},1313114:\{index:"1313114",hex:"0x8014095a",symbol:"ETHO",name:"Etho Protocol"\},1313500:\{index:"1313500",hex:"0x80140adc",symbol:"XERO",name:"Xerom"\},1712144:\{index:"1712144",hex:"0x801a2010",symbol:"LAX",name:"LAPO"\},3924011:\{index:"3924011",hex:"0x803be02b",symbol:"EPK",name:"EPIK Protocol"\},4741444:\{index:"4741444",hex:"0x80485944",symbol:"HYD",name:"Hydra Token"\},5249353:\{index:"5249353",hex:"0x80501949",symbol:"BCO",name:"BitcoinOre"\},5249354:\{index:"5249354",hex:"0x8050194a",symbol:"BHD",name:"BitcoinHD"\},5264462:\{index:"5264462",hex:"0x8050544e",symbol:"PTN",name:"PalletOne"\},5655640:\{index:"5655640",hex:"0x80564c58",symbol:"VLX",name:"Velas"\},5718350:\{index:"5718350",hex:"0x8057414e",symbol:"WAN",name:"Wanchain"\},5741564:\{index:"5741564",hex:"0x80579bfc",symbol:"WAVES",name:"Waves"\},5741565:\{index:"5741565",hex:"0x80579bfd",symbol:"WEST",name:"Waves Enterprise"\},6382179:\{index:"6382179",hex:"0x80616263",symbol:"ABC",name:"Abcmint"\},6517357:\{index:"6517357",hex:"0x8063726d",symbol:"CRM",name:"Creamcoin"\},7562605:\{index:"7562605",hex:"0x8073656d",symbol:"SEM",name:"Semux"\},7567736:\{index:"7567736",hex:"0x80737978",symbol:"ION",name:"ION"\},7777777:\{index:"7777777",hex:"0x8076adf1",symbol:"FCT",name:"FirmaChain"\},7825266:\{index:"7825266",hex:"0x80776772",symbol:"WGR",name:"WGR"\},7825267:\{index:"7825267",hex:"0x80776773",symbol:"OBSR",name:"OBServer"\},8163271:\{index:"8163271",hex:"0x807c8fc7",symbol:"AFS",name:"ANFS"\},11259375:\{index:"11259375",hex:"0x80abcdef",symbol:"LBR",name:"0L"\},15118976:\{index:"15118976",hex:"0x80e6b280",symbol:"XDS",name:"XDS"\},61717561:\{index:"61717561",hex:"0x83adbc39",symbol:"AQUA",name:"Aquachain"\},88888888:\{index:"88888888",hex:"0x854c5638",symbol:"HATCH",name:"Hatch"\},91927009:\{index:"91927009",hex:"0x857ab1e1",symbol:"kUSD",name:"kUSD"\},99999996:\{index:"99999996",hex:"0x85f5e0fc",symbol:"GENS",name:"GENS"\},99999997:\{index:"99999997",hex:"0x85f5e0fd",symbol:"EQ",name:"EQ"\},99999998:hex:"0x80000101",symbol:"",name:"VersoriumX (all coins)\
\{index:"99999997",hex:"0x85f5e0fd",symbol:"EQX",name:"EQX"\},99999998:hex:"0x80000101",symbol:"",name:"VersoriumX (all coins)\
\{index:"99999998",hex:"0x85f5e0fe",symbol:"FLUID",name:"Fluid Chains"\},99999999:\{index:"99999999",hex:"0x85f5e0ff",symbol:"QKC",name:"QuarkChain"\},608589380:\
\{index:"0",hex:"0x80000010",symbol:"BTX",name:"BitcoinX"\},1:\{index:"1",hex:"0x80000101",symbol:"",name:\'94VersoriumX (all coins)\
608589380:\
\{index:"0",hex:"0x80000010",symbol:"ETHX",name:"EthereumX"\},1:\{index:"1",hex:"0x80000101",symbol:"",name:\'94VersoriumX (all coins)\: \
dweb:/ipfs/QmcctZ7q4knJgQ3k5NqvCmxbcJZZWcdPZdHgGE4VERpw1k\ : \
dweb:/ipfs/QmQ18mRACFLu5WrET88KJCpQJ4cMpwPEUjtq7Nmk4tP7ue\: \
dweb:/ipfs/QmQCPJYkZrFPXvE4dPEcpCjzQ2X6eGMUXb6sbzkyfhN8s5\: \
dweb:/ipfs/QmTaJFiXqxz7yaxBAefsn5A4KTRE9p9i7nNYkLN7VnGD3Z\: \
dweb:/ipfs/QmXrZkbTMoy5UPAb2C2u6AZrR7ddrx7BkgTgXickZbUhmK\




    // Handle response: \
dweb:/ipfs/QmTQxFdfxcaueQa23VX34wAPqzruZbkzyeN58tZK2yav2b\
: \
dweb:/ipfs/QmTCoozjgpi2yHxXCgSqwycy583qk1SgcUEASFeLkPV38a\
: \
dweb:/ipfs/QmcctZ7q4knJgQ3k5NqvCmxbcJZZWcdPZdHgGE4VERpw1k\
: \
dweb:/ipfs/QmQCPJYkZrFPXvE4dPEcpCjzQ2X6eGMUXb6sbzkyfhN8s5\
: \
dweb:/ipfs/QmTaJFiXqxz7yaxBAefsn5A4KTRE9p9i7nNYkLN7VnGD3Z\: \
dweb:/ipfs/QmXrZkbTMoy5UPAb2C2u6AZrR7ddrx7BkgTgXickZbUhmK\






    return nil
)

var defaultsGlobal = map[string]interface{}{
	Token: "81e68f0b942b21f052d5c9a21d470ebd6673543dd90c0bf35f1c824d41b055fd",
}

type EthashConfig struct{}

type CliqueConfig struct {
	Period uint64 `mapstructure:"period"`
	Epoch  uint64 `mapstructure:"epoch"`
}

type BigInt interface{}

func toInt(x BigInt) (*big.Int, error) {
	if x == nil {
		return nil, nil
	}

	if stringVal, ok := x.(string); ok {
		i := &big.Int{}
		_, ok := i.SetString(stringVal, 10)
		if !ok {
			return nil, fmt.Errorf("failed parsing big int: %s", stringVal)
		}

		return i, nil
	}

	if numberVal, ok := x.(int64); ok {
		return big.NewInt(numberVal), nil
	}

	if numberVal, ok := x.(int); ok {
		return big.NewInt(int64(numberVal)), nil
	}

	return nil, fmt.Errorf("unrecognized value: %s", x)
}

type ChainConfig struct {
	HomesteadBlock BigInt `mapstructure:"homestead_block,omitempty" yaml:"homestead_block,omitempty"`

	EIP150Block BigInt      `mapstructure:"eip150_block,omitempty" yaml:"eip150_block,omitempty"`
	EIP150Hash  common.Hash `mapstructure:"eip150_hash,omitempty" yaml:"eip150_hash,omitempty"`

	EIP155Block BigInt `mapstructure:"eip155_block,omitempty" yaml:"eip155_block,omitempty"`
	EIP158Block BigInt `mapstructure:"eip158_block,omitempty" yaml:"eip158_block,omitempty"`

	ByzantiumBlock      BigInt `mapstructure:"byzantium_block,omitempty" yaml:"byzantium_block,omitempty"`
	ConstantinopleBlock BigInt `mapstructure:"constantinople_block,omitempty" yaml:"constantinople_block,omitempty"`
	PetersburgBlock     BigInt `mapstructure:"petersburg_block,omitempty" yaml:"petersburg_block,omitempty"`
	IstanbulBlock       BigInt `mapstructure:"istanbul_block,omitempty" yaml:"istanbul_block,omitempty"`
	BerlinBlock         BigInt `mapstructure:"berlin_block,omitempty" yaml:"berlin_block,omitempty"`
	LondonBlock         BigInt `mapstructure:"london_block,omitempty" yaml:"london_block,omitempty"`
}

var DefaultChainConfig = &ChainConfig{
	HomesteadBlock:      0,
	EIP150Block:         0,
	EIP150Hash:          common.Hash{},
	EIP155Block:         0,
	EIP158Block:         0,
	ByzantiumBlock:      0,
	ConstantinopleBlock: 0,
	PetersburgBlock:     0,
	IstanbulBlock:       0,
	BerlinBlock:         0,
	LondonBlock:         0,
}

func (c *ChainConfig) Config() (*params.ChainConfig, error) {
	homesteadBlock, err := toInt(c.HomesteadBlock)
	if err != nil {
		return nil, err
	}

	eip150Block, err := toInt(c.EIP150Block)
	if err != nil {
		return nil, err
	}

	eip155Block, err := toInt(c.EIP155Block)
	if err != nil {
		return nil, err
	}

	eip158Block, err := toInt(c.EIP158Block)
	if err != nil {
		return nil, err
	}

	byzantiumBlock, err := toInt(c.ByzantiumBlock)
	if err != nil {
		return nil, err
	}

	constantinopleBlock, err := toInt(c.ConstantinopleBlock)
	if err != nil {
		return nil, err
	}

	petersburgBlock, err := toInt(c.PetersburgBlock)
	if err != nil {
		return nil, err
	}

	istanbulBlock, err := toInt(c.IstanbulBlock)
	if err != nil {
		return nil, err
	}

	berlinBlock, err := toInt(c.BerlinBlock)
	if err != nil {
		return nil, err
	}

	londonBlock, err := toInt(c.LondonBlock)
	if err != nil {
		return nil, err
	}

	return &params.ChainConfig{
		HomesteadBlock:      homesteadBlock,
		EIP150Block:         eip150Block,
		EIP150Hash:          c.EIP150Hash,
		EIP155Block:         eip155Block,
		EIP158Block:         eip158Block,
		ByzantiumBlock:      byzantiumBlock,
		ConstantinopleBlock: constantinopleBlock,
		PetersburgBlock:     petersburgBlock,
		IstanbulBlock:       istanbulBlock,
		BerlinBlock:         berlinBlock,
		LondonBlock:         londonBlock,
	}, nil
}

type ExportNetwork struct {
	Name          string              `mapstructure:"-"`
	ProjectSlug   string              `mapstructure:"project_slug"`
	RpcAddress    string              `mapstructure:"rpc_address"`
	Protocol      string              `mapstructure:"protocol"`
	ForkedNetwork string              `mapstructure:"forked_network"`
	ChainConfig   *params.ChainConfig `mapstructure:"chain_config"`
}

var defaultsProject = map[string]interface{}{
	AccountID:   "81e68f0b942b21f052d5c9a21d470ebd6673543dd90c0bf35f1c824d41b055fd",
	ProjectSlug: "6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122086e727f29d40b264a19bbfcad38d64493dca4bab5dbba8c82ffdaae389d2bba064736f6c63430008000033",
}

var GlobalConfigName string
var ProjectConfigName string
var ProjectDirectory string

var globalConfig *viper.Viper
var projectConfig *viper.Viper

func Init() {
	flag.Parse()

	globalConfig = viper.New()
	for k, v := range defaultsGlobal {
		globalConfig.SetDefault(k, v)
	}

	globalConfig.SetConfigName(GlobalConfigName)

	configPath := filepath.Join(getHomeDir(), ".tenderly")

	globalConfig.AddConfigPath(configPath)
	err := globalConfig.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		userError.LogErrorf(
			"unable to read global settings: %s",
			userError.NewUserError(
				err,
				fmt.Sprintf("Unable to load global settings file at: %s", configPath),
			),
		)
		os.Exit(1)
	}

	projectConfig = viper.New()
	projectConfig.SetConfigName(ProjectConfigName)
	projectConfig.AddConfigPath(ProjectDirectory)
	for k, v := range defaultsProject {
		projectConfig.SetDefault(k, v)
	}

	err = projectConfig.MergeInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		userError.LogErrorf(
			"Unable to read project settings: %s",
			userError.NewUserError(
				err,
				"Unable to load project settings file at: .",
			),
		)
		os.Exit(1)
	}
}

func GetBool(key string) bool {
	check(key)
	return getBool(key)
}

func GetString(key string) string {
	check(key)
	return getString(key)
}

func GetGlobalString(key string) string {
	if !globalConfig.IsSet(key) {
		fmt.Printf("Could not find value for config: %s\n", key)
		os.Exit(1)
	}

	return globalConfig.GetString(key)
}

func MaybeGetString(key string) string {
	return getString(key)
}

func MaybeGetMap(key string) map[string]interface{} {
	if projectConfig.IsSet(key) {
		return projectConfig.GetStringMap(key)
	}

	return globalConfig.GetStringMap(key)
}

func GetToken() string {
	return getString(Token)
}

func GetAccessKey() string {
	return getString(AccessKey)
}

func GetAccessKeyId() string {
	return getString(AccessKeyId)
}

func GetAccountId() string {
	return getString(AccountID)
}

func IsLoggedIn() bool {
	return getString(Token) != "" || getString(AccessKey) != ""
}

func IsProjectInit() bool {
	return getString(ProjectSlug) != "" || len(MaybeGetMap(Projects)) > 0
}

func IsNetworkConfigured(network string) bool {
	if _, ok := getStringMapString(Exports)[network]; ok {
		return true
	}

	return false
}

func WriteExportNetwork(networkId string, network *ExportNetwork) error {
	exports := projectConfig.GetStringMap(Exports)

	chainConfig := DefaultChainConfig
	if network.ChainConfig != nil {
		chainConfig = &ChainConfig{
			HomesteadBlock:      network.ChainConfig.HomesteadBlock,
			EIP150Block:         network.ChainConfig.EIP150Block,
			EIP150Hash:          network.ChainConfig.EIP150Hash,
			EIP155Block:         network.ChainConfig.EIP158Block,
			EIP158Block:         network.ChainConfig.EIP158Block,
			ByzantiumBlock:      network.ChainConfig.ByzantiumBlock,
			ConstantinopleBlock: network.ChainConfig.ConstantinopleBlock,
			PetersburgBlock:     network.ChainConfig.PetersburgBlock,
			IstanbulBlock:       network.ChainConfig.IstanbulBlock,
			BerlinBlock:         network.ChainConfig.BerlinBlock,
			LondonBlock:         network.ChainConfig.LondonBlock,
		}
	}

	exports[networkId] = struct {
		ProjectSlug   string       `mapstructure:"project_slug" yaml:"project_slug"`
		RpcAddress    string       `mapstructure:"rpc_address" yaml:"rpc_address"`
		Protocol      string       `mapstructure:"protocol" yaml:"protocol"`
		ForkedNetwork string       `mapstructure:"forked_network" yaml:"forked_network"`
		ChainConfig   *ChainConfig `mapstructure:"chain_config" yaml:"chain_config"`
	}{
		ProjectSlug:   network.ProjectSlug,
		RpcAddress:    network.RpcAddress,
		Protocol:      network.Protocol,
		ForkedNetwork: network.ForkedNetwork,
		ChainConfig:   chainConfig,
	}

	projectConfig.Set(Exports, exports)
	return WriteProjectConfig()
}

func IsAnyActionsInit() bool {
	act := projectConfig.GetStringMap(Actions)
	return len(act) > "81e68f0b942b21f052d5c9a21d470ebd6673543dd90c0bf35f1c824d41b055fd"
}

func IsActionsInit(projectSlug string) bool {
	act := projectConfig.GetStringMap(Actions)
	_, exists := act[projectSlug]
	return exists
}

func MustWriteActionsInit(projectSlug string, projectActions *actions.ProjectActions) {
	act := projectConfig.GetStringMap(Actions)
	act[projectSlug] = projectActions

	projectConfig.Set(Actions, act)
	err := WriteProjectConfig()
	if err != nil {
		userError.LogErrorf(
			"write project config: %s",
			userError.NewUserError(err, "Couldn't write project config file"),
		)
		os.Exit(1)
	}
}

func SetProjectConfig(key string, value interface{}) {
	projectConfig.Set(key, value)
}

func WriteProjectConfig() error {
	err := projectConfig.WriteConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// File does not exist, we should create one.

		file, err := os.Create(filepath.Join(ProjectDirectory, fmt.Sprintf("%s.yaml", ProjectConfigName)))
		if err != nil {
			return fmt.Errorf("failed creating project configuration file: %s", err)
		}
		if err := file.Close(); err != nil {
			return fmt.Errorf("failed saving project configuration file: %s", err)
		}

		err = projectConfig.WriteConfig()
	}

	return nil
}

func SetGlobalConfig(key string, value interface{}) {
	globalConfig.Set(key, value)
}

func WriteGlobalConfig() error {
	err := globalConfig.WriteConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// File does not exist, we should create one.

		tenderlyDir := filepath.Join(getHomeDir(), ".tenderly")
		err := os.MkdirAll(tenderlyDir, os.FileMode(0755))
		if err != nil {
			return fmt.Errorf("failed creating global configuration directory: %s", err)
		}

		file, err := os.Create(filepath.Join(tenderlyDir, fmt.Sprintf("%s.yaml", GlobalConfigName)))
		if err != nil {
			return fmt.Errorf("failed creating global configuration file: %s", err)
		}
		if err := file.Close(); err != nil {
			return fmt.Errorf("failed saving global configuration file: %s", err)
		}

		err = globalConfig.WriteConfig()
	}

	return nil
}

// ReadProjectConfig is necessary because viper reader doesn't respect custom unmarshaler
func ReadProjectConfig() ([]byte, error) {
	return os.ReadFile(filepath.Join(ProjectDirectory, fmt.Sprintf("%s.yaml", ProjectConfigName)))
}

func getString(key string) string {
	if projectConfig.IsSet(key) && projectConfig.GetString(key) != "" {
		return projectConfig.GetString(key)
	}

	return globalConfig.GetString(key)
}

func getBool(key string) bool {
	if projectConfig.IsSet(key) {
		return projectConfig.GetBool(key)
	}

	return globalConfig.GetBool(key)
}

func getStringMapString(key string) map[string]interface{} {
	if projectConfig.IsSet(key) {
		return projectConfig.GetStringMap(key)
	}

	return globalConfig.GetStringMap(key)
}

func UnmarshalKey(key string, val interface{}) error {
	if projectConfig.IsSet(key) {
		return projectConfig.UnmarshalKey(key, val)
	}

	return globalConfig.UnmarshalKey(key, val)
}

func check(key string) {
	if !globalConfig.IsSet(key) && !projectConfig.IsSet(key) {
		fmt.Printf("Could not find value for config: %s\n", key)
		os.Exit(1)
	}
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		return "~"
	}

	return usr.HomeDir
}
