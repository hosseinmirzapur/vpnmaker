# VPN Maker

## Quickstart

First make sure you have Go version 1.20 or higher installed on your system

```bash
# easily install it on ubuntu via snapcraft
# try running it with --classic flag if any error happened
sudo snap install go

```

Then clone this repository and after navigating into the cloned folder, run below commands.

```bash
# clone this repository
git clone https://github.com/hosseinmirzapur/vpnmaker

# cd into it
cd vpnmaker

# make vpn.sh executable
chmod +x vpn.sh && ./vpn.sh

# build and run the vpnmaker
go build -o bin/vpnmaker
./bin/vpnmaker

# OR
# Build and run the project using one make command
make run

```

> Note: choose first option by entering number **1**, if prompted (which you will be!).

After the commands being successfully done, you shall get your config in the created.json file on the path you ran the commands!

## Hiddify Config Usage

Be sure to check the [Latest Hiddify-Next App For Your Desired Platform](https://github.com/hiddify/hiddify-next/releases) and download the file, then easily copy/paste the config from previous level inside the app, and enjoy your experience :)