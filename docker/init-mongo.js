// Switch to movie_streaming database
db = db.getSiblingDB('movie_streaming');

// Import genres
print('Importing genres...');
try {
  const genresData = JSON.parse(require('fs').readFileSync('/seed/genres.json', 'utf8'));
  db.genres.insertMany(genresData);
  print('✓ Genres imported: ' + genresData.length + ' documents');
} catch (e) {
  print('! Genres import skipped or already exists: ' + e.message);
}

// Import movies
print('Importing movies...');
try {
  const moviesData = JSON.parse(require('fs').readFileSync('/seed/movies.json', 'utf8'));
  db.movies.insertMany(moviesData);
  print('✓ Movies imported: ' + moviesData.length + ' documents');
} catch (e) {
  print('! Movies import skipped or already exists: ' + e.message);
}

// Import users
print('Importing users...');
try {
  const usersData = JSON.parse(require('fs').readFileSync('/seed/users.json', 'utf8'));
  db.users.insertMany(usersData);
  print('✓ Users imported: ' + usersData.length + ' documents');
} catch (e) {
  print('! Users import skipped or already exists: ' + e.message);
}

// Import rankings
print('Importing rankings...');
try {
  const rankingsData = JSON.parse(require('fs').readFileSync('/seed/rankings.json', 'utf8'));
  db.rankings.insertMany(rankingsData);
  print('✓ Rankings imported: ' + rankingsData.length + ' documents');
} catch (e) {
  print('! Rankings import skipped or already exists: ' + e.message);
}

print('\n✓ Seed data initialization completed!');
