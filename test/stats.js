const chai = require("chai");
const chaiHttp = require("chai-http");
const expect = chai.expect;
const server = require("../bin/www");
chai.use(chaiHttp);

let token = "";

describe("Stats", () => {
    before((done) => {
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
                done();
            });
    });

    describe("List stats", () => {
        it("it should get stats", (done) => {
            chai.request(server)
                .get("/api/stats")
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    expect(res).to.be.json;
                    done();
                });
        });

        it("it shouldn't get stats", (done) => {
            chai.request(server)
                .get("/api/stats")
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });
    });
});