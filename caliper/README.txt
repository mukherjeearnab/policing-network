1. First install NPM v10.22.1
2. Run these commands to install Hyperledger Caliper CLI tool and bind to Fabric v1.4.4 SUT.
    
    $ npm install -g --only=prod @hyperledger/caliper-cli@0.3.2
    $ caliper bind --caliper-bind-sut fabric:1.4.4 --caliper-bind-args=-g

3. Run the generate.sh to generate the network_config.json file.

    $ bash command.sh

4. Run the run.sh to run the benchmark.

    $ bash run.sh

5. Check report.html once the test runs successfully!