geth --dev --http --http.api personal,eth,net,web3,web3 --rpc.enabledeprecatedpersonal --http.corsdomain "https://remix.ethereum.org" --datadir "/Users/sinancskn/Desktop/eth_notify"

geth --datadir "/Users/sinancskn/Desktop/eth_notify" account new

geth attach http://127.0.0.1:8545

eth.accounts

personal.newAccount("password")

eth.getBalance(eth.accounts[0])
eth.getBalance("0x411ee650A394b22a1D684834f2728b6b71E0fE50")

eth.sendTransaction({
    from: eth.accounts[0],
    to: "0x411ee650A394b22a1D684834f2728b6b71E0fE50",
    value: web3.toWei(1, "ether")
})

eth.sendTransaction({
    from: "0xbe327837c0DE653BDCcabFbF588c4e0f01e04097",
    to: "0x411ee650A394b22a1D684834f2728b6b71E0fE50",
    value: web3.toWei(5, "ether")
})

web3.personal.unlockAccount("0xbe327837c0DE653BDCcabFbF588c4e0f01e04097", "password", 600)

geth --datadir "/Users/sinancskn/Desktop/eth_notify" --unlock "0xbe327837c0DE653BDCcabFbF588c4e0f01e04097" --password "password"

# delete all accounts
cd /Users/sinancskn/Desktop/eth_notify
rm UTC--*