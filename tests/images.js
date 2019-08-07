const chai = require("chai");
const chaiHttp = require("chai-http");
const expect = chai.expect;
const server = require("../bin/www");
const util = require("../lib/util/cmd");
chai.use(chaiHttp);

let token = "";

describe("Images", () => {
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

    after(async () => {
        let command = "docker image rm hello-world:latest";
        await util.execCommand(command);
    });

    describe("List images", () => {
        it("it should get images", (done) => {
            chai.request(server)
                .get("/api/images")
                .set("authorization", `Bearer ${token}`)
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(200);
                    expect(res).to.be.json;
                    done();
                });
        });

        it("it shouldn't get images", (done) => {
            chai.request(server)
                .get("/api/images")
                .end((err, res) => {
                    expect(err).to.be.null;
                    expect(res).to.have.status(401);
                    done();
                });
        });
    });
});