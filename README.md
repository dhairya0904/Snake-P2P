# Snake-P2P

# Peer to peer Snake game

## Steps to play
* go build
* Run `./snake --port {portNumber}`
This will return connection string like this `/ip4/192.168.29.69/tcp/6666/p2p/QmUuR16jLd44NvCcFLAdygsMymGjfdan3XhkHQUAuuw8wE`
* Now, run `./snake --port {portNumber} --peer /ip4/192.168.29.69/tcp/6666/p2p/QmUuR16jLd44NvCcFLAdygsMymGjfdan3XhkHQUAuuw8wE` on another terminal. Other device should be on same network.
* Play and enjoy.
     
     <img width="1512" alt="Screenshot 2023-08-02 at 6 32 56 PM" src="https://github.com/dhairya0904/Snake-P2P/assets/19638959/c4283c7c-990a-4152-9007-123f72b8616c">
