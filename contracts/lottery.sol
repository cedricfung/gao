// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.4 <0.9.0;

import {BytesLib} from './bytes.sol';
import {BLS} from './bls.sol';

contract MixinProcess {
  using BytesLib for bytes;
  using BLS for uint256[2];
  using BLS for bytes;

  event MixinTransaction(bytes);
  event MixinEvent(address indexed sender, uint256 nonce, uint128 asset, uint256 amount, uint64 timestamp, bytes extra);

  uint256[4] public GROUP = [
    0x2f741961cea2e88cfa2680eeaac040d41f41f3fedb01e38c06f4c6058fd7e425, // x.y
    0x007d68aef83f9690b04f463e13eadd9b18f4869041f1b67e7f1a30c9d1d2c42c, // x.x
    0x2a32fa1736807486256ad8dc6a8740dfb91917cf8d15848133819275be92b673, // y.y
    0x257ad901f02f8a442ccf4f1b1d0d7d3a8e8fe791102706e575d36de1c2a4a40f  // y.x
  ];

  // PID is the app id in Mixin which the contract will process, e.g. c6d0c728-2624-429b-8e0d-d9d19b6592fa
  // The app id will add 0x as prefix and delete '-'
  uint128 public constant PID = 0xaa26639c756f358ca184651de77fbefd;
  uint64 public NONCE = 0;
  mapping(uint128 => uint256) public custodian;
  mapping(address => bytes) public members;

  uint128 public ASSET = 0xc94ac88f46713976b60a09064f1811e8;
  uint256 public AMOUNT = 1000000;
  uint64 public TIMESTAMP = 0;
  uint64 public GRACE = 24 * 3600 * 1000000000;
  mapping(address => uint256) public ROUND;
  address[] participants;

  function stats() public view returns (uint256) {
    uint256 score = 0;
    for (uint i = 0; i < participants.length; i++) {
      score = score + ROUND[participants[i]];
    }
    return score;
  }

  function work(address sender, uint64 nonce, uint128 asset, uint256 amount, uint64 timestamp, bytes memory sig) internal returns (bool) {
    require(timestamp > 0, "invalid timestamp");

    if (asset != ASSET) {
      return false;
    }
    if (amount / AMOUNT < 1) {
      return false;
    }
    if (TIMESTAMP == 0) {
      TIMESTAMP = timestamp;
    }
    if (TIMESTAMP + GRACE < timestamp) {
      finish(nonce, sig);
      TIMESTAMP = timestamp;
    }

    if (ROUND[sender] == 0) {
      participants.push(sender);
    }
    ROUND[sender] = ROUND[sender] + amount / AMOUNT;
    return true;
  }

  function finish(uint64 nonce, bytes memory sig) internal returns (bool) {
    uint256 rand = sig.toUint256(0);
    uint256 score = stats();
    if (score == 0) {
      return false;
    }

    uint256 seq = 0;
    uint256 winner = rand % score;
    for (uint i = 0; i < participants.length; i++) {
      if (winner >= seq && winner < seq + ROUND[participants[i]]) {
        winner = i;
        break;
      }
      seq = seq + ROUND[participants[i]];
    }

    uint256 amount = score * AMOUNT;
    address receiver = participants[winner];
    bytes memory log = encodeMixinEvent(nonce, ASSET, amount, "GAO", members[receiver]);
    emit MixinTransaction(log);

    for (uint i = 0; i < participants.length; i++) {
      delete ROUND[participants[i]];
    }
    delete participants;
    return true;
  }

  // process || nonce || asset || amount || extra || timestamp || members || threshold || sig
  function mixin(bytes calldata raw) public returns (bool) {
    require(raw.length >= 141, "event data too small");

    uint256 size = 0;
    uint256 offset = 0;
    uint128 process = raw.toUint128(offset);
    require(process == PID, "invalid process");
    offset = offset + 16;

    uint64 nonce = raw.toUint64(offset);
    require(nonce == NONCE, "invalid nonce");
    NONCE = NONCE + 1;
    offset = offset + 8;

    uint128 asset = raw.toUint128(offset);
    offset = offset + 16;

    size = raw.toUint16(offset);
    offset = offset + 2;
    require(size <= 32, "integer out of bounds");
    uint256 amount = new bytes(32 - size).concat(raw.slice(offset, size)).toUint256(0);
    offset = offset + size;

    size = raw.toUint16(offset);
    offset = offset + 2;
    bytes memory extra = raw.slice(offset, size);
    offset = offset + size;

    uint64 timestamp = raw.toUint64(offset);
    offset = offset + 8;

    size = raw.toUint16(offset);
    size = 2 + size * 16 + 2;
    bytes memory sender = raw.slice(offset, size);
    offset = offset + size;

    offset = offset + 2;
    bytes memory sig = raw.slice(offset, 64);
    require(verifySignature(raw, offset), "invalid signature");

    offset = offset + 64;
    require(raw.length == offset, "malformed event encoding");

    custodian[asset] = custodian[asset] + amount;
    members[mixinSenderToAddress(sender)] = sender;

    emit MixinEvent(mixinSenderToAddress(sender), nonce, asset, amount, timestamp, extra);
    return work(mixinSenderToAddress(sender), nonce, asset, amount, timestamp, sig);
  }

  // pid || nonce || asset || amount || extra || timestamp || members || threshold || sig
  function encodeMixinEvent(uint64 nonce, uint128 asset, uint256 amount, bytes memory extra, bytes memory receiver) internal returns (bytes memory) {
    require(extra.length < 128, "extra too large");
    require(custodian[asset] >= amount, "insufficient custodian");
    custodian[asset] = custodian[asset] - amount;
    bytes memory raw = uint128ToFixedBytes(PID);
    raw = raw.concat(uint64ToFixedBytes(nonce));
    raw = raw.concat(uint128ToFixedBytes(asset));
    (bytes memory ab, uint16 al) = uint256ToVarBytes(amount);
    raw = raw.concat(uint16ToFixedBytes(al));
    raw = raw.concat(ab);
    raw = raw.concat(uint16ToFixedBytes(uint16(extra.length)));
    raw = raw.concat(extra);
    raw = raw.concat(new bytes(8));
    raw = raw.concat(receiver);
    raw = raw.concat(new bytes(2));
    return raw;
  }

  function verifySignature(bytes memory raw, uint256 offset) internal view returns (bool) {
    uint256[2] memory sig = [raw.toUint256(offset), raw.toUint256(offset+32)];
    uint256[2] memory message = raw.slice(0, offset-2).concat(new bytes(2)).hashToPoint();
    return sig.verifySingle(GROUP, message);
  }

  function mixinSenderToAddress(bytes memory sender) internal pure returns (address) {
    return address(uint160(uint256(keccak256(sender))));
  }

  function uint16ToFixedBytes(uint16 x) internal pure returns (bytes memory) {
    bytes memory c = new bytes(2);
    bytes2 b = bytes2(x);
    for (uint i=0; i < 2; i++) {
      c[i] = b[i];
    }
    return c;
  }

  function uint64ToFixedBytes(uint64 x) internal pure returns (bytes memory) {
    bytes memory c = new bytes(8);
    bytes8 b = bytes8(x);
    for (uint i=0; i < 8; i++) {
      c[i] = b[i];
    }
    return c;
  }

  function uint128ToFixedBytes(uint128 x) internal pure returns (bytes memory) {
    bytes memory c = new bytes(16);
    bytes16 b = bytes16(x);
    for (uint i=0; i < 16; i++) {
      c[i] = b[i];
    }
    return c;
  }

  function uint256ToVarBytes(uint256 x) internal pure returns (bytes memory, uint16) {
    bytes memory c = new bytes(32);
    bytes32 b = bytes32(x);
    uint16 offset = 0;
    for (uint16 i=0; i < 32; i++) {
      c[i] = b[i];
      if (c[i] > 0 && offset == 0) {
        offset = i;
      }
    }
    uint16 size = 32 - offset;
    return (c.slice(offset, 32-offset), size);
  }
}
