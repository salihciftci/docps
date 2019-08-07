const chai = require("chai");
const chaiHttp = require("chai-http");
const expect = chai.expect;
const server = require("../bin/www");
const knex = require("../db").knex;
chai.use(chaiHttp);

let token = "";

describe("Users", () => {
    describe("Login", () => {
        it("it should login", (done) => {
            let user = {
                "password": "test"
            };

            chai.request(server)
                .post("/api/users/test")
                .send(user)
                .end((err, res) => {
                    token = res.body.token;
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    done();
                });
        });

        it("it shouldn't login", (done) => {
            let user = {};
            chai.request(server)
                .post("/api/users/test")
                .send(user)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });
    });

    describe("User information", () => {
        it("it should get user information", (done) => {
            chai.request(server)
                .get("/api/users/test")
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    expect(res).to.be.json;
                    done();
                });
        });

        it("it shouldn't get user information", (done) => {
            chai.request(server)
                .get("/api/users/test")
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });
    });

    describe("Create new user", () => {
        it("it should create new user", (done) => {
            let user = {
                "username": "new",
                "password": "test"
            };

            chai.request(server)
                .post("/api/users/")
                .set("authorization", `Bearer ${token}`)
                .send(user)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    done();
                });
        });

        it("it shouldn't create same user", (done) => {
            let user = {
                "username": "new",
                "password": "test"
            };

            chai.request(server)
                .post("/api/users/")
                .set("authorization", `Bearer ${token}`)
                .send(user)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(409);
                    done();
                });
        });


        it("it shouldn't create new user", (done) => {
            let user = {
                "username": "new",
                "password": "test"
            };

            chai.request(server)
                .post("/api/users")
                .send(user)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });

        after((done) => {
            knex("users").where("username", "new").del().then();
            done();
        });
    });

    describe("Edit User", () => {
        after((done) => {
            let user = {
                "username": "test",
                "password": "test",
                "email": "test@test.com"
            };

            chai.request(server)
                .patch("/api/users/test")
                .send(user)
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    done();
                });
        });

        it("it should change user email", (done) => {
            let user = {
                "email": "new@test.com"
            };

            chai.request(server)
                .patch("/api/users/test")
                .send(user)
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    done();
                });
        });

        it("it should change user username", (done) => {
            let user = {
                "username": "test"
            };

            chai.request(server)
                .patch("/api/users/test")
                .send(user)
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    done();
                });
        });

        it("it should change user password", (done) => {
            let user = {
                "password": "123"
            };

            chai.request(server)
                .patch("/api/users/test")
                .send(user)
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    done();
                });
        });
    });
});