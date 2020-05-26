#docker-compose -f docker-compose-couchdb.yaml up -d
docker-compose -f docker-compose-cli.yaml up -d

docker exec -it cli bash ./scripts/channel/create-channel.sh

docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 police PoliceMSP 8051 1.0
docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 forensics ForensicsMSP 9051 1.0
docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 court CourtMSP 10051 1.0
docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 identityprovider IdentityProviderMSP 11051 1.0

docker cp ~/go/src/fcc cli:/opt/gopath/src/fcc

docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh fcc peer0 citizen CitizenMSP 7051 1.0
docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh fcc peer0 police PoliceMSP 8051 1.0
docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh fcc peer0 forensics ForensicsMSP 9051 1.0
docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh fcc peer0 court CourtMSP 10051 1.0
docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh fcc peer0 identityprovider IdentityProviderMSP 11051 1.0

docker exec -it cli bash ./scripts/install-cc/instantiate.sh fcc
