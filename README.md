# hashmeme

It's the year 2099 and historians are debating who first published this and that meme.

If only there was a system that could trace origins of memes...

Introducing `HashMeme`, a tool for keeping an immutable logs of memes using Hedera services.

![](./image_processor/resources/hashmeme.png)

# Nix Installation (optional)

This uses `nix` to get a development environment up.

This is optional but highly recommended to get up a reproducible development environment.

For more detailed instructions for your system see: https://nixos.org/download.html

```sh
make install
```

## Start nix daemon

Start a daemon in one terminal window
```sh
sudo /nix/var/nix/profiles/default/bin/nix-daemon
```

## Start a nix-shell

This is your development environment
```
nix-shell --pure
```

# Building

```sh

make build

# Check ./target for runnable
```

# Making a GUI frontend

```sh
make gui

```

```sh
# or if not using nix:
make gui-nonix
```

# Running tests

```
make tests
```

