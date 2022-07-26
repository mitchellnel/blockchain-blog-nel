# blockchainblognel

**blockchainblognel** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Goals for this repository:

-   Create a blockchain with a module that lets us write to and read data from the blockchain

## Get started

First, clone this repository:

```
git clone https://github.com/mitchellnel/blockchain-blog-nel
```

Then, invoke:

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

After the blockchain has started, open a separate terminal window, and invoke:

```
blockchain-blog-neld tx blockchainblognel create-post "Lewis Hamilton" "7-time world champion" \
  --from alice --chain-id blockchainblognel
```

This sends a transaction from the account with name `alice` to create a post with `Title: "Lewis Hamilton"` and `Body: "7-time world champion"`.

To display all posts currently on the blockchain blog, invoke:

```
blockchain-blog-neld q blockchainblognel posts
```

And this will currently output:

```
Post:
- body: 7-time world champion
  createdAt: "3221"
  creator: cosmos1fs03pn8htxzhk5r3qr3c7hp5c6sp7uhxswrlc8
  id: "0"
  title: Lewis Hamilton
body: ""
pagination:
  next_key: null
  total: "1"
title: ""
```

Then, to create a comment on this post, invoke:

```
blockchain-blog-neld tx blockchainblognel create-comment 0 "Number of Wins" "103" \
  --from alice --chain-id blockchainblognel
```

This sends a transaction from the account with name `alice` to create a post on the post with `id: 0` with `Title: "Number of Wins"` and `Body: "103"`.

Then, to display the posts on a post of a particular ID, invoke:

```
blockchain-blog-neld q blockchainblognel comments 0
```

And this will currently output:

```
comments:
- body: "103"
  createdAt: "3302"
  creator: cosmos1fs03pn8htxzhk5r3qr3c7hp5c6sp7uhxswrlc8
  id: "0"
  postID: "0"
  title: Number of Wins
pagination:
  next_key: null
  total: "1"
post:
  body: 7-time world champion
  createdAt: "3221"
  creator: cosmos1fs03pn8htxzhk5r3qr3c7hp5c6sp7uhxswrlc8
  id: "0"
  title: Lewis Hamilton
```
