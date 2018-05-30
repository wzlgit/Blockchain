import React, {
  Component
}
from 'react'
import VotingContract from '../build/contracts/Voting.json'
import getWeb3 from './utils/getWeb3'

import './css/oswald.css'
import './css/open-sans.css'
import './css/pure-min.css'
import './App.css'
import './main.css'

// 合约的实例
var votingContractInstance;
// 投票合约地址
const contractAddr = "0x7829abb7d424f118f1243ec13207a02f5bffc114"

class App extends Component {
  constructor(props) {
    super(props)

        this.state = {
          storageValue: 0,
          web3: null,
          candidateList: [{
            "candidateId": 1,
            "candidateName": "Jack Chen",
            "voteNumber": 0
          }, {
            "candidateId": 2,
            "candidateName": "Andy Lau",
            "voteNumber": 0
          }, {
            "candidateId": 3,
            "candidateName": "Stephen Chow",
            "voteNumber": 0
          }, {
            "candidateId": 4,
            "candidateName": "Wenzil",
            "voteNumber": 0
          }]
        }
  }

  componentWillMount() {
    // Get network provider and web3 instance.
    // See utils/getWeb3 for more info.

    getWeb3
      .then(results => {
        this.setState({
          web3: results.web3
        })

        // Instantiate contract once web3 provided.
        this.instantiateContract()
      })
      .catch(() => {
        console.log('Error finding web3.')
      })
  }

  instantiateContract() {
    /*
     * SMART CONTRACT EXAMPLE
     *
     * Normally these functions would be called in the context of a
     * state management library, but for convenience I've placed them here.
     */

    const contract = require('truffle-contract')
    const votingContract = contract(VotingContract)
    votingContract.setProvider(this.state.web3.currentProvider)

    // Declaring this for later so we can chain functions on SimpleStorage.

    // Get accounts.
    this.state.web3.eth.getAccounts((error, accounts) => {
      votingContract.at(contractAddr).then((instance) => {
        votingContractInstance = instance
        // 读取合约成功后，遍历所有候选人的票数，并更新到前端
        var candidateListTemp = this.state.candidateList;
        for (let i = 0; i < this.state.candidateList.length; i++) {
          let object = candidateListTemp[i];
          votingContractInstance.totalVotesFor(object.candidateName).then(result => {
              object.voteNumber = result.c[0]
              this.setState({
                candidateList: candidateListTemp
              })
          });
        }        
      })
    })
  }

  render() {
    return ( 
      <div className="App" style={{margin:20}}>
        <div className="pageTitle">以太坊DApp投票系统</div> 
        <div className="dataContent">
          <div className="dataItem">候选人</div>
          <div className="dataItem">票数</div>
        </div>
        {/* 读取当前页面的候选人列表，并显示 */}
        {
         this.state.candidateList.map((item) => {
           return (
              <div className="dataContent" key={item.candidateId}>
                <div className="dataItem">{item.candidateName}</div>
                <div className="dataItem">{item.voteNumber}</div>
              </div>
           )
         })
        }
        <input className="inputName" placeholder="请输入候选人名字" ref="inputName" />
        {/* 获取输入框的内容 */}
        <button className="voteButton"
          onClick={() => {
            let candidateName = this.refs.inputName.value; 
            let account = this.state.web3.eth.accounts[0];
            {/* 为输入的候选人投票，写入到区块链中 */}
            votingContractInstance.voteForCandidate(candidateName,{from: account}).then((result) => {
              var currentCandidateIndex = 0;
              {/* 获取输入候选人的数组下标 */}

              for(let i = 0; i < this.state.candidateList.length; i++) {
                let currentCandidate = this.state.candidateList[i]; 
                if (currentCandidate.candidateName === candidateName) {
                   currentCandidateIndex = i;
                  break; 
                }
              }
              {/* 根据合约读取区块中的候选人列表数据，并刷新到前端页面 */}
              votingContractInstance.totalVotesFor(candidateName).then(result => {
                var candidateListTemp = this.state.candidateList;
                {/* 根据输入候选人的数组下标来更新列表数据 */}
                for (let i = 0; i < candidateListTemp.length; i++) {
                  if (candidateName === candidateListTemp[currentCandidateIndex].candidateName) {
                    candidateListTemp[currentCandidateIndex].voteNumber = result.c[0]
                  }
                }
                this.setState({
                  candidateList: candidateListTemp
                })
              })

            })
          }}>投票</button>
      </div>
    );
  }
}

export default App
