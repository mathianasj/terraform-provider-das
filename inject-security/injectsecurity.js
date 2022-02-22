import fs from 'fs'
import YAML from 'yaml'
import fetch from 'node-fetch';

let url = "https://docs.styra.com/openapi/v2";

let settings = { method: "Get" };

fetch(url, settings)
    .then(res => res.json())
    .then((json) => {
       // add security
       json.securityDefinitions = {
           basicAuth: {
               type: "basic"
           }
       };

       // iterate each path and method adding security requirements
       for(const path in json.paths) {
        for(const method in json.paths[path]) {
            json.paths[path][method].security = [
                {
                    basicAuth: []
                }
            ]
        }
       }

       var output = YAML.stringify(json);
       fs.writeFileSync("../swagger.yaml", output);
    });
