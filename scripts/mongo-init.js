print("Started adding the users.");

db = db.getSiblingDB("admin");

db.createUser({
    user: "dev",
    pwd: "passdev",
    roles: ["root"]
});
print("End adding the user roles.");


