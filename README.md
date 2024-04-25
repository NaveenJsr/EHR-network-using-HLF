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

org1:
```bash
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/HospitalMSP.tx --channelID $CHANNEL_NAME -asOrg Hospital
```
org2:
```bash
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/PatientMSP.tx --channelID $CHANNEL_NAME -asOrg Patient
```
org3:
```bash
  configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/LabsMSP.tx --channelID $CHANNEL_NAME -asOrg Labs
```
org4:
```bash
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
  peer channel create -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem --outputBlock ./channel-artifacts/ehrchannel.block

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

org1:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/users/Admin@Hospital-ehr.com/msp 
  export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051
  export CORE_PEER_LOCALMSPID=HospitalMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/HospitalMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem
```
org2:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/users/Admin@Patient-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Patient-ehr.com:7051
  export CORE_PEER_LOCALMSPID=PatientMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/PatientMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem
```
org3:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/users/Admin@Labs-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Labs-ehr.com:7051
  export CORE_PEER_LOCALMSPID=LabsMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/LabsMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem
```
org4:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/users/Admin@Insurance-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Insurance-ehr.com:7051
  export CORE_PEER_LOCALMSPID=InsuranceMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt
  peer channel update -o orderer.ehr.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/InsuranceMSP.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/msp/tlscacerts/tlsca.ehr.com-cert.pem
```

### Install ChainCode

```bash
  cd ../../..
  cd chaincode/go
  go mod init
  go mod tidy
  go build
```

```bash
  peer lifecycle chaincode package ehrcc.tar.gz \
  --path . \
  --lang golang \
  --label ehrcc_1.0
```

org1:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/users/Admin@Hospital-ehr.com/msp 
  export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051
  export CORE_PEER_LOCALMSPID=HospitalMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt
  peer chaincode install -n ehrcc -v 1.0 -p /opt/gopath/src/github.com/chaincode/ehr

```
org2:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/users/Admin@Patient-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Patient-ehr.com:7051
  export CORE_PEER_LOCALMSPID=PatientMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt
  peer chaincode install -n ehrcc -v 1.0 -p /opt/gopath/src/github.com/chaincode/ehr

```
org3:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/users/Admin@Labs-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Labs-ehr.com:7051
  export CORE_PEER_LOCALMSPID=LabsMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt
  peer chaincode install -n ehrcc -v 1.0 -p /opt/gopath/src/github.com/chaincode/ehr

```
org4:
```bash
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/users/Admin@Insurance-ehr.com/msp
  export CORE_PEER_ADDRESS=peer0.Insurance-ehr.com:7051
  export CORE_PEER_LOCALMSPID=InsuranceMSP
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt
  peer chaincode install -n ehrcc -v 1.0 -p /opt/gopath/src/github.com/chaincode/ehr

```

### Aprove chaincode :

org1:
```bash
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt
  export CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/server.crt
  export CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/server.key
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/users/Admin@Hospital-ehr.com/msp
  export CORE_PEER_LOCALMSPID=HospitalMSP
  export CORE_PEER_ADDRESS=peer0.Hospital-ehr.com:7051
  
  peer lifecycle chaincode approveformyorg \
    --channelID ehrchannel \
    --name ehrcc \
    --version 1.0 \
    --package-id <id> \
    --sequence 1 \
    --init-required \
    --signature-policy "OR('HospitalMSP.peer','PatientMSP.peer','LabsMSP.peer','InsuranceMSP.peer')" \
    --orderer orderer.ehr.com:7050 \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/tls/ca.crt \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt

```

org2:
```bash
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt
  export CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/server.crt
  export CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/server.key
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/users/Admin@Patient-ehr.com/msp
  export CORE_PEER_LOCALMSPID=PatientMSP
  export CORE_PEER_ADDRESS=peer0.Patient-ehr.com:7051
  
  peer lifecycle chaincode approveformyorg \
    --channelID ehrchannel \
    --name ehrcc \
    --version 1.0 \
    --package-id <id> \
    --sequence 1 \
    --init-required \
    --signature-policy "OR('HospitalMSP.peer','PatientMSP.peer','LabsMSP.peer','InsuranceMSP.peer')" \
    --orderer orderer.ehr.com:7050 \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/tls/ca.crt \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt

```

org3:
```bash
  export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt
  export CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/server.crt
  export CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/server.key
  export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/users/Admin@Labs-ehr.com/msp
  export CORE_PEER_LOCALMSPID=LabsMSP
  export CORE_PEER_ADDRESS=peer0.Labs-ehr.com:7051
  
  peer lifecycle chaincode approveformyorg \
    --channelID ehrchannel \
    --name ehrcc \
    --version 1.0 \
    --package-id <id> \
    --sequence 1 \
    --init-required \
    --signature-policy "OR('HospitalMSP.peer','PatientMSP.peer','LabsMSP.peer','InsuranceMSP.peer')" \
    --orderer orderer.ehr.com:7050 \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/tls/ca.crt \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt

```

org4:
```bash
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt
export CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/server.crt
export CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/server.key
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/users/Admin@Insurance-ehr.com/msp
export CORE_PEER_LOCALMSPID=InsuranceMSP
export CORE_PEER_ADDRESS=peer0.Insurance-ehr.com:7051

peer lifecycle chaincode approveformyorg \
  --channelID ehrchannel \
  --name ehrcc \
  --version 1.0 \
  --package-id <id> \
  --sequence 1 \
  --init-required \
  --signature-policy "OR('HospitalMSP.peer','PatientMSP.peer','LabsMSP.peer','InsuranceMSP.peer')" \
  --orderer orderer.ehr.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/tls/ca.crt \
  --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt

```


### check commit readiness

```bash
  peer lifecycle chaincode checkcommitreadiness \
    --channelID ehrchannel \
    --name ehrcc \
    --version 1.0 \
    --sequence 1 \
    --output json \
    --init-required \
    --signature-policy "OR('HospitalMSP.peer','PatientMSP.peer','LabsMSP.peer','InsuranceMSP.peer')" \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/tls/ca.crt \
    --peerAddresses peer0.Insurance-ehr.com:7051 \
    --tls \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt
```

### commit chaincode

```bash
  peer lifecycle chaincode commit \
    --channelID ehrchannel \
    --name ehrcc \
    --version 1.0 \
    --sequence 1 \
    --init-required \
    --signature-policy "OR('HospitalMSP.peer','PatientMSP.peer','LabsMSP.peer','InsuranceMSP.peer')" \
    --peerAddresses peer0.Hospital-ehr.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Hospital-ehr.com/peers/peer0.Hospital-ehr.com/tls/ca.crt \
    --peerAddresses peer0.Patient-ehr.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Patient-ehr.com/peers/peer0.Patient-ehr.com/tls/ca.crt \
    --peerAddresses peer0.Labs-ehr.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Labs-ehr.com/peers/peer0.Labs-ehr.com/tls/ca.crt \
    --peerAddresses peer0.Insurance-ehr.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Insurance-ehr.com/peers/peer0.Insurance-ehr.com/tls/ca.crt \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/ehr.com/orderers/orderer.ehr.com/tls/ca.crt \
    --orderer orderer.ehr.com:7050

```
