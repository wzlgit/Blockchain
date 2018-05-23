import React, {Component} from 'react'
import SimpleStorageContract from '../build/contracts/SimpleStorage.json'
import getWeb3 from './utils/getWeb3'

import './css/oswald.css'
import './css/open-sans.css'
import './css/pure-min.css'
import './App.css'

const ipfsAPI = require('ipfs-api');
const ipfs = ipfsAPI({
  host: 'localhost',
  port: '5001',
  protocol: 'http'
});

const contract = require('truffle-contract')
const simpleStorage = contract(SimpleStorageContract)
let account;

// Declaring this for later so we can chain functions on SimpleStorage.
let contractInstance;

let saveImageToIPFS = (reader) => {
  return new Promise(function(resolve, reject) {
    const buffer = Buffer.from(reader.result);
    ipfs.add(buffer).then((response) => {
      console.log(response)
      resolve(response[0].hash);
    }).catch((err) => {
      console.error(err)
      reject(err);
    })
  })
}

class App extends Component {
  constructor(props) {
    super(props)

    this.state = {
      blockChainHash: null,
      web3: null,
      address: null,
      imageHash: null,
      isSuccess: false
    }
  }

  componentWillMount() {

    ipfs.swarm.peers(function(err, res) {
      if (err) {
        console.error(err);
      } else {
        // var numPeers = res.Peers === null ? 0 : res.Peers.length;
        // console.log("IPFS - connected to " + numPeers + " peers");
        console.log(res);
      }
    });

    getWeb3.then(results => {
      this.setState({web3: results.web3})

      // Instantiate contract once web3 provided.
      this.instantiateContract()
    }).catch(() => {
      console.log('Error finding web3.')
    })
  }

  instantiateContract = () => {

    simpleStorage.setProvider(this.state.web3.currentProvider);
    this.state.web3.eth.getAccounts((error, accounts) => {
      account = accounts[0];
      simpleStorage.at('0xb177d6cf6916f128c9931e610c63208d9c5a40f3').then((contract) => {
        console.log(contract.address);
        contractInstance = contract;
        this.setState({address: contractInstance.address});
        return;
      });
    })

  }
  render() {
    return (
      <div className="App">
        <div style={{marginTop:10}}>智能合约地址：</div>
        <div>{this.state.address}</div>
        <div style={{marginTop:10}}>上传图片到IPFS：</div>
        <div>
          <label id="file">选择图片</label>
          <input type="file" ref="file" id="file" name="file" multiple="multiple"/>
        </div>
        <button style={{marginTop:10}} onClick={() => {
            var file = this.refs.file.files[0];
            var reader = new FileReader();
            reader.readAsArrayBuffer(file)
            reader.onloadend = function(e) {
              console.log(reader);
              saveImageToIPFS(reader).then((hash) => {
                console.log(hash);
                this.setState({imageHash: hash})
              });

          }.bind(this);
      }}>开始上传</button>

      <div style={{marginTop:10}}>图片哈希值：{this.state.imageHash}</div>
      <button onClick={() => {
           contractInstance.set(this.state.imageHash, {from: account}).then(() => {
              console.log('图片的hash已经写入到区块链！');
              this.setState({isSuccess: true});
          })
      }}>图片哈希写入区块链</button>

      {
        this.state.isSuccess
          ? <div style={{marginTop:10}}>
              <div>图片哈希成功写入区块链！</div>
              <button onClick={() => {
                  contractInstance.get({from: account}).then((data) => {
                    console.log(data);
                    this.setState({blockChainHash: data});
                  })
                }}>从区块链读取图片哈希</button>
            </div>
          : <div/>
      }
      {
        this.state.blockChainHash
          ? <div style={{marginTop:10}}>
              <div>从区块链读取到的哈希值：{this.state.blockChainHash}</div>
            </div>
          : <div/>
      }
      {
        this.state.blockChainHash
          ? <div style={{marginTop:10}}>
              <div>访问本地文件：</div>
              <div>{"http://localhost:8082/ipfs/" + this.state.imageHash}</div>
              <div>访问PFS网关：</div>
              <div>{"https://ipfs.io/ipfs/" + this.state.imageHash}</div>
              <img alt="" style={{width: 100, height: 100 }} src={"https://ipfs.io/ipfs/" + this.state.imageHash}/>
            </div>
          : <img alt=""/>
      }

    </div>);
  }
}

export default App
