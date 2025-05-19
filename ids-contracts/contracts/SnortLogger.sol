// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SnortLogger {
    event AlertLogged(address indexed sender, string message, uint256 timestamp);

    struct Alert {
        address sender;
        string message;
        uint256 timestamp;
    }

    address public owner;

    mapping(uint256 => Alert) public alerts; // timestamp => Alert
    uint256[] public timestamps;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can call this");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function logAlert(uint256 timestamp, string memory message) public onlyOwner {
        require(alerts[timestamp].timestamp == 0, "Alert with this timestamp already exists");
        alerts[timestamp] = Alert(msg.sender, message, timestamp);
        timestamps.push(timestamp);
        emit AlertLogged(msg.sender, message, timestamp);
    }

    function getAlert(uint256 timestamp) public view returns (address, string memory, uint256) {
        require(alerts[timestamp].timestamp != 0, "No alert at this timestamp");
        Alert memory alert = alerts[timestamp];
        return (alert.sender, alert.message, alert.timestamp);
    }

    function getAllTimestamps() public view returns (uint256[] memory) {
        return timestamps;
    }

    function getAlertCount() public view returns (uint256) {
        return timestamps.length;
    }
}