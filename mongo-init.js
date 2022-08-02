dbStori = db.getSiblingDB(_getEnv("MONGO_INITDB_DATABASE"));

dbStori.auth({
  user: _getEnv("MONGO_INITDB_ROOT_USERNAME"),
  pwd: _getEnv("MONGO_INITDB_ROOT_PASSWORD"),
  mechanisms: ["SCRAM-SHA-1"],
  digestPassword: true,
});

// Create DB and collection
db = new Mongo().getDB(_getEnv("MONGO_INITDB_DATABASE"));
db.createCollection(_getEnv("INITDB_COLLECTION"), { capped: false });
