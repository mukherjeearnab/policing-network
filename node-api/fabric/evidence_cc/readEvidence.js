const { FileSystemWallet, Gateway } = require("fabric-network");
const path = require("path");

const ccp = require("../ccp/connection-citizen.json");

ReadEvidence = async (user, ID) => {
    const walletPath = path.join(process.cwd(), "wallets");
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: user,
        discovery: { enabled: true, asLocalhost: true },
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork("mainchannel");

    // Get the contract from the network.
    const contract = network.getContract("evidence_cc");

    // Evaluate the specified transaction.
    const result = await contract.evaluateTransaction("readEvidence", ID);

    return JSON.parse(result.toString());
};

module.exports = ReadEvidence;
