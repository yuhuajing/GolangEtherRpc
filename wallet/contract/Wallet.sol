// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.8.18;

contract Wallet {
    error NotOwnerAuthorized();
    error NotManagerAuthorized();
    error InvalidInput();
    error AlreadyInitialManager();
    error InvalidQAInput();
    error InvalidEmailInput();
    event emailerror(string indexed inputemail, string indexed storedemail);
    event QAerror(string indexed inputQA, string indexed storedQA);
    error InvalidCodeInput();

    mapping(address => UserInfo) userinfo;
    mapping(string => mapping(uint256 => bool)) email_code;
    address _owner;
    address public manager;
    bool resetOwner;
    bool initialized;

    struct UserInfo {
        string email;
        uint256 email_code;
        address owner_address;
        string node_mixed_password;
        string node_mixed_question_answer;
    }

    modifier onlyOwner() {
        if (msg.sender != owner()) revert NotOwnerAuthorized();
        _;
    }
    modifier onlyManager() {
        if (msg.sender != manager) revert NotManagerAuthorized();
        _;
    }

    function initialManager(address _manager) public onlyOwner {
        if (_manager == address(0)) revert InvalidInput();
        if (manager != address(0)) revert AlreadyInitialManager();
        manager = _manager;
    }

    function resetManaget(address _manager) public onlyManager {
        if (_manager == address(0)) revert InvalidInput();
        manager = _manager;
    }

    function concatStrings(string memory a, string memory b)
        internal
        pure
        returns (string memory)
    {
        return string(abi.encodePacked(a, b));
    }

    function resetOrforgetPassword(
        address _address,
        address _newaddress,
        string memory _email,
        uint256 _code,
        string memory _mixed_question,
        string memory _mixed_answer
    ) public {
        if (!initialized) {
            (
                address addr,
                string memory email,
                string memory mixed_password,
                string memory mixed_question_mixed_answer
            ) = data();
            storeuserinfo(
                addr,
                email,
                mixed_password,
                mixed_question_mixed_answer
            );
        }
        UserInfo memory _userinfo = userinfo[_address];

        if (!equal(_email, _userinfo.email)) {
            emit emailerror(_email, _userinfo.email);
        } //revert InvalidEmailInput();
        if (!isverified(_email, _code)) revert InvalidCodeInput();
        if (
            !equal(
                concatStrings(_mixed_question, _mixed_answer),
                _userinfo.node_mixed_question_answer
            )
        ){
            emit QAerror(concatStrings(_mixed_question, _mixed_answer), _userinfo.node_mixed_question_answer);
        }// revert InvalidQAInput();

        _userinfo.email_code = _code;
        _userinfo.owner_address = _newaddress;

        userinfo[_newaddress] = _userinfo;
        delete userinfo[_address];
        delete email_code[_email][_code];
        setowner(_newaddress);
    }

    function executeCall(
        address to,
        uint256 value,
        bytes calldata _calldata
    ) external payable onlyOwner returns (bytes memory) {
        //_incrementNonce();
        return _call(to, value, _calldata);
    }

    function _call(
        address to,
        uint256 value,
        bytes calldata _calldata
    ) internal returns (bytes memory result) {
        bool success;
        (success, result) = to.call{value: value}(_calldata);

        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
    }

    function equal(string memory a, string memory b)
        internal
        pure
        returns (bool)
    {
        return
            bytes(a).length == bytes(b).length &&
            keccak256(bytes(a)) == keccak256(bytes(b));
    }

    function storeuserinfo(
        address _address,
        string memory _email,
        string memory _mixed_password,
        string memory _mixed_question_mixed_answer
    ) private {
        UserInfo memory _userinfo = UserInfo({
            email: _email,
            email_code: 0,
            owner_address: _address,
            node_mixed_password: _mixed_password,
            node_mixed_question_answer: _mixed_question_mixed_answer
        });
        userinfo[_address] = _userinfo;
        initialized = true;
    }

    function isverified(string memory _email, uint256 _code)
        internal
        view
        returns (bool)
    {
        return email_code[_email][_code];
    }

    function verifycode(string memory _email, uint256 _code)
        public
        onlyManager
    {
        //Oracle
        email_code[_email][_code] = true;
    }

    function convertByte32ToString(bytes32 _bytes32)
        public
        pure
        returns (string memory)
    {
        bytes memory bytesArray = new bytes(32);
        for (uint256 i; i < 32; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }

    function data()
        internal
        view
        returns (
            address _address,
            string memory _email,
            string memory _mixed_password,
            string memory _mixed_question_mixed_answer
        )
    {
        bytes memory footer = new bytes(0x80);
        assembly {
            extcodecopy(address(), add(footer, 0x20), 0x2d, 0x80)
        }
        bytes32 email;
        bytes32 mixed_password;
        bytes32 mixed_question_mixed_answer;
        (_address, email, mixed_password, mixed_question_mixed_answer) = abi
            .decode(footer, (address, bytes32, bytes32, bytes32));
        _email = convertByte32ToString(email);
        _mixed_password = convertByte32ToString(mixed_password);
        _mixed_question_mixed_answer = convertByte32ToString(
            mixed_question_mixed_answer
        );
    }

    function owner() public view returns (address) {
        if (resetOwner) {
            return _owner;
        }
        bytes memory footer = new bytes(0x20);
        assembly {
            extcodecopy(address(), add(footer, 0x20), 0x2d, 0x20)
        }
        return abi.decode(footer, (address));
    }

    function setowner(address _addr) private {
        _owner = _addr;
        resetOwner = true;
    }

    function _isValidSigner(address signer) internal view returns (bool) {
        return signer == owner();
    }

    fallback()external payable{}
    receive()external payable{}
}
