1. First install NPM v10.22.1
2. Run these commands to install Hyperledger Caliper CLI tool and bind to Fabric v1.4.4 SUT.
    
    $ npm install -g --only=prod @hyperledger/caliper-cli@0.3.2
    $ caliper bind --caliper-bind-sut fabric:1.4.4 --caliper-bind-args=-g

3. Update "tlsCACerts", "signedCert", "adminPrivateKey" in ./networks/network_config.json

    NOTE: "tlsCACerts" can be fetched from the connection profiles JSON created in ../connections/
    NOTE: "signedCert", "adminPrivateKey" are located in ../backend/crypto_config/ once generate.sh is run to init. the network.

4. Run the command.sh to run the benchmark.

    $ bash command.sh

5. Check report.html once the test runs successfully!