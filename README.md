# EHR-network-using-HLF

## How to run

### Create crypto Materials :

```bash
  cryptogen generate --config=./crypto-config.yaml
```
### Generate Genesis Block :

```bash
  configtxgen -profile FourOrgsOrdererGenesis -channelID system-channel -configPath=./ -outputBlock ./channel-artifacts/genesis.tx
```

### Generate channel configuration transaction :

```bash
  export CHANNEL_NAME=ehrchannel
  configtxgen -profile FourOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx --channelID $CHANNEL_NAME
```

### Generate define anchor peers for each organization

```bash
  org1:
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/HospitalMSP.tx --channelID $CHANNEL_NAME -asOrg Hospital

  org2:
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/PatientMSP.tx --channelID $CHANNEL_NAME -asOrg Patient

  org3:
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/LabsMSP.tx --channelID $CHANNEL_NAME -asOrg Labs

  org4:
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/InsuranceMSP.tx --channelID $CHANNEL_NAME -asOrg Insurance
```

### Start the Hyperledger Fabric blockchain network (Start the Docker Containers) :

```bash
  docker-compose -f docker-compose-cli.yaml up -d
  docker exec -it cli bash
```

### Create the channel (org1) :

```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/users/Admin@Hospital-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051
  export CORE_PEER_LOCALMSPID="HospitalMSP"
  export CHANNEL_NAME=ehrchannel
  channel create -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem --outputBlock ./channel-artifacts/ehrchannel.block

```


### Join Channel (org1) :
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/users/Admin@Hospital-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051
  export CORE_PEER_LOCALMSPID=HospitalMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt
  peer channel join -b ./channel-artifacts/ehrchannel.block
```

### Join Channel (org2) :
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/users/Admin@Patient-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Patient-ehr.com:7051
  export CORE_PEER_LOCALMSPID=PatientMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt
  peer channel join -b ./channel-artifacts/ehrchannel.block
```

### Join Channel (org3) :
  ```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/users/Admin@Labs-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Labs-ehr.com:7051
  export CORE_PEER_LOCALMSPID=LabsMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt
  peer channel join -b ./channel-artifacts/ehrchannel.block
```

### Join Channel (org4) :
  ```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/users/Admin@Insurance-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Insurance-ehr.com:7051
  export CORE_PEER_LOCALMSPID=InsuranceMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt
  peer channel join -b ./channel-artifacts/ehrchannel.block
```

### Update anchor peers
```bash
  org1:
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/users/Admin@Hospital-ehr.com/msp 
  export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051
  export CORE_PEER_LOCALMSPID=HospitalMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/HospitalMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem

  org2:
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/users/Admin@Patient-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Patient-ehr.com:7051
  export CORE_PEER_LOCALMSPID=PatientMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/PatientMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem

  org3:
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/users/Admin@Labs-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Labs-ehr.com:7051
  export CORE_PEER_LOCALMSPID=LabsMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/LabsMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem

  org4:
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/users/Admin@Insurance-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Insurance-ehr.com:7051
  export CORE_PEER_LOCALMSPID=InsuranceMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/InsuranceMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem

```

### Install ChainCode
