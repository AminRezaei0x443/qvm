import * as networking from "networking";


let cli = new networking.HttpClient("GET", "https://google.com");
print(cli.download());