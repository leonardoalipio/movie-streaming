#!/bin/bash
set -e

# Wait for MongoDB to start
echo "Waiting for MongoDB to start..."
sleep 5

# Use mongosh to create databases and collections, then import data
mongosh --eval "
  // Switch to movie_streaming database
  db = db.getSiblingDB('movie_streaming');

  // Import genres
  print('Importing genres...');
  db.genres.insertMany($(cat /seed/genres.json));

  // Import movies
  print('Importing movies...');
  db.movies.insertMany($(cat /seed/movies.json));

  // Import users
  print('Importing users...');
  db.users.insertMany($(cat /seed/users.json));

  // Import rankings
  print('Importing rankings...');
  db.rankings.insertMany($(cat /seed/rankings.json));

  print('Seed data imported successfully!');
"
