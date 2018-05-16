var Person = artifacts.require("./Person.sol");

module.exports = function(deployer) {
  deployer.deploy(Person);
};
