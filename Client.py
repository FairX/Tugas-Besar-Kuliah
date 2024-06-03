import sys
from socket import *

serverName = sys.argv[1]
serverPort = int(sys.argv[2])
filename = sys.argv[3]

clientSocket = socket(AF_INET, SOCK_STREAM)
clientSocket.connect((serverName, serverPort))

request = f'GET /{filename} HTTP/1.1\r\nHost: {serverName}\r\n\r\n'
clientSocket.send(request.encode())

print('From Server: ')

response = clientSocket.recv(4096).decode()
while response:
    print(response)
    response = clientSocket.recv(4096).decode()

input("Press Enter to close the connection...")

clientSocket.close()
