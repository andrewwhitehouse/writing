// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.14;

contract MyContract {
    mapping(uint => Person) public people;
    uint256 public peopleCount;

    address owner;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    struct Person {
        uint _id;
        string _firstName;
        string _lastName;
    }

    constructor() {
        owner = msg.sender;
    }

    function addPerson(string memory _firstName, string memory _lastName) public onlyOwner {
        incrementCount();
        people[peopleCount] = Person(peopleCount, _firstName, _lastName);
    }

    function incrementCount() internal {
         peopleCount += 1;
    }
}
