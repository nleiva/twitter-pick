# Pick a Winner from a Tweet

## Make Tweeter API Token available

You can get one from the [Developer Portal](https://developer.twitter.com/).

```bash
export TWITTER_TOKEN="..."
```

## Select Conversation ID

And a user to pull it from.

```bash
# Conversation ID
export TWITTER_CID="..."

# User ID
export TWITTER_USER="..."
```

## Running

```go
 go run main.go
```