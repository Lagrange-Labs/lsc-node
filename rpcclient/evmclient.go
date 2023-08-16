package rpcclient

import (
	"context"
	"math/big"
        "errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
        rlp "github.com/ethereum/go-ethereum/rlp"
        common "github.com/ethereum/go-ethereum/common"
        hexutil "github.com/ethereum/go-ethereum/common/hexutil"

	_ "github.com/Lagrange-Labs/lagrange-node/governance/arbitrum"
	_ "github.com/Lagrange-Labs/lagrange-node/governance/optimism"
	//"github.com/Lagrange-Labs/lagrange-node/config"
)

type EvmClient struct {
	ethClient *ethclient.Client
}

var _ RpcClient = (*EvmClient)(nil)

// CreateRPCClient creates a new rpc client.
func CreateRPCClient(chain, rpcURL string) (RpcClient, error) {
	switch chain {
	case "arbitrum":
		return NewEvmClient(rpcURL)
	case "optimism":
		return NewEvmClient(rpcURL)
	default:
		return nil, nil
	}
}

// NewEvmClient creates a new EvmClient instance.
func NewEvmClient(rpcURL string) (*EvmClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &EvmClient{
		ethClient: client,
	}, nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *EvmClient) GetBlockHashByNumber(blockNumber uint64) (string, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err == ethereum.NotFound {
		return "", ErrBlockNotFound
	}

	return header.Hash().Hex(), err
}

// GetExtraDataByNumber returns the block extradata by the given block number.
func (c *EvmClient) GetExtraDataByNumber(blockNumber uint64) (string, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err == ethereum.NotFound {
		return "", ErrBlockNotFound
	}
	extraData := header.Extra

	return hexutil.Encode(extraData), err
}

// GetBlockHashByNumber returns the block number by the given block hash.
func (c *EvmClient) GetBlockNumberByHash(blockHash string) (int, error) {
	header, err := c.ethClient.HeaderByHash(context.Background(), common.HexToHash(blockHash))
	if err == ethereum.NotFound {
		return 0, ErrBlockNotFound
	}

	return int(header.Number.Int64()), err
}

// GetChainID returns the chain ID.
func (c *EvmClient) GetChainID() (uint32, error) {
	chainID, err := c.ethClient.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	return uint32(chainID.Int64()), err
}

func GetRawAttestBlockHeader(blockNum int) (string, error) {
        return "0x00",nil
/*
	cfg,err := config.Default()
	if err != nil { panic(err) }
	optClient,err := NewEvmClient(cfg.RPCEndpoint)
	if err != nil { panic(err) }
	hex,err := optClient.GetRawBlockHeader(blockNum)
	return hex,err
*/
}

func (c *EvmClient) GetRawBlockHeader(blockNum int) (string, error) {
        header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil { return "",err }
	rlpBytes, err := rlp.EncodeToBytes(header)
	if err != nil { return "",err }	
        hex := hexutil.Encode(rlpBytes)
	return hex,nil
}

// Retrieve RLP-encoded block headers, boundary-inclusive.
func (c *EvmClient) GetRawBlockHeaders(startblock *big.Int, endblock *big.Int) (map[*big.Int]string, error) {
    headers := make(map[*big.Int]string)
    // Iterate block numbers
    for i := (*startblock).Int64(); i <= (*endblock).Int64(); i++ {
        hex,err := c.GetRawBlockHeader(int(i))
        if err != nil {
	    return headers,err
	}
        // Collect raw block header
        headers[big.NewInt(i)] = hex
    }
    return headers,nil
}

func GetExtraDataByNetwork(blockNum int) (string, common.Hash, error) {
	    return "", common.HexToHash("0x00"),errors.New("GetExtraDataByNetwork(): Unsupported network")
	    /*
	cfg,err := config.Default()
	network := cfg.Chain
	if network == "arbitrum" {
	    proofCfg := arbitrum.ProofConfig{cfg.EthereumURL, cfg.RPCendpoint, cfg.Outbox}
	    l2Hash,err := arbitrum.GetL2Hash(proofCfg, blockNum)
	    if err != nil { panic(err) }
	    eth,err := NewEvmClient(cfg.EthereumURL)
	    if err != nil { panic(err) }
	    extra,err := eth.GetExtraDataByNumber(uint64(blockNum))
	    if err != nil { panic(err) }
	    return extra, common.HexToHash(l2Hash), nil
	} else if network == "optimism" {
	    proofCfg := optimism.ProofConfig{cfg.EthereumURL, cfg.RPCendpoint, cfg.L2OutputOracle}
	    proof, err := optimism.GetProof(proofCfg, blockNum)
	    if err != nil { panic(err) }
	    proofHex, err := proof.Hex()
	    if err != nil { panic(err) }
	    return proofHex, proof.LatestBlockhash, nil
	} else {
	    return "", common.HexToHash("0x00"),errors.New("GetExtraDataByNetwork(): Unsupported network")
	}
	*/
}