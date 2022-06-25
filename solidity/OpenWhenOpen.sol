// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.14;

contract MyContract {
    mapping(uint => Person) public people;
    uint256 public peopleCount;

    uint256 openingTime = 1656179058;

    modifier onlyWhileOpen() {
        require(block.timestamp >= openingTime);
        _;
    }

    struct Person {
        uint _id;
        string _firstName;
        string _lastName;
    }

    function addPerson(string memory _firstName, string memory _lastName) public onlyWhileOpen {
        incrementCount();
        people[peopleCount] = Person(peopleCount, _firstName, _lastName);
    }

    function incrementCount() internal {
         peopleCount += 1;
    }
}
