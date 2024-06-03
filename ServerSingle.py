from socket import *
import sys

serverSocket = socket(AF_INET, SOCK_STREAM)
serverPort = 6789
serverSocket.bind(('', serverPort))
serverSocket.listen(1)

print('Server is up')
connectionSocket, addr =  serverSocket.accept()
try:
    message = connectionSocket.recv(1024).decode()
    filename = message.split()[1]
    f = open(filename[1:])
    outputdata = f.read()
    connectionSocket.send('Connected to server'.encode())
    connectionSocket.send(outputdata.encode())

    print('Connection Succeed')
    connectionSocket.close()

except IOError:
    connectionSocket.send('Failed to connect'.encode())

    print('Connection Failed')
    connectionSocket.close()

serverSocket.close()
sys.exit()