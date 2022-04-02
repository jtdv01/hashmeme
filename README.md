# hashmeme

It's the year 2099 and historians are debating who first published this and that meme.

If only there was a system that could trace origins of memes...

Introducing `hashmeme`, an image watermarking tool that makes use of Hedera services.


# Start backend server

```sh
TODO
```

# Bootstrap token

TODO

# Start frontend

TODO

# Nix Installation (optional)

This uses `nix` to get a development environment up.

This is optional but highly recommended to get up a reproducible development environment.

For more detailed instructions for your system see: https://nixos.org/download.html
```sh
# For linux multi-user installation:
sh <(curl -L https://nixos.org/nix/install) --daemon
```

## Start nix daemon

Start a daemon in one terminal window
```sh
sudo sudo /nix/var/nix/profiles/default/bin/nix-daemon
```

## Start a nix-shell

This is your development environment
```
nix-shell --pure
```
