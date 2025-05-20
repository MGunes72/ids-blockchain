// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SnortLogger {
    event AlertLogged(uint256 id, address sender, string hash);

    struct Alert {
        address sender;
        string hash;
    }

    mapping(uint256 => Alert) public alerts;
    uint256[] public ids;
    address public owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can log alerts");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function logAlert(uint256 id, string memory hash) public onlyOwner {
        require(bytes(alerts[id].hash).length == 0, "Alert ID already logged");
        alerts[id] = Alert(msg.sender, hash);
        ids.push(id);
        emit AlertLogged(id, msg.sender, hash);
    }

    function getAlert(uint256 id) public view returns (address, string memory) {
        Alert memory a = alerts[id];
        return (a.sender, a.hash);
    }

    function getAllIds() public view returns (uint256[] memory) {
        return ids;
    }
}