# Blog

A simple blockchain application that allows users to create, update, and delete blog posts. Only contains one custom x/ module, `blog`.

## Getting started

Starting from the `blog` directory:

- Install the chain binary with `make install`
- Initialize the genesis files and run a set of tests `make script-test` 
   - Contains a complete script test of all the actions that blog module provides.
- To run unit tests `make test-unit`
  - Contains all the unit tests for additional features added.

## Tasks and Fixes

- Module Visibility 
  - Problem: The blog module does not show up in `blogd query` and `blogd tx` commands. 
  - Solution: Add the blog module to app/app.go and app/app_config.go to wire it into the application. 
- Post Overwriting 
  - Problem: New blog posts overwrite existing posts. 
  - Solution: Increment the `post count` to ensure each post has a unique key.
- Post Update Issue
  - Problem: Updating a post doesn't work correctly. 
  - Solution: Fix the issue by ensuring the correct key is used to store the updated blog.
- Timestamps 
  - Requirement: Add `created_at` and `last_updated_at` timestamps to the blog post. 
  - Implementation:
    - Add `created_at` and `last_updated_at` fields in `proto/blog/blog/post.proto`. 
    - Update `created_at` when a post is created and `last_updated_at` everytime when a post is updated.
- Grant Authorization (Bonus)
  - Requirement: Allow post creators to grant other addresses the ability to update or delete their posts. 
  - Implementation:
    - Add an `editors` field (a slice of strings) to the Post struct.
    - Set the creator as the default editor.
    - Implement two new actions:
      - `AddEditor(signer, id, editor_address)`: Allows the creator to add additional editors.
      - `DeleteEditor(signer, id, editor_address)`: Allows the creator to remove editors (excluding themselves).
    - Editors can perform `UpdatePost` and `DeletePost` actions.

## Useful commands

- `rm -rf ~/blog/` - Remove the chain data

### Transactions

- `blogd tx blog create-post hello world --from alice --chain-id blog` - Create a new post
- `blogd tx blog update-post "Hello" "Cosmos" 1 --from alice --chain-id blog` - Update a post
- `blogd tx blog delete-post 1 --from alice  --chain-id blog` - Delete a post
- `blogd tx blog add-editor 1$(blogd keys show $BOB --keyring-backend $KEYRING --output json | jq -r '.address') --from alice --chain-id blog` - Add Editor
- `blogd tx blog update-post "Hello from Editor" "Cosmos is the best ecosystem to develop in as it can give fine control on what action can be baked into the blockchain" 1 --from bob --chain-id blog` - Update a post from editor (bob)
- `blogd tx blog delete-editor 1$(blogd keys show $BOB --keyring-backend $KEYRING --output json | jq -r '.address') --from alice --chain-id blog` - Delete Editor

## Queries

- `blogd q blog show-post 0` - Show a post
- `blogd q blog list-post` - List all posts
