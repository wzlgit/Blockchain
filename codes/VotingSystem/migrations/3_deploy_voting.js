var Voting = artifacts.require("./Voting.sol");

module.exports = function(deployer) {
  deployer.deploy(Voting, ['Jack Chen', 'Andy Lau', 'Stephen Chow', 'Wenzil'], {
    gas: 6700000
  });
};
