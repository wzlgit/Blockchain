var TutorialToken = artifacts.require("./TutorialToken.sol");

module.exports = function(deployer) {
  deployer.deploy(TutorialToken);
};
