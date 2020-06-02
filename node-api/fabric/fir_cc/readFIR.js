const { FileSystemWallet, Gateway } = require("fabric-network");
const path = require("path");

ReadFIR = async (user, ID) => {
    const ccp = require(`../ccp/connection-${user.group}.json`);
    const walletPath = path.join(process.cwd(), `wallet_${user.group}`);
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: user.username,
        discovery: { enabled: true, asLocalhost: true },
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork("mainchannel");

    // Get the contract from the network.
    const contract = network.getContract("fir_cc");

    // Evaluate the specified transaction.
    const result = await contract.evaluateTransaction("readFIR", ID);

    return JSON.parse(result.toString());
};

module.exports = ReadFIR;
