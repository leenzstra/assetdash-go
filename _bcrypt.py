import argparse
import sys
import bcrypt
import json
import base64
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad

# Helper script for AssetDash arcade game data encryptions

def run():
    parser = argparse.ArgumentParser()
    parser.add_argument('-i, --session', default="", dest='session')
    parser.add_argument('-c', '--coins', default=0, dest='coins')
    parser.add_argument('-s', '--score', default=0, dest='score')

    args = parser.parse_args()
    if args.session == "":
        sys.stdout.write(json.dumps(
            {"err": True, "msg": "missing session_id"}))
        sys.exit(1)

    session_id = args.session
    earned_coins = int(args.coins)
    score = int(args.score)

    try:
        session_hash, data = encrypt(session_id, earned_coins, score)
        sys.stdout.write(json.dumps(
            {"err": False, "msg": "ok", "session_hash": session_hash, "data": data}))
    except Exception as e:
        sys.stdout.write(json.dumps(
            {"err": True, "msg": str(e)}))
        sys.exit(1)


def encrypt(session_id: str, coins: int, score: int) -> tuple[str, str]:
    # Encrypt game results with AES
    BLOCK_SIZE = 32
    keyhex = '603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4'
    data = json.dumps({"earned_coins": coins,
                       "score": score}, separators=(',', ':'))

    aescipher = AES.new(bytes.fromhex(keyhex), AES.MODE_ECB)

    enc_data = aescipher.encrypt(pad(data.encode(), BLOCK_SIZE))
    enc_data_b64 = base64.b64encode(enc_data).decode()

    # Encrypt game session data with bcrypt
    salt = bytes([0x24, 0x32, 0x62, 0x24, 0x31, 0x32, 0x24, 0x41,
                  0x4d, 0x61, 0x6b, 0x35, 0x34, 0x70, 0x48, 0x57, 0x4c, 0x76,
                  0x51, 0x6a, 0x2f, 0x67, 0x6a, 0x72, 0x6d, 0x30, 0x6b, 0x4e, 0x4f])

    hashme = (session_id+enc_data_b64).replace("'", '')

    session_hash = bcrypt.hashpw(hashme.encode(), salt)[-31:].decode()

    return (session_hash, enc_data_b64)


if __name__ == "__main__":
    run()
