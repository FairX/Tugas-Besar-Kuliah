from socket import *
import sys
import threading

serverSocket = socket(AF_INET, SOCK_STREAM)
serverPort = 6789
serverSocket.bind(('', serverPort))
serverSocket.listen(5)

def handle_client(connectionSocket):
    try:
        message = connectionSocket.recv(1024).decode()
        filename = message.split()[1]
        f = open(filename[1:])
        outputdata = f.read()
        connectionSocket.send('Connected to server'.encode())
        connectionSocket.send(outputdata.encode())

        print('Connection Succeed')

    except IOError:
        connectionSocket.send('Failed to connect'.encode())

        print('Connection Failed')
    
    connectionSocket.close()

print('Server is up: ')
no = 1
while True:
    connectionSocket, addr = serverSocket.accept()
    client_thread = threading.Thread(target=handle_client, args=(connectionSocket,))
    client_thread.start()
    print('Client', no)
    no=no+1
