#!/bin/bash

./scripts/init.sh

sleep 10

# Utility function to check if a value matches the expected value
check_equal() {
  local key="$1"
  local value="$2"
  local expected="$3"

  if [ "$value" == "$expected" ]; then
    echo "✅ $key is correct: $expected."
  else
    echo "❌ $key mismatch. Expected: $expected, Got: $value"
    exit 1
  fi
}

# Variables for post titles, bodies, and addresses
TITLE="test title"
TITLE_UPDATED="test title updated"
TITLE2="test title 2"
BODY="test body of the first blog post"

# Constants for users
ALICE="alice"
BOB="bob"
CHAIN_ID="blog"
KEYRING="test"

# Create the first post as Alice
echo "\n 📨 Creating the first post by Alice..."
TEMP=$(blogd tx blog create-post "$TITLE" "$BODY" --from $ALICE --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

# Validate the first post body
BODY_NEW=$(blogd q blog show-post 1 --output json | jq -r '.post.body')
check_equal "Post Body" "$BODY_NEW" "$BODY"

# Create the second post as Alice
echo "\n 📨 Creating the second post by Alice..."
TEMP=$(blogd tx blog create-post "$TITLE2" "$BODY" --from $ALICE --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

# Validate the second post title
TITLE_NEW=$(blogd q blog show-post 2 --output json | jq -r '.post.title')
check_equal "Second Post Title" "$TITLE_NEW" "$TITLE2"

# Update the first post's title as Alice
echo "\n 🔄 Updating the first post title by Alice..."
TEMP=$(blogd tx blog update-post "$TITLE_UPDATED" "$BODY" 1 --from $ALICE --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

# Validate the updated title
TITLE_NEW_UPDATED=$(blogd q blog show-post 1 --output json | jq -r '.post.title')
check_equal "Updated Post Title" "$TITLE_NEW_UPDATED" "$TITLE_UPDATED"

# Add Bob as an editor to the first post
echo "\n ➕ Adding Bob as an editor..."
BOB_ADDRESS=$(blogd keys show $BOB --keyring-backend $KEYRING --output json | jq -r '.address')
TEMP=$(blogd tx blog add-editor 1 "$BOB_ADDRESS" --from $ALICE --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

# Validate that Bob was added as an editor
ADDED_EDITOR_ADDRESS=$(blogd q blog show-post 1 --output json | jq -r '.post.editors[-1]')
check_equal "Bob's Address as Editor" "$ADDED_EDITOR_ADDRESS" "$BOB_ADDRESS"

# Update the first post as Bob
echo "\n 🔄 Updating the post title as Bob..."
TEMP=$(blogd tx blog update-post "$TITLE_NEW_UPDATED" "$BODY" 1 --from $BOB --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

# Validate Bob's update
EDITOR_UPDATED_TITLE=$(blogd q blog show-post 1 --output json | jq -r '.post.title')
check_equal "Editor Updated Title" "$EDITOR_UPDATED_TITLE" "$TITLE_NEW_UPDATED"

# Remove Bob as an editor from the first post
echo "\n Removing Bob as an editor..."
TEMP=$(blogd tx blog delete-editor 1 "$BOB_ADDRESS" --from $ALICE --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

# Validate Alice is the last editor
ALICE_ADDRESS=$(blogd keys show $ALICE --keyring-backend $KEYRING --output json | jq -r '.address')
EDITOR_ADDRESS=$(blogd q blog show-post 1 --output json | jq -r '.post.editors[-1]')
check_equal "Alice's Address as Editor" "$EDITOR_ADDRESS" "$ALICE_ADDRESS"

# Delete the first post as Alice
echo "\n 🗑️ Deleting the first post..."
TEMP=$(blogd tx blog delete-post 1 --from $ALICE --keyring-backend $KEYRING --chain-id $CHAIN_ID -y)
sleep 6

## Validate the post was deleted
RES=$(blogd q blog list-post --output json | jq '.post | length')
EXPECTED=1
check_equal "Total posts check" "$EXPECTED" "$RES"

echo "\n 🎉 All operations completed successfully!"
