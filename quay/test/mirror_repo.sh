#!/bin/bash

# Old and new repository info
old_repo="quay.io/dlawton/kuadrant-operator"
new_repo="quay.io/test_org_123/kuadrant-operator"

# Get list of tags from the old repository
tags=$(curl -s -u $ROBOT_USER:$ROBOT_PASS "https://quay.io/api/v1/repository/dlawton/kuadrant-operator/tag/?onlyActiveTags=true" | jq -r '.tags[].name')

# Loop through each tag and migrate it
for tag in $tags; do
  echo "Migrating tag: $tag"
  
  docker pull $old_repo:$tag
  docker tag $old_repo:$tag $new_repo:$tag
  docker push $new_repo:$tag
  
  echo "Successfully migrated tag: $tag"
done
