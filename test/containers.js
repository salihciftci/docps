const chai = require("chai");
const chaiHttp = require("chai-http");
const expect = chai.expect;
const server = require("../bin/www");
const util = require("../lib/util/cmd");
chai.use(chaiHttp);

let token = "";
let container = "";

describe("Containers", () => {
    before(async () => {
        let user = {
            "password": "test"
        };

        chai.request(server)
            .post("/api/users/test")
            .send(user)
            .end((err, res) => {
                expect(err).to.be.null;
                token = res.body.token;
                expect(res).to.have.status(200);
            });


        let command = "docker run --name test hello-world";
        await util.execCommand(command);
    });

    after(async () => {
        let command = "docker rm -f test";
        await util.execCommand(command);
    });

    describe("List containers", () => {
        it("it should get containers", (done) => {
            chai.request(server)
                .get("/api/containers")
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    expect(res).to.be.json;
                    container = res.body[0].id;
                    done();
                });
        });

        it("it shouldn't get containers", (done) => {
            chai.request(server)
                .get("/api/containers")
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });
    });

    describe("Container Information", () => {
        it("it should get container information", (done) => {
            chai.request(server)
                .get(`/api/containers/${container}`)
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    expect(res).to.be.json;
                    done();
                });
        });

        it("it shouldn't get container information", (done) => {
            chai.request(server)
                .get(`/api/containers/${container}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });

    });

    describe("Logs", () => {
        it("it should get container logs", (done) => {
            chai.request(server)
                .get(`/api/containers/${container}/logs`)
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    expect(res).to.be.json;
                    done();
                });
        });

        it("it shouldn't get container logs", (done) => {
            chai.request(server)
                .get(`/api/containers/${container}/logs`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });
    });
});