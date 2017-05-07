#!/usr/bin/env python3

# https://docs.python.org/3.6/howto/sockets.html#using-a-socket
# https://docs.python.org/3/howto/unicode.html

import socket

MSGLEN = 1024


def convert_to_bytes(msg):
    msg_as_bytes = msg if isinstance(msg, bytes) else msg.encode('utf-8')
    if len(msg_as_bytes) != len(msg):
        fmt = 'Binary conversion added {} bytes.'
        print(fmt.format(len(msg_as_bytes) - len(msg)))
    return msg_as_bytes


class MySocket:
    """demonstration class only - coded for clarity, not efficiency"""

    def __init__(self, sock=None):
        self.sock = sock or socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def connect(self, host='127.0.0.1', port=8081):
        self.sock.connect((host, int(port)))

    def mysend(self, msg='Mr. Watson--come here--I want to see you.  '):
        assert msg, 'You gotta give me something to work with here.'
        msg_as_bytes = convert_to_bytes(msg)
        MSGLEN = len(msg_as_bytes)
        print('Sending {} bytes: {}', MSGLEN, msg)
        totalsent = 0
        while totalsent < len(msg_as_bytes):
            sent = self.sock.send(msg[totalsent:])
            if sent == 0:
                raise RuntimeError("socket connection broken")
            totalsent += sent
        print('Sent {} bytes'.format(totalsent))

    def myreceive(self):
        chunks = []
        bytes_recd = 0
        while bytes_recd < MSGLEN:
            chunk = self.sock.recv(min(MSGLEN - bytes_recd, 2048))
            if chunk == b'':
                raise RuntimeError("socket connection broken")
            chunks.append(chunk)
            bytes_recd = bytes_recd + len(chunk)
        return b''.join(chunks).decode('utf-8')


s = 'Lüsai'
print(len(s), type(s), s)  # 5 class 'str'> Lüsai
# s = s.encode('utf-8')
s = convert_to_bytes(s)
print(len(s), type(s), s)  # 6 <class 'bytes'> b'L\xc3\xbcsai'
# exit()

# ===

print(0)
sock = MySocket()
print(1)
sock.connect()  # ConnectionRefusedError: [Errno 61] Connection refused
print(2)
sock.mysend()
print(3)
print(sock.myreceive())
