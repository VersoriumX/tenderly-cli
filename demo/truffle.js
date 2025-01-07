const HDWalletProvider = require('truffle-hdwallet-provider-privkey');

module.exports = {
    networks: {
        tenderly: {
            provider: function() {
                return new HDWalletProvider(["f15ed637fce3841e13d17484f522653dc587b03d407bac1f529fa450a6170216"], `http://127.0.0.1:8545`);
            },
            network_id: '5777',
            gasPrice: 2000000000 // 2 GWei
        },
    }
};
